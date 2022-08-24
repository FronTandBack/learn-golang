package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const marsDistance = 62100000
	const secondsPerDay = 86400

	stacions := [3]string{"Space Adventures", "SpaceX", "Vigring Galactic"}
	tripType := [2]string{"One-Way", "Round-Trip"}
	fmt.Println("Spaceline        Days    Trip type    Price")
	fmt.Println("===========================================")
	for count := 10; count >= 0; {

		stacion := stacions[rand.Intn(3)]

		spaceSpeed := rand.Intn(15) + 16

		howManyDays := marsDistance / spaceSpeed / secondsPerDay

		whatsTrip := tripType[rand.Intn(2)]

		price := rand.Intn(15) + 36

		if spaceSpeed > 22 && whatsTrip == "Round-Trip" {
			price *= 2
		}
		// str := stacion + "    " + string(howManyDays) + "   " + whatsTrip + "    " + string(price)

		// fmt.Println(str)
		fmt.Printf("%-16v %4v %-10v $%4v\n", stacion, howManyDays, whatsTrip, price)

		count--
	}

}
