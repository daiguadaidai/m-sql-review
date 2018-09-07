package service

import (
	"fmt"
	"github.com/daiguadaidai/m-sql-review/parser"
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/reviewer"
	"github.com/juju/errors"
)

/* 开始审核 SQL 语句
Params:
    _sql: 需要审核的sql语句
Return:
	int: 审核状态码
	string: 审核相关信息, 如果成功是成功信息, 如果失败是失败信息
 */
func Start(_sql string) (string, error) {
	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(_sql, "", "")
	if err != nil {
		errMSG := fmt.Sprintf("Syntax Error: %v", err)
		return "", errors.New(errMSG)
	}


	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	reviewMSGs := make([]*reviewer.ReviewMSG, 0, 1)
	for _, stmtNode := range stmtNodes {
		reviewMSG := ReviewStmt(stmtNode)
		reviewMSGs = append(reviewMSGs, reviewMSG)
	}

	// 将审核信息转化为 JSON 字符串
	return reviewer.ReviewMSGs2JSON(reviewMSGs)
}

/*
1. 解析语句的每个节点并生成, 相关方便审核的数据
2. 对不通的语句进行审核
Params:
	_stmtNode: 相关类型语句的主节点

 */
func ReviewStmt(_stmtNode ast.StmtNode) (*reviewer.ReviewMSG) {
	review := reviewer.NewReviewer(_stmtNode)

	return review.Review()
}
