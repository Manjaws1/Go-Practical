package purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	var best string
	if option1 < option2 {
   		 best = option1
	} else {
 	   best = option2
	}
	return best + " is clearly the better choice."

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	} else if age >= 10 {
		return originalPrice * 0.5
	}
	return originalPrice * 0.7
}
