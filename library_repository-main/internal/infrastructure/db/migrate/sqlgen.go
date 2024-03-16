package migrate

import (
	"fmt"
	"github.com/geejjoo/library_repository/internal/db/adapter/dao"
	"github.com/geejjoo/library_repository/internal/db/tabler"
	"strings"
)

type SQLGenerator interface {
	CreateTableSQL(table tabler.Tabler) string
}

type SQLiteGenerator struct{}

func (sg *SQLiteGenerator) CreateTableSQL(table tabler.Tabler) string {
	tableName := table.TableName()
	structInfo := dao.GetStructInfo(table)
	sqlQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(structInfo.Fields, ", "))
	return sqlQuery
}

type PostgreSQLGenerator struct {
}

func (pg *PostgreSQLGenerator) CreateTableSQL(table tabler.Tabler) string {
	tableName := table.TableName()
	structInfo := dao.GetStructInfo(table)
	sqlQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(Zipper(structInfo.Fields, structInfo.FieldsTypes, structInfo.FieldsDefault), ", "))
	return sqlQuery
}

func Zipper(names, types, defaults []string) []string {
	if len(names) != len(types) {
		return nil
	}
	var out = make([]string, len(names))
	for i := range out {
		out[i] = fmt.Sprintf("%v %v %v", names[i], types[i], defaults[i])
	}
	return out
}
