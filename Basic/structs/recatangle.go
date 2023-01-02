package structs

import "fmt"

type Reactangle struct {
	Width, Height float64
}

func (c Reactangle) Area() float64 {
	return c.Height * c.Width
}

func (c Reactangle) Perimeter() float64 {
	return 2 * (c.Width + c.Height)
}

type GeoMetry interface {
	Area() float64
	Perimeter() float64
}

func NewShape(c *Reactangle) GeoMetry {
	return &Reactangle{
		Width:  c.Width,
		Height: c.Height,
	}

}

func CalCulateGeoMetry() {
	rect := Reactangle{
		Width:  34.6,
		Height: 23.7,
	}

	c := NewShape(&rect)

	fmt.Printf("Area of Rectangle : %.2f ", c.Area())
	fmt.Printf("Perimeter of Rectangle : %.2f ", c.Perimeter())
}
