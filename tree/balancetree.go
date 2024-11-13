package tree

type Tree struct {
	Left    *Tree
	Right   *Tree
	Data    int
	IsEmpty bool
}

func (t *Tree) Height() int {
	if t.IsEmpty {
		return 0
	}
	l := t.Left.Height()
	r := t.Right.Height()
	return max(l, r) + 1
}

func Insert(data int, root *Tree) *Tree {
	if root.IsEmpty {
		root.Data = data
		root.IsEmpty = false
		root.Right = &Tree{
			IsEmpty: true,
		}
		root.Left = &Tree{
			IsEmpty: true,
		}
		return root
	}
	if data <= root.Data {
		root.Left = Insert(data, root.Left)
	} else if data > root.Data {
		root.Right = Insert(data, root.Right)
	}
	lh := root.Left.Height()
	rh := root.Right.Height()
	if lh - rh >= 2 || rh - lh >= 2 {
		if data <= root.Data {
			if root.Left.IsEmpty {
				root = LR(root)
			} else if !root.Left.IsEmpty {
				root = LL(root)
			}
		} else {
			if root.Right.IsEmpty {
				root = RL(root)
			} else if !root.Right.IsEmpty {
				root = RR(root)
			}
		}
	}
	return root
}

func LL(root *Tree) *Tree {
	mid := root.Left
	root.Left = mid.Right
	mid.Right = root
	return mid
}

func LR(root *Tree) *Tree {
	root.Right = RR(root.Right)
	return LL(root)
}

func RR(root *Tree) *Tree {
	mid := root.Right
	root.Right = mid.Left
	mid.Left = root
	return mid
}

func RL(root *Tree) *Tree {
	root.Left = LL(root.Left)
	return RR(root)
}

func max(l, r int) int {
	if l >= r {
		return l
	}
	return r
}
