package handle

import "github.com/daiguadaidai/m-sql-review/config"

type RequestReviewParam struct {
	ReviewConfig *config.ReviewConfig // 指定的数据库规则

	DBConfig *config.DBConfig // 链接数据库的参数

	Sqls string // 用于审核的sql

	/////////////////////////////////////////////////////
	// 以下参数主要是为了识别是否使用自定义的规则, 都是bool类型
	/////////////////////////////////////////////////////
	// 是否使用自定义, 通用名字长度
	CustomRuleNameLength bool
	// 是否使用自定义, 通用名字命名规则 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	CustomRuleNameReg bool
	// 是否使用自定义, 通用字符集检测
	CustomRuleCharSet bool
	// 是否使用自定义, 通用 COLLATE
	CustomRuleCollate bool
	// 是否使用自定义, 是否允许创建数据库
	CustomRuleAllowCreateDatabase bool
	// 是否使用自定义, 是否允许删除数据库
	CustomRuleAllowDropDatabase bool
	// 是否使用自定义, 是否允许删除表
	CustomRuleAllowDropTable bool
	// 是否使用自定义, 是否允许 rename table
	CustomRuleAllowRenameTable bool
	// 是否使用自定义, 是否允许 truncate table
	CustomRuleAllowTruncateTable bool
	// 是否使用自定义, 允许的存储引擎
	CustomRuleTableEngine bool
	// 是否使用自定义, 不允许使用的字段
	CustomRuleNotAllowColumnType bool
	// 是否使用自定义, 表是否需要注释
	CustomRuleNeedTableComment bool
	// 是否使用自定义, 字段需要有注释
	CustomRuleNeedColumnComment bool
	// 是否使用自定义, 主键自增
	CustomRulePKAutoIncrement bool
	// 是否使用自定义, 是否使用自定义, 必须要要有主键
	CustomRuleNeedPK bool
	// 是否使用自定义, 索引字段个数
	CustomRuleIndexColumnCount bool
	// 是否使用自定义, 表名 命名规范
	CustomRuleTableNameReg bool
	// 是否使用自定义, 索引命名规范
	CustomRuleIndexNameReg bool
	// 是否使用自定义, 唯一所有命名规范
	CustomRuleUniqueIndexNameReg bool
	// 是否使用自定义, 所有字段都必须为 NOT NULL
	CustomRuleAllColumnNotNull bool
	// 是否使用自定义, 是否允许使用外键
	CustomRuleAllowForeignKey bool
	// 是否使用自定义, 是否允许有全文索引
	CustomRuleAllowFullText bool
	// 是否使用自定义, 必须为NOT NULL的字段
	CustomRuleNotNullColumnType bool
	// 是否使用自定义, 必须为NOT NULL 的字段名
	CustomRuleNotNullColumnName bool
	// 是否使用自定义, text字段允许使用个数
	CustomRuleTextTypeColumnCount bool
	// 是否使用自定义, 必须有索引的字段名
	CustomRuleNeedIndexColumnName bool
	// 是否使用自定义, 必须包含的字段名
	CustomRuleHaveColumnName bool
	// 是否使用自定义, 字段定义必须要有默认值
	CustomRuleNeedDefaultValue bool
	// 是否使用自定义, 必须有默认值的字段名字
	CustomRuleNeedDefaultValueName bool
}

func (this *RequestReviewParam) ReSetReviewConfig(_reviewConfig *config.ReviewConfig) {
	// 是否使用自定义, 通用名字长度
	if this.CustomRuleNameLength {
		_reviewConfig.RuleNameLength = this.ReviewConfig.RuleNameLength
	}
	// 是否使用自定义, 通用名字命名规则 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	if this.CustomRuleNameReg {
		_reviewConfig.RuleNameReg = this.ReviewConfig.RuleNameReg
	}
	// 是否使用自定义, 通用字符集检测
	if this.CustomRuleCharSet {
		_reviewConfig.RuleCharSet = this.ReviewConfig.RuleCharSet
	}
	// 是否使用自定义, 通用 COLLATE
	if this.CustomRuleCollate {
		_reviewConfig.RuleCollate = this.ReviewConfig.RuleCollate
	}
	// 是否使用自定义, 是否允许创建数据库
	if this.CustomRuleAllowCreateDatabase {
		_reviewConfig.RuleAllowCreateDatabase = this.ReviewConfig.RuleAllowCreateDatabase
	}
	// 是否使用自定义, 是否允许删除数据库
	if this.CustomRuleAllowDropDatabase {
		_reviewConfig.RuleAllowDropDatabase = this.ReviewConfig.RuleAllowDropDatabase
	}
	// 是否使用自定义, 是否允许删除表
	if this.CustomRuleAllowDropTable {
		_reviewConfig.RuleAllowDropTable = this.ReviewConfig.RuleAllowDropTable
	}
	// 是否使用自定义, 是否允许 rename table
	if this.CustomRuleAllowRenameTable {
		_reviewConfig.RuleAllowRenameTable = this.ReviewConfig.RuleAllowRenameTable
	}
	// 是否使用自定义, 是否允许 truncate table
	if this.CustomRuleAllowTruncateTable {
		_reviewConfig.RuleAllowTruncateTable = this.ReviewConfig.RuleAllowTruncateTable
	}
	// 是否使用自定义, 允许的存储引擎
	if this.CustomRuleTableEngine {
		_reviewConfig.RuleTableEngine = this.ReviewConfig.RuleTableEngine
	}
	// 是否使用自定义, 不允许使用的字段
	if this.CustomRuleNotAllowColumnType {
		_reviewConfig.RuleNotAllowColumnType = this.ReviewConfig.RuleNotAllowColumnType
	}
	// 是否使用自定义, 表是否需要注释
	if this.CustomRuleNeedTableComment {
		_reviewConfig.RuleNeedTableComment = this.ReviewConfig.RuleNeedTableComment
	}
	// 是否使用自定义, 字段需要有注释
	if this.CustomRuleNeedColumnComment {
		_reviewConfig.RuleNeedColumnComment = this.ReviewConfig.RuleNeedColumnComment
	}
	// 是否使用自定义, 主键自增
	if this.CustomRulePKAutoIncrement {
		_reviewConfig.RulePKAutoIncrement = this.ReviewConfig.RulePKAutoIncrement
	}
	// 是否使用自定义, 是否使用自定义, 必须要要有主键
	if this.CustomRuleNeedPK {
		_reviewConfig.RuleNeedPK = this.ReviewConfig.RuleNeedPK
	}
	// 是否使用自定义, 索引字段个数
	if this.CustomRuleIndexColumnCount {
		_reviewConfig.RuleIndexColumnCount = this.ReviewConfig.RuleIndexColumnCount
	}
	// 是否使用自定义, 表名 命名规范
	if this.CustomRuleTableNameReg {
		_reviewConfig.RuleTableNameReg = this.ReviewConfig.RuleTableNameReg
	}
	// 是否使用自定义, 索引命名规范
	if this.CustomRuleIndexNameReg {
		_reviewConfig.RuleIndexNameReg = this.ReviewConfig.RuleIndexNameReg
	}
	// 是否使用自定义, 唯一所有命名规范
	if this.CustomRuleUniqueIndexNameReg {
		_reviewConfig.RuleUniqueIndexNameReg = this.ReviewConfig.RuleUniqueIndexNameReg
	}
	// 是否使用自定义, 所有字段都必须为 NOT NULL
	if this.CustomRuleAllColumnNotNull {
		_reviewConfig.RuleAllColumnNotNull = this.ReviewConfig.RuleAllColumnNotNull
	}
	// 是否使用自定义, 是否允许使用外键
	if this.CustomRuleAllowForeignKey {
		_reviewConfig.RuleAllowForeignKey = this.ReviewConfig.RuleAllowForeignKey
	}
	// 是否使用自定义, 是否允许有全文索引
	if this.CustomRuleAllowFullText {
		_reviewConfig.RuleAllowFullText = this.ReviewConfig.RuleAllowFullText
	}
	// 是否使用自定义, 必须为NOT NULL的字段
	if this.CustomRuleNotNullColumnType {
		_reviewConfig.RuleNotNullColumnType = this.ReviewConfig.RuleNotNullColumnType
	}
	// 是否使用自定义, 必须为NOT NULL 的字段名
	if this.CustomRuleNotNullColumnName {
		_reviewConfig.RuleNotNullColumnName = this.ReviewConfig.RuleNotNullColumnName
	}
	// 是否使用自定义, text字段允许使用个数
	if this.CustomRuleTextTypeColumnCount {
		_reviewConfig.RuleTextTypeColumnCount = this.ReviewConfig.RuleTextTypeColumnCount
	}
	// 是否使用自定义, 必须有索引的字段名
	if this.CustomRuleNeedIndexColumnName {
		_reviewConfig.RuleNeedIndexColumnName = this.ReviewConfig.RuleNeedIndexColumnName
	}
	// 是否使用自定义, 必须包含的字段名
	if this.CustomRuleHaveColumnName {
		_reviewConfig.RuleHaveColumnName = this.ReviewConfig.RuleHaveColumnName
	}
	// 是否使用自定义, 字段定义必须要有默认值
	if this.CustomRuleNeedDefaultValue {
		_reviewConfig.RuleNeedDefaultValue = this.ReviewConfig.RuleNeedDefaultValue
	}
	// 是否使用自定义, 必须有默认值的字段名字
	if this.CustomRuleNeedDefaultValueName {
		_reviewConfig.RuleNeedIndexColumnName = this.ReviewConfig.RuleNeedIndexColumnName
	}

}
