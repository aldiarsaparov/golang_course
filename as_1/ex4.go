package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func swap(str1 string, str2 string) (string, string) {
    return str2, str1
}

func quotientRemainder(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    sum := add(5, 3)
    fmt.Printf("Sum: %d\n", sum)

    str1, str2 := swap("swap", "me")
    fmt.Printf("Swapped: %s, %s\n", str1, str2)

    quotient, remainder := quotientRemainder(10, 3)
    fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}


// In Go, a function can return multiple values by specifying the return types in parentheses. 
// func swap(a string, b string) (string, string) {
//     return b, a
// }

// Named return values in Go allow you to predefine the names of the variables that will store the return values.
// This simplifies the code inside the function, as you can assign values directly to these names and omit the explicit return statement with arguments.

// In Go, you can ignore return values using the underscore (_). This is useful when a function returns multiple values, but you only care about some of them. Example:
// quotient, _ := divide(10, 3)

