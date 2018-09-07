package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
)

type RenameTableReviewer struct {
	StmtNode *ast.RenameTableStmt
}

func (this *RenameTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !RULE_ALLOW_RENAME_TABLE {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = MSG_FORBIDEN_RENAME_TABLE_ERROR
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}
