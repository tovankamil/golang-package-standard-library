package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	} else if id == "eko" {
		return NotFoundError
	} else {
		return nil
	}
}
func main() {
	err := GetById("eko")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found Error")
		} else {
			fmt.Println("unknown error")
		}

	}
}
