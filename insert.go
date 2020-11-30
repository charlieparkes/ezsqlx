package ezsqlx

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/tovala/go-helpers"
)

func Insert(db *sqlx.DB, table string, model interface{}, excludedFields []string) (*sqlx.Rows, error) {
	fields := Fields(model)
	filteredFields := helpers.Filter(fields, excludedFields)
	formattedFields := strings.Join(helpers.PrependStrings(filteredFields, ":"), ", ")
	sql := "INSERT INTO " + table + " (" + strings.Join(helpers.WrapStrings(filteredFields, "\""), ", ") + ") VALUES (" + formattedFields + ") RETURNING *"
	return db.NamedQuery(sql, model)
}

// Maybe go generics will make InsertMany feasible?

// func InsertManyQuery(db *sqlx.DB, table string, model interface{}, excludedFields []string) string {
// 	fields := Fields(model)
// 	filteredFields := helpers.Filter(fields, excludedFields)
// 	formattedFields := strings.Join(helpers.PrependStrings(filteredFields, ":"), ", ")
// 	sql := "INSERT INTO " + table + " (" + strings.Join(helpers.WrapStrings(filteredFields, "\""), ", ") + ") VALUES (" + formattedFields + ")"
// 	return sql
// }

// func InsertMany(db *sqlx.DB, table string, models []interface{}, excludedFields []string) (*sqlx.Rows, error) {
// 	sql := InsertManyQuery(db, table, models, excludedFields)
// 	return db.NamedQuery(sql, models)
// }
