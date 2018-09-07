package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
)

type TruncateTableReviewer struct {
	StmtNode *ast.TruncateTableStmt
}

func (this *TruncateTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !RULE_ALLOW_TRUNCATE_TABLE {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = MSG_FORBIDEN_TRUNCATE_TABLE_ERROR
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}
