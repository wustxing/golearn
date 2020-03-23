package tree

import "fmt"

type treeNode struct {
	name  string
	left  *treeNode
	right *treeNode
}

func (p *treeNode) AddLeftNode(value int) {
	children := treeNode{
		name: fmt.Sprintf("子节点%d", value),
	}
	p.left = &children
}

func (p *treeNode) AddRightNode(value int) {
	children := treeNode{
		name: fmt.Sprintf("子节点%d", value),
	}
	p.right = &children
}

//前序遍历
func (p *treeNode) aheadTraversal() {
	if p == nil {
		return
	}

	fmt.Println("node name", p.name)
	p.left.aheadTraversal()
	p.right.aheadTraversal()
}

//中序遍历
func (p *treeNode) middleTraversal() {
	if p == nil {
		return
	}

	p.left.middleTraversal()
	fmt.Println("node name", p.name)
	p.right.middleTraversal()
}

//后序遍历
func (p *treeNode) lateTraversal() {
	if p == nil {
		return
	}

	p.left.lateTraversal()
	p.right.lateTraversal()
	fmt.Println("node name", p.name)
}
