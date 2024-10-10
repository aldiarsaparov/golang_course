package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func PrintArea(s Shape) {
    fmt.Printf("The area is: %.2f\n", s.Area())
}

func main() {
    circle := Circle{Radius: 15}
    rectangle := Rectangle{Width: 7, Height: 10}

    PrintArea(circle)
    PrintArea(rectangle)
}

// Define an Interface: Use the type keyword followed by the interface name and a method signature. 
// The methods listed in the interface must be implemented by any type that satisfies the interface.
// Implement an Interface: Define a struct (or other type) and provide methods with the same signatures as those declared in the interface. 
// The implementation must match exactly, including method names and parameter/return types.

// Interfaces in Go enable polymorphism by allowing different types to be treated uniformly through a common interface type. 
// This means you can write functions or methods that operate on an interface type, and they can handle any concrete type that implements that interface, regardless of the specific details of the type.

// Compile-time Check: You can use a type assertion to check if a type implements an interface. 
// This check is typically performed at compile time, and if you mistakenly try to assign a type that does not implement the interface, the compiler will generate an error.


