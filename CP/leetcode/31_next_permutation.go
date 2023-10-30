package leetcode

/*
Logic
Step 1 : first find the breakPoint (i.e the number from the last which is less then rightmost)
Step 2 : If no any breakpoint then reverse the number i.e. the answer
step 3 : If yes now from last search the first number which is grater than breakPoint element swap it
step 4 : now from breakPoint + 1 to last number reverse it. i.e. the number.
*/

import "fmt"

func NextPermutation(nums []int) {
	ind := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] > nums[i] {
			ind = i
			break
		}
	}
	if ind == -1 { // reverse the number
		reverse(nums, 0, len(nums)-1)
		return
	}
	// now search from breakPoint from last if any number is grather than
	for i := len(nums) - 1; i > ind; i-- {
		if nums[ind] < nums[i] {
			swap(nums, ind, i)
			break
		}
	}
	reverse(nums, ind+1, len(nums)-1)

	fmt.Println(nums)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--

	}
}

func swap(arr []int, ind1, ind2 int) {
	arr[ind1], arr[ind2] = arr[ind2], arr[ind1]
}
