package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/config"
	"strings"
)

type CreateTableReviewer struct {
	StmtNode *ast.CreateTableStmt
	ReviewConfig *config.ReviewConfig
	ColumnNames map[string]bool
	PKNames map[string]bool // 所有主键名
	AutoIncrementName string  // 子增字段名
	IndexColumnNames [][]string // 保存了所有索引所有的字段名
	hasTableComment bool // 有表注释
}

func (this *CreateTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 检测数据库名称长度
	reviewMSG = this.DetectDBNameLength()
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测数据库命名规则
	reviewMSG = this.DetectDBNameReg()
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测表名称长度
	reviewMSG = this.DetectTableNameLength()
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "表名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测表命名规则
	reviewMSG = this.DetectTableNameReg()
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "表名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测建表选项
	reviewMSG = this.DetectTableOptions()
	if reviewMSG != nil {
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测表字段信息
	reviewMSG = this.DetectColumns()
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
func (this *CreateTableReviewer) DetectDBNameLength() *ReviewMSG {
	return DetectNameLength(this.StmtNode.Table.Schema.String(), this.ReviewConfig.RuleNameLength)
}

// 检测数据库命名规范
func (this *CreateTableReviewer) DetectDBNameReg() *ReviewMSG {
	return DetectNameReg(this.StmtNode.Table.Schema.String(), this.ReviewConfig.RuleNameReg)
}

// 检测表名长度
func (this *CreateTableReviewer) DetectTableNameLength() *ReviewMSG {
	return DetectNameLength(this.StmtNode.Table.Name.String(), this.ReviewConfig.RuleNameLength)
}

// 检测表命名规范
func (this *CreateTableReviewer) DetectTableNameReg() *ReviewMSG {
	return DetectNameReg(this.StmtNode.Table.Name.String(), this.ReviewConfig.RuleNameReg)
}

// 检测创建数据库其他选项值
func (this *CreateTableReviewer) DetectTableOptions() *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, option := range this.StmtNode.Options {
		switch option.Tp {
		case ast.TableOptionEngine:
			reviewMSG = DetectEngine(option.StrValue, this.ReviewConfig.RuleTableEngine)
		case ast.TableOptionCharset:
			reviewMSG = DetectCharset(option.StrValue, this.ReviewConfig.RuleCharSet)
		case ast.TableOptionCollate:
			reviewMSG = DetectCollate(option.StrValue, this.ReviewConfig.RuleCollate)
		case ast.TableOptionComment:
			// 有设置表注释, 并且不是空字符串, 才代表有设置注释
			if strings.Trim(option.StrValue, " ") != "" {
				this.hasTableComment = true
			}
		}
		// 一检测到有问题键停止下面检测, 返回检测错误值
		if reviewMSG != nil {
			break
		}
	}

	// 检测表是否有注释
	if this.ReviewConfig.RuleNeedTableComment {
		if !this.hasTableComment {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("表: %v 检测失败. %v", this.StmtNode.Table.Name.String(),
				config.MSG_NEED_TABLE_COMMENT_ERROR)
		}
	}

	return reviewMSG
}

// 循环检测表的字段
func (this *CreateTableReviewer) DetectColumns() *ReviewMSG {
	var reviewMSG *ReviewMSG

	this.ColumnNames = make(map[string]bool) // 保存所有字段名

	for _, column := range this.StmtNode.Cols {
		// 1. 检测字段名是否有重复
		_, ok := this.ColumnNames[column.Name.String()]
		if ok {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("字段: %v. %v",
				column.Name.String(), config.MSG_TABLE_COLUMN_DUP_ERROR)

			return reviewMSG
		}
		this.ColumnNames[column.Name.String()] = true // 缓存字段名

		// 2. 检测字段名长度
		reviewMSG = DetectNameLength(column.Name.String(), this.ReviewConfig.RuleNameLength)
		if reviewMSG != nil {
			reviewMSG.MSG = fmt.Sprintf("字段名 %v", reviewMSG.MSG)
			return reviewMSG
		}

		// 3. 检测字段名规则
		reviewMSG = DetectNameReg(column.Name.String(), this.ReviewConfig.RuleNameReg)
		if reviewMSG != nil {
			reviewMSG.MSG = fmt.Sprintf("字段名 %v", reviewMSG.MSG)
			return reviewMSG
		}

		// 4.检测不允许的字段类型
		reviewMSG = DetectNotAllowColumnType(column.Tp.Tp, this.ReviewConfig.RuleNotAllowColumnType)
		if reviewMSG != nil {
			reviewMSG.MSG = fmt.Sprintf("字段: %v, 类型: %v 不被允许. %v",
				column.Name.String(), column.Tp.String(), reviewMSG.MSG)
			return reviewMSG
		}

		// 5. 字段定义选项
		reviewMSG = this.DetectColumnOptions(column)
		if reviewMSG != nil {
			return reviewMSG
		}
	}

	return reviewMSG
}

func (this *CreateTableReviewer) DetectColumnOptions(_column *ast.ColumnDef) *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, option := range _column.Options {
		switch option.Tp {
		case ast.ColumnOptionPrimaryKey:
			this.PKNames[_column.Name.String()] = true
		case ast.ColumnOptionNotNull:
		case ast.ColumnOptionAutoIncrement:
			this.AutoIncrementName = _column.Name.String()
		case ast.ColumnOptionDefaultValue:
		case ast.ColumnOptionUniqKey:
		case ast.ColumnOptionNull:
		case ast.ColumnOptionOnUpdate:
		case ast.ColumnOptionFulltext:
		case ast.ColumnOptionComment:
		case ast.ColumnOptionGenerated:
		case ast.ColumnOptionReference:
		}
	}

	return reviewMSG
}

// 循环检测数据库的相关索引信息
func (this *CreateTableReviewer) DetectConstraints() *ReviewMSG {
	var reviewMSG *ReviewMSG

	return reviewMSG
}


