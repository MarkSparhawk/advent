package main

func fuelNeeded(masses []int) int {
	fuel := 0
	for _, mass := range masses {
		fuel = fuel + (mass / 3 - 2)
	} 
	return fuel
}

func getFuel(mass int) int {
	fuel := mass / 3 - 2
	if (fuel) <= 0 {
		fuel = 0
		return fuel
	}
	return fuel + getFuel(mass / 3 - 2)
}
