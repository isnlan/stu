package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		"root",
		"root",
		"www.snlan.top",
		3306,
		"block",
	)
	db, err := sql.Open("mysql", dsn)
	check(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT TABLE_NAME, COLUMN_NAME, DATA_TYPE FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ?")
	check(err)
	defer stmt.Close()

	rows, err := stmt.Query("block")

	columns, err := rows.Columns()

	fmt.Println(columns)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)

	for rows.Next() {
		entry := fetchRow(rows.Scan, &columns, count)
		tableData = append(tableData, entry)
		for k, v := range entry {
			fmt.Println("k,v", k, v)
		}
		fmt.Println("------------")
	}




}
type scan func(dest ...interface{}) error
func fetchRow(scan_proc scan, columns *[]string, count int) map[string]interface{} {
	values := make([]interface{}, count)
	valueParams := make([]interface{}, count)
	for i := 0; i < count; i++ {
		valueParams[i] = &values[i]
	}
	scan_proc(valueParams...)
	entry := make(map[string]interface{})
	for i, col := range *columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		entry[col] = v
	}
	return entry
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}