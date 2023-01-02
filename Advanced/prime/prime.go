package prime

import "fmt"

type GeoMetry interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}
type Traingle struct {
	Base   float64
	Height float64
}
type Rectangle struct {
	Lenght float64
	Width  float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c *Traingle) Area() float64 {
	return 0.5 * c.Base * c.Height
}
func (c *Rectangle) Area() float64 {
	return c.Lenght * c.Width
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c *Traingle) Perimeter() float64 {
	return 0.99
}
func (c *Rectangle) Perimeter() float64 {
	return 2 * (c.Lenght + c.Width)
}

func NewTraingle() GeoMetry {
	return &Traingle{
		Base:   10,
		Height: 12,
	}
}

func Measure() float64 {
	traingle := NewTraingle()

	fmt.Println(traingle.Area())
	return traingle.Area()
}

func main() {

	Measure()

}
