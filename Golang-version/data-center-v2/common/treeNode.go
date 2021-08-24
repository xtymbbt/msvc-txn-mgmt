package common

type TreeNode struct {
	Name         string
	ParentName   string
	Parent       *TreeNode
	ChildrenName map[string]bool
	Children     []*TreeNode
	DbName       string
	SqlStr       string
}
