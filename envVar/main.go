package main

import "fmt"

func main() {
	fmt.Println("Hello World !")
	model := &Model{
		name:  "VW",
		value: "Up",
	}
	fmt.Println(model.GetAbsoluteInstallationPath())
}
