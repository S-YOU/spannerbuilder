package spannerbuilder

import (
	"log"
	"strings"

	"cloud.google.com/go/spanner"
)

func GetInsertSQL(table string, cols, keys []string, values []interface{}) spanner.Statement {
	setColumns := make([]string, 0, len(cols)-len(keys))
	for _, c := range cols {
		if stringSliceContains(keys, c) {
			continue
		}
		setColumns = append(setColumns, c+" = @"+c)
	}

	sql := `INSERT ` + kwQuoted(table) + ` (` + strings.Join(cols, ", ") + `) VALUES (@` + strings.Join(cols, ", @") + ")"
	stmt := spanner.NewStatement(sql)
	for i, col := range cols {
		stmt.Params[col] = values[i]
	}

	if debug {
		log.Printf("SQL: `%s`, Params: %+v\n", sql, stmt.Params)
	}
	return stmt
}

func GetUpdateSQL(table string, cols, keys []string, values []interface{}) spanner.Statement {
	where := make([]string, len(keys))
	for i, k := range keys {
		where[i] = k + " = @" + k
	}
	setColumns := make([]string, 0, len(cols)-len(keys))
	for _, c := range cols {
		if stringSliceContains(keys, c) {
			continue
		}
		setColumns = append(setColumns, c+" = @"+c)
	}

	sql := `UPDATE ` + kwQuoted(table) + ` SET ` + strings.Join(setColumns, ", ") + ` WHERE ` + strings.Join(where, " AND ")
	stmt := spanner.NewStatement(sql)
	for i, col := range cols {
		stmt.Params[col] = values[i]
	}

	if debug {
		log.Printf("SQL: `%s`, Params: %+v\n", sql, stmt.Params)
	}
	return stmt
}

func GetDeleteSQL(table string, keys []string, values []interface{}) spanner.Statement {
	where := make([]string, len(keys))
	for i, k := range keys {
		where[i] = k + " = @" + k
	}

	sql := `DELETE FROM ` + kwQuoted(table) + ` WHERE ` + strings.Join(where, " AND ")
	stmt := spanner.NewStatement(sql)
	for i, col := range keys {
		stmt.Params[col] = values[i]
	}

	if debug {
		log.Printf("SQL: `%s`, Params: %+v\n", sql, stmt.Params)
	}
	return stmt
}
