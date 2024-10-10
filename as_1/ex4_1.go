package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

func ProductToJSON(p Product) (string, error) {
    jsonData, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}

func JSONToProduct(jsonStr string) (Product, error) {
    var p Product
    err := json.Unmarshal([]byte(jsonStr), &p)
    if err != nil {
        return p, err
    }
    return p, nil
}

func main() {
    product := Product{
        Name:     "Laptop",
        Price:    999.99,
        Quantity: 10,
    }

    jsonStr, err := ProductToJSON(product)
    if err != nil {
        fmt.Println("Error converting to JSON:", err)
        return
    }
    fmt.Println("Product JSON:", jsonStr)

    decodedProduct, err := JSONToProduct(jsonStr)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }
    fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}

// To work with JSON in Go, you primarily use the encoding/json package.
// Encoding (Marshaling) JSON: Convert Go data structures (like structs) into JSON format using json.Marshal.
// Decoding (Unmarshaling) JSON: Convert JSON data into Go data structures using json.Unmarshal.

// Struct tags in Go provide metadata about struct fields, which helps the encoding/json package understand how to encode/decode data.

// Errors during JSON encoding/decoding are common and should be handled gracefully. 
// Typical errors include invalid JSON format or issues with type mismatches.


