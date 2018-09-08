package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/config"
)

type CreateDatabaseReviewer struct {
	StmtNode *ast.CreateDatabaseStmt
	ReviewConfig *config.ReviewConfig
}

func (this *CreateDatabaseReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 检测名称长度
	reviewMSG = this.DetectDBNameLength()
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测命名规则
	reviewMSG = this.DetectDBNameReg()
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测创建数据库其他选项
	reviewMSG = this.DetectDBOptions()
	if reviewMSG != nil {
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 能走到这里说明写的语句审核成功
	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"
	reviewMSG.Sql = this.StmtNode.Text()

	return reviewMSG
}

// 检测数据库名长度
func (this *CreateDatabaseReviewer) DetectDBNameLength() *ReviewMSG {
	return DetectNameLength(this.StmtNode.Name, this.ReviewConfig.RuleNameLength)
}

// 检测数据库命名规范
func (this *CreateDatabaseReviewer) DetectDBNameReg() *ReviewMSG {
	return DetectNameReg(this.StmtNode.Name, this.ReviewConfig.RuleNameReg)
}

// 检测创建数据库其他选项值
func (this *CreateDatabaseReviewer) DetectDBOptions() *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, option := range this.StmtNode.Options {
		switch option.Tp {
		case ast.DatabaseOptionCharset:
			reviewMSG = DetectCharset(option.Value, this.ReviewConfig.RuleCharSet)
		case ast.DatabaseOptionCollate:
			reviewMSG = DetectCollate(option.Value, this.ReviewConfig.RuleCollate)
		}

		// 一检测到有问题键停止下面检测, 返回检测错误值
		if reviewMSG != nil {
			break
		}
	}

	return reviewMSG
}
