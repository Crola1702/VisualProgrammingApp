package main

import (
	"Backend/VPANodes"
	//"./VPANodes"
	"fmt"
)

var variables = VPANodes.VariableManager{SystemVariables: make(map[string]VPANodes.Variable)}

func main() {
	fmt.Println("Test init \n")

	fmt.Println("\nTest finish \n")
}
