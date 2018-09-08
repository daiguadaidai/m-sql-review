package config


type ReviewConfig struct {
	// 通用名字长度
	RuleNameLength int
	// 通用名字命名规则 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	RuleNameReg string
	// 通用字符集检测
	RuleCharSet string
	// 通用 COLLATE
	RuleCollate string
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
}

func NewReviewConfig() *ReviewConfig {
	reviewConfig := new(ReviewConfig)

	reviewConfig.RuleNameLength = RULE_NAME_LENGTH
	reviewConfig.RuleNameReg = RULE_NAME_REG
	reviewConfig.RuleCharSet = RULE_CHARSET
	reviewConfig.RuleCollate = RULE_COLLATE
	reviewConfig.RuleAllowDropDatabase = RULE_ALLOW_DROP_DATABASE
	reviewConfig.RuleAllowDropTable = RULE_ALLOW_DROP_TABLE
	reviewConfig.RuleAllowRenameTable = RULE_ALLOW_RENAME_TABLE
	reviewConfig.RuleAllowTruncateTable = RULE_ALLOW_TRUNCATE_TABLE
	reviewConfig.RuleTableEngine = RULE_TABLE_ENGINE
	reviewConfig.RuleNotAllowColumnType = RULE_NOT_ALLOW_COLUMN_TYPE
	reviewConfig.RuleNeedTableComment = RULE_NEED_TABLE_COMMENT

	return reviewConfig
}