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
	layoutTZ = "Monday, January 2, 2006 15:04:05"
)
func convert(date string) { // Main funcition to convert the timestamp
	t, _ := time.Parse(layoutUS, date) // Formatting time to layoutUS
	if t.Unix() == -62135596800 {  // Chekcks if the format entered is valid
		fmt.Println("Your date is invalid! Check the format, and try again!")
		fmt.Println("------------")
		return // Returns if the format is invalid
	} else {
	fmt.Println("------------")
	fmt.Println("UNIX:", t.Unix()) // Prints UNIX timestamp 
	fmt.Println("ISO:", t)   //Prints ISO timestamp
	fmt.Println("RFC1123:", t.Format(time.RFC1123)) // Prints RFC1123 timestamp
	loc, err := time.LoadLocation("UTC")
    if err != nil {
        fmt.Println(err)
    }
    t = t.In(loc)
    fmt.Println("UTC:", t.Format(layoutTZ))  // Prints UTC timestamp using layoutTZ format
    loc, err = time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println(err)
    }  
    t = t.In(loc)
	fmt.Println("EST:", t.Format(layoutTZ)) // Prints EST timestamp using layoutTZ format
	//PST
    loc, err = time.LoadLocation("America/Los_Angeles")
    if err != nil {
        fmt.Println(err)
    }
    t = t.In(loc)
    fmt.Println("PST:", t.Format(layoutTZ)) // Prints PST timestamp using layoutTZ format
	fmt.Println("------------")
}}
func main() {
	for { // Keeps program running untill closed so the user can format multiple timestamps
	scanner := bufio.NewScanner(os.Stdin) 
	fmt.Println("Please enter a timestamp in this format: \"dd.MM.yy HH:mm:SS Z\" (02.01.06 15:04:00 MST)")
	fmt.Println("------------")
	fmt.Print("> ")
	if scanner.Scan() { // Reads date input
		convert(scanner.Text()) // Calls convert function on the entered input
	}} 
} 
