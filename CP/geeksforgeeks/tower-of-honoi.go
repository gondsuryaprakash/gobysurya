package geeksforgeeks

import "fmt"

func TowerOfHonoi(n int, from, to, aux string) {
	if n == 0 {
		return
	}

	TowerOfHonoi(n-1, from, aux, to)
	fmt.Println("Move Disk Number: ", n, " From ", from, " To :", to, " Aux :", aux)
	TowerOfHonoi(n-1, aux, to, from)
}
