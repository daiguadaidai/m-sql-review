package common

import (
	"github.com/dlclark/regexp2"
)

/* 字符串匹配
Params:
    _str: 需要匹配的字符串
    _reg: 正则表达式
 */
func StrIsMatch(_str string, _reg string) bool {
	var matched bool = false

	re := regexp2.MustCompile(_reg, 0)
	if isMatch, _ := re.MatchString(_str); isMatch {
		matched = true
	}

	return matched
}
