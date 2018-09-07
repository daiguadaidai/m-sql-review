package reviewer

import "github.com/daiguadaidai/m-sql-review/ast"

func NewReviewer(_stmtNode ast.Node) Reviewer {
	switch stmt := _stmtNode.(type) {
	case *ast.CreateDatabaseStmt:
		return &CreateDatabaseReviewer{StmtNode: stmt}
	case *ast.DropDatabaseStmt:
		return &DropDatabaseReviewer{StmtNode: stmt}
	case *ast.CreateTableStmt:
	case *ast.DropTableStmt:
		return &DropTableReviewer{StmtNode: stmt}
	case *ast.RenameTableStmt:
		return &RenameTableReviewer{StmtNode: stmt}
	case *ast.CreateViewStmt:
	case *ast.CreateIndexStmt:
	case *ast.DropIndexStmt:
	case *ast.AlterTableStmt:
		return &AlterTableReviewer{StmtNode: stmt}
	case *ast.TruncateTableStmt:
		return &TruncateTableReviewer{StmtNode: stmt}
	case *ast.SelectStmt:
	case *ast.UnionStmt:
	case *ast.LoadDataStmt:
	case *ast.InsertStmt:
	case *ast.DeleteStmt:
	case *ast.UpdateStmt:
	case *ast.ShowStmt:
	case *ast.TraceStmt:
	case *ast.ExplainStmt:
	case *ast.PrepareStmt:
	case *ast.DeallocateStmt:
	case *ast.ExecuteStmt:
	case *ast.BeginStmt:
	case *ast.BinlogStmt:
	case *ast.CommitStmt:
	case *ast.RollbackStmt:
	case *ast.UseStmt:
	case *ast.FlushStmt:
	case *ast.KillStmt:
	case *ast.SetStmt:
	case *ast.SetPwdStmt:
	case *ast.CreateUserStmt:
	case *ast.AlterUserStmt:
	case *ast.DropUserStmt:
	case *ast.DoStmt:
	case *ast.AdminStmt:
	case *ast.RevokeStmt:
	case *ast.GrantStmt:
	case *ast.AnalyzeTableStmt:
	case *ast.DropStatsStmt:
	case *ast.LoadStatsStmt:
	}

	return nil
}