package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven == 0 {
        panic("RemainingOvenTime not implemented")
    }
    return OvenTime - actualMinutesInOven
    
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    intialPreparationTime := 2
    if numberOfLayers == 0 {
		panic("PreparationTime not implemented")
	}

    return numberOfLayers * intialPreparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	if numberOfLayers == 0 || actualMinutesInOven == 0 {
		panic("ElapsedTime not implemented")
	}
elapsedTime := PreparationTime(numberOfLayers) + actualMinutesInOven
	return elapsedTime
}
