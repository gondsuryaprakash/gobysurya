package adaptermodel

import "fmt"

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (dog *Dog) Bark() {
	fmt.Println(dog.name, " Bark")
}

func (cat *Cat) Meow() {
	fmt.Println(cat.name, " Meow")

}

type NoiseMaker interface {
	noise()
}

type DogAdapter struct {
	dog *Dog
}
type CatAdapter struct {
	cat *Cat
}

func (d *DogAdapter) noise() {
	d.dog.Bark()
}

func (c *CatAdapter) noise() {
	c.cat.Meow()
}

func PlayWithAnimal(n NoiseMaker) {
	n.noise()
}

func CallAdapterPattern() {
	dog := &Dog{
		name: "Kutta",
	}
	cat := &Cat{
		name: "Billi",
	}

	PlayWithAnimal(&DogAdapter{
		dog: dog,
	})

	PlayWithAnimal(&CatAdapter{
		cat: cat,
	})

}
