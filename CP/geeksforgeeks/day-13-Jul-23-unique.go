package geeksforgeeks

/*
Given an array arr of N integers, the task is to check whether the frequency of the elements in the array is unique or not.
Or in other words, there are no two distinct numbers in array with equal frequency.
If all the frequency is unique then return true, else return false.
*/

/*
N = 5
arr = [1, 1, 2, 5, 5]
Output:
false
Explanation:
The array contains 2 (1’s), 1 (2’s) and 2 (5’s), since the number of frequency of 1 and 5 are the same i.e. 2 times.
Therefore, this array does not satisfy the condition.
*/

/*
Input:
N = 10
arr = [2, 2, 5, 10, 1, 2, 10, 5, 10, 2]
Output:
true
Explanation:
Number of 1’s -> 1
Number of 2’s -> 4
Number of 5’s -> 2
Number of 10’s -> 3.
Since, the number of occurrences of elements present in the array is unique. Therefore, this array satisfy the condition.
*/
func IsFrequencyUnique(arr []int) bool {
	records := make(map[int]int)
	uniqueRecords := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if val, ok := records[arr[i]]; ok {
			records[arr[i]] = val + 1
		} else {
			records[arr[i]] = 1
		}
	}
	for _, v := range records {
		if _, ok := uniqueRecords[v]; ok {
			return false
		}
		uniqueRecords[v] = true
	}
	return true
}
