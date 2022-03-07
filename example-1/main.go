package main

import (
	"errors"
	"fmt"
)

func boom() error {
	return errors.New("boom error function")
}

func main() {
	if err := boom(); err != nil {
		fmt.Printf("Boom error occurred!: %s", err)
	} else {
		fmt.Println("All is fine!")
	}
}
