package handleMessage

import (
	"data-center-v2/common"
	"data-center-v2/config"
	"data-center-v2/database"
	myErr "data-center-v2/error"
	"data-center-v2/proto/execTxnRpc"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

/*定义全局变量*/
var (
	timeMap       = make(map[string]chan bool, 0) /*用于超时删除缓存所用*/
	mutex         sync.RWMutex
	receivedNodes = make(map[string]map[string]*common.TreeNode, 0) /*表示目前已接收到的树节点，即子事务节点*/
	tagged        = make(map[string]map[string]bool, 0)             /*表示目前已完成标记的树节点*/
	untagged      = make(map[string]map[string]bool, 0)             /*表示目前尚未标记完成的树节点*/
)

/*事务中的每一个服务均将本服务的事务信息传递至事务管理系统中的该方法中*/
func HandleMessage(message *execTxnRpc.TxnMessage) (err error) {
	/*若当前服务不在线，则立即删除缓存，并返回error*/
	if !message.Online {
		log.Infof("Current Service %s is not online or available. Deleting caches...", message.ServiceUuid)
		mutex.RLock()
		_, ok := receivedNodes[message.TreeUuid]
		mutex.RUnlock()
		if ok {
			/*清除当前事务树的缓存*/
			mutex.Lock()
			delete(receivedNodes, message.TreeUuid)
			delete(tagged, message.TreeUuid)
			delete(untagged, message.TreeUuid)
			delete(timeMap, message.TreeUuid)
			mutex.Unlock()
			log.Info("caches deleted.")
		} else {
			log.Errorf("no caches found. cannot delete caches at TreeUUID: %s", message.TreeUuid)
		}
		return myErr.NewError(200, "Current Service "+message.ServiceUuid+" is not online or available.")
	}
	mutex.RLock()
	_, ok := receivedNodes[message.TreeUuid]
	mutex.RUnlock()
	/*若当前receivedNodes map中包含该事务树ID，则开始向该事务n叉树中添加该节点*/
	if ok {
		currTreeNode := &common.TreeNode{
			Name:         message.ServiceUuid,
			ParentName:   message.ParentUuid,
			ChildrenName: message.Children,
		}
		err := addTreeNode(message, currTreeNode)
		if err != nil {
			return err
		}
	} else { /*若receivedNodes map中不包含该事务树ID，
		则表明该条信息为该事务中的第一条事务信息，
		则需要新建一个事务树，并将该事务树的ID放入receivedNodes map 中，
		其key值即为事务树ID，其value值即为该事务树所对应的哈希表，
		以方便查找该事务树中的任意节点。
		同时，在tagged map和untagged map中新建该事务树ID，
		对应的value值即为当前事务树所对应的已标记节点名称和未标记节点名称的哈希表*/
		mutex.Lock()
		receivedNodes[message.TreeUuid] = make(map[string]*common.TreeNode, 0)
		tagged[message.TreeUuid] = make(map[string]bool, 0)
		untagged[message.TreeUuid] = make(map[string]bool, 0)
		mutex.Unlock()
		currTreeNode := &common.TreeNode{
			Name:         message.ServiceUuid,
			ParentName:   message.ParentUuid,
			ChildrenName: message.Children,
		}
		err := addTreeNode(message, currTreeNode)
		if err != nil {
			return err
		}
		mutex.Lock()
		timeMap[message.TreeUuid] = make(chan bool, 1)
		mutex.Unlock()
		/*新建一个计时器以用于该事务的超时处理，在接收到第一条该事务的节点信息时开始计时*/
		go timeOut(message.TreeUuid, &err)
	}
	/*判断当前事务是否完整，若完整，则写入数据库，若不完整，则结束该方法，然后等待下一条事务信息的到来*/
	if isCompleteTree(message.TreeUuid) {
		log.Info("message all received, writing into database...")
		/*判断是否超时*/
		mutex.RLock()
		_, ok := timeMap[message.TreeUuid]
		mutex.RUnlock()
		if ok { /*未超时，则执行事务的写入，并向timeMap中该
			事务树ID所对应的管道channel中传递true，表示已收到所有信息，
			马上可以执行写入，可以关闭计时器了*/
			mutex.Lock()
			timeMap[message.TreeUuid] <- true
			mutex.Unlock()
			mutex.RLock()
			root := receivedNodes[message.TreeUuid]["root"] /*获取该事务树的根节点*/
			mutex.RUnlock()
			err = database.Write(root) /*从根节点开始，层序遍历写入数据库中*/
			if err != nil {
				log.Error("write into database failed, deleting caches...")
			} else {
				log.Info("write into database success, deleting caches...")
			}
		} else { /*事务已超时，不执行写入*/
			log.Error("current service chain already timed out. deleting caches...")
		}
		ok = false
		mutex.RLock()
		_, ok = receivedNodes[message.TreeUuid]
		mutex.RUnlock()
		if ok { /*清除缓存*/
			mutex.Lock()
			delete(receivedNodes, message.TreeUuid)
			delete(tagged, message.TreeUuid)
			delete(untagged, message.TreeUuid)
			delete(timeMap, message.TreeUuid)
			mutex.Unlock()
			log.Info("caches deleted.")
		} else {
			log.Errorf("no caches found. cannot delete caches at TreeUUID: %s", message.TreeUuid)
		}
	}
	return err
}

/*超时处理方法*/
func timeOut(treeUuid string, err *error) {
	mutex.RLock()
	timeChan := timeMap[treeUuid]
	mutex.RUnlock()
	defer close(timeChan) // 使用完该通道后，必须关闭该通道。GO的GC不会回收通道。
	select {              /*判断是否超时*/
	case <-timeChan:
		log.Info("receiving message succeeded, timeOut function stopped, writing into database.")
	case <-time.After(config.TIMELAPSES): /*清除当前事务树的缓存*/
		log.Error("receiving message timed out, deleting caches...")
		mutex.RLock()
		_, ok := receivedNodes[treeUuid]
		mutex.RUnlock()
		if ok {
			mutex.Lock()
			delete(receivedNodes, treeUuid)
			delete(tagged, treeUuid)
			delete(untagged, treeUuid)
			delete(timeMap, treeUuid)
			mutex.Unlock()
			log.Info("caches deleted.")
			*err = myErr.NewError(500, "receive message timed out.")
		}
	}
}

/*判断该事务n叉树是否完整(即事务是否完整)的方法*/
func isCompleteTree(treeUUID string) bool {
	/*若已标记的节点和未标记的节点长度均为0，则说明还未收到事务信息，返回false*/
	if len(tagged) == 0 && len(untagged) == 0 {
		return false
	}
	mutex.RLock()
	_, ok := receivedNodes[treeUUID]["root"]
	mutex.RUnlock()
	/*若接收到了根节点，且未标记的节点数为0，则说明收到了所有的子事务节点，返回true*/
	if ok && len(untagged[treeUUID]) == 0 {
		return true
	}
	/*其他情况下，要么是root节点的事务信息还未接收到，要么是untagged map中对应
	事务树ID的map的大小不为0，这两种情况显然均不是事务树完整的情况*/
	return false
}

/*添加子事务节点的方法*/
func addTreeNode(message *execTxnRpc.TxnMessage, node *common.TreeNode) error {
	treeUUID := message.TreeUuid
	mutex.RLock()
	_, ok := receivedNodes[treeUUID][node.Name]
	mutex.RUnlock()
	if ok {
		/*若已包含该节点，则说明存在节点重复调用的情况，不符合幂等性，返回错误*/
		// ERROR!!!一个节点重复调用，不符合幂等性。
		return myErr.NewError(500, "Duplicated Nodes!")
	}
	mutex.Lock()
	/*将该节点放入receivedNodes map中对应事务树ID的HashMap中，以方便查找*/
	receivedNodes[treeUUID][node.Name] = node
	mutex.Unlock()
	ok = false
	mutex.RLock()
	parent, ok := receivedNodes[treeUUID][node.ParentName]
	mutex.RUnlock()
	/*若在receivedNodes map中，该事务树ID下的map中包含该节点的父节点，
	则需将当前节点添加到其父节点的子节点数组中，即Children中
	(注：Children区别于ChildrenName，ChildrenName为事务信息传递过来的子节点名称，
	类型为String HashMap，而Children为树节点，类型为TreeNode数组)*/
	if ok {
		if parent.Children == nil {
			parent.Children = make([]*common.TreeNode, 0)
		}
		parent.Children = append(parent.Children, node)
		node.Parent = parent
		if judgeNode(parent) { /*当把本节点放入父节点的Children节点数组中以后，
			就要判断该节点的父节点是否可以被标记为黑色，如若可以划为已标记节点，则将父节
			点名称从untagged map 中移至tagged map中*/
			mutex.Lock()
			tagged[treeUUID][parent.Name] = true
			mutex.Unlock()
			ok = false
			mutex.RLock()
			_, ok := untagged[treeUUID][parent.Name]
			mutex.RUnlock()
			if ok {
				mutex.Lock()
				delete(untagged[treeUUID], parent.Name)
				mutex.Unlock()
			}
		}
	} else { /*若在receivedNodes map中找不到该父节点，
		那么就需要将该父节点名称放入untagged map中，即标记为白色*/
		if node.ParentName != "" {
			mutex.Lock()
			untagged[treeUUID][node.ParentName] = true
			mutex.Unlock()
		}
	}
	/*对于当前节点的子节点们，首先从ChildrenName中获取其子节点们的唯一标识，
	然后在receivedNodes map中寻找该节点，如若在map中找到，
	则可以从receivedNodes map中通过该标识获取到相应的子节点的TreeNode形式，
	然后放入到当前节点的Children数组中*/
	for _, name := range node.ChildrenName {
		ok = false
		mutex.RLock()
		child, ok := receivedNodes[treeUUID][name]
		mutex.RUnlock()
		if ok {
			if node.Children == nil {
				node.Children = make([]*common.TreeNode, 0)
			}
			node.Children = append(node.Children, child)
			if child.Parent == nil {
				child.Parent = node
			}
			/*放入Children数组之后，同样需要判断一下当前子节点是否可以被划为已标记节点。
			如果可以被划为已标记，则将该节点名称从untagged map中转移至tagged map中。*/
			if judgeNode(child) {
				mutex.Lock()
				tagged[treeUUID][child.Name] = true
				mutex.Unlock()
				ok = false
				mutex.RLock()
				_, ok := untagged[treeUUID][child.Name]
				mutex.RUnlock()
				if ok {
					mutex.Lock()
					delete(untagged[treeUUID], child.Name)
					mutex.Unlock()
				}
			}
		} else { /*如若在receivedNodes　map中没有找到该子节点，
			则直接将该子节点名称放入untagged map之中，即划为未标记节点*/
			mutex.Lock()
			untagged[treeUUID][name] = true
			mutex.Unlock()
		}
	}
	if judgeNode(node) { /*最后判断本节点是否可以被标记为黑色*/
		mutex.Lock()
		tagged[treeUUID][node.Name] = true
		mutex.Unlock()
		ok = false
		mutex.RLock()
		_, ok := untagged[treeUUID][node.Name]
		mutex.RUnlock()
		if ok {
			mutex.Lock()
			delete(untagged[treeUUID], node.Name)
			mutex.Unlock()
		}
	} else {
		mutex.Lock()
		untagged[treeUUID][node.Name] = true
		mutex.Unlock()
	}
	/*对于每一个子事务节点，均生成相对应的SQL语句，并存在当前节点的信息中*/
	node.DbName = message.DbName
	database.GoWriteTX(message, &node.SqlStr)
	return nil
}

/*判断当前节点是否可以被标记为黑色的方法*/
func judgeNode(node *common.TreeNode) bool {
	if node.Name == "root" { /*首先判断本节点是否为root节点*/
		/*若是root节点的同时，其子节点与子节点名称均为空，则表明本事务只有一个根节点，
		即，已接收完成，不需要判断root节点的父节点（因为root没有父节点），
		直接可以划为已标记节点，返回true*/
		if node.Children == nil && node.ChildrenName == nil {
			return true
		}
		/*其他情况时，则需要满足子节点与子节点名称均不为空，
		且子节点数组的长度要等于子节点名称数组的长度。其中，
		子节点名称数组是通过事务信息传递过来的，因此子节点名
		称数组的长度是固定的，因此需要使用随后期接收到的信息
		而变化的子节点数组的长度与不变的子节点名称数组的长度进行比较。*/
		return node.Children != nil &&
			node.ChildrenName != nil &&
			len(node.Children) == len(node.ChildrenName)
	}
	/*该节点不是根节点时，必须要有父节点。子节点们是否完整的
	判断逻辑与该节点是根节点时的判断逻辑相同*/
	if node.Parent != nil && node.Children == nil && node.ChildrenName == nil {
		return true
	} else {
		return node.Parent != nil &&
			node.Children != nil &&
			node.ChildrenName != nil &&
			len(node.Children) == len(node.ChildrenName)
	}
}

// Through level order sort, we can have a correct SQL execute order.
//func levelOrder(root *common.TreeNode) []*execTxnRpc.TxnMessage {
//	result := make([]*execTxnRpc.TxnMessage, 0)
//	queue := make([]*common.TreeNode, 0)
//	queue = append(queue, root)
//	var tmp *common.TreeNode
//	for len(queue) != 0 {
//		tmp = queue[0]
//		queue = queue[1:]
//		for _, child := range tmp.Children {
//			queue = append(queue, child)
//		}
//		result = append(result, tmp.Info)
//	}
//	return result
//}
