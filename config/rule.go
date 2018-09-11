package config

const (
	// 名称长度
	RULE_NAME_LENGTH = 100
	// 通用名字命名规则 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	RULE_NAME_REG = `^[a-z\$_][a-z\$\d_]*$`
	// 通用字符集检测
	RULE_CHARSET = "utf8,utf8mb4"
	// 通用 COLLATE
	RULE_COLLATE = "utf8_general_ci,utf8mb4_general_ci"
	// 是否允许删除数据库
	RULE_ALLOW_DROP_DATABASE = false
	// 是否允许删除表
	RULE_ALLOW_DROP_TABLE = false
	// 是否允许 rename table
	RULE_ALLOW_RENAME_TABLE = false
	// 是否允许 truncate table
	RULE_ALLOW_TRUNCATE_TABLE = false
	// 建表允许的存储引擎, 多个以逗号隔开
	RULE_TABLE_ENGINE = "innodb"
	// 是否允许大字段: text, blob
	RULE_NOT_ALLOW_COLUMN_TYPE = "247,248,249,250,251,252"
	// 表是否需要注释
	RULE_NEED_TABLE_COMMENT = true
	// 字段是否需要注释
	RULE_NEED_COLUMN_COMMENT = true
	// 主键需要有子增
	RULE_PK_AUTO_INCREMENT = true
	// 必须有主键
	RULE_NEED_PK = true
	// 索引字段个数
	RULE_INDEX_COLUMN_COUNT = 5
	// 表名  命名规范
	RULE_TABLE_NAME_GRE = `(?i)^(?!taishan)[a-z\$_][a-z\$\d_]*$`
	// 索引命名规范
	RULE_INDEX_NAME_REG = `^idx_[a-z\$\d_]*$`
	// 唯一索引命名规范
	RULE_UNIQUE_INDEX_NAME_REG = `^udx_[a-z\$\d_]*$`
)
