package reviewer

import (
	"testing"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/parser"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/dependency/mysql"
	"github.com/daiguadaidai/m-sql-review/dao"
)

func TestCreateTableReviewer_Review(t *testing.T) {
	var host string = "10.10.10.12"
	var port int = 3306
	var username string = "HH"
	var password string = "oracle"
	var database string = "test"
	sql := `
CREATE TABLE test.t1 (
  id bigint(18) NOT NULL AUTO_INCREMENT COMMENT '主键',
  dep varchar(3) NOT NULL DEFAULT '' Comment '注释',
  arr varchar(3) NOT NULL DEFAULT '' Comment '注释',
  flightNo varchar(10) NOT NULL DEFAULT '' Comment '注释',
  flightDate date NOT NULL DEFAULT '1000-10-10' Comment '注释',
  flightTime varchar(20) NOT NULL DEFAULT '' Comment '注释',
  isCodeShare tinyint(1) Comment '注释',
  tax int(11) NOT NULL DEFAULT '0' Comment '注释',
  yq int(11) NOT NULL DEFAULT '0' Comment '注释',
  cabin char(2) NOT NULL default '' Comment '注释',
  ibe_price int(11) NOT NULL DEFAULT '0' Comment '注释',
  ctrip_price int(11) NOT NULL DEFAULT '0' Comment '注释',
  official_price int(11) NOT NULL DEFAULT '0' Comment '注释',
  uptime datetime NOT NULL DEFAULT '1000-10-10 10:10:10' Comment '注释',
  PRIMARY KEY (id),
  UNIQUE KEY udx_uid (dep, arr, flightNo, flightDate, cabin),
  Index idx_uptime (uptime),
  KEY idx_flight (dep,arr),
  KEY idx_flightdate (flightDate)
) ENGINE=InnoDb  DEFAULT CHARSET=utF8 COLLATE=Utf8mb4_general_ci comment="你号";
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}

	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	dbConfig := config.NewDBConfig(host, port, username ,password, database)
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

		review := NewReviewer(stmtNode, reviewConfig, dbConfig)
		reviewMSG := review.Review()
		reviewMSGs = append(reviewMSGs, reviewMSG)
	}

	for _, reviewMSG := range reviewMSGs {
		if reviewMSG != nil {
			fmt.Printf("Code: %v, MSG: %v", reviewMSG.Code, reviewMSG.MSG)
		}
	}
}


func TestCreateTableReviewer_Review_Partition_range(t *testing.T) {
	sql := `
CREATE TABLE test.mf_fd_cache (
  id bigint(18) NOT NULL AUTO_INCREMENT COMMENT '注释',
  uptime datetime NOT NULL DEFAULT '1000-10-10 10:10:10' Comment '注释',
  PRIMARY KEY (id)
) ENGINE=InnoDb  DEFAULT CHARSET=Utf8mb4 COLLATE=Utf8mb4_general_ci comment="注释"
 PARTITION BY RANGE(uptime)(
     PARTITION p0 VALUES LESS THAN (TO_DAYS('2010-04-15')),
     PARTITION p19 VALUES LESS ThAN  MAXVALUE
 );
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}

	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	for _, stmtNode := range stmtNodes {

		createTableStmt := stmtNode.(*ast.CreateTableStmt)
		fmt.Printf("Schema: %v, Table: %v\n", createTableStmt.Table.Schema.String(),
			createTableStmt.Table.Name.String())
		fmt.Printf("%T, %[1]v \n", createTableStmt.Partition.Expr)
		columnNameExpr := createTableStmt.Partition.Expr.(*ast.ColumnNameExpr)
		fmt.Println("Partition column Name:", columnNameExpr.Name.String())

		fmt.Println("Partition type:", createTableStmt.Partition.Tp.String())
		fmt.Println("Partition cols:", createTableStmt.Partition.ColumnNames)
		// 字段
		for _, columnName := range createTableStmt.Partition.ColumnNames {
			fmt.Printf("%v, ", columnName)
		}

		// partition 定义
		for _, definition := range createTableStmt.Partition.Definitions {
			fmt.Printf("partition definition： %v, %v, %v \n",
				definition.Name, definition.Comment, definition.MaxValue)
		}

		fmt.Println("")
	}
}

func TestCreateTableReviewer_Review_Partition_Range_Columns(t *testing.T) {
	sql := `
CREATE TABLE partition_test ( 
t_id int(11) NOT NULL AUTO_INCREMENT, 
test_date datetime NOT NULL, 
t_key varchar(16), 
test_info varchar(50) DEFAULT 'test', 
PRIMARY KEY (t_id,test_date,t_key) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8 
PARTITION BY RANGE COLUMNS (test_date,t_key) 
( 
PARTITION p201303151 VALUES LESS THAN ('2013-03-15','m2'), 
PARTITION p201303152 VALUES LESS THAN ('2013-03-15','m3'), 
PARTITION p201303161 VALUES LESS THAN ('2013-03-16','m2'), 
PARTITION p201303162 VALUES LESS THAN ('2013-03-16','m3'), 
PARTITION p201303171 VALUES LESS THAN ('2013-03-17','m2'), 
PARTITION p201303172 VALUES LESS THAN ('2013-03-17','m3') 
);
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}

	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	for _, stmtNode := range stmtNodes {
		createTableStmt := stmtNode.(*ast.CreateTableStmt)
		fmt.Printf("Schema: %v, Table: %v\n", createTableStmt.Table.Schema.String(),
			createTableStmt.Table.Name.String())

		if createTableStmt.Partition != nil {
			fmt.Println("是分区表")
			for _, columnName := range createTableStmt.Partition.ColumnNames {
				fmt.Printf("%v, ", columnName)
			}
			fmt.Println("")

			// partition 定义
			for _, definition := range createTableStmt.Partition.Definitions {
				fmt.Printf("partition definition： %v, %v, %v \n",
				definition.Name, definition.Comment, definition.MaxValue)
			}
		} else {
			fmt.Println("不是分区表")
		}
	}
}

func TestCreateTableReviewer_Review_Partition_Range_func(t *testing.T) {
	sql := `
CREATE TABLE partition_test ( 
t_id int(11) NOT NULL AUTO_INCREMENT, 
test_date datetime NOT NULL, 
t_key varchar(16), 
test_info varchar(50) DEFAULT 'test', 
PRIMARY KEY (t_id,test_date,t_key) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8 
PARTITION BY RANGE (to_day(test_date, t_id)) 
( 
PARTITION p201303151 VALUES LESS THAN ('2013-03-15'), 
PARTITION p201303152 VALUES LESS THAN ('2013-03-15'), 
PARTITION p201303161 VALUES LESS THAN ('2013-03-16'), 
PARTITION p201303162 VALUES LESS THAN ('2013-03-16'), 
PARTITION p201303171 VALUES LESS THAN ('2013-03-17'), 
PARTITION p201303172 VALUES LESS THAN ('2013-03-17') 
);
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}

	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	for _, stmtNode := range stmtNodes {
		createTableStmt := stmtNode.(*ast.CreateTableStmt)
		fmt.Printf("Schema: %v, Table: %v\n", createTableStmt.Table.Schema.String(),
			createTableStmt.Table.Name.String())

		if createTableStmt.Partition != nil {
			fmt.Println("是分区表")

			fmt.Printf("%T --- %[1]v \n", createTableStmt.Partition.Expr)
			funcExpr := createTableStmt.Partition.Expr.(*ast.FuncCallExpr)
			fmt.Println("func Name:", funcExpr.FnName.String())

			partitionColumnName := make([]string, 0, 1)
			for _, arg := range funcExpr.Args {
				fmt.Printf("%T --- %[1]v \n", arg)
				columNameExpr := arg.(*ast.ColumnNameExpr)
				partitionColumnName = append(partitionColumnName, columNameExpr.Name.String())
			}
			fmt.Println("column Names:", partitionColumnName)


			// partition 定义
			for _, definition := range createTableStmt.Partition.Definitions {
				fmt.Printf("partition definition： %v, %v, %v \n",
					definition.Name, definition.Comment, definition.MaxValue)
			}
		} else {
			fmt.Println("不是分区表")
		}
	}
}


func TestCreateTableReviewer_Review_Partition(t *testing.T) {
	var host string = "10.10.10.12"
	var port int = 3306
	var username string = "HH"
	var password string = "oracle"
	sql := `
CREATE TABLE test.mf_fd_cache (
  id bigint(18) NOT NULL AUTO_INCREMENT COMMENT '注释',
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
  PRIMARY KEY (id, uptime),
  UNIQUE KEY udx_uid (dep, arr, flightNo, uptime),
  -- UNIQUE KEY udx_uid_1 (cabin),
  Index idx_uptime (uptime),
  KEY idx_flight (dep,arr),
  KEY idx_flightdate (flightDate)
) ENGINE=InnoDb  DEFAULT CHARSET=Utf8mb4 COLLATE=Utf8mb4_general_ci comment="注释"
PARTITION BY RANGE(TO_DAYS (uptime1))
(
    PARTITION p0 VALUES LESS THAN (TO_DAYS('2010-04-15')),
    PARTITION p1 VALUES LESS THAN (TO_DAYS('2010-05-01')),
    PARTITION p2 VALUES LESS THAN (TO_DAYS('2010-05-15')),
    PARTITION p3 VALUES LESS THAN (TO_DAYS('2010-05-31')),
    PARTITION p4 VALUES LESS THAN (TO_DAYS('2010-06-15')),
    PARTITION p19 VALUES LESS ThAN  MAXVALUE
);
    `

	sqlParser := parser.New()
	stmtNodes, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("Syntax Error: %v", err)
	}

	// 循环每一个sql语句进行解析, 并且生成相关审核信息
	dbConfig := dao.NewDBConfig(host, port, username ,password, "")
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



		review := NewReviewer(stmtNode, reviewConfig, dbConfig)
		reviewMSG := review.Review()
		reviewMSGs = append(reviewMSGs, reviewMSG)
	}

	for _, reviewMSG := range reviewMSGs {
		if reviewMSG != nil {
			fmt.Printf("Code: %v, MSG: %v", reviewMSG.Code, reviewMSG.MSG)
		}
	}
}
