package plan

import (
	"github.com/daiguadaidai/m-sql-review/ast"
)

// Trace represents a trace plan.
type Trace struct {
	baseSchemaProducer

	StmtNode ast.StmtNode
}
