package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	layoutISO = "2006-01-02T15:04:05-0700"
	layoutUS  = "02.01.06 15:04:00 MST"
	layoutTZ = "Monday, January 2, 2006 15:04:05"
)

func unixToDate(date string) {
	t, _ := time.Parse(layoutUS, date)
	if t.Unix() == -62135596800 {
		fmt.Println("Your date is invalid! Check the format, and try again!")
		fmt.Println("------------")
		return
	} else {
	fmt.Println("------------")
	// UNIX
	fmt.Println("UNIX:", t.Unix()) 
	// ISO
	fmt.Println("ISO:", t)  
	fmt.Println("RFC1123:", t.Format(time.RFC1123))
	// UTC
	loc, err := time.LoadLocation("UTC")
    if err != nil {
        fmt.Println(err)
    }
    t = t.In(loc)
    fmt.Println("UTC:", t.Format(layoutTZ))  
	// EST
    loc, err = time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println(err)
    }  
    t = t.In(loc)
	fmt.Println("EST:", t.Format(layoutTZ))
	//PST
    loc, err = time.LoadLocation("America/Los_Angeles")
    if err != nil {
        fmt.Println(err)
    }
      
    t = t.In(loc)
    fmt.Println("PST:", t.Format(layoutTZ))   
	fmt.Println("------------")
}}

func main() {
	for {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a timestamp in this format: \"dd.MM.yy HH:mm:SS Z\" (02.01.06 15:04:00 MST)")
	fmt.Println("------------")
	fmt.Print("> ")
	if scanner.Scan() {
		unixToDate(scanner.Text())
	}}
} 
