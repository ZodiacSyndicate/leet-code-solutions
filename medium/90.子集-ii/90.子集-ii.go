/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (53.16%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    6.9K
 * Total Submissions: 12.8K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 * 
 * 说明：解集不能包含重复的子集。
 * 
 * 示例:
 * 
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 * 
 */
 func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
  sort.Ints(nums)

	for i := 1; i <= len(nums); i++ {
		tmp := []int{}
		f12(i,nums,0,0,tmp,&result)
	}
	result = append(result,[]int{})
	return result
}

func f12(k int,choice []int,height int,start int,tmp []int,result *[][]int) {
	if height == k {
		*result = append(*result,tmp)
		return
	}
	if height > k {
		return
	}
	mp := map[int]int{}
	for start < len(choice) && height + 1 <= k {
        if  mp[choice[start]] == 1 {
            start ++
			continue
		}
		tmp1 := make([]int,len(tmp))
		copy(tmp1,tmp)
		tmp1 = append(tmp1, choice[start])
		mp[choice[start]] = 1
		f12(k,choice,height + 1,start + 1,tmp1,result)
		start ++
	}

}

