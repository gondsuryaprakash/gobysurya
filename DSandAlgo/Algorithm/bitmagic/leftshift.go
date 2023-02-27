package bitmagic

import "fmt"

/*
x<<y
1. Remove leading bit from y bit
2. Move remaining 32-y in left
3. Append 0s at the end
*/

func LeftShiftInBit() {

	a := 10
	leftshift := a << 1
	fmt.Println("After Left Shift Bit : ", leftshift)
}
