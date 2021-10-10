package main

import (
	"Backend/VPANodes"
	//"./VPANodes"
	"fmt"
)

var Variables = VPANodes.VariableManager{SystemVariables: make(map[string]VPANodes.Variable)}

var Number1 = "number1"
var Counter = "counter"
var Result = "result"

func main() {
	fmt.Println("Test init \n")

	Variables.CreateVariable(Number1, 1.0)
	Variables.CreateVariable(Counter, 0.0)
	Variables.CreateVariable(Result, 0.0)

	VarN1 := VPANodes.VariableNode{
		VariableName: Number1,
		VarManager:   &Variables,
	}

	VarCounter := VPANodes.VariableNode{
		VariableName: Counter,
		VarManager:   &Variables,
	}

	VarResult := VPANodes.VariableNode{
		VariableName: Result,
		VarManager:   &Variables,
	}

	one := VPANodes.NumberNode{
		Id:  "one",
		In1: 1.0,
	}

	two := VPANodes.NumberNode{
		Id:  "two",
		In1: 2.0,
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
		VarManager:   &Variables,
	}

	assignCounter := VPANodes.AssignNode{
		Id:           "assignCounter",
		VariableName: Counter,
		In1:          &s1,
		VarManager:   &Variables,
	}

	assignResult := VPANodes.AssignNode{
		Id:           "assignResult",
		VariableName: Result,
		In1:          &VarN1,
		VarManager:   &Variables,
	}

	n := VPANodes.NumberNode{
		Id:  "n",
		In1: 5.0,
	}

	condition := VPANodes.BooleanNode{
		Id:            "condition",
		In1:           &VarCounter,
		In2:           &n,
		OperationType: VPANodes.Leq,
	}

	///*
	whileBlock := VPANodes.ExecuteBlock{
		Statements: []interface{ Operation() interface{} }{&assignResult, &assignN1, &assignCounter},
	}

	number1, _ := Variables.GetVariable("number1")
	counter, _ := Variables.GetVariable("counter")
	result, _ := Variables.GetVariable("result")
	fmt.Printf("number1: %v, counter: %v, result: %v, condition: %v\n", number1.Value, counter.Value, result.Value, condition.Operation().(bool))
	for condition.Operation().(bool) {

		for _, v := range whileBlock.Statements {
			v.Operation()
		}

		number1, _ := Variables.GetVariable("number1")
		counter, _ := Variables.GetVariable("counter")
		result, _ := Variables.GetVariable("result")
		fmt.Printf("number1: %v, counter: %v, result: %v, condition: %v\n", number1.Value, counter.Value, result.Value, condition.Operation().(bool))
	}
	//*/
	fmt.Println("\n---- FINAL STATE WITH NEW IMPORTS ----")
	fmt.Printf("2**%v=%v\n", n.In1, VarResult.Operation())
	fmt.Println("\nTest finish \n")
}
