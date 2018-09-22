package config

import (
	"strings"
	"github.com/daiguadaidai/m-sql-review/common"
)

var reviewConfig *ReviewConfig

type ReviewConfig struct {
	// 通用名字长度
	RuleNameLength int
	// 通用名字命名规则 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	RuleNameReg string
	// 通用字符集检测
	RuleCharSet string
	// 通用 COLLATE
	RuleCollate string
	// 是否允许创建数据库
	RuleAllowCreateDatabase bool
	// 是否允许删除数据库
	RuleAllowDropDatabase bool
	// 是否允许删除表
	RuleAllowDropTable bool
	// 是否允许 rename table
	RuleAllowRenameTable bool
	// 是否允许 truncate table
	RuleAllowTruncateTable bool
	// 允许的存储引擎
	RuleTableEngine string
	// 不允许使用的字段
	RuleNotAllowColumnType string
	// 表是否需要注释
	RuleNeedTableComment bool
	// 字段需要有注释
	RuleNeedColumnComment bool
	// 主键自增
	RulePKAutoIncrement bool
	// 必须要要有主键
	RuleNeedPK bool
	// 索引字段个数
	RuleIndexColumnCount int
	// 表名 命名规范
	RuleTableNameReg string
	// 索引命名规范
	RuleIndexNameReg string
	// 唯一所有命名规范
	RuleUniqueIndexNameReg string
	// 所有字段都必须为 NOT NULL
	RuleAllColumnNotNull bool
	// 是否允许使用外键
	RuleAllowForeignKey bool
	// 是否允许有全文索引
	RuleAllowFullText bool
	// 必须为NOT NULL的字段
	RuleNotNullColumnType string
	// 必须为NOT NULL 的字段名
	RuleNotNullColumnName string
	// text字段允许使用个数
	RuleTextTypeColumnCount int
	// 必须有索引的字段名
	RuleNeedIndexColumnName string
	// 必须包含的字段名
	RuleHaveColumnName string
	// 字段定义必须要有默认值
	RuleNeedDefaultValue bool
	// 必须有默认值的字段名字
	RuleNeedDefaultValueName string
}

func NewReviewConfig() *ReviewConfig {
	rc := new(ReviewConfig)

	rc.RuleNameLength = RULE_NAME_LENGTH
	rc.RuleNameReg = RULE_NAME_REG
	rc.RuleCharSet = RULE_CHARSET
	rc.RuleCollate = RULE_COLLATE
	rc.RuleAllowCreateDatabase = RULE_ALLOW_CREATE_DATABASE
	rc.RuleAllowDropDatabase = RULE_ALLOW_DROP_DATABASE
	rc.RuleAllowDropTable = RULE_ALLOW_DROP_TABLE
	rc.RuleAllowRenameTable = RULE_ALLOW_RENAME_TABLE
	rc.RuleAllowTruncateTable = RULE_ALLOW_TRUNCATE_TABLE
	rc.RuleTableEngine = RULE_TABLE_ENGINE
	rc.RuleNotAllowColumnType = RULE_NOT_ALLOW_COLUMN_TYPE
	rc.RuleNeedTableComment = RULE_NEED_TABLE_COMMENT
	rc.RuleNeedColumnComment = RULE_NEED_COLUMN_COMMENT
	rc.RulePKAutoIncrement = RULE_PK_AUTO_INCREMENT
	rc.RuleNeedPK = RULE_NEED_PK
	rc.RuleIndexColumnCount = RULE_INDEX_COLUMN_COUNT
	rc.RuleTableNameReg = RULE_TABLE_NAME_GRE
	rc.RuleIndexNameReg = RULE_INDEX_NAME_REG
	rc.RuleUniqueIndexNameReg = RULE_UNIQUE_INDEX_NAME_REG
	rc.RuleAllColumnNotNull = RULE_ALL_COLUMN_NOT_NULL
	rc.RuleAllowForeignKey = RULE_ALLOW_FOREIGN_KEY
	rc.RuleAllowFullText = RULE_ALLOW_FULL_TEXT
	rc.RuleNotNullColumnType = RULE_NOT_NULL_COLUMN_TYPE
	rc.RuleNotNullColumnName = RULE_NOT_NULL_COLUMN_NAME
	rc.RuleTextTypeColumnCount = RULE_TEXT_TYPE_COLUMN_COUNT
	rc.RuleNeedIndexColumnName = RULE_NEED_INDEX_COLUMN_NAME
	rc.RuleHaveColumnName = RULE_HAVE_COLUMN_NAME
	rc.RuleNeedDefaultValue = RULE_NEED_DEFAULT_VALUE
	rc.RuleNeedDefaultValueName = RULE_NEED_DEFAULT_VALUE_NAME

	return rc
}

/* 设置全局的 reviewconfig
Params:
	_reviewConfig: sql审核配置
 */
func SetReviewConfig(_reviewConfig *ReviewConfig) {
	reviewConfig = _reviewConfig
}

// 获取一个Copy的reviewConfig
func GetReviewConfig() *ReviewConfig {
	rc := new(ReviewConfig)
	common.StructCopy(reviewConfig, rc)

	return rc
}

// 获取不允许的字段类型映射
func (this *ReviewConfig) GetNotAllowColumnTypeMap() map[string]bool {
	notAllowColumnTypeMap := make(map[string]bool)

	notAllowColumnTypes := strings.Split(this.RuleNotAllowColumnType, ",")
	for _, notAllowColumnType := range notAllowColumnTypes {
		notAllowColumnType = strings.ToLower(strings.TrimSpace(notAllowColumnType))
		if notAllowColumnType == "" {
			continue
		}
		//  text 相关类型 要多保存为 blob 类型
		switch notAllowColumnType {
		case "tinytext":
			notAllowColumnTypeMap[TYPE_STR_TINYBLOB] = true
		case "text":
			notAllowColumnTypeMap[TYPE_STR_BLOB] = true
		case "mediumtext":
			notAllowColumnTypeMap[TYPE_STR_MEDIUMBLOB] = true
		case "longtext":
			notAllowColumnTypeMap[TYPE_STR_LONG_BLOB] = true
		}

		notAllowColumnTypeMap[notAllowColumnType] = true
	}

	return notAllowColumnTypeMap
}

// 对必须为not null 的字段类型通过 (逗号分割). 保存到map中
func (this *ReviewConfig) GetNotNullColumnTypeMap() map[string]bool {
	notNullColumnTypeMap := make(map[string]bool)

	notNullColumnTypes := strings.Split(this.RuleNotNullColumnType, ",")
	for _, notNullColumnType := range notNullColumnTypes {
		notNullColumnType = strings.ToLower(strings.TrimSpace(notNullColumnType))
		if notNullColumnType == "" {
			continue
		}
		//  text 相关类型 要多保存为 blob 类型
		switch notNullColumnType {
		case "tinytext":
			notNullColumnTypeMap[TYPE_STR_TINYBLOB] = true
		case "text":
			notNullColumnTypeMap[TYPE_STR_BLOB] = true
		case "mediumtext":
			notNullColumnTypeMap[TYPE_STR_MEDIUMBLOB] = true
		case "longtext":
			notNullColumnTypeMap[TYPE_STR_LONG_BLOB] = true
		}

		notNullColumnTypeMap[notNullColumnType] = true
	}

	return notNullColumnTypeMap
}

// 将必须为not null的字段名规则进行(逗号)分割, 保存到map中
func (this *ReviewConfig) GetNotNullColumnNameMap() map[string]bool {
	return common.SplitString2Map(this.RuleNotNullColumnName, ",")
}

// 获取必须要有索引的字段
func (this *ReviewConfig) GetNeedIndexColumnNameMap() map[string]bool {
	return common.SplitString2Map(this.RuleNeedIndexColumnName, ",")
}

// 获取必须要有的字段名
func (this *ReviewConfig) GetHaveColumnNameMap() map[string]bool {
	return common.SplitString2Map(this.RuleHaveColumnName, ",")
}

// 获取必须要有默认值的字段
func (this *ReviewConfig) GetNeedDefaultValueNameMap() map[string]bool {
	return common.SplitString2Map(this.RuleNeedDefaultValueName, ",")
}

