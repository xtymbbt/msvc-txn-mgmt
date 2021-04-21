package handleMessage

import (
	"../../../resources/config"
	"../database"
	myErr "../error"
	"../proto/commonInfo"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type TreeNode struct {
	Name         string
	ParentName   string
	Parent       *TreeNode
	ChildrenName map[string]bool
	Children     []*TreeNode
	Info         *commonInfo.HttpRequest
}

var (
	timeMap       = make(map[string]chan bool, 0)
	mutex         sync.RWMutex
	receivedNodes = make(map[string]map[string]*TreeNode, 0)
	tagged        = make(map[string]map[string]bool, 0)
	untagged      = make(map[string]map[string]bool, 0)
)

func HandleMessage(message *commonInfo.HttpRequest) (err error) {
	if !message.Online {
		log.Infof("Current Service %s is not online or available. Deleting caches...", message.ServiceUuid)
		mutex.Lock()
		if _, ok := receivedNodes[message.TreeUuid]; ok {
			delete(receivedNodes, message.TreeUuid)
			delete(tagged, message.TreeUuid)
			delete(untagged, message.TreeUuid)
			delete(timeMap, message.TreeUuid)
			log.Info("caches deleted.")
		} else {
			log.Errorf("no caches found. cannot delete caches at TreeUUID: %s", message.TreeUuid)
		}
		mutex.Unlock()
		return myErr.NewError(200, "Current Service "+message.ServiceUuid+" is not online or available.")
	}
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := receivedNodes[message.TreeUuid]; ok {
		currTreeNode := &TreeNode{
			Name:         message.ServiceUuid,
			ParentName:   message.ParentUuid,
			ChildrenName: message.Children,
			Info:         message,
		}
		addTreeNode(message.TreeUuid, currTreeNode)
	} else {
		receivedNodes[message.TreeUuid] = make(map[string]*TreeNode, 0)
		tagged[message.TreeUuid] = make(map[string]bool, 0)
		untagged[message.TreeUuid] = make(map[string]bool, 0)
		currTreeNode := &TreeNode{
			Name:         message.ServiceUuid,
			ParentName:   message.ParentUuid,
			ChildrenName: message.Children,
			Info:         message,
		}
		addTreeNode(message.TreeUuid, currTreeNode)
		timeMap[message.TreeUuid] = make(chan bool, 1)
		go timeOut(message.TreeUuid, &err)
	}

	if isCompleteTree(message.TreeUuid) {
		log.Info("message all received, writing into database...")
		if _, ok := timeMap[message.TreeUuid]; ok {
			mutex.Lock()
			timeMap[message.TreeUuid] <- true
			mutex.Unlock()
			mutex.RLock()
			dataS := levelOrder(receivedNodes[message.TreeUuid]["root"])
			err = database.Write(dataS)
			mutex.RUnlock()
			if err != nil {
				log.Error("write into database failed, deleting caches...")
			} else {
				log.Info("write into database success, deleting caches...")
			}
		} else {
			log.Error("current service chain already timed out. deleting caches...")
		}
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok := receivedNodes[message.TreeUuid]; ok {
			delete(receivedNodes, message.TreeUuid)
			delete(tagged, message.TreeUuid)
			delete(untagged, message.TreeUuid)
			delete(timeMap, message.TreeUuid)
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
	select {
	case <-timeChan:
		log.Info("receiving message succeeded, timeOut function stopped, writing into database.")
	case <-time.After(time.Second * config.TIMELAPSES):
		log.Error("receiving message timed out, deleting caches...")
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok := receivedNodes[treeUuid]; ok {
			delete(receivedNodes, treeUuid)
			delete(tagged, treeUuid)
			delete(untagged, treeUuid)
			delete(timeMap, treeUuid)
			log.Info("caches deleted.")
			*err = myErr.NewError(300, "receive message timed out.")
		}
	}
}

func isCompleteTree(treeUUID string) bool {
	if len(tagged) == 0 && len(untagged) == 0 {
		return false
	}
	if _, ok := receivedNodes[treeUUID]["root"]; ok && len(untagged[treeUUID]) == 0 {
		return true
	}
	return false
}

func addTreeNode(treeUUID string, node *TreeNode) {
	if _, ok := receivedNodes[treeUUID][node.Name]; ok {
		// ERROR!!!一个节点重复调用，不符合幂等性。
		return
	}
	receivedNodes[treeUUID][node.Name] = node
	if parent, ok := receivedNodes[treeUUID][node.ParentName]; ok {
		if parent.Children == nil {
			parent.Children = make([]*TreeNode, 0)
		}
		parent.Children = append(parent.Children, node)
		node.Parent = parent
		if judgeNode(parent) {
			tagged[treeUUID][parent.Name] = true
			if _, ok := untagged[treeUUID][parent.Name]; ok {
				delete(untagged[treeUUID], parent.Name)
			}
		}
	} else {
		if node.ParentName != "" {
			untagged[treeUUID][node.ParentName] = true
		}
	}
	for name := range node.ChildrenName {
		if child, ok := receivedNodes[treeUUID][name]; ok {
			if node.Children == nil {
				node.Children = make([]*TreeNode, 0)
			}
			node.Children = append(node.Children, child)
			if child.Parent == nil {
				child.Parent = node
			}
			if judgeNode(child) {
				tagged[treeUUID][child.Name] = true
				if _, ok := untagged[treeUUID][child.Name]; ok {
					delete(untagged[treeUUID], child.Name)
				}
			}
		} else {
			untagged[treeUUID][name] = true
		}
	}
	if judgeNode(node) {
		tagged[treeUUID][node.Name] = true
		if _, ok := untagged[treeUUID][node.Name]; ok {
			delete(untagged[treeUUID], node.Name)
		}
	} else {
		untagged[treeUUID][node.Name] = true
	}
}

func judgeNode(node *TreeNode) bool {
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
func levelOrder(root *TreeNode) []*commonInfo.HttpRequest {
	result := make([]*commonInfo.HttpRequest, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var tmp *TreeNode
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
