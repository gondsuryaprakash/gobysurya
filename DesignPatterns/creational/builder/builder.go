package builder

type Process interface {
	SetVehicle() 
	
}


type VehicleProduct struct {
	Seats  int
	Wheels int
	Name   string
}


