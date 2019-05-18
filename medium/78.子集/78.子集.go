import "math"

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
func subsets(nums []int) [][]int {
	var ret [][]int
	ret = append(ret, []int{})
	size := len(nums)
	subsize := int(math.Pow(2, float64(size)))
	hash := 1
	for hash < subsize {
		var temp []int
		for k := 0; k < size; k++ {
			a := int(math.Pow(2, float64(k)))
			if (a & hash) != 0 {
				temp = append(temp, nums[k])
			}
		}
		ret = append(ret, temp)
		hash++
	}
	return ret
}

