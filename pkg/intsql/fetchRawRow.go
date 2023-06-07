package intsql

import (
	"database/sql"
)

func FetchRawRows(db *sql.DB, query string) []map[string]string {
	r, _ := db.Query(query)
	defer r.Close()

	data := make(map[string]string)

	var rows []map[string]string

	c := 0
	for r.Next() {
		cols, _ := r.Columns()
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		r.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}
		c++
		rows = append(rows, data)
	}

	return rows
}
