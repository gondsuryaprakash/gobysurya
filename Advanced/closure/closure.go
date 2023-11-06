package closure

import "fmt"

func Adder() func(int) int {

	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// return anonymous function for the closure
func Incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// fibNocci of the number

func FibnociOfNumber() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func CallFibGen() {
	fibGen := FibnociOfNumber()
	for i := 0; i < 10; i++ {
		fmt.Println(fibGen())
	}
}

func MakeGreeter(greeter string) func(string) string {
	return func(name string) string {
		return greeter + ", " + name
	}
}

func CallGreetClosure() {

	helloGreeter := MakeGreeter("Hello")
	namsteGreeter := MakeGreeter("Namaste")

	fmt.Println(helloGreeter("Surya"))
	fmt.Println(namsteGreeter("Surya"))

}
