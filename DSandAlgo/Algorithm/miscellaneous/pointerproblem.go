package miscellaneous

import "fmt"

func PoniterProblem() {

	a := []*int{}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		fmt.Println(&i)
		a = append(a, &i)
	}
	fmt.Println("Value of := ", *a[0], *a[1], *a[2]) 
	fmt.Println("Value of := ", a[0], a[1], a[2])    

}
