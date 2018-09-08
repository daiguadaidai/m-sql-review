package reviewer

import (
	"testing"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/parser"
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
)

func TestRenameTableReviewer_Review(t *testing.T) {
	sql := `
		rename table test.1table to test1.t1, t2 to tt2;
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}


	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	reviewMSGs := make([]*ReviewMSG, 0, 1)
	reviewConfig := config.NewReviewConfig()
	for _, stmtNode := range stmtNodes {
		renameStmt := stmtNode.(*ast.RenameTableStmt)
		for i, subStmt := range renameStmt.TableToTables {
			fmt.Printf(
				"%v: %v -> %v\n",
				i,
				subStmt.OldTable.Name.String(),
				subStmt.NewTable.Name.String(),
			)
		}

		review := NewReviewer(stmtNode, reviewConfig)
		reviewMSG := review.Review()
		reviewMSGs = append(reviewMSGs, reviewMSG)
	}

	for _, reviewMSG := range reviewMSGs {
		if reviewMSG != nil {
			fmt.Printf("Code: %v, MSG: %v", reviewMSG.Code, reviewMSG.MSG)
		}
	}
}
