package tree

// 前缀树


type Trie struct {
	IsWord bool
	Next map[string]*Trie
	Idx int
}


func Constructor() Trie {
	return Trie{
		Next: map[string]*Trie{},
	}
}


func (this *Trie) Insert(word string, idx int)  {
	m := this
	for _, s := range []rune(word) {
		t := string(s)
		val, ok := m.Next[t]
		if !ok {
			val = &Trie{
				Next: map[string]*Trie{},
			}
			m.Next[t] = val
		}
		m = val
	}
	m.IsWord = true
	m.Idx = idx
}


func (this *Trie) Search(word string) bool {
	m := this
	for _, s := range []rune(word) {
		t := string(s)
		val, ok := m.Next[t]
		if !ok {
			return false
		}
		m = val
	}
	if m.IsWord == true {
		return true
	}
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
	m := this
	for _, s := range []rune(prefix) {
		t := string(s)
		val, ok := m.Next[t]
		if !ok {
			return false
		}
		m = val
	}
	return true
}

// 字典排序树

func LexicalOrder(n int) []int {
	var ret []int
	var fun func(int2 int)
	fun = func(num int) {
		if num > n {
			return
		}
		ret = append(ret, num)
		for k := num*10; k < num*10 + 10; k++ {
			if k > n {
				break
			}
			fun(k)
		}
	}
	for i:=1; i <=9; i++ {
		fun(i)
	}
	return ret
}

// 多次匹配
// 题意: 匹配多个子串是否在目标串中出现过。

// 解答： 字典树 + AC 自动机 + KMP

type TriNode struct {
	Next map[string]*TriNode
	Fail *TriNode
	Idx []int
	IsLast bool
	Parent *TriNode
	val string
}

func (m *TriNode) CreateTri(smalls []string) {
	for idx, small := range smalls {
		mid := m
		for _, s := range []rune(small) {
			t := string(s)
			if mid.Next == nil {
				mid.Next = make(map[string]*TriNode)
			}
			val, ok := mid.Next[t]
			if !ok {
				mid.Next[t] = &TriNode{
					Parent: mid,
					Next: make(map[string]*TriNode),
				}
				val = mid.Next[t]
			}
			mid = val
		}
		mid.IsLast = true
		mid.Idx = append(mid.Idx, idx)
	}
}

func (m *TriNode) AcPatternBuild() {
	var nodes []*TriNode
	for key, val := range m.Next {
		val.val = key
		nodes = append(nodes, val)
	}
	for len(nodes) != 0 {
		var rets []*TriNode
		for _, node := range nodes {
			fail := node.Parent.Fail
			for {
				if fail == nil {
					node.Fail = m
					break
				}
				mid, ok := fail.Next[node.val]
				if ok {
					node.Fail = mid
					break
				}
				fail = fail.Fail
			}
			for key, val := range node.Next {
				val.val = key
				rets = append(rets, val)
			}
		}
		nodes = rets
	}
}

// Search ac 自动机
func (m *TriNode) Search(str string, rets [][]int, smalls []string) [][]int {
	strs := []rune(str)
	idx := 0
	for idx < len(strs) {
		mid := m
		for {
			if idx >= len(strs) {
				break
			}
			t := string(strs[idx])
			val, ok := mid.Next[t]
			if ok {
				if val.IsLast {
					for _, i := range val.Idx {
						rets[i] = append(rets[i], idx-len(smalls[i])+1)
					}
				}
				tmp := val.Fail
				for {
					if tmp == m || tmp == nil {
						break
					}
					if tmp.IsLast {
						for _, i := range tmp.Idx {
							rets[i] = append(rets[i], idx-len(smalls[i])+1)
						}
					}
					tmp = tmp.Fail
				}
				mid = val
				idx += 1
				continue
			}
			mid = mid.Fail
			if mid == m || mid == nil{
				idx += 1
				break
			}
		}
	}
	return rets
}


func MultiSearch(big string, smalls []string) [][]int {


	return nil
}
