package VPABooleans

type Operable interface {
	Operation() interface{}
}

const (
	_ int = iota
	Eq
	Dist
	Geq
	Leq
	Le
	Ge
)

type BooleanNode struct {
	Id            string
	In1           Operable
	In2           Operable
	OperationType int
	Not           bool
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

type BooleanOperator struct {
	Id            string
	In1           BooleanNode
	In2           BooleanNode
	OperationType int
	Not           bool
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
