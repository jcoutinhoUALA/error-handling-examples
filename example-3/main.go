package main

import "fmt"

func boom() customErrors {
	return boomError{}
}

func main() {
	if err := boom(); err != nil {
		fmt.Printf("Boom error occurred!: %s \n", err)
		fmt.Printf("Error detail: %s", err.GetMsgDetail())
	} else {
		fmt.Println("All is fine!")
	}
}

type customErrors interface {
	error
	GetMsgDetail() string
}

type boomError struct {
	customErrors
	errorMsgDetail string
}

func (c boomError) Error() string {
	return fmt.Sprintf("boom error function")
}

func (c boomError) GetMsgDetail() string {
	return fmt.Sprintf("error 1001")
}
