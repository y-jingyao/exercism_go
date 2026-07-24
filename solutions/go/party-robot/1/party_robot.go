package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	//panic("Please implement the Welcome function")
    	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	//panic("Please implement the HappyBirthday function")
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	//panic("Please implement the AssignTable function")
    line1 := Welcome(name)
	line2 := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	line3 := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return fmt.Sprintf("%s\n%s\n%s", line1, line2, line3)
}
