package VPAVariables

import (
	"errors"
)

type Operable interface {
	Operation() interface{}
}

type Variable struct {
	Name  string
	Value interface{}
}

type AssignNode struct {
	Id           string
	VariableName string
	In1          Operable
	VarManager   *VariableManager
}
func (an *AssignNode) Operation() interface{} {
	return an.VarManager.AssignVariable(an.VariableName, an.In1.Operation())
}

type VariableNode struct {
	VariableName string
	VarManager *VariableManager
}
func (vn *VariableNode) Operation() interface{} {
	res, ok := vn.VarManager.GetVariable(vn.VariableName)
	if ok != nil {
		return ok
	}
	return res.Value
}

type VariableManager struct {
	SystemVariables map[string]Variable
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
