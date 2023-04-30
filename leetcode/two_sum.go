package leetcode

func TwoSum(nums []int, target int) []int {
	var result []int

	numsLen := len(nums)
	for x := 0; x < numsLen; x++ {
		for y := x + 1; y < numsLen; y++ {
			if nums[x]+nums[y] == target {
				result = []int{x, y}
			}
		}
	}

	return result
}
