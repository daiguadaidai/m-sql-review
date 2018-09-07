package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
)

type DropTableReviewer struct {
	StmtNode *ast.DropTableStmt
}

func (this *DropTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !RULE_ALLOW_DROP_TABLE {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = MSG_FORBIDEN_DROP_TABLE_ERROR
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}
