package main

import (
	"log"
	"math/rand"
	"time"
)

type SomeError string // un tipo básico

const ( // se inicializan constantes, no variables
	unknownErr  = SomeError("something went wrong")
	notFoundErr = SomeError("the resource was not found")
)

func (se SomeError) Error() string { // el tipo declarado cumple con la interface de Error
	return string(se)
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	err := getErr(r1.Intn(10))
	if err != nil {
		log.Fatalln(err)
	}
}

func getErr(dato int) error {
	if dato%2 == 0 {
		return unknownErr
	}
	return notFoundErr
}
