package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
	"strings"
	"fmt"
)

type CreateDatabaseReviewer struct {
	StmtNode *ast.CreateDatabaseStmt
}

func (this *CreateDatabaseReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 检测名称长度
	reviewMSG = this.DetectDBNameLength()
	if reviewMSG != nil {
		return reviewMSG
	}

	// 检测命名规则
	reviewMSG = this.DetectDBName()
	if reviewMSG != nil {
		return reviewMSG
	}

	// 检测创建数据库其他选项
	reviewMSG = this.DetectDBOption()
	if reviewMSG != nil {
		return reviewMSG
	}

	// 能走到这里说明写的语句审核成功
	reviewMSG = new(ReviewMSG)
	reviewMSG.Code = REVIEW_CODE_SUCCESS
	reviewMSG.MSG = "审核成功"

	return reviewMSG
}

// 检测数据库名长度
func (this *CreateDatabaseReviewer) DetectDBNameLength() *ReviewMSG {
	return DetectNameLength(this.StmtNode.Name)
}

// 检测数据库命名规范
func (this *CreateDatabaseReviewer) DetectDBName() *ReviewMSG {
	return DetectNameReg(this.StmtNode.Name)
}

// 检测创建数据库其他选项值
func (this *CreateDatabaseReviewer) DetectDBOption() *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, option := range this.StmtNode.Options {
		switch option.Tp {
		case ast.DatabaseOptionCharset:
			reviewMSG = this.DetectDBCharset(option.Value)
		case ast.DatabaseOptionCollate:
			reviewMSG = this.DetectDBCollate(option.Value)
		}

		// 一检测到有问题键停止下面检测, 返回检测错误值
		if reviewMSG != nil {
			break
		}
	}

	return reviewMSG
}

/* 检测数据库的字符集
Params:
    _charset: 需要审核的字符集
 */
func (this *CreateDatabaseReviewer) DetectDBCharset(_charset string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	allowCharsets := strings.Split(RULE_CHARSET, ",") // 获取允许的字符集数组
	isMatch := false
	// 将需要检测的字符集 和 允许的字符集进行循环比较
	for _, charset := range allowCharsets {
		if _charset == charset {
			isMatch = true
			break
		}
	}

	if !isMatch {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = fmt.Sprintf(
			"字符类型检测失败: %v",
			fmt.Sprintf(MSG_CHARSET_ERROR, RULE_CHARSET),
		)
	}

	return reviewMSG
}

/* 检测数据库的Collate
Params:
    _collate: 需要审核的字符集
 */
func (this *CreateDatabaseReviewer) DetectDBCollate(_collate string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	allowCollate := strings.Split(RULE_COLLATE, ",") // 获取允许的Collate数组
	isMatch := false
	// 将需要检测的collate 和 允许的字符集进行循环比较
	for _, collate := range allowCollate {
		if _collate == collate {
			isMatch = true
			break
		}
	}

	if !isMatch {
		reviewMSG = new(ReviewMSG)
		reviewMSG.Code = REVIEW_CODE_ERROR
		reviewMSG.MSG = fmt.Sprintf(
			"Collate 类型检测失败: %v",
			fmt.Sprintf(MSG_COLLATE_ERROR, RULE_COLLATE),
		)
	}

	return reviewMSG
}
