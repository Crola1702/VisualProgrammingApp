package VPAConditionalsCicles

import (
	"Crola1702/Backend/VisualProgrammingApp/VPANodes/VPABooleans"
)

type Operable interface {
	Operation() interface{}
}

type ExecuteBlock struct {
	Statements []interface{ Operation() interface{} }
}

func (b *ExecuteBlock) Operation() {
	for _, v := range b.Statements {
		v.Operation()
	}
}

type IfNode struct {
	Id        string
	Condition VPABooleans.Operable
	Block     ExecuteBlock
	ElseBlock ExecuteBlock
}

func (n *IfNode) Operation() {
	if n.Condition.Operation().(bool) {
		n.Block.Operation()
	} else {
		n.ElseBlock.Operation()
	}
}

type WhileNode struct {
	Id        string
	Condition VPABooleans.Operable
	Block     Operable
}

func (n *WhileNode) Operation() {
	for n.Condition.Operation().(bool) {
		n.Block.Operation()
	}
}