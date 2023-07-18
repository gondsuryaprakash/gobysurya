package miscellaneous

import "fmt"

func PoniterProblem() {

	a := []*int{}
	fmt.Println("len", len(a))
	fmt.Println("cap", cap(a))
	for i := 0; i < 3; i++ { // i:=0  //0xc00001c810  3
		fmt.Println(i)
		fmt.Println(&i)
		a = append(a, &i)
	}
	fmt.Println("Value of := ", *a[0], *a[1], *a[2])
	fmt.Println("Value of := ", a[0], a[1], a[2])

}
