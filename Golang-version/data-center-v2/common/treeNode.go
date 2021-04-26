package common

import "data-center-v2/proto/commonInfo"

type TreeNode struct {
	Name         string
	ParentName   string
	Parent       *TreeNode
	ChildrenName map[string]bool
	Children     []*TreeNode
	Info         *commonInfo.HttpRequest
	SqlStr       string
}
