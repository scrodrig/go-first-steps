package main

import "fmt"

func main() {
	// initialize variables here

	var smsSendingLimit int = 0
	var costPerSMS float64 = 0.0
	var hasPermission bool = true
	var username string = "wagslane"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
