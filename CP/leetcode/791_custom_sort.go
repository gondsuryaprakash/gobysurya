package leetcode

func CustomSortString(order string, s string) string {

	rc := make(map[rune]int)

	for _, v := range s {
		if _, ok := rc[v]; !ok {
			rc[v] = 1
		} else {
			rc[v] += 1
		}
	}
	result := ""
	for _, v := range order {
		if val, ok := rc[v]; ok {
			for val > 0 {
				result += string(v)
				rc[v] -= 1
				val--
			}

		}
	}

	for k, v := range rc {
		for v > 0 {
			result += string(k)
			v--
		}
	}

	return result

}
