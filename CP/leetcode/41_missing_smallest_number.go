package leetcode

// import "sort"

// // O(n)
// // space O(1)

// func FirstMissingPositive(nums []int) int {
//     records:= make(map[int]bool)

//     startIndex:= 1
// 	for i:=0;i<len(nums);i++ {
//         records[nums[i]]= true

// 	}

// 	sort.Ints(nums)
// 	// base condition
// 	if nums[0] > 1 {
// 		return 1
// 	}
// 	var index int
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] < 1 {
// 			index++
// 		}
// 	}

// 	nums = nums[:index+1]

// 	// Finding the missing number
// 	for i:=0;i<len(nums)-1;i++ {
// 		if nums[i+1]-

// 	}

// 	//[1,2,0]

// 	return 0
// }
