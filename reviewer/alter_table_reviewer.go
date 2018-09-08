package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
)

type AlterTableReviewer struct {
	StmtNode *ast.AlterTableStmt
	ReviewConfig *config.ReviewConfig
}

func (this *AlterTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	return reviewMSG
}
