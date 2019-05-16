package stack

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	head *Node
}

func NewStack() *Stack {
	s := &Stack{nil}
	return s
}

func (s *Stack) Push(data interface{}) {
	top := &Node{
		data: data,
		next: s.head,
	}
	s.head = top
}

func (s *Stack) Pop() (interface{}, bool) {
	top := s.head
	if top == nil {
		return nil, false
	}

	s.head = top.next
	return top.data, true
}

func (s *Stack) Top() (interface{}, bool) {
	top := s.head
	if top == nil {
		return nil, false
	}

	return top.data, true
}
