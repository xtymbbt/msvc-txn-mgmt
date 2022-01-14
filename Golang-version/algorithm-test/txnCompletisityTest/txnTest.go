package txnCompletisityTest

import (
	"algorithm-test/common"
	myErr "algorithm-test/error"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"sync"
	"time"
)

type TreeNode struct {
	name     string
	parent   *TreeNode
	children []*TreeNode
}

/*定义全局变量*/
var (
	mutex         sync.RWMutex
	receivedNodes = make(map[string]*common.TreeNode, 0) /*表示目前已接收到的树节点，即子事务节点*/
	blackSet      = make(map[string]bool, 0)             /*表示目前已完成标记的树节点*/
	//whiteSet      = make(map[string]map[string]bool, 0)             /*表示目前尚未标记完成的树节点*/
)

func tree(层数 int, 分叉数目 int) {
	// 二叉树式
	// 两层
	// 三层
	// 四层
	messages := make([]*common.Message, 0, 0)
	root := &TreeNode{
		name:     "root",
		children: make([]*TreeNode, 0, 0),
	}
	k := 1
	lastLevel := make([]*TreeNode, 0, 0)
	for i := 1; i < 层数; i++ {
		thisLevel := make([]*TreeNode, 0, 0)
		for j := 0; j < int(math.Pow(float64(分叉数目), float64(i))); j++ {
			var child *TreeNode
			if i == 1 {
				child = &TreeNode{
					name:     "child" + strconv.Itoa(k),
					parent:   root,
					children: make([]*TreeNode, 0, 0),
				}
				root.children = append(root.children, child)
			} else {
				child = &TreeNode{
					name:     "child" + strconv.Itoa(k),
					children: make([]*TreeNode, 0, 0),
				}
				for p := 0; p < len(lastLevel); p++ {
					if len(lastLevel[p].children) < 分叉数目 {
						child.parent = lastLevel[p]
						lastLevel[p].children = append(lastLevel[p].children, child)
						break
					}
				}
			}
			k++
			thisLevel = append(thisLevel, child)
		}
		fmt.Println("thisLevel:")
		for i := 0; i < len(thisLevel); i++ {
			fmt.Print(thisLevel[i].name + " ")
		}
		fmt.Println()
		fmt.Println("lastLevel:")
		for i := 0; i < len(lastLevel); i++ {
			fmt.Print(lastLevel[i].name + " ")
		}
		fmt.Println()
		lastLevel = make([]*TreeNode, 0, 0)
		for i := 0; i < len(thisLevel); i++ {
			lastLevel = append(lastLevel, thisLevel[i])
		}
	}
	log.Infof("k: %d", k)

	messages = levelOrder2(root)

	log.Infof("messages: %d", len(messages))

	start := time.Now().UnixNano()
	log.Infof("start: %v ns", start)
	for i := 0; i < len(messages); i++ {
		handleMessage(messages[i])
	}
	end := time.Now().UnixNano()
	log.Infof("end: %v ns", end)
	log.Infof("total: %d ns", end-start)
}

/*事务中的每一个服务均将本服务的事务信息传递至事务管理系统中的该方法中*/
func handleMessage(message *common.Message) (err error) {
	currTreeNode := &common.TreeNode{
		Name:         message.ServiceUuid,
		ParentName:   message.ParentUuid,
		ChildrenName: message.Children,
	}
	addTreeNode(currTreeNode)
	/*判断当前事务是否完整，若完整，则写入数据库，若不完整，则结束该方法，然后等待下一条事务信息的到来*/
	if isCompleteTree() {
		log.Info("事务树已完整, writing into database...")

		root := receivedNodes["root"] /*获取该事务树的根节点*/
		levelOrder(root)
		log.Debugf("%v", root)
		if err != nil {
			log.Error("write into database failed, deleting caches...")
		} else {
			log.Info("write into database success, deleting caches...")
		}
	} else { /*事务已超时，不执行写入*/
		log.Error("事务树尚未完整")
	}

	return err
}

/*判断该事务n叉树是否完整(即事务是否完整)的方法*/
func isCompleteTree() bool {
	if len(receivedNodes) == len(blackSet) {
		return true
	}
	return false
}

/*添加子事务节点的方法*/
func addTreeNode(node *common.TreeNode) error {
	_, ok := receivedNodes[node.Name]
	if ok {
		/*若已包含该节点，则说明存在节点重复调用的情况，不符合幂等性，返回错误*/
		// ERROR!!!一个节点重复调用，不符合幂等性。
		return myErr.NewError(500, "Duplicated Nodes!")
	}
	/*将该节点放入receivedNodes map中对应事务树ID的HashMap中，以方便查找*/
	receivedNodes[node.Name] = node
	ok = false
	parent, ok := receivedNodes[node.ParentName]
	/*若在receivedNodes map中，该事务树ID下的map中包含该节点的父节点，
	则需将当前节点添加到其父节点的子节点数组中，即Children中
	(注：Children区别于ChildrenName，ChildrenName为事务信息传递过来的子节点名称，
	类型为String HashMap，而Children为树节点，类型为TreeNode数组)*/
	if ok {
		if parent.Children == nil {
			parent.Children = make(map[*common.TreeNode]bool, 0)
		}
		ok = false
		ok = parent.Children[node]
		if ok {
			return myErr.NewError(500, "节点重复传入，不符合幂等性！")
		}
		parent.Children[node] = true
		node.Parent = parent
		if judgeNode(parent) { /*当把本节点放入父节点的Children节点数组中以后，
			就要判断该节点的父节点是否可以被标记为黑色，如若可以划为已标记节点，则将父节
			点名称从white set 中移至black set中*/
			blackSet[parent.Name] = true
			//ok = false
			//mutex.RLock()
			//_, ok := whiteSet[treeUUID][parent.Name]
			//mutex.RUnlock()
			//if ok {
			//	mutex.Lock()
			//	delete(whiteSet[treeUUID], parent.Name)
			//	mutex.Unlock()
			//}
		}
	}
	//else { /*若在receivedNodes map中找不到该父节点，
	//	那么就需要将该父节点名称放入white set中，即标记为白色*/
	//if node.ParentName != "" {
	//	mutex.Lock()
	//	whiteSet[treeUUID][node.ParentName] = true
	//	mutex.Unlock()
	//}
	//}
	/*对于当前节点的子节点们，首先从ChildrenName中获取其子节点们的唯一标识，
	然后在receivedNodes map中寻找该节点，如若在map中找到，
	则可以从receivedNodes map中通过该标识获取到相应的子节点的TreeNode形式，
	然后放入到当前节点的Children数组中*/
	for _, name := range node.ChildrenName {
		ok = false
		child, ok := receivedNodes[name]
		if ok {
			if node.Children == nil {
				node.Children = make(map[*common.TreeNode]bool, 0)
			}
			ok = false
			ok = node.Children[child]
			if ok {
				return myErr.NewError(500, "节点重复传入，不符合幂等性！")
			}
			node.Children[child] = true
			if child.Parent == nil {
				child.Parent = node
			}
			/*放入Children数组之后，同样需要判断一下当前子节点是否可以被划为已标记节点。
			如果可以被划为已标记，则将该节点名称从white set中转移至black set中。*/
			if judgeNode(child) {
				blackSet[child.Name] = true
				//ok = false
				//mutex.RLock()
				//_, ok := whiteSet[treeUUID][child.Name]
				//mutex.RUnlock()
				//if ok {
				//	mutex.Lock()
				//	delete(whiteSet[treeUUID], child.Name)
				//	mutex.Unlock()
				//}
			}
		}
		//else { /*如若在receivedNodes　map中没有找到该子节点，
		//	则直接将该子节点名称放入white set之中，即划为未标记节点*/
		//mutex.Lock()
		//whiteSet[treeUUID][name] = true
		//mutex.Unlock()
		//}
	}
	if judgeNode(node) { /*最后判断本节点是否可以被标记为黑色*/
		blackSet[node.Name] = true
		ok = false
		//mutex.RLock()
		//_, ok := whiteSet[treeUUID][node.Name]
		//mutex.RUnlock()
		//if ok {
		//	mutex.Lock()
		//	delete(whiteSet[treeUUID], node.Name)
		//	mutex.Unlock()
		//}
	}
	//else {
	//mutex.Lock()
	//whiteSet[treeUUID][node.Name] = true
	//mutex.Unlock()
	//}
	/*对于每一个子事务节点，均生成相对应的SQL语句，并存在当前节点的信息中*/
	time.Sleep(100 * time.Millisecond)
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
		if node.Children != nil && node.ChildrenName == nil {
			return len(node.Children) == 0
		}
		if node.ChildrenName != nil && node.Children == nil {
			return len(node.ChildrenName) == 0
		}
		return node.Children != nil &&
			node.ChildrenName != nil &&
			len(node.Children) == len(node.ChildrenName)
	}
	/*该节点不是根节点时，必须要有父节点。子节点们是否完整的
	判断逻辑与该节点是根节点时的判断逻辑相同*/
	if node.Parent != nil {
		if node.Children == nil && node.ChildrenName == nil {
			return true
		}
		if node.Children != nil && node.ChildrenName == nil {
			return len(node.Children) == 0
		}
		if node.ChildrenName != nil && node.Children == nil {
			return len(node.ChildrenName) == 0
		}
		return node.Parent != nil &&
			node.Children != nil &&
			node.ChildrenName != nil &&
			len(node.Children) == len(node.ChildrenName)
	} else {
		return false
	}
}

// Through level order sort, we can have a correct SQL execute order.
func levelOrder(root *common.TreeNode) []*common.Message {
	result := make([]*common.Message, 0)
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	var tmp *common.TreeNode
	for len(queue) != 0 {
		tmp = queue[0]
		queue = queue[1:]
		for child := range tmp.Children {
			queue = append(queue, child)
		}
		x := &common.Message{
			ServiceUuid: tmp.Name,
			ParentUuid:  tmp.ParentName,
			Children:    tmp.ChildrenName,
		}
		result = append(result, x)
		time.Sleep(100 * time.Millisecond)
	}
	return result
}

func levelOrder2(root *TreeNode) []*common.Message {
	result := make([]*common.Message, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var tmp *TreeNode
	for len(queue) != 0 {
		tmp = queue[0]
		queue = queue[1:]
		for _, child := range tmp.children {
			queue = append(queue, child)
		}
		x := &common.Message{
			ServiceUuid: tmp.name,
			ParentUuid:  "",
			Children:    make([]string, 0, 0),
		}
		fmt.Printf("tmp.name: %s", tmp.name)
		if tmp.parent != nil {
			x.ParentUuid = tmp.parent.name
			fmt.Printf(" tmp.parent: %s", tmp.parent.name)
		}
		if tmp.children == nil {
			x.Children = nil
		} else {
			fmt.Println()
			fmt.Println("tmp.children:")
			for i := 0; i < len(tmp.children); i++ {
				x.Children = append(x.Children, tmp.children[i].name)
				fmt.Print(tmp.children[i].name + " ")
			}
		}
		fmt.Println()
		fmt.Println()
		result = append(result, x)
	}
	return result
}
