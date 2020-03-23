package tree

import (
	"fmt"
	"testing"
)

func Test_Tree(t *testing.T) {
	var node = treeNode{
		name: "root",
		left: &treeNode{
			name: "it",
			left: nil,
		},
		right: &treeNode{name: "it2", right: nil},
	}
	node.AddLeftNode(12)
	node.left.AddLeftNode(23)
	node.left.AddRightNode(24)
	node.left.right.AddLeftNode(31)
	node.left.right.AddRightNode(32)
	node.AddRightNode(15)
	fmt.Println(node)

	node.lateTraversal()

}
