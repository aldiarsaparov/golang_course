package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Aldiar", Age: 21}
    person.Greet()
}

// A struct in Go is defined using the type keyword followed by the struct name and the fields it contains. Each field has a name and a type.

// Methods in Go are similar to regular functions, but they are associated with a specific receiver type, 
// typically a struct. The receiver is written before the method name and allows the method to access and modify the fields of the struct.

// Yes, in Go, methods can be associated with any user-defined type, not just structs. This includes types like int, float, arrays, slices, and even pointers.