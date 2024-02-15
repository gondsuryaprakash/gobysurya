package leetcode

func RearrangeArray(nums []int) []int {

	// Simple approach

	positivArray := []int{}
	negativArray := []int{}

	for _, v := range nums {
		if v > 0 {
			positivArray = append(positivArray, v)
		} else {
			negativArray = append(negativArray, v)
		}
	}

	result := make([]int, len(nums))

	for i, v := range positivArray {
		result[i*2] = v
	}

	for i, v := range negativArray {
		result[i*2+1] = v
	}
	return result

	// With two pointers 

	
}
