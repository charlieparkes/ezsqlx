package ezsqlx

import (
	"bytes"
	"strings"

	"github.com/charlieparkes/go-helpers"
	"github.com/charlieparkes/go-structs"
	"github.com/jmoiron/sqlx"
)

func Insert(db *sqlx.DB, table string, model interface{}, excludedFields []string) (*sqlx.Rows, error) {
	fields := Columns(structs.Fields(model))
	filteredFields := helpers.Filter(fields, excludedFields)
	formattedFields := strings.Join(helpers.PrependStrings(filteredFields, ":"), ", ")
	sql := "INSERT INTO " + table + " (" + strings.Join(helpers.WrapStrings(filteredFields, "\""), ", ") + ") VALUES (" + formattedFields + ") RETURNING *"
	return db.NamedQuery(sql, model)
}

func InsertQuery(table string, model interface{}) string {
	columns := Columns(structs.Fields(model))
	cols := strings.Join(helpers.WrapStrings(columns, "\""), ", ")
	vals := strings.Join(helpers.PrependStrings(columns, ":"), ", ")
	sql := "INSERT INTO " + table + " (" + cols + ") VALUES (" + vals + ")"
	return sql
}

func UpsertQuery(table string, model interface{}, constraint string) string {
	var sql bytes.Buffer

	sql.WriteString(InsertQuery(table, model))
	sql.WriteString(" ON CONFLICT ON CONSTRAINT " + constraint + " DO UPDATE SET ")

	sf := structs.Fields(model)
	columns := Columns(sf)
	for _, c := range primaryKey(sf) {
		columns = remove(columns, c)
	}
	fields := []string{}
	for _, c := range columns {
		fields = append(fields, c+` = EXCLUDED.`+c)
	}
	sql.WriteString(strings.Join(fields, ", "))

	return sql.String()
}
