package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now().Format(time.RFC3339)
	dateNow,_ := time.Parse(time.RFC3339,date)
	fmt.Println(dateNow.Second())
}


// ? Create time zone from date and time
func TajikistanTimeZone() time.Time {
	NewZone := time.FixedZone("tjk", 5*3600)
	Date := time.Now().UTC().In(NewZone)

	return Date
}