package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func check(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}

func QueryMysql()  {
	db, err := sql.Open("mysql", "5gai:5G_ai_2019@tcp(10.17.35.114:3306)/5gai")
    check(err)
	defer db.Close()

	juneRows, err := db.Query("SELECT * FROM `WID_KPI_DWPRB_JULY` LIMIT 0, 1")
	check(err)
	for juneRows.Next() {
		columns, _ := juneRows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
 
		for i := range values {
			scanArgs[i] = &values[i]
		}

		row := make(map[string]float64)
        err = juneRows.Scan(scanArgs...)
        for i, value := range values {
            if value != nil {
				row[columns[i]], _ = strconv.ParseFloat(string(value.([]byte)), 64)
            }
		}

		for k, v := range row {
			fmt.Println(k, v)
		}
	}
}