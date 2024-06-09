package leetcode

func NumJewelsInStones(jewels string, stones string) int {

	rc := make(map[byte]int)

	for i := 0; i < len(stones); i++ {
		rc[stones[i]]++
	}
	sum := 0
	for i := 0; i < len(jewels); i++ {
		sum += rc[jewels[i]]
	}
	return sum
}
