package purchase

func NeedsLicense(kind string) bool {
	if kind =="car" || kind == "truck"{
    	return true
    }
        return false
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
        return option1 + " is clearly the better choice."
    }
    return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {	
    var estimatePrice float64
    
	if age < 3 {
        estimatePrice = originalPrice * 0.8
        return estimatePrice 
    } else if age >= 10 {
        estimatePrice = originalPrice * 0.5
        return estimatePrice
    } else {
        estimatePrice = originalPrice * 0.7
        return estimatePrice
    }
}
