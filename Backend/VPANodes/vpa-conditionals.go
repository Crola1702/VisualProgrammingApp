// Package VPANodes provides a set of blocks to create any type of
// program in this Visual Programming app
package VPANodes

// ExecuteBlocks represents a list of statements (Nodes) that should be
// executed sequentially
type ExecuteBlock struct {
	Statements []interface{ Operation() interface{} }
}

func (b *ExecuteBlock) Operation() {
	for _, v := range b.Statements {
		v.Operation()
	}
}

// IfNode represents an if control statement. It receives a condition (BooleanNode
// or BooleanOperator), an ExecuteBlock and an ElseBlock
type IfNode struct {
	Id        int           `json:"id"`
	Condition Operable      `json:"condition"`
	Block     *ExecuteBlock `json:"block"`
	ElseBlock *ExecuteBlock `json:"elseBlock"`
}

func (n *IfNode) Operation() {
	if n.Condition.Operation().(bool) {
		n.Block.Operation()
	} else {
		n.ElseBlock.Operation()
	}
}

// IfNode represents a while control statement. It receives a condition (BooleanNode
// or BooleanOperator) and an ExecuteBlock
type WhileNode struct {
	Id        int           `json:"id"`
	Condition Operable      `json:"condition"`
	Block     *ExecuteBlock `json:"block"`
}

func (n *WhileNode) Operation() {
	for n.Condition.Operation().(bool) {
		n.Block.Operation()
	}
}
