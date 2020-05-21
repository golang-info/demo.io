package _0_二叉查找树

type Node struct {
	Key int
	Left *Node
	Right *Node
}

func (n *Node) Search(key int) bool {
	// 这是我们的基本情况， 如果 n == nil,
	// 则`key`在二叉查找树中不存在
	if n == nil {
		return false
	}
	if n.Key < key {
		//向右走
		return n.Right.Search(key)
	}
	if n.Key > key {
		//向左走
		return n.Left.Search(key)
	}
	// 如果 n.Key == key, 就说明找到了
	return true
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			// 我们找到了一个空位，结束
			n.Right = &Node{Key: key}
			return
		}
		// 向右边找
		n.Right.Insert(key)
		return
	}
	if n.Key > key {
		if n.Left == nil {
			// 我们找到了一个空位，结束
			n.Left = &Node{Key: key}
			return
		}
		//向左边找
		n.Left.Insert(key)
	}
	// 如果 n.Key == key，则什么都不做
}

func (n *Node) Delete(key int) *Node {
	// 按 `key` 查找
	if n.Key < key {
		n.Right = n.Right.Delete(key)
		return n
	}
	if n.Key > key {
		n.Left = n.Left.Delete(key)
		return n
	}

	// n.Key == `key`
	if n.Left == nil {
		// 只指向反向的节点
		return n.Right
	}
	if n.Right == nil {
		// 只指向反向的节点
		return n.Left
	}

	// 如果 `n` 有两个子节点， 则需要确定下一个放在位置n的最大值
	// 使得二叉查找树保持正确的性质
	min := n.Right.Min()

	// 如果只是用最小节点来更新 `n` 的 key
	// 因此 n 的直接子节点不再为空
	n.Key = min
	n.Right = n.Right.Delete(min)
	return n
}

func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}
