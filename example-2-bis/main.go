package main

import (
	"errors"
	"fmt"
)

var (
	boomError     = errors.New("boom error function")
	harmlessError = errors.New("harmless error function")
)

func boom() error {
	// Un programador desprevenido puede devolver un error.New, dado que no queda explícito en la firma qué error se quiere retornar
	// Un new de un error, en lugar del var, hará que el manejo del error en capas superiores falle y produzca errores inesperados
	// return errors.New("boom error function")
	return boomError
}

func harmless() error {
	return harmlessError
}

func main() {
	err := boom()
	// El método Is es más para errores estáticos devueltos que puedo referenciarlos 1 a 1
	// Pero tiene sus riesgos: Bajo el principio de caja negra, no tengo garantía de que el errors.Is funcione correctamente
	// Un developer que modifique la función "boom()" sin saber cómo se manipulan los errores más arriba, puede romper la app
	if isBoomError := errors.Is(err, boomError); isBoomError {
		fmt.Printf("Boom error occurred!: %s \n", err)
		// Se pierde el "beneficio" de tener funciones adicionales como "GetMsgDetail"
		//fmt.Printf("Error detail: %s", err.GetMsgDetail())
	} else {
		fmt.Println("All is fine!")
	}
}
