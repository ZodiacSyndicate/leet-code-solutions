/*
 * @lc app=leetcode.cn id=1 lang=rust
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.37%)
 * Total Accepted:    238.5K
 * Total Submissions: 537.5K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */
use std::collections::HashMap;
impl Solution {
	fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
		let mut map: HashMap<i32, i32> = HashMap::new();
		let mut res: Vec<i32> = vec![];
		for (i, &n) in nums.iter().enumerate() {
			let t = target - n;
			if map.get(&t).is_some() {
				return vec![*map.get(&t).unwrap(), i as i32];
			} else {
				map.insert(n, i as i32);
			}
		}
		vec![]
	}
}
