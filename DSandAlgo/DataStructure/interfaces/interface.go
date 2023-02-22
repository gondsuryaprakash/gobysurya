package interfaces

import "fmt"

type animal interface {
	breathe()
	walk()
}

type Lion struct {
	age int
}

func (l Lion) breathe() {
	fmt.Println("Lion Breathe ")
}

func (l Lion) walk() {
	fmt.Println("Lion walk")
}

func callingAnimal(a animal) {
	a.breathe()
	a.walk()
}

func CallInterface() {
	callingAnimal(Lion{age: 25})
	callingAnimal(Dog{age: 23})
}

type Dog struct {
	age int
}

func (d Dog) breathe() {
	fmt.Println("Dog breathe : ")
}

func (d Dog) walk() {
	fmt.Println("Dog run : ")

}
