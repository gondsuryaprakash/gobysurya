package leetcode

func FindSubstring3(s string, words []string) []int {

	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	indexArray := []int{}
	record := make(map[string]int)

	for _, val := range words {
		if count, ok := record[val]; !ok {
			record[val] = 1
		} else {
			record[val] = count + 1
		}
	}

	windowSize := len(words) * len(words[0])
	wordSize := len(words[0])

	for i := 0; i <= len(s)-windowSize; i++ {

		// Seen Word record will check new seen
		seenWordRecord := make(map[string]int)

		for j := 0; j <= windowSize; j += 1 {

			if j+i+wordSize > len(s) {
				break
			}

			newWords := s[j+i : j+i+wordSize]
			// Check newWord is present in record map or not
			if _, isValid := record[newWords]; !isValid {
				break
			}

			// Update in seenRecordMap

			if seenCount, isSeenBefore := seenWordRecord[newWords]; !isSeenBefore {
				seenWordRecord[newWords] = 1
			} else {
				seenWordRecord[newWords] = seenCount + 1
			}

			// Now check if the seenCount is greater or not
			if record[newWords] < seenWordRecord[newWords] {
				break
			}
			j += wordSize - 1

			if j+1 == windowSize {
				indexArray = append(indexArray, i)
			}

		}

	}

	return indexArray
}
