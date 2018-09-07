package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
)

type DropTableReviewer struct {
	StmtNode *ast.DropTableStmt
}

func (this *DropTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_ERROR
	reviewMSG.MSG = MSG_FORBIDEN_DROP_TABLE_ERROR

	return reviewMSG
}
