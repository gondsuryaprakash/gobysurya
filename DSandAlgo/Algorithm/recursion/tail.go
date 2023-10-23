package recursion

import "fmt"

func PrintNTo1(n int) {
	if n > 0 {
		fmt.Println("n :", n)

		PrintNTo1(n - 1)
	}
}

/*
n : 5
n : 4
n : 3
n : 2
n : 1
*/
