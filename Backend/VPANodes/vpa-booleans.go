// Package VPANodes provides a set of blocks to create any type of
// program in this Visual Programming app
package VPANodes

const (
	_ int = iota
	Eq
	Dist
	Geq
	Leq
	Le
	Ge
)

// BooleanNode represents a boolean type. It compares its inputs
// by a boolean operator like 'and', 'or', 'equals'
type BooleanNode struct {
	Id            int      `json:"id"`
	In1           Operable `json:"in1"`
	In2           Operable `json:"in2"`
	OperationType int      `json:"operationType"`
	Not           bool     `json:"not"`
}

func (b *BooleanNode) Operation() interface{} {
	var Ret bool = false
	switch b.OperationType {
	case Eq:
		Ret = b.In1.Operation().(float64) == b.In2.Operation().(float64)
	case Dist:
		Ret = b.In1.Operation().(float64) != b.In2.Operation().(float64)
	case Geq:
		Ret = b.In1.Operation().(float64) >= b.In2.Operation().(float64)
	case Leq:
		Ret = b.In1.Operation().(float64) <= b.In2.Operation().(float64)
	case Le:
		Ret = b.In1.Operation().(float64) < b.In2.Operation().(float64)
	case Ge:
		Ret = b.In1.Operation().(float64) > b.In2.Operation().(float64)
	default:
		Ret = b.In1.Operation().(float64) == b.In2.Operation().(float64)
	}

	if b.Not {
		return !Ret
	} else {
		return Ret
	}
}

const (
	_ int = iota
	And
	Or
	Xor
)

// BooleanOperator represents an 'and', 'or', 'xor' operators between
// two BooleanNodes.
type BooleanOperator struct {
	Id            int         `json:"id"`
	In1           BooleanNode `json:"in1"`
	In2           BooleanNode `json:"in2"`
	OperationType int         `json:"operationType"`
	Not           bool        `json:"not"`
}

func (b *BooleanOperator) Operation() interface{} {
	var Ret bool = false
	switch b.OperationType {
	case And:
		Ret = b.In1.Operation().(bool) && b.In2.Operation().(bool)
	case Or:
		Ret = b.In1.Operation().(bool) || b.In2.Operation().(bool)
	case Xor:
		Ret = b.In1.Operation().(bool) != b.In2.Operation().(bool)
	default:
		Ret = b.In1.Operation().(bool) && b.In2.Operation().(bool)
	}

	if b.Not {
		Ret = !Ret
	}
	return Ret
}
