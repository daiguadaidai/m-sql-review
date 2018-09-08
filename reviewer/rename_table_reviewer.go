package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
)

type RenameTableReviewer struct {
	StmtNode *ast.RenameTableStmt
	ReviewConfig *config.ReviewConfig
}

func (this *RenameTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 禁止使用 rename
	if !this.ReviewConfig.RuleAllowRenameTable {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = config.MSG_FORBIDEN_RENAME_TABLE_ERROR

		return reviewMSG
	}

	// 允许使用rename
	// 循环一个语句中需要rename的所有表, 如: rename t1 to tt2, t2 to tt2;
	for _, tableToTable := range this.StmtNode.TableToTables{
		// 检测名称长度
		reviewMSG = this.DetectToTableNameLength(tableToTable.NewTable.Name.String())
		if reviewMSG != nil {
			return reviewMSG
		}

		// 检测命名规则
		reviewMSG = this.DetectToTableNameReg(tableToTable.NewTable.Name.String())
		if reviewMSG != nil {
			return reviewMSG
		}
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}

/* 检测数据库名长度
Params:
    _name: 需要检测的名称
 */
func (this *RenameTableReviewer) DetectToTableNameLength(_name string) *ReviewMSG {
	return DetectNameLength(_name, this.ReviewConfig.RuleNameLength)
}

/* 检测数据库命名规范
Params:
    _name: 需要检测的名称
 */
func (this *RenameTableReviewer) DetectToTableNameReg(_name string) *ReviewMSG {
	return DetectNameReg(_name, this.ReviewConfig.RuleNameReg)
}
