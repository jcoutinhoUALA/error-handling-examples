package main

import (
	"fmt"
)

func operation1() (string, customErrors) {
	return "operation1 result", &boomError{}
}

func operation2() (string, customErrors) {
	return "operation2 result", &harmlessError{}
}

func operation3() (string, customErrors) {
	return "operation3 result", nil
}

func main() {
	// Nombrando siempre a los errores como err, nos "obliga" a hacer un tratamiento inmediato de los errores (recomendación de go)
	// Durante la ejecución en una función, la variable err puede ser reutilizada o reasignada N veces

	_, err := operation1()
	//Otras líneas de código random
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------



	// Al mismo nivel una asignacion de err con :=, pisa el valor original
	//ToDo: Ejemplificar que teniendo una sola variable retornada por operation3(), el compilador da error al redeclarar otra variable de igual nombre con :=
	// Pero si son varios retornos, el compilador no da error y ésto puede ser engañoso, porque puede parecer "inofensiva" agregar
	// al medio de un código un := creyendo que estoy creando "de 0" una variable
	// Pero sin querer puedo estar pisando una variable de más arriba y explotando la lógica de más abajo y en definitiva explotando la app

	//operation3Result, err := operation3()
	// Agregado de lógica "inofensiva" al medio de una función ya existente
	//fmt.Println(operation3Result)



	//Otras líneas de código random
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------
	// -----------------------------



	fmt.Printf("Error antes del if: %s \n", err)
	if _, isBoomError := err.(*boomError); isBoomError {
		// Referencia al err externo
		fmt.Printf("Error dentro del if: %s \n", err)

		// Dentro de un if no se pisa el valor de err original con ":="
		// Termina siendo MUY fina la línea entre "ups se me escapó un :=" o "ups se me olvidó un :=", pudiendo cambiar el comportamiento deseado
		// Este tipo de confusiones las he visto en otros proyectos y llegan a generar BUGS con todas las letras
		_, err := operation2()
		// POSTERIORMENTE A LO ANTERIOR, YA NO PUEDO REFERENCIAR AL err "original" dentro del if
		// Referencia al err interno
		fmt.Printf("Error dentro del if sobrescrito: %s \n", err)
	}

	// Por fuera del if, volvemos a tener acceso al "err" original
	fmt.Printf("Error fuera del if: %s \n", err)

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
	return fmt.Sprintf("operation1 error function")
}

func (c *boomError) GetMsgDetail() string {
	return fmt.Sprintf("error 1001")
}

func (c *harmlessError) Error() string {
	return fmt.Sprintf("operation2 error function")
}

func (c *harmlessError) GetMsgDetail() string {
	return fmt.Sprintf("little error, but all is fine")
}
