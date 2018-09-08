package reviewer

import (
	"testing"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/parser"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/dependency/mysql"
)

func TestCreateTableReviewer_Review(t *testing.T) {
	sql := `
CREATE TABLE test.mf_fd_cache (
  id bigint(18) NOT NULL AUTO_INCREMENT COMMENT '主键',
  dep varchar(3) NOT NULL DEFAULT '',
  arr varchar(3) NOT NULL DEFAULT '',
  flightNo varchar(10) NOT NULL DEFAULT '',
  flightDate date NOT NULL DEFAULT '1000-10-10',
  flightTime varchar(20) NOT NULL DEFAULT '',
  isCodeShare tinyint(1) NOT NULL DEFAULT '0',
  tax int(11) NOT NULL DEFAULT '0',
  yq int(11) NOT NULL DEFAULT '0',
  cabin char(2) NOT NULL DEFAULT '',
  ibe_price int(11) NOT NULL DEFAULT '0',
  ctrip_price int(11) NOT NULL DEFAULT '0',
  official_price int(11) NOT NULL DEFAULT '0',
  uptime datetime NOT NULL DEFAULT '1000-10-10 10:10:10',
  PRIMARY KEY (id),
  UNIQUE KEY uid (dep, arr, flightNo, flightDate,cabin),
  KEY uptime (uptime),
  KEY flight (dep,arr),
  KEY flightDate (flightDate)
) ENGINE=InnoDb  DEFAULT CHARSET=utF8 COLLATE=Utf8mb4_general_ci comment="你号";
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}

	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	reviewConfig := config.NewReviewConfig()
	reviewMSGs := make([]*ReviewMSG, 0, 1)
	for _, stmtNode := range stmtNodes {
		createTableStmt := stmtNode.(*ast.CreateTableStmt)
		fmt.Printf("Schema: %v, Table: %v\n", createTableStmt.Table.Schema.String(),
			createTableStmt.Table.Name.String())
		// 建表 option
		for _, option := range createTableStmt.Options {
			fmt.Printf("type: %v, value: %v, int: %v\n", option.Tp, option.StrValue, option.UintValue)
		}

		for i, constraint := range createTableStmt.Constraints {
			fmt.Println(i, constraint)
		}

		for i, column := range createTableStmt.Cols {
			fmt.Println(i, column.Name.String(), column.Tp.Tp, column.Tp.Tp == mysql.TypeBlob)
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
