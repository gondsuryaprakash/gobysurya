package interviewbit

func RepeatedNumber(A []int) []int {

	// finiding duplicate

	xor := 0

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		xor = xor ^ A[i]
	}

	totalSum := (len(A) * (len(A) + 1)) / 2
	duplicate := xor ^ len(A)

	missingNumber := (totalSum - sum) + duplicate

	return []int{duplicate, missingNumber}
}
