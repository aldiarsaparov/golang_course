package main

import "fmt"

type Employee struct {
    Name string
    ID   int
}

type Manager struct {
    Employee    
    Department string
}

func (e Employee) Work() {
    fmt.Printf("Name: %s, ID: %d.\n", e.Name, e.ID)
}

func main() {
    manager := Manager{
        Employee: Employee{Name: "Aldiar", ID: 1},
        Department: "HR",
    }

    manager.Work()
}


// Embedding in Go is a way to include one struct inside another, allowing the outer struct to inherit the fields and methods of the embedded struct. 
// This allows a form of composition, where complex types are built by combining simpler types.
// Composition: Instead of using inheritance like in traditional object-oriented languages, 
// Go relies on composition through embedding. By embedding structs, Go promotes code reuse and flexibility, 
// allowing structs to behave as a combination of various types.

// Go promotes the fields and methods of an embedded type to the outer struct, 
// meaning that you can call the embedded type’s methods directly on the outer struct. 
// The embedded type’s methods are available on instances of the outer struct without needing to specify the embedded struct explicitly.

// No, an embedded type cannot override a method in the outer struct. 
// If both the outer struct and the embedded struct have methods with the same name, the method on the outer struct takes precedence when called on an instance of the outer struct.


