package service

import (
	"fmt"
	"github.com/daiguadaidai/m-sql-review/parser"
	"github.com/daiguadaidai/m-sql-review/reviewer"
	"github.com/juju/errors"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/common"
)

func Run(_config *config.ReviewConfig) error {
	sql := `
        DROP TABLE test1.t1
    `

	_, err := StartReview(sql, _config)
	return err
}

/* 开始审核 SQL 语句
Params:
    _sql: 需要审核的sql语句
Return:
	int: 审核状态码
	string: 审核相关信息, 如果成功是成功信息, 如果失败是失败信息
 */
func StartReview(_sql string, _config *config.ReviewConfig) (string, error) {
	// 拷贝一份 需要审核额配置, 这个拷贝的配置用于每次审核, 如果有需要修改的审核的项目可以修改副本中的值
	reviewConfig := new(config.ReviewConfig)
	common.StructCopy(reviewConfig, _config)

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(_sql, "", "")
	if err != nil {
		errMSG := fmt.Sprintf("Syntax Error: %v", err)
		return "", errors.New(errMSG)
	}


	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	reviewMSGs := make([]*reviewer.ReviewMSG, 0, 1)
	for _, stmtNode := range stmtNodes {
		review := reviewer.NewReviewer(stmtNode, reviewConfig)
		reviewMSG := review.Review()
		if reviewMSG != nil {
			reviewMSG.Sql = stmtNode.Text()
		}
		reviewMSGs = append(reviewMSGs, reviewMSG)
	}

	for _, reviewMSG := range reviewMSGs {
		fmt.Printf("Code: %v, MSG: %v \n", reviewMSG.Code, reviewMSG.MSG)
	}

	// 将审核信息转化为 JSON 字符串
	return reviewer.ReviewMSGs2JSON(reviewMSGs)
}
