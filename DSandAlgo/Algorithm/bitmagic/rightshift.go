package bitmagic

import "fmt"

/*
1. Remove and Ignore trainling y bits
2. Move remaining 32-y bits in rightmost
3. Add y 0s at rightmost
*/

func RightShiftInGo() {

	a := 32
	rightShift := a >> 2

	fmt.Println("Right Shift in Bit :", rightShift)

}
