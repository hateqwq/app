package main

import (
    "fmt"
    "example.com/vehicles"
)

func main() {
    v := vehicles.Vehicle{
        Brand:  "Toyota",
        Model: "Corolla",
        Year:  2020,
    }
    fmt.Println(v.GetInfo())
}

