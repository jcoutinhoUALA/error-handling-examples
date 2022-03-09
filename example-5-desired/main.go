package main

import "fmt"

func boom() customErrors {
	return &boomError{}
}

func harmless() customErrors {
	return &harmlessError{}
}

func main() {
	err := boom()
	// Poniéndolos como punteros, de hecho me obliga en el assert colocar puntero
	// A su vez, los tipos de datos en el assert que puedo colocar se reducen al dominio de errores interfaceados como "customErrors"
	// Es una opción amigable y menos propensa a malentendidos en el manejo de los errores
	// El assert es independiente de si llamando a la función, se retorna un error con New o se utiliza un var
	if _, isBoomError := err.(*boomError); isBoomError {
		fmt.Printf("Boom error occurred!: %s \n", err)
		fmt.Printf("Error detail: %s", err.GetMsgDetail())
	} else if _, isHarmlessError := err.(*harmlessError); isHarmlessError {
		fmt.Println("Error but all is fine!")
	} else {
		fmt.Println("Without errors!")
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

func (c *harmlessError) Error() string {
	return fmt.Sprintf("harmless error function")
}

func (c *harmlessError) GetMsgDetail() string {
	return fmt.Sprintf("little error, but all is fine")
}
