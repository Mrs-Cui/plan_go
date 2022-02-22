package tree

import (
	"fmt"
	"math/rand"
)

var (
	Black string = "black"
	Red string = "red"
)

type Node struct {
	Color string
	LeftNode *Node
	RightNode *Node
	Val float64
}

func (m *Node) toString() string {
	return fmt.Sprintf("Node:[Color:%s Val:%v LeftNode:%v RightNode:%v]", m.Color, m.Val, m.LeftNode, m.RightNode)
}

func (m *Node) IsRed() bool {
	if m == nil {
		return false
	}
	if m.Color == Red {
		return true
	}
	return false
}

func (m *Node) leftHanded() *Node {
	node := m.RightNode
	m.RightNode = node.LeftNode
	node.LeftNode = m
	node.Color = m.Color
	m.Color = Red
	return node
}

func (m *Node) rightHanded() *Node {
	node := m.LeftNode
	m.LeftNode = node.RightNode
	node.RightNode = m
	node.Color = m.Color
	m.Color = Red
	return node
}

func (m *Node) exchangeColor() {
	m.Color = Red
	m.LeftNode.Color = Black
	m.RightNode.Color = Black
}

func (m *Node) CreateNode(val float64) *Node {
	if m == nil {
		return &Node{
			Val:       val,
			Color:     Red,
			RightNode: nil,
			LeftNode:  nil,
		}
	}
	if val <= m.Val {
		m.LeftNode = m.LeftNode.CreateNode(val)
	}
	if val > m.Val {
		m.RightNode = m.RightNode.CreateNode(val)
	}

	if !m.LeftNode.IsRed() && m.RightNode.IsRed() {
		m = m.leftHanded()
	}
	if m.LeftNode.IsRed() && !m.RightNode.IsRed() && m.LeftNode.LeftNode.IsRed() {
		m = m.rightHanded()
	}
	if m.LeftNode.IsRed() && m.RightNode.IsRed() {
		m.exchangeColor()
	}
	return m
}

type RedBlackTree struct {
	Root *Node
}

func (m *RedBlackTree) CreateNode(val float64) {
	node := m.Root.CreateNode(val)
	m.Root = node
	m.Root.Color = Black
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		Root: nil,
	}
}


func CheckRedBlackTree() {
	count := 10
	redBlackTree := NewRedBlackTree()
	nums := make([]int, 0)
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Intn(count))
	}
	for _, val := range nums {
		redBlackTree.CreateNode(float64(val))
	}
	fmt.Printf("tree:%v", redBlackTree)
}