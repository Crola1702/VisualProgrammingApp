// Package VPANodes provides a set of blocks to create any type of
// program in this Visual Programming app
package VPANodes

import (
	"math"
)

// NumberNode represents a node that contains an input
// related to a float number. The operation returns
// the asociated number with it.
type NumberNode struct {
	Id  int     `json:"id"`
	In1 float64 `json:"in1"`
}

func (n *NumberNode) Operation() interface{} {
	return n.In1
}

const (
	_ int = iota
	Plus
	Minus
	Times
	Divided
	Power
	Mod
)

// AritmethicNode represents an aritmetic operation
// like '+' '-' '*' '/' '^' and '%' between two other
// nodes.
type AritmethicNode struct {
	Id            int      `json:"id"`
	In1           Operable `json:"in1"`
	In2           Operable `json:"in2"`
	OperationType int      `json:"operationType"`
}

func (n *AritmethicNode) Operation() interface{} {
	switch n.OperationType {
	case Plus:
		return n.In1.Operation().(float64) + n.In2.Operation().(float64)
	case Minus:
		return n.In1.Operation().(float64) - n.In2.Operation().(float64)
	case Times:
		return n.In1.Operation().(float64) * n.In2.Operation().(float64)
	case Divided:
		return n.In1.Operation().(float64) / n.In2.Operation().(float64)
	case Power:
		return math.Pow(n.In1.Operation().(float64), n.In2.Operation().(float64))
	case Mod:
		return float64(int(n.In1.Operation().(float64)) % int(n.In2.Operation().(float64)))
	default:
		return n.In1.Operation().(float64) + n.In2.Operation().(float64)
	}
}
