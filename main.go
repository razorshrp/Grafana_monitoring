package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	//"time"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "paras",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "3.87.248.27:3306",
		DBName:               "project",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("200 SUCCESS")

	rows, err := db.Query(`
		SELECT
			CAST(event_timestamp AS DATE) as date,
			COUNT(*) AS cnt
		FROM payments
		GROUP BY 1
		ORDER BY 1 DESC;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []output

	for rows.Next() {
		var result output
		if err := rows.Scan(&result.Time, &result.Cnt); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

type output struct {
	Time string `json:"time"`
	Cnt  int    `json:"cnt"`
}
