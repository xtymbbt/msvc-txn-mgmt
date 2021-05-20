package handleMessage

import (
	"data-center-v2/common"
	"data-center-v2/config"
	"data-center-v2/database"
	myErr "data-center-v2/error"
	"data-center-v2/proto/commonInfo"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	timeMap       = make(map[string]chan bool, 0)
	mutex         sync.RWMutex
	receivedNodes = make(map[string]map[string]*common.TreeNode, 0)
	tagged        = make(map[string]map[string]bool, 0)
	untagged      = make(map[string]map[string]bool, 0)
)

func HandleMessage(message *commonInfo.HttpRequest) (err error) {
	if !message.Online {
		log.Infof("Current Service %s is not online or available. Deleting caches...", message.ServiceUuid)
		mutex.RLock()
		_, ok := receivedNodes[message.TreeUuid]
		mutex.RUnlock()
		if ok {
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
	if ok {
		currTreeNode := &common.TreeNode{
			Name:         message.ServiceUuid,
			ParentName:   message.ParentUuid,
			ChildrenName: message.Children,
			Info:         message,
		}
		err := addTreeNode(message.TreeUuid, currTreeNode)
		if err != nil {
			return err
		}
	} else {
		mutex.Lock()
		receivedNodes[message.TreeUuid] = make(map[string]*common.TreeNode, 0)
		tagged[message.TreeUuid] = make(map[string]bool, 0)
		untagged[message.TreeUuid] = make(map[string]bool, 0)
		mutex.Unlock()
		currTreeNode := &common.TreeNode{
			Name:         message.ServiceUuid,
			ParentName:   message.ParentUuid,
			ChildrenName: message.Children,
			Info:         message,
		}
		err := addTreeNode(message.TreeUuid, currTreeNode)
		if err != nil {
			return err
		}
		mutex.Lock()
		timeMap[message.TreeUuid] = make(chan bool, 1)
		mutex.Unlock()
		go timeOut(message.TreeUuid, &err)
	}

	if isCompleteTree(message.TreeUuid) {
		log.Info("message all received, writing into database...")
		mutex.RLock()
		_, ok := timeMap[message.TreeUuid]
		mutex.RUnlock()
		if ok {
			mutex.Lock()
			timeMap[message.TreeUuid] <- true
			mutex.Unlock()
			mutex.RLock()
			root := receivedNodes[message.TreeUuid]["root"]
			mutex.RUnlock()
			err = database.Write(root)
			if err != nil {
				log.Error("write into database failed, deleting caches...")
			} else {
				log.Info("write into database success, deleting caches...")
			}
		} else {
			log.Error("current service chain already timed out. deleting caches...")
		}
		ok = false
		mutex.RLock()
		_, ok = receivedNodes[message.TreeUuid]
		mutex.RUnlock()
		if ok {
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

func timeOut(treeUuid string, err *error) {
	mutex.RLock()
	timeChan := timeMap[treeUuid]
	mutex.RUnlock()
	defer close(timeChan) // 使用完该通道后，必须关闭该通道。GO的GC不会回收通道。
	select {
	case <-timeChan:
		log.Info("receiving message succeeded, timeOut function stopped, writing into database.")
	case <-time.After(config.TIMELAPSES):
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

func isCompleteTree(treeUUID string) bool {
	if len(tagged) == 0 && len(untagged) == 0 {
		return false
	}
	mutex.RLock()
	_, ok := receivedNodes[treeUUID]["root"]
	mutex.RUnlock()
	if ok && len(untagged[treeUUID]) == 0 {
		return true
	}
	return false
}

func addTreeNode(treeUUID string, node *common.TreeNode) error {
	mutex.RLock()
	_, ok := receivedNodes[treeUUID][node.Name]
	mutex.RUnlock()
	if ok {
		// ERROR!!!一个节点重复调用，不符合幂等性。
		return myErr.NewError(500, "Duplicated Nodes!")
	}
	mutex.Lock()
	receivedNodes[treeUUID][node.Name] = node
	mutex.Unlock()
	ok = false
	mutex.RLock()
	parent, ok := receivedNodes[treeUUID][node.ParentName]
	mutex.RUnlock()
	if ok {
		if parent.Children == nil {
			parent.Children = make([]*common.TreeNode, 0)
		}
		parent.Children = append(parent.Children, node)
		node.Parent = parent
		if judgeNode(parent) {
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
	} else {
		if node.ParentName != "" {
			mutex.Lock()
			untagged[treeUUID][node.ParentName] = true
			mutex.Unlock()
		}
	}
	for name := range node.ChildrenName {
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
		} else {
			mutex.Lock()
			untagged[treeUUID][name] = true
			mutex.Unlock()
		}
	}
	if judgeNode(node) {
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
	database.GoWriteTX(node.Info, &node.SqlStr)
	return nil
}

func judgeNode(node *common.TreeNode) bool {
	if node.Name == "root" {
		if node.Children == nil && node.ChildrenName == nil {
			return true
		}
		return node.Children != nil &&
			node.ChildrenName != nil &&
			len(node.Children) == len(node.ChildrenName)
	}
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
func levelOrder(root *common.TreeNode) []*commonInfo.HttpRequest {
	result := make([]*commonInfo.HttpRequest, 0)
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	var tmp *common.TreeNode
	for len(queue) != 0 {
		tmp = queue[0]
		queue = queue[1:]
		for _, child := range tmp.Children {
			queue = append(queue, child)
		}
		result = append(result, tmp.Info)
	}
	return result
}
