package reviewer

import (
	"regexp"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/config"
)

/* 检测名称长度是否合法
Params:
	_name: 需要检测的名字
 */
func DetectNameLength(_name string, _length int) *ReviewMSG {
	if len(_name) > _length {
		return &ReviewMSG{
			Code: REVIEW_CODE_ERROR,
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
			Code: REVIEW_CODE_ERROR,
			MSG: fmt.Sprintf("检测失败. %v. 名称: %v, ", config.MSG_NAME_REG_ERROR, _name),
		}
	}

	return nil
}
