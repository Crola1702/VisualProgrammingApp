package main

import (
	"Backend/VPANodes/VPAArithmetic"
	"Backend/VPANodes/VPABooleans"
	"Backend/VPANodes/VPAConditionalsCicles"
	"Backend/VPANodes/VPAVariables"
	//"./VPANodes/VPAArithmetic"
	//"./VPANodes/VPABooleans"
	//"./VPANodes/VPAConditionalsCicles"
	//"./VPANodes/VPAVariables"
	"fmt"
)

var Variables = VPAVariables.VariableManager{SystemVariables: make(map[string]VPAVariables.Variable)}

var Number1 = "number1"
var Counter = "counter"
var Result = "result"

func main() {
	fmt.Println("Test init \n")

	Variables.CreateVariable(Number1, 1.0)
	Variables.CreateVariable(Counter, 0.0)
	Variables.CreateVariable(Result, 0.0)

	VarN1 := VPAVariables.VariableNode{
		VariableName: Number1,
		VarManager:   &Variables,
	}

	VarCounter := VPAVariables.VariableNode{
		VariableName: Counter,
		VarManager:   &Variables,
	}

	VarResult := VPAVariables.VariableNode{
		VariableName: Result,
		VarManager:   &Variables,
	}

	one := VPAArithmetic.NumberNode{
		Id:  "one",
		In1: 1.0,
	}

	two := VPAArithmetic.NumberNode{
		Id:  "two",
		In1: 2.0,
	}

	m1 := VPAArithmetic.AritmethicNode{
		Id:            "m1",
		In1:           &VarN1,
		In2:           &two,
		OperationType: VPAArithmetic.Times,
	}

	s1 := VPAArithmetic.AritmethicNode{
		Id:            "s1",
		In1:           &VarCounter,
		In2:           &one,
		OperationType: VPAArithmetic.Plus,
	}

	assignN1 := VPAVariables.AssignNode{
		Id:           "assignN1",
		VariableName: Number1,
		In1:          &m1,
		VarManager:   &Variables,
	}

	assignCounter := VPAVariables.AssignNode{
		Id:           "assignCounter",
		VariableName: Counter,
		In1:          &s1,
		VarManager:   &Variables,
	}

	assignResult := VPAVariables.AssignNode{
		Id:           "assignResult",
		VariableName: Result,
		In1:          &VarN1,
		VarManager:   &Variables,
	}

	n := VPAArithmetic.NumberNode{
		Id:  "n",
		In1: 5.0,
	}

	condition := VPABooleans.BooleanNode{
		Id:            "condition",
		In1:           &VarCounter,
		In2:           &n,
		OperationType: VPABooleans.Leq,
	}

	///*
	whileBlock := VPAConditionalsCicles.ExecuteBlock{
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
	fmt.Println("---- FINAL STATE ----")
	fmt.Printf("2**%v=%v\n", n.In1, VarResult.Operation())
	fmt.Println("\nTest finish \n")
}
