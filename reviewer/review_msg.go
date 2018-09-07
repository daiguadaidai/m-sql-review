package reviewer

// 定义返回代码
const (
	REVIEW_CODE_SUCCESS = iota
	REVIEW_CODE_WARNING
	REVIEW_CODE_ERROR
)

type ReviewMSG struct {
	Sql string
	Code int
	MSG string
}

/* 将每一条审核消息数组转化称 Json 字符串

 */
func ReviewMSGs2JSON(_reviewMSGs []*ReviewMSG) (string, error) {

	return "", nil
}

