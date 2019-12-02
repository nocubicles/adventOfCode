package main

import (
	"math"
)

func getModuleFuelFuelRequirement(mass int) int {
	var moduleFuelFuelRequirementSum int = 0
	var requiredFuel int = mass
	stop := false

	for !stop {

		if (requiredFuel/3 - 2) > 0 {
			requiredFuel = int(math.Floor(float64(requiredFuel)/3) - 2)
			moduleFuelFuelRequirementSum += requiredFuel
		} else {
			stop = true
		}
	}

	return moduleFuelFuelRequirementSum
}

//GetAllFuelRequirement Returns all fuel requirement for the Rocket
func GetAllFuelRequirement() (int, int) {
	var moduleFuelRequirementSum int = 0
	var moduleFuelFuelRequirementSum int = 0

	for index := 0; index < len(DataDay1); index++ {
		fuelRequired := int(math.Floor(float64(DataDay1[index]/3)) - 2)
		moduleFuelFuelRequirement := getModuleFuelFuelRequirement(fuelRequired)

		moduleFuelRequirementSum = moduleFuelRequirementSum + fuelRequired
		moduleFuelFuelRequirementSum += moduleFuelFuelRequirement

	}

	return moduleFuelRequirementSum, moduleFuelRequirementSum + moduleFuelFuelRequirementSum
}
