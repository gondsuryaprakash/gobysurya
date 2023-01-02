package structs

import "fmt"

type Customer struct {
	name    string
	address string
	bal     float64
}

func GetCustomerInfo(c Customer) {
	fmt.Printf("%s owes us %.2f\n ", c.name, c.bal)
	fmt.Printf("%s Customer address\n", c.address)
}

func NuewCustomerAddress(c *Customer, address string) {
	c.address = address
}

func CreateStruct() {
	var ts Customer
	ts.name = "Surya"
	ts.address = "Vill RampurMeer"
	ts.bal = 23.3

	ss := Customer{"Surya", "Maharajgan", 23.0}

	fmt.Println(ss.address)

	GetCustomerInfo(ts)
	NuewCustomerAddress(&ts, "Village Rampurmeer Post Patrengawa")
	GetCustomerInfo(ts)
}
