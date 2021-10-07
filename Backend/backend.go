package main

import (
	"Crola1702/Backend/VisualProgrammingApp/VPANodes/VPAArithmetic"
	"Crola1702/Backend/VisualProgrammingApp/VPANodes/VPABooleans"
	"fmt"
)

func main() {
	fmt.Println("Test init \n")

	nn1 := VPAArithmetic.NumberNode{
		Id:  "number",
		In1: 1.0,
	}

	one := VPAArithmetic.NumberNode{
		Id:  "one",
		In1: 1.0,
	}

	two := VPAArithmetic.NumberNode{
		Id:  "two",
		In1: 2.0,
	}

	counter := VPAArithmetic.NumberNode{
		Id:  "counter",
		In1: 0.0,
	}

	n := VPAArithmetic.NumberNode{
		Id:  "n",
		In1: 5.0,
	}

	m1 := VPAArithmetic.AritmethicNode{
		Id:            "m1",
		In1:           &nn1,
		In2:           &two,
		OperationType: VPAArithmetic.Times,
	}

	s1 := VPAArithmetic.AritmethicNode{
		Id:            "s1",
		In1:           &counter,
		In2:           &one,
		OperationType: VPAArithmetic.Plus,
	}

	condition := VPABooleans.BooleanNode{
		Id:            "condition",
		In1:           &counter,
		In2:           &n,
		OperationType: VPABooleans.Le,
	}

	fmt.Printf("nn1: %v, counter: %v, condition: %v\n", nn1.In1, counter.In1, condition.Operation().(bool))
	for condition.Operation().(bool) {
		nn1.In1 = m1.Operation().(float64)
		counter.In1 = s1.Operation().(float64)
		fmt.Printf("nn1: %v, counter: %v, condition: %v\n", nn1.In1, counter.In1, condition.Operation().(bool))
	}
	fmt.Println("---- FINAL STATE ----")
	fmt.Printf("2**%v=%v\n", n.In1, nn1.In1)

	fmt.Println("\nTest finish \n")
}
