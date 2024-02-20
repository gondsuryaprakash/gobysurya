package leetcode

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	start := 0
	end := len(heights) - 1
	bricksUsed := bricks
	ladderUsed := ladders

	for start < end {
		next := start + 1
		// Case 1  if start greater then < next one
		diff := heights[next] - heights[start]
		if heights[start] >= heights[next] {
			start = next
		} else if heights[start] < heights[next] {
			// First check can we use bricks

			if diff <= bricksUsed {

				// Here we have two possibility either we can use ladder or bricks
				// to ensure which should we use

				start = next
				bricksUsed -= diff
			} else if ladderUsed > 0 {
				start = next
				ladderUsed -= 1
			} else {
				return start
			}
		}
	}
	return start
}

func WhichCanUse(start, next int, heights []int) {

	// lets start with ladders

}
