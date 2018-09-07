package reviewer

import (
	"testing"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/parser"
	"github.com/daiguadaidai/m-sql-review/ast"
)

func TestRenameTableReviewer_Review(t *testing.T) {
	sql := `
		rename table test.table to test1.t1, t2 to tt2;
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}


	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	reviewMSGs := make([]*ReviewMSG, 0, 1)
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

		review := NewReviewer(stmtNode)
		reviewMSG := review.Review()
		reviewMSGs = append(reviewMSGs, reviewMSG)
	}

	for _, reviewMSG := range reviewMSGs {
		if reviewMSG != nil {
			fmt.Printf("Code: %v, MSG: %v", reviewMSG.Code, reviewMSG.MSG)
		}
	}
}
