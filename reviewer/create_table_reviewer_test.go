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
  dep varchar(3) NOT NULL DEFAULT '' Comment '注释',
  arr varchar(3) NOT NULL DEFAULT '' Comment '注释',
  flightNo varchar(10) NOT NULL DEFAULT '' Comment '注释',
  flightDate date NOT NULL DEFAULT '1000-10-10' Comment '注释',
  flightTime varchar(20) NOT NULL DEFAULT '' Comment '注释',
  isCodeShare tinyint(1) Comment '注释',
  tax int(11) NOT NULL DEFAULT '0' Comment '注释',
  yq int(11) NOT NULL DEFAULT '0' Comment '注释',
  cabin char(2) NOT NULL DEFAULT '' Comment '注释',
  ibe_price int(11) NOT NULL DEFAULT '0' Comment '注释',
  ctrip_price int(11) NOT NULL DEFAULT '0' Comment '注释',
  official_price int(11) NOT NULL DEFAULT '0' Comment '注释',
  uptime datetime NOT NULL DEFAULT '1000-10-10 10:10:10' Comment '注释',
  PRIMARY KEY (id),
  UNIQUE KEY udx_uid (dep, arr, flightNo, flightDate, cabin),
  Index idx_uptime (uptime),
  KEY idx_flight (dep,arr),
  KEY idx_flightdate (flightDate),
  FULLTEXT KEY full_asdfaf (flightTime)
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
			switch constraint.Tp {
			case ast.ConstraintPrimaryKey:
				fmt.Println(i, "ConstraintPrimaryKey")
			case ast.ConstraintKey:
				fmt.Println(i, "ConstraintKey")
			case ast.ConstraintIndex:
				fmt.Println(i, "ConstraintIndex")
			case ast.ConstraintUniq:
				fmt.Println(i, "ConstraintUniq")
			case ast.ConstraintUniqKey:
				fmt.Println(i, "ConstraintUniqKey")
			case ast.ConstraintUniqIndex:
				fmt.Println(i, "ConstraintUniqIndex")
			case ast.ConstraintForeignKey:
				fmt.Println(i, "ConstraintForeignKey")
			case ast.ConstraintFulltext:
				fmt.Println(i, "ConstraintFulltext")
			default:
				fmt.Println(i, "Default")
			}

		}

		// 字段option
		for i, column := range createTableStmt.Cols {
			fmt.Println(i, column.Name.String(), column.Tp.Tp, column.Tp.Tp == mysql.TypeBlob)
			optionTypes := make([]string, 0, 1)
			for _, option := range column.Options {
				switch option.Tp {
				case ast.ColumnOptionPrimaryKey:
					optionTypes = append(optionTypes, "PrimaryKey")
				case ast.ColumnOptionNotNull:
					optionTypes = append(optionTypes, "NotNull")
				case ast.ColumnOptionAutoIncrement:
					optionTypes = append(optionTypes, "AutoIncrement")
				case ast.ColumnOptionDefaultValue:
					optionTypes = append(optionTypes, fmt.Sprintf("DefaultValue:%v", option.Expr.GetValue()))
				case ast.ColumnOptionUniqKey:
					optionTypes = append(optionTypes, "UniqKey")
				case ast.ColumnOptionNull:
					optionTypes = append(optionTypes, "NULL")
				case ast.ColumnOptionOnUpdate:
					optionTypes = append(optionTypes, "OnUpdate")
				case ast.ColumnOptionFulltext:
					optionTypes = append(optionTypes, "Fulltext")
				case ast.ColumnOptionComment:
					optionTypes = append(optionTypes, fmt.Sprintf("Comment:%v", option.Expr.GetValue()))
				case ast.ColumnOptionGenerated:
					optionTypes = append(optionTypes, "Generated")
				case ast.ColumnOptionReference:
					optionTypes = append(optionTypes, "Reference")
				}
			}
			fmt.Println("column Name:", column.Name.String(), optionTypes)
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
