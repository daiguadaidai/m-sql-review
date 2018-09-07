package reviewer

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
)
