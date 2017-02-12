package main

import (
	"time"
	"os"
)

var ApiKey string = os.Args[1]
var PhoneNumber string = os.Args[2]
var PhoneNumber2 string = os.Args[3]
var SummonerId string = os.Args[4]


func main() {

	if isUserInGame(SummonerId) {
		text()
	}

	time.Sleep(1 * time.Hour)
	main()

}
