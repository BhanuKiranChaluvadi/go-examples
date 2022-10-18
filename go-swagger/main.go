package main

import (
	"fmt"
	"swaggertutorial/models"
)

func main() {
	fmt.Println("Hello World")
	error := models.Error{
		Result:  0,
		Code:    "",
		Message: nil,
		Details: "",
		Target:  &models.Target{},
	}
	fmt.Printf("%+v\n", error)
}
