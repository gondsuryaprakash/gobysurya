package array

import "fmt"

func CreatingArray() {

	// Initialising array
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1)
	var a [5]int
	a[4] = 100
	fmt.Println("Arr a :", a)
	fmt.Println("Arr a[4] :", a[4])
	fmt.Println("Arr len", len(a))

	var twoDArray [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDArray[i][j] = i + j
		}
	}

	fmt.Println("TwoD Array: ", twoDArray)
}
func CreatingSlice() {

	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// Get all value
	fmt.Println("s :", s)

	// Get Indexed value
	fmt.Println("s[1] :", s[1])
	fmt.Println("s[2] :", s[2])

}

func main() {
	CreatingArray()
	CreatingSlice()
}
