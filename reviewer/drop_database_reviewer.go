package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
)

type DropDatabaseReviewer struct {
	StmtNode *ast.DropDatabaseStmt
}

func (this *DropDatabaseReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !RULE_ALLOW_DROP_DATABASE {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = MSG_FORBIDEN_DROP_DATABASE_ERROR
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}
