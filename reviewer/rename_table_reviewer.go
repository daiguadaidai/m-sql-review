package reviewer

import (
"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/dao"
)

type RenameTableReviewer struct {
	StmtNode *ast.RenameTableStmt
	ReviewConfig *config.ReviewConfig
	DBConfig *config.DBConfig
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

		// 检测数据库名称长度
		reviewMSG = this.DetectDBNameLength(tableToTable.NewTable.Schema.String())
		if reviewMSG != nil {
			reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
			reviewMSG.Code = REVIEW_CODE_ERROR
			return reviewMSG
		}

		// 检测数据库命名规则
		if tableToTable.NewTable.Schema.String() != "" {
			reviewMSG = this.DetectDBNameReg(tableToTable.NewTable.Schema.String())
			if reviewMSG != nil {
				reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
				reviewMSG.Code = REVIEW_CODE_ERROR
				return reviewMSG
			}
		}

		// 检测表名称长度
		reviewMSG = this.DetectToTableNameLength(tableToTable.NewTable.Name.String())
		if reviewMSG != nil {
			reviewMSG.MSG = fmt.Sprintf("%v %v", "表名", reviewMSG.MSG)
			reviewMSG.Code = REVIEW_CODE_ERROR
			return reviewMSG
		}

		// 检测表命名规则
		reviewMSG = this.DetectToTableNameReg(tableToTable.NewTable.Name.String())
		if reviewMSG != nil {
			reviewMSG.MSG = fmt.Sprintf("%v %v", "表名", reviewMSG.MSG)
			reviewMSG.Code = REVIEW_CODE_ERROR
			return reviewMSG
		}
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

/* 检测数据库名长度
Params:
_name: 需要检测的名称
*/
func (this *RenameTableReviewer) DetectDBNameLength(_name string) *ReviewMSG {
	return DetectNameLength(_name, this.ReviewConfig.RuleNameLength)
}

/* 检测数据库命名规范
Params:
_name: 需要检测的名称
*/
func (this *RenameTableReviewer) DetectDBNameReg(_name string) *ReviewMSG {
	return DetectNameReg(_name, this.ReviewConfig.RuleNameReg)
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
	var reviewMSG *ReviewMSG

	reviewMSG = DetectNameReg(_name, this.ReviewConfig.RuleTableNameReg)
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("检测失败. %v 表名: %v",
			fmt.Sprintf(config.MSG_TABLE_NAME_GRE_ERROR, this.ReviewConfig.RuleTableNameReg),
			_name)
	}

	return reviewMSG
}

// 链接指定实例检测相关表信息(所有)
func (this *RenameTableReviewer) DetectInstanceTables() *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, tableStmt := range this.StmtNode.TableToTables {
		reviewMSG = this.DetectInstanceTable(tableStmt.OldTable.Name.String(),
			tableStmt.NewTable.Name.String())
		if reviewMSG != nil {
			return reviewMSG
		}
	}

	return reviewMSG
}

/* 链接指定实例检测相关表信息
Params:
    _tableName: 原表名
    _toTablename: 目标表名
 */
func (this *RenameTableReviewer) DetectInstanceTable(_tableName string, _toTableName string) *ReviewMSG {
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


	// 检测表是否不存在
	reviewMSG = DetectTableNotExistsByName(tableInfo, _tableName)
	if reviewMSG != nil {
		return reviewMSG
	}

	// 检测目标表是否存在
	reviewMSG = DetectTableExistsByName(tableInfo, _toTableName)
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
