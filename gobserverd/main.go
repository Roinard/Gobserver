package main

import (
	"fmt"
	"gobserver/gobserverd/src"
)

func main() {
	h := src.NewHost("10.11.1.108")
	fmt.Println("Scanning host: " + h.IP)
}
