package ezsqlx

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

func Insert(db *sqlx.DB, table string, model interface{}, excludedFields []string) (*sqlx.Rows, error) {
	fields := Fields(model)
	filteredFields := filter(fields, excludedFields)
	formattedFields := strings.Join(prependStrings(filteredFields, ":"), ", ")
	sql := "INSERT INTO " + table + " (" + strings.Join(wrapStrings(filteredFields, "\""), ", ") + ") VALUES (" + formattedFields + ") RETURNING *"
	return db.NamedQuery(sql, model)
}

func InsertMany(db *sqlx.DB, table string, models []interface{}, excludedFields []string) (sql.Result, error) {
	fields := Fields(models[0])
	filteredFields := filter(fields, excludedFields)
	formattedFields := strings.Join(prependStrings(filteredFields, ":"), ", ")
	sql := "INSERT INTO " + table + " (" + strings.Join(wrapStrings(filteredFields, "\""), ", ") + ") VALUES (" + formattedFields + ")"
	return db.NamedExec(sql, models)
}