/*
 * @lc app=leetcode.cn id=475 lang=javascript
 *
 * [475] 供暖器
 *
 * https://leetcode-cn.com/problems/heaters/description/
 *
 * algorithms
 * Easy (26.72%)
 * Total Accepted:    1.9K
 * Total Submissions: 7.1K
 * Testcase Example:  '[1,2,3]\n[2]'
 *
 * 冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
 *
 * 现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。
 *
 * 所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。
 *
 * 说明:
 *
 *
 * 给出的房屋和供暖器的数目是非负数且不会超过 25000。
 * 给出的房屋和供暖器的位置均是非负数且不会超过10^9。
 * 只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
 * 所有供暖器都遵循你的半径标准，加热的半径也一样。
 *
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,3],[2]
 * 输出: 1
 * 解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,2,3,4],[1,4]
 * 输出: 1
 * 解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
 *
 *
 */
/**
 * @param {number[]} houses
 * @param {number[]} heaters
 * @return {number}
 */
var findRadius = function(houses, heaters) {
  heaters = heaters.sort((a, b) => a - b)
  let res = 0

  for (const house of houses) {
    let index = binarySearch(heaters, house)
    if (index < 0) {
      index = ~index
      const dis1 =
        index - 1 >= 0 ? house - heaters[index - 1] : Number.MAX_SAFE_INTEGER
      const dis2 =
        index < heaters.length
          ? heaters[index] - house
          : Number.MAX_SAFE_INTEGER
      res = Math.max(res, Math.min(dis1, dis2))
    }
  }
  return res
}

function binarySearch(arr, target) {
  let l = 0
  let r = arr.length - 1
  if (target > arr[r]) {
    return -arr.length - 1
  }
  while (l <= r) {
    const mid = Math.floor(l + (r - l) / 2)
    if (arr[mid] === target) {
      return mid
    } else if (arr[mid] < target && arr[mid + 1] > target) {
      return -mid - 2
    } else if (arr[mid] < target) {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  return -1
}
