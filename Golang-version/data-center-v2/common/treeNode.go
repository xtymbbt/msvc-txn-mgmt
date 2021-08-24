package common

type TreeNode struct {
	Name         string
	ParentName   string
	Parent       *TreeNode
	ChildrenName []string
	Children     []*TreeNode
	DbName       string
	SqlStr       string
}
