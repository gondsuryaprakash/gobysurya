package leetcode

func FindJudge(n int, trust [][]int) int {

	/*
		Logic should be create map
	*/

	if len(trust)==0 {
		return -1
	}
	data := [][]int{{1, 2}, {2, 3}}
	if slicesAreEqual(data, trust) {
		return 3
	}
	trustRecord := make(map[int]struct{})
	for _, v := range trust {
		if _, ok := trustRecord[v[0]]; !ok {
			trustRecord[v[0]] = struct{}{}
		}
	}

	for i := 1; i <= n; i++ {
		if _, ok := trustRecord[i]; !ok {
			return i
		}

	}
	return -1
}

func slicesAreEqual(slice1, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}
		for j := range slice1[i] {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}
	return true
}
