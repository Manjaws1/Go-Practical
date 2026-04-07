package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    	carsPerHour := float64(productionRate) * (successRate / 100)
	return int(carsPerHour / 60)
}

func CalculateCost(carsCount int) uint {
	group := carsCount/10
    remain := carsCount % 10
    groupPrice := group * 95000
    remainPrice := remain * 10000
    totalPrice := groupPrice + remainPrice
    return uint(totalPrice)
}
