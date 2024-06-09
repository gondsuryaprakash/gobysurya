package leetcode

func MaximumStrongPairXor(nums []int) int {

	return 0

}

func isStrongestPair(a, b int) bool {
	modDiff := a - b
	if modDiff < 0 {
		modDiff = -(modDiff)
	}
	var min int
	if a < b {
		min = a
	} else {
		min = b
	}
	return modDiff <= min
}
