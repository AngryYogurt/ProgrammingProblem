/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 > 0 {
			res++
		}
		num >>= 1
	}
	return res
}

// @lc code=end

