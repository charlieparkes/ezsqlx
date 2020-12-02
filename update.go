package ezsqlx

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/tovala/go-helpers"
)

func Update(db *sqlx.DB, table string, model interface{}, where string, excludedFields []string) (sql.Result, error) {
	fields := helpers.Filter(Fields(model), excludedFields)
	quotedFields := helpers.WrapStrings(fields, "\"")
	namedFields := helpers.PrependStrings(fields, ":")
	updateFields := []string{}
	for i, _ := range fields {
		updateFields = append(updateFields, fmt.Sprintf("%v=%v", quotedFields[i], namedFields[i]))
	}
	sql := "UPDATE " + table + " SET " + strings.Join(updateFields, ", ") + " WHERE " + where
	return db.NamedExec(sql, model)
}
