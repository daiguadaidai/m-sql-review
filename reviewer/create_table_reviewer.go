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
	/* 定义所有索引
	map {
		idx_xxx: ["id", "name"]
	}
	*/
	Indexes map[string][]string
	hasTableComment bool // 有表注释
}

// 初始化一些变量
func (this *CreateTableReviewer) Init() {
	this.ColumnNames = make(map[string]bool)
	this.PKNames = make(map[string]bool)
	this.Indexes = make(map[string][]string)
}

func (this *CreateTableReviewer) Review() *ReviewMSG {
	this.Init()
	var reviewMSG *ReviewMSG

	// 检测数据库名称长度
	reviewMSG = this.DetectDBNameLength(this.StmtNode.Table.Schema.String())
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测数据库命名规则
	reviewMSG = this.DetectDBNameReg(this.StmtNode.Table.Schema.String())
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "数据库名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测表名称长度
	reviewMSG = this.DetectTableNameLength(this.StmtNode.Table.Name.String())
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("%v %v", "表名", reviewMSG.MSG)
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测表命名规则
	reviewMSG = this.DetectTableNameReg(this.StmtNode.Table.Name.String())
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

	// 检测字段中定义了多个主键, 这中定义是在字段定义后面添加 primary key, 而不是在添加索引中定义的
	reviewMSG = this.DetectColumnPKReDefine()
	if reviewMSG != nil {
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测表的约束
	reviewMSG = this.DetectConstraints()
	if reviewMSG != nil {
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测是否有主键
	reviewMSG = this.DetectHasPK()
	if reviewMSG != nil {
		reviewMSG.Sql = this.StmtNode.Text()
		reviewMSG.Code = REVIEW_CODE_ERROR
		return reviewMSG
	}

	// 检测主键是否有使用 auto increment
	reviewMSG = this.DetectPKAutoIncrement()
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
func (this *CreateTableReviewer) DetectDBNameLength(_name string) *ReviewMSG {
	return DetectNameLength(_name, this.ReviewConfig.RuleNameLength)
}

// 检测数据库命名规范
func (this *CreateTableReviewer) DetectDBNameReg(_name string) *ReviewMSG {
	return DetectNameReg(_name, this.ReviewConfig.RuleNameReg)
}

// 检测表名长度
func (this *CreateTableReviewer) DetectTableNameLength(_name string) *ReviewMSG {
	return DetectNameLength(_name, this.ReviewConfig.RuleNameLength)
}

// 检测表命名规范
func (this *CreateTableReviewer) DetectTableNameReg(_name string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	reviewMSG = DetectNameReg(_name, this.ReviewConfig.RuleTableNameReg)
	if reviewMSG != nil {
		reviewMSG.MSG = fmt.Sprintf("检测失败. %v 表名: %v",
			fmt.Sprintf(config.MSG_TABLE_NAME_GRE_ERROR, this.ReviewConfig.RuleTableNameReg),
			_name)
	}

	return reviewMSG
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

	for _, column := range this.StmtNode.Cols {
		// 1. 检测字段名是否有重复
		if _, ok := this.ColumnNames[column.Name.String()]; ok {
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
	var hasColumnComment bool = false // 用于检测字段的注释是否指定

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
			if strings.Trim(option.Expr.GetValue().(string), " ") != "" {
				hasColumnComment = true
			}
		case ast.ColumnOptionGenerated:
		case ast.ColumnOptionReference:
		}
	}

	// 检测字段是否有注释
	if this.ReviewConfig.RuleNeedColumnComment {
		if !hasColumnComment {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("字段: %v 检测失败. %v", _column.Name.String(),
				config.MSG_NEED_COLUMN_COMMENT_ERROR)
		}
	}

	return reviewMSG
}

// 检测在定义字段中有多个 primary key出现
func (this *CreateTableReviewer) DetectColumnPKReDefine() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if len(this.PKNames) > 1 {
		columnNames := make([]string, 0, 1)
		for name, _ := range this.PKNames {
			columnNames = append(columnNames, name)
		}
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf("检测失败. 有两个字段都定义了主键(%v). " +
			"请考虑使用定于约束字句定义组合主键", strings.Join(columnNames, ", "))
	}

	return reviewMSG
}

// 检测是否有主键
func (this *CreateTableReviewer) DetectHasPK() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if this.ReviewConfig.RuleNeedPK {
		if len(this.PKNames) < 1 {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("检测失败. 表: %v. 没有主键. %v",
				this.StmtNode.Table.Name.String(), config.MSG_NEED_PK)
		}
	}

	return reviewMSG
}

// 检测主键需要自增
func (this *CreateTableReviewer) DetectPKAutoIncrement() *ReviewMSG {
	var reviewMSG *ReviewMSG

	if this.ReviewConfig.RulePKAutoIncrement {
		// 有主键才检查主键中需要有 auto_increment 选项
		if len(this.PKNames) > 0 { // 有主键字段
			var pkHasAutoIncrement bool = false // 主键中是否有 auto_increment
			if strings.Trim(this.AutoIncrementName, " ") != "" {
				if _, ok := this.PKNames[this.AutoIncrementName]; ok {
					pkHasAutoIncrement = true
				}
			}
			if !pkHasAutoIncrement {
				columnNames := make([]string, 0, 1)
				for name, _ := range this.PKNames {
					columnNames = append(columnNames, name)
				}
				reviewMSG = new(ReviewMSG)
				reviewMSG.MSG = fmt.Sprintf("检测失败. 主键字段: %v. %v",
					strings.Join(columnNames, ", "), config.MSG_PK_AUTO_INCREMENT_ERROR)
			}
		}
	}

	return reviewMSG
}

// 循环检测数据库的相关索引信息
func (this *CreateTableReviewer) DetectConstraints() *ReviewMSG {
	var reviewMSG *ReviewMSG

	for _, constraint := range this.StmtNode.Constraints {
		// 检测索引/约束名是否重复
		if _, ok := this.Indexes[constraint.Name]; ok {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("检测失败. 有索引/约束名称重复: %v", constraint.Name)
			return reviewMSG
		}
		indexColumnNameMap := make(map[string]bool)
		this.Indexes[constraint.Name] = make([]string, 0, 1)

		// 检测一个 索引/约束中不能有重复字段, 并保存 索引/约束 中
		for _, indexName := range constraint.Keys {
			// 检测 索引/约束 中有重复字段
			if _, ok := indexColumnNameMap[indexName.Column.String()]; ok {
				reviewMSG = new(ReviewMSG)
				reviewMSG.MSG = fmt.Sprintf("检测失败. 同一个 索引/约束 中有同一个重复字段. " +
					"索引/约束: %v, 重复的字段名: %v",
					constraint.Name, indexName.Column.String())
				return reviewMSG
			}
			this.Indexes[constraint.Name] = append(this.Indexes[constraint.Name], indexName.Column.String())
			indexColumnNameMap[indexName.Column.String()] = true // 保存 索引/约束中的字段名

			// 检测索引字段需要在表的字段中
			if _, ok := this.ColumnNames[indexName.Column.String()]; !ok {
				reviewMSG = new(ReviewMSG)
				reviewMSG.MSG = fmt.Sprintf("检测失败. 索引字段没定义. 索引/约束: %v, " +
					"字段: %v, 不存在表: %v 中 ",
					constraint.Name, indexName.Column.String(), this.StmtNode.Table.Name.String())
				return reviewMSG
			}
		}

		// 检测索引中字段个数是否符合 指定
		if len(indexColumnNameMap) > this.ReviewConfig.RuleIndexColumnCount {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("检测失败. 索引/约束: %v. %v", constraint.Name,
				fmt.Sprintf(config.MSG_INDEX_COLUMN_COUNT_ERROR, this.ReviewConfig.RuleIndexColumnCount))
			return reviewMSG
		}

		// 约束名称长度
		reviewMSG = DetectNameLength(constraint.Name, this.ReviewConfig.RuleNameLength)
		if reviewMSG != nil {
			reviewMSG = new(ReviewMSG)
			reviewMSG.MSG = fmt.Sprintf("检测失败. %v. 索引/约束: %v",
				fmt.Sprintf(config.MSG_NAME_LENGTH_ERROR, this.ReviewConfig.RuleNameLength),
				constraint.Name)
			return reviewMSG
		}

		switch constraint.Tp {
		case ast.ConstraintPrimaryKey:
			reviewMSG = this.DectectConstraintPrimaryKey(constraint)
			if reviewMSG != nil {
				return reviewMSG
			}
			// 赋值主键列名
			for _, pkName := range constraint.Keys {
				this.PKNames[pkName.Column.String()] = true
			}
		case ast.ConstraintKey, ast.ConstraintIndex:
			reviewMSG = this.DectectConstraintIndex(constraint)
			if reviewMSG != nil {
				return reviewMSG
			}
		case ast.ConstraintUniq, ast.ConstraintUniqKey, ast.ConstraintUniqIndex:
			reviewMSG = this.DectectConstraintUniqIndex(constraint)
			if reviewMSG != nil {
				return reviewMSG
			}
		case ast.ConstraintForeignKey:
		case ast.ConstraintFulltext:
		}
	}

	return reviewMSG
}

/* 检测主键约束相关东西
Params:
	_constraint: 约束信息
 */
func (this *CreateTableReviewer) DectectConstraintPrimaryKey(_constraint *ast.Constraint) *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 检测在字段定义字句中和约束定义字句中是否有重复定义 主键
	if len(this.PKNames) > 0 {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf("检测失败. 表: %v 主键有重复定义. ",
			this.StmtNode.Table.Name.String())
		return reviewMSG
	}

	return reviewMSG
}

/* 检测索引相关信息
	_constraint: 约束信息
 */
func (this *CreateTableReviewer) DectectConstraintIndex(_constraint *ast.Constraint) *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 检测索引命名规范
	reviewMSG = DetectNameReg(_constraint.Name, this.ReviewConfig.RuleIndexNameReg)
	if reviewMSG != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf("检测失败. %v 索引/约束: %v",
			fmt.Sprintf(config.MSG_INDEX_NAME_REG_ERROR, this.ReviewConfig.RuleIndexNameReg),
			_constraint.Name)
		return reviewMSG
	}

	return reviewMSG
}

/* 检测索引相关信息
	_constraint: 约束信息
 */
func (this *CreateTableReviewer) DectectConstraintUniqIndex(_constraint *ast.Constraint) *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 间隔唯一索引命名规范
	reviewMSG = DetectNameReg(_constraint.Name, this.ReviewConfig.RuleUniqueIndexNameReg)
	if reviewMSG != nil {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf("检测失败. %v 唯一索引: %v",
			fmt.Sprintf(config.MSG_UNIQUE_INDEX_NAME_REG_ERROR, this.ReviewConfig.RuleUniqueIndexNameReg),
			_constraint.Name)
		return reviewMSG
	}

	return reviewMSG
}
