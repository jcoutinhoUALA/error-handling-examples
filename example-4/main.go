package main

import "fmt"

func boom() customErrors {
	return boomError{}
}

func harmless() customErrors {
	return harmlessError{}
}

func main() {
	err := boom()
	// Si no declaro los errores custom como punteros, no puedo discriminarlos correctamente por type assertions
	if _, isBoomError := err.(boomError); isBoomError {
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

type harmlessError struct {
	customErrors
	errorMsgDetail string
}

func (c boomError) Error() string {
	return fmt.Sprintf("boom error function")
}

func (c boomError) GetMsgDetail() string {
	return fmt.Sprintf("error 1001")
}

func (c harmlessError) Error() string {
	return fmt.Sprintf("harmless error function")
}

func (c harmlessError) GetMsgDetail() string {
	return fmt.Sprintf("little error, but all is fine")
}
