package slice

import "fmt"

func CreatingSlice() {
	// Creating Slice using make
	sl1 := make([]int, 0, 0)
	fmt.Println("len ", len(sl1))
	fmt.Println("cap ", cap(sl1))
	fmt.Println("Hello, 世界")

	sl2 := []int{1, 2, 3}
	sl1 = append(sl1, sl2...)

	fmt.Println("sl1 ", sl1)

	fmt.Println("len ", len(sl1))
	fmt.Println("cap ", cap(sl1))

	// Creating slice using new
	sl3 := new([]int)

	// new return the pointer of nil
	fmt.Println("sl3 ", sl3) // &[]

	fmt.Println(*sl3)              // []
	fmt.Println("len ", len(*sl3)) //len  0
	fmt.Println("cap ", cap(*sl3)) //cap  0

	*sl3 = append(*sl3, 1)
	fmt.Println(*sl3) // [1]

	fmt.Println("len ", len(*sl3)) //len 1
	fmt.Println("cap ", cap(*sl3)) //cap 1

}

// Accessing and Modification in Slice

func AccessingAndModification() {
	array := []int{1, 2, 3, 4, 5}
	slice := array[:]

	fmt.Println("array ", array)
	fmt.Println("slice ", slice)

	//Pointer value of slice and Array
	fmt.Println("&array[0] ", &array[0]) // &array  0xc000122038
	fmt.Println("&slice[0] ", &slice[0]) // &slice  0xc000122038

	array[1] = 7

	fmt.Println("After Modification in array")
	fmt.Println("array ", array)
	fmt.Println("slice ", slice)

	slice[1] = 2

	fmt.Println("After Modification slice")
	fmt.Println("array ", array)
	fmt.Println("slice ", slice)

}

