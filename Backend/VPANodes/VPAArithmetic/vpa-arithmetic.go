package VPAArithmetic

import (
	"math"
)

type Operable interface {
	Operation() interface{}
}

type NumberNode struct {
	Id  string
	In1 float64
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

type AritmethicNode struct {
	Id            string
	In1           Operable
	In2           Operable
	OperationType int
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
