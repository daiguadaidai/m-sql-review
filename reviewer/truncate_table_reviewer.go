package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/dao"
	"fmt"
)

type TruncateTableReviewer struct {
	StmtNode *ast.TruncateTableStmt
	ReviewConfig *config.ReviewConfig
	DBConfig *config.DBConfig
}

func (this *TruncateTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !this.ReviewConfig.RuleAllowTruncateTable {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = config.MSG_FORBIDEN_TRUNCATE_TABLE_ERROR

		return reviewMSG
	}

	// 链接数据库检测实例相关信息
	reviewMSG = this.DetectInstanceTable()
	if reviewMSG != nil {
		return reviewMSG
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}

// 链接到实例检测相关信息
func (this *TruncateTableReviewer) DetectInstanceTable() *ReviewMSG {
	var reviewMSG *ReviewMSG

	tableInfo := dao.NewTableInfo(this.DBConfig, this.StmtNode.Table.Name.String())
	err := tableInfo.OpenInstance()
	if err != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_WARNING
		reviewMSG.MSG = fmt.Sprintf("警告: 无法链接到指定实例. 无法检测数据库是否存在.")
		return reviewMSG
	}

	// 检测表是否不存在
	reviewMSG = DetectTableNotExists(tableInfo)
	if reviewMSG != nil {
		return reviewMSG
	}

	err = tableInfo.CloseInstance()
	if err != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_WARNING
		reviewMSG.MSG = fmt.Sprintf("警告: 链接实例检测数据库相关信息. 关闭连接出错.")
		return reviewMSG
	}
	return reviewMSG
}
