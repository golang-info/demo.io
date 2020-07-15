package main

/*
	Go实战--实现一个单向链表(The way to go)
	https://blog.csdn.net/wangshubo1989/article/details/71515933
*/

type Node struct {
	Value interface{}
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type SinglyLinkList struct {
	front  *Node
	length int
}

func (s *SinglyLinkList) Init() *SinglyLinkList {
	s.length = 0
	return s
}

func New() *SinglyLinkList {
	return new(SinglyLinkList).Init()
}

func (s *SinglyLinkList) Front() *Node {
	return s.front
}

func (s *SinglyLinkList) Back() *Node {
	currentNode := s.front
	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (s *SinglyLinkList) Append(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		currentNode := s.front

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = n
	}

	s.length++
}

func (s *SinglyLinkList) Prepend(n *Node) {
	if s.front == nil {
		s.front = n
	} else {
		n.next = s.front
		s.front = n
	}

	s.length++
}

func (s *SinglyLinkList) InsertBefore(insert *Node, before *Node) {
	if s.front == before {
		insert.next = before
		s.front = insert
		s.length++
	} else {
		currentNode := s.front
		for currentNode != nil && currentNode.next != nil && currentNode.next != before {
			currentNode = currentNode.next
		}

		if currentNode.next == before {
			insert.next = before
			currentNode.next = insert
			s.length++
		}
	}
}

func (s *SinglyLinkList) InsertAfter(insert *Node, after *Node) {
	currentNode := s.front
	for currentNode != nil && currentNode != after && currentNode.next != nil {
		currentNode = currentNode.next
	}

	if currentNode == after {
		insert.next = after.next
		after.next = insert
		s.length++
	}
}

func (s *SinglyLinkList) Remove(n *Node) {
	if s.front == n {
		s.front = n.next
		s.length--
	} else {
		currentNode := s.front

		// search for node n
		for currentNode != nil && currentNode.next != nil && currentNode.next != n {
			currentNode = currentNode.next
		}

		// see if current's next node is n
		// if it's not n, then node n wasn't found in list s
		if currentNode.next == n {
			currentNode.next = currentNode.next.next
			s.length--
		}
	}
}
