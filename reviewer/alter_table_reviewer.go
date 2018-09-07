package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
)

type AlterTableReviewer struct {
	StmtNode *ast.AlterTableStmt
}

func (this *AlterTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	return reviewMSG
}
