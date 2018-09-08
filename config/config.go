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

	return reviewConfig
}