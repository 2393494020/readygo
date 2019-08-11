package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "strconv"
)

func check(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}

func QueryMysql()  {
	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/test")
    check(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `sys_userinfo` LIMIT 0, 2")
	check(err)
	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
 
		for i := range values {
			scanArgs[i] = &values[i]
		}

		row := make(map[string]string)
        err = rows.Scan(scanArgs...)
        for i, value := range values {
            if value != nil {
				row[columns[i]] = string(value.([]byte))
            }
		}

		for k, v := range row {
			fmt.Println(k, v)
		}
		fmt.Println()
	}
}