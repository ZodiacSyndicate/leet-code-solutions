/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
func combine(n int, k int) [][]int {
	var res [][]int
	if k > n || k == 0 {
		return res
	}
	if k == 1 {
		for i := 0; i < n; i++ {
			res = append(res, []int{i + 1})
		}
		return res
	}
	if k == n {
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, i+1)
		}
		res = append(res, arr)
		return res
	}
	res = combine(n-1, k)
	for _, item := range combine(n-1, k-1) {
		item = append(item, n)
		res = append(res, item)
	}
	return res
}

