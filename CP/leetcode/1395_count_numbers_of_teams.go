package leetcode

func NumTeams(rating []int) int {
	count := 0
	for i := 0; i < len(rating)-2; i++ {
		for j := i + 1; j < len(rating)-1; j++ {
			for k := j + 1; k < len(rating); k++ {
				if (rating[i] > rating[j] && rating[j] > rating[k]) || (rating[k] > rating[j] && rating[j] > rating[i]) {
					count++
				}
			}
		}
	}
	return count
}

func NumTeamsV2(rating []int) int {

	count := 0
	smaller := make([]int, len(rating))
	bigger := make([]int, len(rating))

	for i := 0; i < len(rating); i++ {

		for k := i + 1; k < len(rating); k++ {
			if rating[i] > rating[k] {
				smaller[i]++
			}
			if rating[i] < rating[k] {
				bigger[i]++
			}
		}
	}

	for i := 0; i < len(rating); i++ {

		for k := i + 1; k < len(rating); k++ {

			if rating[i] < rating[k] {
				count += bigger[i]
			}

			if rating[i] > rating[k] {
				count += smaller[i]
			}
		}
	}
	return count
}
