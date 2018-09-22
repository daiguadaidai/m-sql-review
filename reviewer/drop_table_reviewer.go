package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/dao"
	"fmt"
)

type DropTableReviewer struct {
	StmtNode *ast.DropTableStmt
	ReviewConfig *config.ReviewConfig
	DBConfig *config.DBConfig
}

func (this *DropTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !this.ReviewConfig.RuleAllowDropTable {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = config.MSG_FORBIDEN_DROP_TABLE_ERROR

		return reviewMSG
	}

	// 链接实例检测表相关信息(所有)
	reviewMSG = this.DetectInstanceTables()
	if reviewMSG != nil {
		return reviewMSG
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}

// 链接指定实例检测相关表信息(所有)
func (this *DropTableReviewer) DetectInstanceTables() *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, tableStmt := range this.StmtNode.Tables {
		reviewMSG = this.DetectInstanceTable(tableStmt.Name.String())
		if reviewMSG != nil {
			return reviewMSG
		}
	}

	return reviewMSG
}

/* 链接指定实例检测相关表信息
Params:
    _tableName: 原表名
 */
func (this *DropTableReviewer) DetectInstanceTable(_tableName string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	tableInfo := dao.NewTableInfo(this.DBConfig, _tableName)
	err := tableInfo.OpenInstance()
	if err != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_WARNING
		reviewMSG.MSG = fmt.Sprintf("警告: 无法链接到指定实例. 无法删除表[%v]. %v",
			_tableName, err)
		return reviewMSG
	}

	// 检测表是否存在
	reviewMSG = DetectTableNotExists(tableInfo)
	if reviewMSG != nil {
		return reviewMSG
	}

	err = tableInfo.CloseInstance()
	if err != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_WARNING
		reviewMSG.MSG = fmt.Sprintf("警告: 链接实例检测表相关信息. 关闭连接出错. %v", err)
		return reviewMSG
	}
	return reviewMSG
}
