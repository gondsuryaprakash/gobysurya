package bitmagic

import "fmt"

/*
XOR Table
a  b  a^b
0  0   0
1  1   0
1  0   1
0  1   1
*/

func XORinBit() {
	a := 2
	b := 3
	xor := a ^ b
	fmt.Println("Numbers after XOR :", xor)
}
