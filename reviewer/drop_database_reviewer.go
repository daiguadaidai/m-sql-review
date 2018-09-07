package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
)

type DropDatabaseReviewer struct {
	StmtNode *ast.DropDatabaseStmt
}

func (this *DropDatabaseReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_ERROR
	reviewMSG.MSG = MSG_FORBIDEN_DROP_DATABASE_ERROR

	return reviewMSG
}
