package panic

import "fmt"

func CreatePanic(arr []string, index int) {

	if index > len(arr)-1 {
		panic("Array out of Boud")
	}
	fmt.Println(arr[index])
}

func CreatePanicWithDefer() {
	defer fmt.Println("In Defer")
	panic("Panic in function ")
	fmt.Println("After panic")
}

func CreatePanicWithDifferentFunction() {
	f1()
}

func CreatePanicAndRecover(arr []string, index int) {

	defer handleRecover()

	if index > len(arr)-1 {
		panic("Out of bound")
	}

	fmt.Println(arr[index])
}

func CreatePanicWithConcurrency(arr []string, index int){
	


}

func handleRecover() {
	r := recover()
	if r != nil {
		fmt.Print("recovered from Panic")
	}
}

func f1() {
	defer fmt.Println("Defer in f1")
	f2()
	fmt.Println("After Panic")
}
func f2() {
	defer fmt.Println("Defer in f2")
	panic("Panic in f2")
	fmt.Println("After Panic")
}
