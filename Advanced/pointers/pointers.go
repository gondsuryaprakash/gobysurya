package pointers

// It doesn't Double the number 
func Double(x int)  {
	x+=x 
}

// It will double the number
func DoubleWithPtr(x *int) {
	*x+=*x
	x= nil  // Here Nothing will affect the 
}



