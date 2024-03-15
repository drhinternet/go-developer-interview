package treenode

type TreeNode struct {
	Key   string
	Value int
}

func NewTreeNode(key string, value int) *TreeNode {
	return &TreeNode{}
}

func (node *TreeNode) Equal(other *TreeNode) bool {
	return false
}

func (node *TreeNode) AddChild(key string, value int) *TreeNode {
	return &TreeNode{}
}

func (node *TreeNode) GetChild(key string) *TreeNode {
	return nil
}
