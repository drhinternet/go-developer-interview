package treenode

type TreeNode struct {
	Key   string
	Value int
}

func NewTreeNode(key string, value int) *TreeNode {
	return &TreeNode{}
}

func (node *TreeNode) Equal(other *TreeNode) bool {
	// TreeNodes are considered equal if their keys are equal.
	return false
}

func (node *TreeNode) AddChild(key string, value int) *TreeNode {
	// Add the child with the given key to this +node+. Each TreeNode may
	// have up to one direct child with a given key.
	return &TreeNode{}
}

func (node *TreeNode) GetChild(key string) *TreeNode {
	// Find and return a direct child of +node+. Do not look down the tree.
	return nil
}
