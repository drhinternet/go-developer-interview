package treenode_test

import (
	"testing"
	"treenode"
)

func TestTreeNodesHaveKeyAndValue(t *testing.T) {
	node := treenode.NewTreeNode("my_key", 1)

	if node.Key != "my_key" {
		t.Errorf("treenode.Key expected=%v got=%v", "my_key", node.Key)
	}

	if node.Value != 1 {
		t.Errorf("treenode.Key expected=%v got=%v", 1, node.Key)
	}
}

func TestTreeNodeEquality(t *testing.T) {
	node := treenode.NewTreeNode("key1", 1)
	nodeWithSameKeySameValue := treenode.NewTreeNode("key1", 1)
	nodeWithSameKeyDifferentValue := treenode.NewTreeNode("key1", 2)
	nodeWithDifferentKeySameValue := treenode.NewTreeNode("key2", 1)
	nodeWithDifferentKeyDifferentValue := treenode.NewTreeNode("key2", 2)

	if !node.Equal(node) {
		t.Errorf("treenode is not equal to itself: %+v", node)
	}

	if !node.Equal(nodeWithSameKeySameValue) {
		t.Errorf(
			"treenode is not equal to treenode with same key and same value: %+v != %+v",
			node,
			nodeWithSameKeySameValue,
		)
	}

	if !node.Equal(nodeWithSameKeyDifferentValue) {
		t.Errorf(
			"treenode is not equal to treenode with same key and different value: %+v != %+v",
			node,
			nodeWithSameKeyDifferentValue,
		)
	}

	if node.Equal(nodeWithDifferentKeySameValue) {
		t.Errorf(
			"treenode should not be equal to treenode with different key and same value: %+v == %+v",
			node,
			nodeWithDifferentKeySameValue,
		)
	}

	if node.Equal(nodeWithDifferentKeySameValue) {
		t.Errorf(
			"treenode should not be equal to treenode with different key and different value: %+v == %+v",
			node,
			nodeWithDifferentKeyDifferentValue,
		)
	}
}

func TestTreeNodesCanHaveChildren(t *testing.T) {
	root := treenode.NewTreeNode("key1", 1)
	child1 := root.AddChild("key2", 10)
	child2 := root.AddChild("key3", 20)
	child1.AddChild("key4", 25)

	if !treenode.NewTreeNode("key2", 0).Equal(root.GetChild("key2")) {
		t.Errorf(
			"could not find child key2 in treenode: %+v",
			root,
		)
	}

	if !treenode.NewTreeNode("key3", 0).Equal(root.GetChild("key3")) {
		t.Errorf(
			"could not find child key3 in treenode: %+v",
			root,
		)
	}

	if child := root.GetChild("key4"); child != nil {
		t.Errorf(
			"should not find key4 in root, but did: %+v",
			child,
		)
	}

	if !treenode.NewTreeNode("key4", 0).Equal(child1.GetChild("key4")) {
		t.Errorf(
			"could not find child key4 in treenode: %+v",
			child1,
		)
	}

	if child := child2.GetChild("key4"); child != nil {
		t.Errorf(
			"should not find key4 in child2, but did: %+v",
			child,
		)
	}
}
