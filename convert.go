package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	// Basic time formats i will be using throught the script
	layoutISO = "2006-01-02T15:04:05-0700"
	layoutUS  = "02.01.06 15:04:00 MST"
	layoutTZ = "Monday, January 02, 2006 15:04:05 MST"
)

// Function to convert the timezone of the timestamp
func timezone(zone, date string) {

	t, _ := time.Parse(layoutUS, date)
	
		 loc, err := time.LoadLocation(zone)
		if err != nil {
			fmt.Println(err)
		}
		t = t.In(loc)
		fmt.Println(zone + ":", t.Format(layoutTZ))  // Prints zone timestamp using layoutTZ format
	}

func convert(date string) { // Main funcition to convert the timestamp
    t, err := time.Parse(layoutUS, date) // Formatting time to layoutUS
    if err != nil {  // Chekcks if the format entered is valid
    	fmt.Println("Your date is invalid! Check the format, and try again!")
        fmt.Println("------------")
        return // Returns if the format is invalid
	} else {
	scanner1 := bufio.NewScanner(os.Stdin) 
	fmt.Println("Please enter a timezone in this format: \"America/New_York\" (Supports all IANA timezone database zones)")
	fmt.Println("------------")
	fmt.Print("> ")
	if scanner1.Scan() { // Reads date input
	fmt.Println("------------")
	fmt.Println("UNIX:", t.Unix()) // Prints UNIX timestamp 
	fmt.Println("ISO:", t)   //Prints ISO timestamp
	fmt.Println("RFC1123:", t.Format(time.RFC1123)) // Prints RFC1123 timestamp
	timezone("Local", date) // Prints local timestamp
	timezone("UTC", date) // Prints UTC timestamp
	timezone(scanner1.Text(), date) // Calls convert function on the entered input	
}}}


func main() {
	for { // Keeps program running untill closed so the user can format multiple timestamps
	scanner := bufio.NewScanner(os.Stdin) 
	fmt.Println("------------")
	fmt.Println("Please enter a timestamp in this format: \"dd.MM.yy HH:mm:SS Z\" (02.01.06 15:04:00 MST)")
	fmt.Println("------------")
	fmt.Print("> ")
	if scanner.Scan() { // Reads date input
		fmt.Println("------------")
		convert(scanner.Text()) // Calls convert function on the entered input
	}} }
