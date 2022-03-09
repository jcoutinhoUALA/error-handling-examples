package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type customErr struct {
	HTTPCode int
	Issue    string
}

func (c customErr) Error() string {
	return c.Issue
}

var (
	internalError = customErr{
		HTTPCode: http.StatusInternalServerError,
		Issue:    "something went wrong on the DB",
	}
	badRequest = customErr{
		HTTPCode: http.StatusBadRequest,
		Issue:    "the validation was not successful",
	}
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if err := action(r1.Intn(10)); err != nil {
		var ce customErr
		if errors.As(err, &ce) {
			fmt.Println(ce.HTTPCode)
			log.Printf(ce.Issue)
			return
		}
		log.Fatalln("the error is not the expected")
	}
}

func action(data int) error {
	if data%2 == 0 {
		return internalError
	}
	return badRequest
}
