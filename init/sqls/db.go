package sqls

import _ "embed"

var (
	//go:embed branch.sql
	branchSql string
	//go:embed global.sql
	global string
)

func Sql() []string {
	return []string{branchSql, global}
}
