package config

const (
	// ------------- 通用规则 --------------
	// 通用名称字符长度
	MSG_NAME_LENGTH_ERROR = "名称长度不能超过 %v"
	// 通用名称的规范
	MSG_NAME_REG_ERROR = "命名规则: 以 (字母/_/$) 开头, 之后可用字符 (字母/数字/_/$)"
	// 通用字符集
	MSG_CHARSET_ERROR = "使用的字符集只允许 %v"
	// 通用COLLATE
	MSG_COLLATE_ERROR = "使用的collate只允许 %v"
	// 禁用DROP DATABASE 操作
	MSG_FORBIDEN_DROP_DATABASE_ERROR = "禁止删除数据库"
	// 禁止DROP TABLE操作
	MSG_FORBIDEN_DROP_TABLE_ERROR = "禁止删除表"
	// 禁止 truncate 表
	MSG_FORBIDEN_TRUNCATE_TABLE_ERROR = "禁止 truncate 表操作"
	// 禁止 rename 表
	MSG_FORBIDEN_RENAME_TABLE_ERROR = "禁止 rename 表操作"
	// 允许的存储引擎
	MSG_TABLE_ENGINE_ERROR = "允许的存储引擎 %v"
	// 重复定义字段名
	MSG_TABLE_COLUMN_DUP_ERROR = "一个字段名字不能被定义多次"
	// 不允许的字段
	MSG_NOT_ALLOW_COLUMN_TYPE_ERROR = "这些字段不能使用 %v"
	// 表需要注释
	MSG_NEED_TABLE_COMMENT_ERROR = "新建表必须要有注释"
	// 字段需要注释
	MSG_NEED_COLUMN_COMMENT_ERROR = "字段必须要有注释"
	// 主键需要自增
	MSG_PK_AUTO_INCREMENT_ERROR = "主键字段中必须有一个字段有自增属性"
	// 必须有主键
	MSG_NEED_PK = "必须要有主键"
	// 索引字段个数
	MSG_INDEX_COLUMN_COUNT_ERROR = "索引字段不能超过 %v 个"
	// 表名命名规则
	MSG_TABLE_NAME_GRE_ERROR = "表名, 命名规则不符合规范. 规范为(正则): %v"
	// 所有名称命名规范
	MSG_INDEX_NAME_REG_ERROR = "索引名称命名不符合规范, 索引命名规范为(正则): %v"
	// 唯一索引命名规范
	MSG_UNIQUE_INDEX_NAME_REG_ERROR = `唯一索引名称命名不符合规范, 索引命名规范为(正则): %v`
)