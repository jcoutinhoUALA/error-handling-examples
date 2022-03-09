package main

import "fmt"

func boom() error {
	errorMsgDetail := "error 1001"
	return fmt.Errorf("boom error function: %s", errorMsgDetail)
}

func main() {
	if err := boom(); err != nil {
		fmt.Printf("Boom error occurred!: %s", err)
	} else {
		fmt.Println("All is fine!")
	}
}
