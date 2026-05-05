package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	card := []int{2,6,9}
    return card
}

// GetItem retrieves an item from a slice at given position.
func GetItem(slice []int, index int) int {
        if index < 0 || index >= len(slice) {
            return -1
        }
    return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	} 
    slice = append(slice[:index], slice[index+1:]...)
    return slice
    
}
