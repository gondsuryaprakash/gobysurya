package bitmagic

import "fmt"

/*
OR Table
a  b  a|b
0  0   0
1  1   1
1  0   1
0  1   1
*/
func ORinBit() {
	a := 2
	b := 3
	c := a | b
	fmt.Println("Numbers after OR :", c)
}
