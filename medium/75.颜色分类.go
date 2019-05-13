/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */
func sortColors(nums []int) {
	l := len(nums)
	i, j, k := 0, 0, l-1
	for j <= k {
		if nums[j] == 1 {
			j++
		} else if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}

