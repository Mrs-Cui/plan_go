package tree

import "fmt"



func KMP(source string, pattern string) {
	next := BuildNext(pattern)
	ls := len(source)
	lp := len(pattern)
	var i, j int
	for i < ls && j < lp {
		if j == -1 || source[i] == pattern[j] {
			i += 1
			j += 1
		} else {
			j = next[j]
		}
	}
	if i < ls && j == lp {

	}
}


func BuildNext(pattern string) []int {
	l := len(pattern)
	next := make([]int, l)
	next[0] = -1
	i, j := -1, 0
	for j < l - 1 {  // ababcabd
		if i == -1 || pattern[i] == pattern[j] {
			i += 1
			j += 1
			next[j] = i
		} else {
			i = next[i]
		}
	}
	return next

}


func GetBadChar(pattern string) map[string]int {
	pos := map[string]int{}
	for idx, ch := range pattern {
		pos[string(ch)] = idx
	}
	return pos
}

func MovePos(pattern string) map[string]int {
	lp := len(pattern)
	pos := map[string]int{}
	for i := 0; i < lp; i++ {
		suffix := pattern[lp-i-1:]
		for j := 0; j < lp - 1 - i; j++ {
			ch := pattern[j:j+i+1]
			if ch == suffix {
				pos[suffix] = lp - i - 1 - j
			}
		}
	}
	fmt.Println(pos)
	for i := 0; i < lp; i++ {
		suffix := pattern[lp-i-1:]
		if _, ok := pos[suffix]; ok {
			continue
		}
		pos[suffix] = pos[suffix[1:]]
	}
	return pos
}