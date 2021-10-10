package main

import (
	"Backend/VPANodes"
	"fmt"
)

var variables = VPANodes.VariableManager{SystemVariables: make(map[string]VPANodes.Variable)}

var Number1 = "number1"
var Counter = "counter"
var Result = "result"

func main() {
	fmt.Println("Test init \n")

	variables.CreateVariable(Number1, 1.0)
	variables.CreateVariable(Counter, 0.0)
	variables.CreateVariable(Result, 0.0)

	VarN1 := VPANodes.VariableNode{
		VariableName: Number1,
		VarManager:   &variables,
	}

	VarCounter := VPANodes.VariableNode{
		VariableName: Counter,
		VarManager:   &variables,
	}

	VarResult := VPANodes.VariableNode{
		VariableName: Result,
		VarManager:   &variables,
	}

	one := VPANodes.NumberNode{
		Id:  "one",
		In1: 1.0,
	}

	two := VPANodes.NumberNode{
		Id:  "two",
		In1: 2.0,
	}

	n := VPANodes.NumberNode{
		Id:  "n",
		In1: 16,
	}

	m1 := VPANodes.AritmethicNode{
		Id:            "m1",
		In1:           &VarN1,
		In2:           &two,
		OperationType: VPANodes.Times,
	}

	s1 := VPANodes.AritmethicNode{
		Id:            "s1",
		In1:           &VarCounter,
		In2:           &one,
		OperationType: VPANodes.Plus,
	}

	assignN1 := VPANodes.AssignNode{
		Id:           "assignN1",
		VariableName: Number1,
		In1:          &m1,
		VarManager:   &variables,
	}

	assignCounter := VPANodes.AssignNode{
		Id:           "assignCounter",
		VariableName: Counter,
		In1:          &s1,
		VarManager:   &variables,
	}

	assignResult := VPANodes.AssignNode{
		Id:           "assignResult",
		VariableName: Result,
		In1:          &VarN1,
		VarManager:   &variables,
	}

	condition := VPANodes.BooleanNode{
		Id:            "condition",
		In1:           &VarCounter,
		In2:           &n,
		OperationType: VPANodes.Leq,
	}

	whileBlock := VPANodes.ExecuteBlock{
		Statements: []interface{ Operation() interface{} }{&assignResult, &assignN1, &assignCounter},
	}

	while := VPANodes.WhileNode{
		Id:        "while",
		Condition: &condition,
		Block:     &whileBlock,
	}

	while.Operation()

	fmt.Println("\n---- FINAL STATE ----")
	fmt.Printf("2**%v=%v\n", n.In1, VarResult.Operation())
	fmt.Println("\nTest finish \n")
}
