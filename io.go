package main

import "fmt"

func keyboard_input() {
	var make string
	fmt.Println("Make")
	fmt.Scanln(&make)

	var model string
	fmt.Println("Model")
	fmt.Scanln(&model)

	fmt.Println("Make and model are: " + make + " " + model)
}

func file_output() {

}

func main() {
	// keyboard_input()
}
