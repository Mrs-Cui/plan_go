package dynamic_program

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/mathx"
	"plan_go/tree"
	"sort"
	"strconv"
	"strings"
)

// 动态规划，以空间换时间


// 最长回文字符串，动态规划解法
func LongestPalindrome(s string) string {
	ss := []rune(s)
	le := len(ss)
	dp := make([][]bool, le)
	for idx := range dp {
		dp[idx] = make([]bool, le)
		dp[idx][idx] = true
	}
	var maxLen int
	var left int
	var right int
	for r := 1; r < le; r++ {
		for l := 0; l < r; l++ {
			if string(ss[l]) == string(ss[r]) {
				if l+1 > r-1 || dp[l+1][r-1] {
					dp[l][r] = true
					if r - l >= maxLen {
						maxLen = r - l
						left, right = l, r
					}
				}
			}
		}
	}
	return s[left:right+1]
}

func Massage(nums []int) int {
	le := len(nums)
	if le == 0 {
		return 0
	}
	if le == 1 {
		return nums[0]
	}
	if le == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	dp := make([]int, le)
	dp[0] = nums[0]
	if nums[0] >= nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i:= 2; i < le; i++ {
		if nums[i]+dp[i-2] >= dp[i-1] {
			dp[i] = nums[i] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[le-1]
}

// 字符串断句

func Respace(dictionary []string, sentence string) int {
	tre := tree.Constructor()
	for i, s := range dictionary {
		tre.Insert(s, i)
	}
	ss := []rune(sentence)
	le := len(ss)
	dp := make([]int, le+1)
	for i := 1; i <= le; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			flag := tre.Search(sentence[j : i])
			if flag {
				dp[i] = mathx.MinInt(dp[i], dp[j])
			}
		}
	}
	return dp[le]
}

// 输入括号对数，给出有效的括号组 []

func GenerateParenthesis(n int) []string {
	var ret [][]string
	var left func(int, int, []string)
	var right func(int, int, []string)
	left = func(l, r int, dp []string) {
		if l <= 0 || r <= 0 {
			return
		}
		if l == 0 && r != 0 {
			right(l, r, dp)
			return
		}
		dp[n*2-l-r] = "("
		dp2 := make([]string, n*2)
		copy(dp2, dp)
		left(l-1, r, dp2)
		dp1 := make([]string, n*2)
		copy(dp1, dp)
		right(l-1, r, dp1)
	}
	right = func(l int, r int, dp []string) {
		if l < 0 || r < 0 {
			return
		}
		if l == 0 && r == 0 {
			ret = append(ret, dp)
		}
		if l > r {
			return
		}
		if l == r {
			left(l, r, dp)
			return
		}
		idx := n*2 - l - r
		dp[idx] = ")"
		dp0 := make([]string, n*2)
		copy(dp0, dp)
		left(l, r-1, dp0)
		dp1 := make([]string, n*2)
		copy(dp1, dp)
		right(l, r-1, dp1)

	}
	dp := make([]string, n*2)
	left(n, n, dp)
	rets := make(map[string]bool)
	for _, r := range ret {
		rets[strings.Join(r, "")] = true
	}
	var retss []string
	for key, _ := range rets {
		retss = append(retss, key)
	}
	return retss
}

// 不同路径 II

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	var low, row int
	endR := len(obstacleGrid)
	endL := len(obstacleGrid[0])
	var num int
	var find func(int, int)
	find = func(low int, row int) {
		if low >= endL || row >= endR {
			return
		}
		if obstacleGrid[row][low] == 1 {
			return
		}
		if low == endL-1 && row == endR-1 {
			num += 1
			return
		}
		find(low+1, row)
		find(low, row+1)
	}
	find(low, row)
	return num
}

// 优化， 使用动态规划

func UniquePathsWithObstacles1(obstacleGrid [][]int) int {
	lr := len(obstacleGrid)
	ll := len(obstacleGrid[0])
	dp := make([][]int, lr+1)
	for i := 0; i <= lr; i++ {
		dp[i] = make([]int, ll+1)
	}
	if obstacleGrid[0][0] == 0 {
		dp[1][1] = 1
	}
	for i:=2; i <= ll; i++ {
		if obstacleGrid[0][i-1] == 1 {
			dp[1][i] = 0
			continue
		}
		dp[1][i] = dp[0][i] + dp[1][i-1]
	}
	for i:=2; i <= lr; i++ {
		for j:=1; j <= ll; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[lr][ll]
}

// 编辑距离




// 不同的二叉搜索树 II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	var data []int
	for i := 1; i <= n; i++ {
		data = append(data, i)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {

		}
	}

	return nil
}

// 不同的二叉搜索树
func NumTrees(n int) int {
	var ret []int
	for i:=1; i<=n; i++ {
		ret = append(ret, i)
	}
	cache := make(map[string]int)
	sumFunc := func(rets []int) string {
		str, _ := json.Marshal(rets)
		return string(str)
	}
	var fun func(ret []int) int
	fun = func(ret []int) int {
		if len(ret) == 0 {
			return 1
		}
		if val, ok := cache[sumFunc(ret)]; ok {
			return val
		}
		var sum int
		for idx, _ := range ret {
			leftL := ret[:idx]
			rightL := ret[idx+1:]
			left := fun(leftL)
			right := fun(rightL)
			sum += left * right
			cache[sumFunc(leftL)] = left
			cache[sumFunc(rightL)] = right
		}
		return sum
	}
	re := fun(ret)
	return re
}

// 交错字符串
func IsInterleave(s1 string, s2 string, s3 string) bool {
	ss1, ss2, ss3 := []rune(s1), []rune(s2), []rune(s3)
	ls1, ls2, ls3 := len(ss1), len(ss2), len(ss3)
	if ls1 + ls2 != ls3 {
		return false
	}
	dp := make([][]bool, ls1+1)
	for idx, _ := range dp {
		dp[idx] = make([]bool, ls2+1)
	}
	dp[0][0] = true
	for i := 0; i <= ls1; i++ {
		for j := 0; j <= ls2; j++ {
			is3 := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && string(ss1[i-1]) == string(ss3[is3]))
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && string(ss2[j-1]) == string(ss3[is3]))
			}
		}
	}
	return dp[ls1][ls2]
}

// 分割回文串

func Partition(s string) [][]string {
	ss := []rune(s)
	var st []string
	for _, v := range ss {
		st = append(st, string(v))
	}
	le := len(ss)
	dp := make([][]bool, le)
	for idx := range dp {
		dp[idx] = make([]bool, le)
		dp[idx][idx] = true
	}
	for r := 1; r < le; r++ {
		for l := 0; l < r; l++ {
			if string(ss[l]) == string(ss[r]) {
				if l+1 > r-1 || dp[l+1][r-1] {
					dp[l][r] = true
				}
			}
		}
	}
	var rets [][]string
	var fun func(int, []string)
	fun = func(idx int, ret []string) {
		if idx >= le {
			rets = append(rets, ret)
		}
		for i:=idx; i < le; i++ {
			if dp[idx][i] {
				re :=  make([]string, len(ret))
				copy(re, ret)
				re = append(re, strings.Join(st[idx:i+1], ""))
				fun(i+1, re)
			}
		}
	}
	fun(0, []string{})
	return rets
}

// 为运算表达式设计优先级

func DiffWaysToCompute(expression string) []int {
	ss := []rune(expression)
	var flag bool
	for _, s := range ss {
		c := string(s)
		if c == "-" || c == "+" || c == "*" || c == "/" {
			flag = true
		}
	}
	if !flag {
		v, _ := strconv.ParseInt(expression, 10, 64)
		return []int{int(v)}
	}
	var rets []int
	sum := func(c string, left, right int) int {
		if c == "-" {
			return left - right
		} else if c == "+" {
			return left + right
		} else if c == "*" {
			return left * right
		} else if c == "/" {
			return left / right
		}
		return 0
	}
	var fun func(ss []rune) []int
	fun = func(ss []rune) []int {
		if len(ss) == 0 {
			return nil
		}
		var flag bool
		var expr string
		for _, s := range ss {
			c := string(s)
			if c == "-" || c == "+" || c == "*" || c == "/" {
				flag = true
			}
			expr += c
		}
		if !flag {
			v, _ := strconv.ParseInt(expr, 10, 64)
			return []int{int(v)}
		}
		var ret []int
		for idx, s := range ss {
			c := string(s)
			if c == "-" || c == "+" || c == "*" || c == "/" {
				lefts := fun(ss[:idx])
				rights := fun(ss[idx+1:])
				for _, i := range lefts {
					for _, j := range rights {
						ret = append(ret, sum(c, i, j))
					}
				}
			}
		}
		return ret
	}
	rets = fun(ss)
	return rets
}

// 打家劫舍 III

func Rob(root *TreeNode) int {
	var find func(root *TreeNode) []int
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	find = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		left := find(root.Left)
		right := find(root.Right)
		ret := []int{left[0]+right[0], left[1]+right[1]}
		sel := ret[1] + root.Val
		noSel := max(left[0], left[1]) + max(right[0], right[1])
		return []int{sel, noSel}
	}
	ret := find(root)
	if ret[0] > ret[1] {
		return ret[0]
	}
	return ret[1]
}

// 整数拆分

func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	for i:=1; i<=n; i++ {
		dp[i] = 1
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for i:=2; i<=n; i++ {
		for j:=1; j < i;j ++ {
			ret := max(j * (i-j), j * dp[i-j])
			dp[i] = max(dp[i], ret)
		}
	}
	return dp[n]
}

// 猜数字大小 II
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = j + dp[i][j-1]
			for k := i; k < j; k++ {
				c := k + max(dp[i][k-1], dp[k+1][j])
				if c < dp[i][j] {
					dp[i][j] = c
				}
			}
		}
	}
	return dp[1][n]
}

// 组合综合 II
func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target +1)
	for i:=1; i<= target; i++ {
		for _, num := range nums {
			if i == num {
				dp[i] += 1
				continue
			}
			if i - num <= 0 {
				continue
			}
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

// 旋转函数
func MaxRotateFunction(nums []int) int {
	le := len(nums)
	dp := make([]int, le)
	front := make([]int, le)
	last := make([]int, le)
	front[0], last[0] = nums[0], nums[le-1]
	for i := 1; i < le; i++ {
		front[i] = nums[i] + front[i-1]
	}
	for i := le - 2; i >= 0; i-- {
		last[le-i-1] = nums[i] + last[le-i-2]
	}
	var t int
	for i:=0; i < le; i++ {
		t += i * nums[i]
	}
	dp[0] = t
	idx := le -1
	max := dp[0]
	for i := 1; i < le; i++ {
		mid := (le-1) * nums[idx]
		dp[i] = dp[i-1] - mid + front[idx-1]
		if i - 2 >= 0 {
			dp[i] += last[i-2]
		}
		idx -= 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 整数替换 (待优化)
func IntegerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}
	var ret int
	mid := 1
	last := 1
	var tmp int
	for i := 3; i <= n; i++ {
		if i % 2 != 0 {
			if tmp != 0 {
				mid = tmp
			}
			v1 := last + 1
			v2 := mid + 2
			ret = min(v1, v2)
			tmp = ret
		} else {
			ret = mid + 1
			last = ret
		}
	}
	return ret
}

// 无重叠区间

func EraseOverlapIntervals(intervals [][]int) int {
	le := len(intervals)
	if le == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	mid := intervals[0][1]
	var ret int
	for i:=1; i < le; i++ {
		if intervals[i][0] < mid {
			ret += 1
			continue
		}
		mid = intervals[i][1]
	}
	return ret
}

// 我能赢吗
func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	for i:=1; i <= maxChoosableInteger; i ++ {
		mid := desiredTotal - i
		if mid <= maxChoosableInteger {
			return false
		}
		t := mid / maxChoosableInteger
		if t % 2 != 0 {
			return true
		}
	}
	return false
}


func FindSubstringInWraproundString(p string) int {
	le := len(p)
	if le == 0 || le == 1 {
		return le
	}
	dp := make([]int32, 30)
	zero := "a"
	idx := 1
	dp[rune(p[0]) - rune(zero[0])] = 1
	var fun func(idx int, pre int32)
	fun = func(idx int, pre int32) {
		if idx >= le {
			return
		}
		var mid int32 = 1
		if string(p[idx]) == "a" && string(p[idx-1]) == "z" {
			mid += pre
		}
		i := rune(p[idx]); j := rune(p[idx-1])
		if i - j == 1 {
			mid += pre
		}
		k := rune(p[idx]) - rune(zero[0])
		if dp[k] < mid {
			dp[k] = mid
		}
		fun(idx+1, mid)
	}
	fun(idx, 1)
	var num int32
	for _, val := range dp {
		num += val
	}
	return int(num)
}

// Makesquare 火柴拼正方形
func Makesquare(matchsticks []int) bool {
	var num int
	sides := make(map[int]int)
	var max int
	for _, val := range matchsticks {
		num += val
		sides[val] += 1
		if val >= max {
			max = val
		}
	}
	if num % 4 != 0 {
		return false
	}
	side := num / 4
	if max > side {
		return false
	}
	le := len(matchsticks)
	if le < 4 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] >= matchsticks[j]
	})
	dp := make([]int, 4)
	dp[0] = matchsticks[0]
	sides[dp[0]] -= 1
	var mid []int
	for key, _ := range sides {
		mid = append(mid, key)
	}
	sort.Slice(mid, func(i, j int) bool {
		return mid[i] >= mid[j]
	})
	var fun func(ls int, use map[int]int) bool
	fun = func(ls int, use map[int]int) bool {
		var flag bool
		for _, key := range mid {
			if key > ls || sides[key] - use[key] <= 0 {
				continue
			}
			if key == ls {
				use[key] += 1
				return true
			}
			use[key] += 1
			flag = fun(ls-key, use)
			if flag {
				if sides[key] < use[key] {
					return false
				}
				break
			} else {
				use[key] -= 1
			}
		}
		if !flag {
			return false
		}
		return true
	}
	for i, v := range dp {
		if v == side {
			continue
		}
		use := make(map[int]int)
		flag := fun(side-v, use)
		if flag {
			for key, val := range use {
				sides[key] -= val
			}
		} else {
			return false
		}
		if i >= 3 {
			continue
		}
		var t int
		for key, val := range sides {
			if val != 0 && key >= t {
				t = key
			}
		}
		sides[t] -= 1
		dp[i+1] = t
	}
	return true
}

// FindMaxForm 0 和 1
func FindMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i:=0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := func(i, j int) int {
		if i >= j {
			return i
		}
		return j
	}
	for _, val := range strs {
		zero := strings.Count(val, "0")
		one := strings.Count(val, "1")
		if zero > m || one > n {
			continue
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one] + 1)
			}
		}
	}
	return dp[m][n]
}

// 预测最大赢家

func PredictTheWinner(nums []int) bool {
	le := len(nums)
	if le == 1 || le == 2 {
		return true
	}
	max := func(l, r int) int {
		if nums[l] >= nums[r] {
			return l
		}
		return r
	}
	var fun func(l, r, idx int, dp [2]int) bool
	fun = func(l, r, idx int, dp [2]int) bool {
		if l > r {
			if dp[0] != 0 && dp[1] != 0 && dp[0] >= dp[1] {
				return true
			}
			return false
		}
		mid := idx % 2
		dp1 := dp
		dp2 := dp
		si := max(l, r)
		var flag bool
		if si == l {
			dp1[mid] += nums[l]
			flag = fun(l+1, r, idx+1, dp1)
		} else {
			dp2[mid] += nums[r]
			flag = fun(l, r-1, idx+1, dp2)
		}
		if flag {
			return true
		}
		dp1 = dp
		dp2 = dp
		if si == l {
			dp1[mid] += nums[r]
			flag = fun(l, r-1, idx+1, dp2)
		} else {
			dp2[mid] += nums[l]
			flag = fun(l+1, r, idx+1, dp1)
		}
		return flag
	}
	var dp [2]int
	return fun(0, le-1, 0, dp)
}

// 动态规划优化

func PredictTheWinner1(nums []int) bool {
	le := len(nums)
	dp := make([][]int, le)
	for idx:=0; idx < le; idx++ {
		dp[idx] = make([]int, le)
	}
	for i :=0; i < le; i++ {
		dp[i][i] = nums[i]
	}
	max := func(i, j int) int {
		if i >= j {
			return i
		}
		return j
	}
	for i := le - 2; i >= 0; i-- {
		for j := i + 1; j < le; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][le-1] >= 0
}

func FindTargetSumWays(nums []int, target int) int {
	le := len(nums)
	var count int
	var fun func(idx, sum int)
	fun = func(idx, sum int) {
		if idx >= le {
			if sum == target {
				count += 1
			}
			return
		}
		sum1 := sum + nums[idx] * -1
		fun(idx+1, sum1)
		sum2 := sum + nums[idx]
		fun(idx+1, sum2)
	}
	fun(0, 0)
	return count
}

// 最长回文子序列

func LongestPalindromeSubseq(s string) int {
	le := len(s)
	dp := make([][]int, le)
	for i := 0; i < le; i++ {
		dp[i] = make([]int, le)
	}
	max := func(i, j int) int {
		if i >= j {
			return i
		}
		return j
	}
	for i := le-1; i >=0; i-- {
		dp[i][i] = 1
		for j := i+1; j < le; j++ {
			if string(s[i]) == string(s[j]) {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1] +2)
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][le-1]
}

// 兑换零钱II

func Change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// 优美的排列

func CountArrangement(n int) int {

	return 0
}

/*
	有一堆石头，用整数数组stones表示。其中stones[i]表示第i块石头的重量。
	每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x和y，
	且x<=y。那么粉碎的可能结果如下:
	如果x==y，那么两块石头都会被完全粉碎; 如果x!=y，那么重量为x的石头将会完全粉碎，
	而重量为y的石头新重量为y-x。 最后，最多只会剩下一块石头。返回此石头最小的可能重量。
	如果没有石头剩下，就返回 0。
*/

func minWeight(stones []int) {
	var sum int
	for _, stone := range stones {
		sum += stone
	}
	weight := sum / 2
	if sum % 2 != 0 {
		weight += 1
	}
	dp := make([][]int, len(stones))
	for i:=0; i < len(stones); i ++ {
		dp[i] = make([]int, weight)
	}
	for i:=0; i < len(stones); i ++ {
		if weight >= stones[i] {
			dp[0][stones[i]] = 1
		}
	}
	for i:=1; i < len(stones); i++ {
		for j := stones[i]; j <= weight; j++ {

		}
	}
}


func RegularPattern(s string, p string) {
	sl := len(s)
	pl := len(p)
	dp := make([][]bool, sl+1)
	for i := 0; i < sl; i++ {
		dp[i] = make([]bool, pl+1)
	}
	dp[0][0] = true
	for i := 0; i < pl; i++ {
		if string(p[i]) == "*" && dp[0][i] {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < sl; i++ {
		for j := 0; j < pl; j++ {
			if s[i] == p[j] || string(p[j]) == "." {
				dp[i+1][j+1] = dp[i][j]
			} else if string(p[j]) == "*" && j > 0  {
				if string(s[i]) == string(p[j-1]) || string(p[j-1]) == "." {

				} else if string(s[i]) != string(p[j-1]) {

				}
			}
		}
	}
}

