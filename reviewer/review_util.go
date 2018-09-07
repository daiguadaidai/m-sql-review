package reviewer

import (
	"regexp"
	"fmt"
)

/* 检测名称长度是否合法
Params:
	_name: 需要检测的名字
 */
func DetectNameLength(_name string) *ReviewMSG {
	if len(_name) > RULE_NAME_LENGTH {
		return &ReviewMSG{
			Code: REVIEW_CODE_ERROR,
			MSG: fmt.Sprintf(
				"名称: %v, 检测失败. %v",
				_name,
				fmt.Sprintf(MSG_NAME_LENGTH_ERROR, RULE_NAME_LENGTH),
			),
		}
	}

	return nil
}

/* 检测名字是否合法
Params:
	_name: 需要检测的名字
 */
func DetectNameReg(_name string) *ReviewMSG {
	// 正则规则: 以(字母/$/_)开头, 之后任意多个(字母/数字/_/$)
	match, err := regexp.MatchString(RULE_NAME_REG, _name)
	if err != nil || !match {
		return &ReviewMSG{
			Code: REVIEW_CODE_ERROR,
			MSG: fmt.Sprintf("名称: %v, 检测失败. %v", _name, MSG_NAME_REG_ERROR),
		}
	}

	return nil
}
