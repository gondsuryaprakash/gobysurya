package goquiz

import "fmt"

func Op1() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x)

}

func Op2() {
	var x = []string{"A", "B", "C"}
	for i, s := range x {
		fmt.Println(x)
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		fmt.Println("After Append", x)
		x[i+1] = "Z"
		fmt.Println("After Append Z", x)
	}
}
