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

	_, operation1Error := operation1()

	// Al mismo nivel una asignacion de err con :=, pisa el valor original
	//ToDo: Si en la práctica cada error lo nombramos de forma "custom" específica, hay mucho menos probabilidad de "pisar sin querer" una variable llamada "err"
	// Si voy a llamar "operation2Error" al medio de una función, siendo operation2 una función que antes no existía, me quedo tranquilo de que no piso ni rompo nada
	// posterior a la línea que agrego

	//operation2Result, operation2Error := operation2()
	//fmt.Println(operation2Result)

	//ToDo: Puedo declarar una variable operation3Error por fuera del if sin ningun problema si quisiera darle un tratamiento por fuera del if
	// nombrando a todos los errores como "err", el compilador no me dejaría hacer ésto

	if _, isBoomError := operation1Error.(*boomError); isBoomError {

		// Referencia al err externo
		fmt.Printf("Error a la entrada del if: %s \n", operation1Error)

		// Dentro de un if no se pisa el valor de err original con ":="
		// Termina siendo MUY fina la línea entre "ups se me escapó un :=" o "ups se me olvidó un :=", pudiendo cambiar el comportamiento deseado
		// Este tipo de confusiones las he visto en otros proyectos y llegan a generar BUGS con todas las letras
		_, operation3Error := operation3()
		// POSTERIORMENTE A LO ANTERIOR, YA NO PUEDO REFERENCIAR AL err "original" dentro del if
		// Referencia al err interno
		fmt.Printf("Error dentro del if sobrescrito: %s \n", operation3Error)


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
	}

	// Por fuera del if, volvemos a tener acceso al "err" original
	fmt.Printf("Error fuera del if: %s \n", operation1Error)

	if _, isBoomError := operation1Error.(*boomError); isBoomError {
		fmt.Printf("Boom error occurred!: %s \n", operation1Error)
		fmt.Printf("Error detail: %s", operation1Error.GetMsgDetail())
	} else if _, isHarmlessError := operation1Error.(*harmlessError); isHarmlessError {
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
