package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/dao"
	"fmt"
)

type DropDatabaseReviewer struct {
	StmtNode *ast.DropDatabaseStmt
	ReviewConfig *config.ReviewConfig
	DBConfig *config.DBConfig
}

func (this *DropDatabaseReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if !this.ReviewConfig.RuleAllowDropDatabase {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = config.MSG_FORBIDEN_DROP_DATABASE_ERROR

		return reviewMSG
	}

	// 链接数据库检测实例相关信息
	reviewMSG = this.DetectInstanceDatabase()
	if reviewMSG != nil {
		return reviewMSG
	}

	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}

// 链接到实例检测相关信息
func (this *DropDatabaseReviewer) DetectInstanceDatabase() *ReviewMSG {
	var reviewMSG *ReviewMSG

	tableInfo := dao.NewTableInfo(this.DBConfig, "")
	tableInfo.DBName = this.StmtNode.Name
	err := tableInfo.OpenInstance()
	if err != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_WARNING
		reviewMSG.MSG = fmt.Sprintf("警告: 无法链接到指定实例. 无法检测数据库是否存在.")
		return reviewMSG
	}

	reviewMSG = DetectDatabaseNotExists(tableInfo)
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
