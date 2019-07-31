/*
 * @lc app=leetcode.cn id=213 lang=javascript
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (33.55%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    7.1K
 * Total Submissions: 21.2K
 * Testcase Example:  '[2,3,2]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [2,3,2]
 * 输出: 3
 * 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 */
/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
  if (!nums.length) return 0
  const dp1 = new Array(nums.length + 1)
  const dp2 = new Array(nums.length + 1)
  dp1[1] = nums[0]
  dp1[2] = Math.max(nums[1], nums[0])
  for (let i = 3; i < nums.length; i++) {
    dp1[i] = Math.max(nums[i - 1] + dp1[i - 2], dp1[i - 1])
  }
  dp2[2] = nums[1]
  dp2[3] = Math.max(nums[1], nums[2])
  for (let i = 4; i < nums.length; i++) {
    dp2[i] = Math.max(nums[i - 1] + dp2[i - 2], dp2[i - 1])
  }
  return Math.max(dp1[nums.length - 1], dp2[nums.length - 1])
}
