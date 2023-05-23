package structs

import (
	"fmt"
	"unsafe"
)

func StructWithPadding() {

	type A struct {
		a int8
		b int64
		c int16
	}

	type B struct {
		a int8
		b int16
		c int64
	}

	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
}
