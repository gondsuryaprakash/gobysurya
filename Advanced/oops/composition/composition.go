package composition

import "fmt"

type Engine struct {
	name string
}

func (e Engine) Start() {
	fmt.Println("Start")
}

type Car struct {
	Engine
	Wheel int
}

func NewCar() {
	c := Car{
		Engine: Engine{
			name: "PowerFull",
		},
		Wheel: 4,
	}
	c.Start()
}
