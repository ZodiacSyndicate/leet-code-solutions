/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */
import "strconv"

func getPermutation(n int, k int) string {
	arr := []int{}
	fac := 1
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
		fac *= i
	}
	res := []int{}
	pos := k - 1
	for j := n; j > 0; j-- {
		fac /= j
		index := pos / fac
		res = append(res, arr[index])
		arr = append(arr[:index], arr[index+1:]...)
		pos %= fac
	}
	result := ``
	for _, n := range res {
		result += strconv.Itoa(n)
	}
	return result
}

