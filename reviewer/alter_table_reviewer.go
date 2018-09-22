package reviewer

import (
	"github.com/daiguadaidai/m-sql-review/ast"
	"github.com/daiguadaidai/m-sql-review/config"
	"fmt"
)

type AlterTableReviewer struct {
	StmtNode *ast.AlterTableStmt
	ReviewConfig *config.ReviewConfig
	DBConfig *config.DBConfig
	AddColumns map[string]bool
	DropColumns map[string]bool
	AddIndexes map[string][]string
	DropIndexes map[string]bool
	AddUniqueIndexes map[string][]string
	DropUniqueIndexes map[string]bool
	IsDropPrimaryKey bool
	PKName string
	PKColumns map[string]bool
}

func (this *AlterTableReviewer) Init() {
	this.AddColumns = make(map[string]bool)
	this.DropColumns = make(map[string]bool)
	this.AddIndexes = make(map[string][]string)
}

func (this *AlterTableReviewer) Review() *ReviewMSG {
	var reviewMSG *ReviewMSG

	// 循环每个
	for _, spec := range this.StmtNode.Specs {
		fmt.Println(spec)
	}

	return reviewMSG
}
