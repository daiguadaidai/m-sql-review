package reviewer

import (
	"regexp"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/config"
	"strings"
	"strconv"
)

/* 检测名称长度是否合法
Params:
	_name: 需要检测的名字
 */
func DetectNameLength(_name string, _length int) *ReviewMSG {
	if len(_name) > _length {
		return &ReviewMSG{
			MSG: fmt.Sprintf(
				"检测失败: %v. 名称: %v",
				fmt.Sprintf(config.MSG_NAME_LENGTH_ERROR, _length),
				_name,
			),
		}
	}

	return nil
}

/* 检测名字是否合法
Params:
	_name: 需要检测的名字
 */
func DetectNameReg(_name string, _reg string) *ReviewMSG {
	// 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	match, err := regexp.MatchString(_reg, _name)
	if err != nil || !match {
		return &ReviewMSG{
			MSG: fmt.Sprintf("检测失败. %v. 名称: %v, ", config.MSG_NAME_REG_ERROR, _name),
		}
	}

	return nil
}

/* 检测数据库的字符集
Params:
    _charset: 需要审核的字符集
    _allowCharsetStr: 允许的字符集 字符串 "utf8,gbk,utf8mb4"
 */
func DetectCharset(_charset string, _allowCharsetStr string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	allowCharsets := strings.Split(_allowCharsetStr, ",") // 获取允许的字符集数组
	isMatch := false
	// 将需要检测的字符集 和 允许的字符集进行循环比较
	for _, allowCharset := range allowCharsets {
		if strings.ToLower(_charset) == allowCharset {
			isMatch = true
			break
		}
	}

	if !isMatch {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf(
			"字符类型检测失败: %v",
			fmt.Sprintf(config.MSG_CHARSET_ERROR, _allowCharsetStr),
		)
	}

	return reviewMSG
}

/* 检测数据库的Collate
Params:
    _collate: 需要审核的字符集
    _allowCollateStr: 允许的 collate 字符串 "utf8_general_ci,utf8mb4_general_ci"
 */
func DetectCollate(_collate string, _allowCollateStr string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	allowCollates := strings.Split(_allowCollateStr, ",") // 获取允许的Collate数组
	isMatch := false
	// 将需要检测的collate 和 允许的字符集进行循环比较
	for _, allowCollate := range allowCollates {
		if strings.ToLower(_collate) == allowCollate {
			isMatch = true
			break
		}
	}

	if !isMatch {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf(
			"Collate 类型检测失败: %v",
			fmt.Sprintf(config.MSG_COLLATE_ERROR, _allowCollateStr),
		)
	}

	return reviewMSG
}

/* 检测数据库允许的存储引擎
Params:
    _engine: 需要审核的存储引擎
    _allowEngineStr: 允许的存储引擎
 */
func DetectEngine(_engine string, _allowEngineStr string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	allowEngines := strings.Split(_allowEngineStr, ",") // 获取允许的存储引擎
	isMatch := false
	// 将需要检测的collate 和 允许的字符集进行循环比较
	for _, allowEngine := range allowEngines {
		if strings.ToLower(_engine) == allowEngine {
			isMatch = true
			break
		}
	}

	if !isMatch {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf(
			"存储引擎 类型检测失败: %v",
			fmt.Sprintf(config.MSG_TABLE_ENGINE_ERROR, _allowEngineStr),
		)
	}

	return reviewMSG
}

/* 检测不允许的字段类型
Params:
    _type: 需要审核的字段类型
    _notAllowTypeStr: 不允许的字段类型
 */
func DetectNotAllowColumnType(_type byte, _notAllowTypeSrt string) *ReviewMSG {
	var reviewMSG *ReviewMSG

	notAllowTypes := strings.Split(_notAllowTypeSrt, ",") // 获取不允许的字段类型
	isMatch := false
	fmt.Println(strconv.Itoa(int(_type)), notAllowTypes)
	// 将需要检测的collate 和 允许的字符集进行循环比较
	for _, notAllowType := range notAllowTypes {
		if strconv.Itoa(int(_type)) == notAllowType {
			isMatch = true
			break
		}
	}

	if isMatch {
		reviewMSG = new(ReviewMSG)
		reviewMSG.MSG = fmt.Sprintf(
			"字段检测失败: %v",
			fmt.Sprintf(config.MSG_NOT_ALLOW_COLUMN_TYPE_ERROR, _notAllowTypeSrt),
		)
	}

	return reviewMSG
}
