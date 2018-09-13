package reviewer

import (
	"fmt"
	"github.com/daiguadaidai/m-sql-review/config"
	"strings"
	"strconv"
	"github.com/dlclark/regexp2"
	"crypto/md5"
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
	// 使用正则表达式匹配名称
	re := regexp2.MustCompile(_reg, 0)
	if isMatch, _ := re.MatchString(_name); !isMatch {
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

/* 通过所有所以和所有的唯一索引获取所有的普通索引
Params:
    _indexes: 所有的索引
	_uniqueIndex: 所有的普通索引
 */
func GetNoUniqueIndexes(_indexes map[string][]string, _uniqueIndex map[string][]string) map[string][]string{
	normalIndexes := make(map[string][]string)

	for indexName, index := range _indexes {
		if _, ok := _uniqueIndex[indexName]; ok { // 过滤掉索引中的 唯一索引
			continue
		}

		normalIndex := make([]string, 0, 1)
		for _, columnName := range index {
			normalIndex = append(normalIndex, columnName)
		}

		normalIndexes[indexName] = normalIndex
	}

	return normalIndexes
}

/* 将索引的字段转化成 hash过后的值
Params:
    _indexes: 需要转化的索引
 */
func GetIndexesHashColumn(_indexes map[string][]string) map[string]string {
	hashIndexes := make(map[string]string)

	for indexName, index := range _indexes {
		hashIndex := make([]string, 0, 1)
		for _, columnName := range index {
			data := []byte(columnName)
			has := md5.Sum(data)
			hashColumn := fmt.Sprintf("%x", has)
			hashIndex = append(hashIndex, hashColumn)
		}

		hashIndexes[indexName] = strings.Join(hashIndex, ",")
	}

	return hashIndexes
}
