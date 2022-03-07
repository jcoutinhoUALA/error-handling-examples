package main

import (
	"errors"
	"fmt"
)

func boom() customErrors {
	return &boomError{}
}

func harmless() customErrors {
	return &harmlessError{}
}

func main() {
	err := boom()
	// Tengo que asegurarme de haber sobrescrito "Is" en cada error custom. Uno que no lo tenga va a producir un comportamiento inesperado
	// Se hace más "tediosa" la implementación de errores de ésta manera, y propenso a errores si nos olvidamos en algún error custom sobreescribir la función "Is"
	if isBoomError := errors.Is(err, &boomError{}); isBoomError {
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

func (c *boomError) Error() string {
	return fmt.Sprintf("boom error function")
}

func (c *boomError) GetMsgDetail() string {
	return fmt.Sprintf("error 1001")
}

func (c *boomError) Is(err error) bool {
	_, isBoomError := err.(*boomError)
	return isBoomError
}

func (c *harmlessError) Error() string {
	return fmt.Sprintf("harmless error function")
}

func (c *harmlessError) GetMsgDetail() string {
	return fmt.Sprintf("little error, but all is fine")
}

func (c *harmlessError) Is(err error) bool {
	_, isHarmlessError := err.(*harmlessError)
	return isHarmlessError
}
