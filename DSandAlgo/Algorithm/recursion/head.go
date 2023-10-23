package recursion

import "fmt"

func Print1ToN(n int) {

	if n > 0 {
		Print1ToN(n - 1)
		fmt.Println(" n:", n)
	}

}

/*
n : 1
n : 2
n : 3
n : 4
n : 5
*/
