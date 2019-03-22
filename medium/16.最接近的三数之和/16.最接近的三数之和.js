/*
 * @lc app=leetcode.cn id=16 lang=javascript
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (39.31%)
 * Total Accepted:    18.9K
 * Total Submissions: 48.1K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function(nums, target) {
  nums = nums.sort((a, b) => a - b)
  let closest = nums[0] + nums[1] + nums[2]
  for (let [i, num] of nums.entries()) {
    let l = i + 1,
      r = nums.length - 1
    while (l < r) {
      const threeSum = num + nums[l] + nums[r]
      if (Math.abs(threeSum - target) < Math.abs(closest - target)) {
        closest = threeSum
      }
      if (threeSum > target) {
        r--
      } else if (threeSum < target) {
        l++
      } else {
        return threeSum
      }
    }
  }
  return closest
}
