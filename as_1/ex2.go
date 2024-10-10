package main 

import "fmt"

func main() {
	var day int = 6
	var temperature float64 = 25.4
	var name string = "Kevin"
	var isLegit bool = false

	height := 34.7
	block := 'B'
	flag := true

	fmt.Printf("Day: %d, Type: %T\n", day, day)
	fmt.Printf("Temperature: %.2f, Type: %T\n", temperature, temperature)
    fmt.Printf("Name: %s, Type: %T\n", name, name)
    fmt.Printf("Is Legit: %t, Type: %T\n", isLegit, isLegit)
    fmt.Printf("Height: %.1f, Type: %T\n", height, height)
    fmt.Printf("Block: %c, Type: %T\n", block, block)
    fmt.Printf("Flag: %t, Type: %T\n", flag, flag)
}

// var allows to specify the type of a variable, can be used inside and outside the function
// := is a short version of declaration and can be used only inside a function

// You can print the type of a variable using fmt.Printf with the %T format specifier

// No. Because Go is a statically typed language, and the type of a variable is set when declared