package leetcode

func CombinationSum(candidates []int, target int) [][]int {

	ans := [][]int{}
	tmp := []int{}
	findNumbers(candidates, target, &ans, tmp, 0)
	return ans
}

func findNumbers(candidates []int, target int, ans *[][]int, tmp []int, index int) {
	if target == 0 {
		*ans = append(*ans, tmp)
	}
	for i := index; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			tmp = append(tmp, candidates[i])
			findNumbers(candidates, target-candidates[i], ans, tmp, i)
			// removing element from the temp
			tmp = tmp[:len(tmp)-1]
		}
	}

}
