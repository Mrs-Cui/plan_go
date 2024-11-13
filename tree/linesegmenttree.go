package tree

type NumArray struct {
	LeftChild *NumArray
	RightChild *NumArray
	Range []int
	Sum int
}


func Constructor1(nums []int) NumArray {
	start, end := 0, len(nums) - 1
	var fun = func(root *NumArray, start, end int) int{
		return 0
	}
	fun = func(root *NumArray, start, end int) int {
		if start > end {
			return 0
		}
		if len(root.Range) == 0 {
			root.Range = []int{start, end}
			root.LeftChild, root.RightChild = new(NumArray), new(NumArray)
		}
		if start == end {
			root.Sum = nums[start]
			return root.Sum
		}
		mid := (start + end) / 2
		left := fun(root.LeftChild, start, mid)
		right := fun(root.RightChild, mid+1, end)
		root.Sum = left + right
		return root.Sum
	}
	root := new(NumArray)
	fun(root, start, end)
	return *root
}


func (this *NumArray) Update(index int, val int)  {
	mid := (this.Range[0] + this.Range[1]) / 2
	if index == this.Range[0] && index == this.Range[1] {
		this.Sum = val
		return
	}
	if index <= mid {
		this.LeftChild.Update(index, val)
	}
	if index > mid {
		this.RightChild.Update(index, val)
	}
	this.Sum = this.LeftChild.Sum + this.RightChild.Sum
}


func (this *NumArray) SumRange(left int, right int) int {
	if this.Range[0] == left && this.Range[1] == right {
		return this.Sum
	}
	mid := (this.Range[0] + this.Range[1]) / 2
	if right <= mid {
		return this.LeftChild.SumRange(left, right)
	}
	if left > mid {
		return this.RightChild.SumRange(left, right)
	}
	if left <= mid || right > mid {
		return this.LeftChild.SumRange(left, mid) + this.RightChild.SumRange(mid+1, right)
	}
	return 0
}
