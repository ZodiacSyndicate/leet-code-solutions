/*
 * @lc app=leetcode.cn id=475 lang=golang
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
import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	maxValue := (1 << 31) - 1
	sort.Sort(sort.IntSlice(heaters))
	result := 0

	for _, house := range houses {
		index := binarySearch(heaters, house)
		if index < 0 {
			index = ^index
			var dis1, dis2 int
			if index-1 >= 0 {
				dis1 = house - heaters[index-1]
			} else {
				dis1 = maxValue
			}
			if index < len(heaters) {
				dis2 = heaters[index] - house
			} else {
				dis2 = maxValue
			}
			result = max(result, min(dis1, dis2))
		}
	}
	return result
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	if target > arr[r] {
		return -len(arr) - 1
	}
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target && arr[mid+1] > target {
			return -mid - 2
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

