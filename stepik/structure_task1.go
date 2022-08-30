package main

import "fmt"

type Bike struct {
	On    bool
	Ammo  int
	Power int
}

func (bike *Bike) Shoot() bool {
	return bike.operation(&bike.Ammo)
}

func (bike *Bike) RideBike() bool {
	return bike.operation(&bike.Power)
}

func (bike *Bike) operation(observed *int) bool {
	if !bike.On || *observed <= 0 {
		return false
	} else {
		*observed--
		return true
	}
}

func main() {
	testStruct := &Bike{On: true, Ammo: 3, Power: 5}

	fmt.Println(testStruct.RideBike())
}
