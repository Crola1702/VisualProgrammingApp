// IfNode represents an if control statement. It receives a condition (BooleanNode
// or BooleanOperator), an ExecuteBlock and an ElseBlock
package VPANodes

import (
	"errors"
)

type Operable interface {
	Operation() interface{}
}

// Variable represents a variable with a Name and a Value with any type
type Variable struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// AssignNode represents an asignment operation of a variable.
type AssignNode struct {
	Id           int              `json:"id"`
	VariableName string           `json:"variableName"`
	In1          Operable         `json:"in1"`
	VarManager   *VariableManager `json:"varManager"`
}

func (an *AssignNode) Operation() interface{} {
	return an.VarManager.AssignVariable(an.VariableName, an.In1.Operation())
}

// VariableNode is the node representation of a variable. It requires the name of
// an existing variable and the system VariableManager
type VariableNode struct {
	Id           int              `json:"id"`
	VariableName string           `json:"variableName"`
	VarManager   *VariableManager `json:"varManager"`
}

func (vn *VariableNode) Operation() interface{} {
	res, ok := vn.VarManager.GetVariable(vn.VariableName)
	if ok != nil {
		return ok
	}
	return res.Value
}

// VariableManager is a map of variable names and variable values
type VariableManager struct {
	SystemVariables map[string]Variable `json:"systemVariables"`
}

func (vm *VariableManager) GetVariable(name string) (Variable, error) {
	v, ok := vm.SystemVariables[name]
	if !ok {
		return Variable{}, errors.New("there is no variable with such name")
	}
	return v, nil
}
func (vm *VariableManager) CreateVariable(name string, value ...interface{}) error {
	if _, ok := vm.SystemVariables[name]; ok {
		return errors.New("variable already exists")
	}

	if len(value) > 0 {
		vm.SystemVariables[name] = Variable{name, value[0]}
	} else {
		vm.SystemVariables[name] = Variable{name, nil}
	}
	return nil
}
func (vm *VariableManager) AssignVariable(name string, value interface{}) error {

	if _, ok := vm.SystemVariables[name]; !ok {
		return errors.New("no such variable exists")
	}
	vm.SystemVariables[name] = Variable{name, value}
	return nil
}
