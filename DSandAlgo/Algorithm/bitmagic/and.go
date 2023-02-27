package bitmagic

import "fmt"

/*
AND Table
a  b  a&b
0  0   0
1  1   1
1  0   0
0  1   0
*/

func AndInBit() {
	a := 2 // 000010
	b := 3 // 000011

	c := b & a // 2
	fmt.Println("After and in Numbers :", c)
}
