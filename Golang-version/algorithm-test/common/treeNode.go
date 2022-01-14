package common

type TreeNode struct {
	Name         string
	ParentName   string
	Parent       *TreeNode
	ChildrenName []string
	Children     map[*TreeNode]bool
	DbName       string
	SqlStr       string
}
