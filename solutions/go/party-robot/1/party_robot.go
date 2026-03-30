package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	wel := fmt.Sprintf("Welcome to my party, %s!", name)
        return wel
}

// H birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	wish := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return wish
}

// AssignTable assigns a table to each guest.

func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	message := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	return message
}
