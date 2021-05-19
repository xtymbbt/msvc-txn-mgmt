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
	if _, ok := receivedNodes[message.TreeUuid]; ok {
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
		receivedNodes[message.TreeUuid] = make(map[string]*common.TreeNode, 0)
		tagged[message.TreeUuid] = make(map[string]bool, 0)
		untagged[message.TreeUuid] = make(map[string]bool, 0)
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
		timeMap[message.TreeUuid] = make(chan bool, 1)
		go timeOut(message.TreeUuid, &err)
	}
	mutex.Unlock()

	if isCompleteTree(message.TreeUuid) {
		log.Info("message all received, writing into database...")
		if _, ok := timeMap[message.TreeUuid]; ok {
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
	case <-time.After(config.TIMELAPSES):
		log.Error("receiving message timed out, deleting caches...")
		mutex.Lock()
		if _, ok := receivedNodes[treeUuid]; ok {
			delete(receivedNodes, treeUuid)
			delete(tagged, treeUuid)
			delete(untagged, treeUuid)
			delete(timeMap, treeUuid)
			log.Info("caches deleted.")
			*err = myErr.NewError(500, "receive message timed out.")
		}
		mutex.Unlock()
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

func addTreeNode(treeUUID string, node *common.TreeNode) error {
	if _, ok := receivedNodes[treeUUID][node.Name]; ok {
		// ERROR!!!一个节点重复调用，不符合幂等性。
		return myErr.NewError(500, "Duplicated Nodes!")
	}
	receivedNodes[treeUUID][node.Name] = node
	if parent, ok := receivedNodes[treeUUID][node.ParentName]; ok {
		if parent.Children == nil {
			parent.Children = make([]*common.TreeNode, 0)
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
				node.Children = make([]*common.TreeNode, 0)
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
