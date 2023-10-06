package main

func main() {
	// init patient on which execution is done
	// this object is passed via multiple chain
	// of processors
	patient := Patient{name: "Doremon"}

	// init all stages of checkup
	reception := &Reception{}
	doctor := &Doctor{}
	medical := &Medical{}
	cashier := &Cashier{}

	// set next stages of all stages
	reception.setNext(doctor)
	doctor.setNext(medical)
	medical.setNext(cashier)

	reception.execute(&patient)
}

//Done Registration
//Doctor checkup done
//Medicine Done
//Payment Done
