package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

type output struct {
	time string `json:"time"`
	cnt  int    `json:"count"`
}

func main() {

	router := gin.Default()
	router.GET("/metrics", func(c *gin.Context) {

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
			if err := rows.Scan(&result.time, &result.cnt); err != nil {
				log.Fatal(err)
			}
			results = append(results, result)
		}
		fmt.Println("Results: ", len(results))
		fmt.Println("results value: ", results)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		// jsonData, err := json.Marshal(results)
		// if err != nil {
		// 	log.Fatal("Error marshaling JSON:", err)
		// 	return
		// }
		// fmt.Println("M I blank?: ", string(jsonData))
		c.JSON(http.StatusOK, gin.H{
			"time": results[0].time,
			"cnt":  results[0].cnt,
		})
	})

	router.Run(":8080")
}
