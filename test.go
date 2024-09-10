// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/go-sql-driver/mysql"
// 	"log"
// 	//"time"
// )

// var db *sql.DB

// func main() {
// 	cfg := mysql.Config{
// 		User:                 "paras",
// 		Passwd:               "password",
// 		Net:                  "tcp",
// 		Addr:                 "3.87.248.27:3306",
// 		DBName:               "project",
// 		AllowNativePasswords: true,
// 	}
// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("200 SUCCESS")

// 	rows, err := db.Query("SELECT DATE_ADD( CAST(event_timestamp AS DATE), INTERVAL (FLOOR(HOUR(event_timestamp) * 60 + MINUTE(event_timestamp) / 5) * 5) MINUTE ), COUNT(*) FROM payments GROUP BY 1 ORDER BY 1 DESC;")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var id string
// 	var eventTimestamp int

// 	for rows.Next() {
// 		if err := rows.Scan(&id, &eventTimestamp); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("ID: %d, Cnt: %s\n", id, eventTimestamp)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	type output struct {
// 		time string `json:"time"`
// 		cnt  int    `json:"cnt"`
// 	}

// 	final := Final{
// 		time: id,
// 		cnt:  eventTimestamp,
// 	}

// 	jsonData, err := json.Marshal(final)
// 	if err != nil {
// 		fmt.Print("error")
// 		return
// 	}
// 	fmt.Println(string(jsonData))
// }

// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	//"time"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-sql-driver/mysql"
// 	"net/http"
// )

// var db *sql.DB

// func main() {

// 	router := gin.Default()
// 	router.GET("/metrics", func(c *gin.Context) {

// 	c.JSON(http.StatusOK, gin.H{
// 		"time": time,
// 		"count": cnt,
// 	})
// 	}

// 	cfg := mysql.Config{
// 		User:                 "paras",
// 		Passwd:               "password",
// 		Net:                  "tcp",
// 		Addr:                 "3.87.248.27:3306",
// 		DBName:               "project",
// 		AllowNativePasswords: true,
// 	}

// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("200 SUCCESS")

// 	rows, err := db.Query(`
// 		SELECT
// 			CAST(event_timestamp AS DATE) as date,
// 			COUNT(*) AS cnt
// 		FROM payments
// 		GROUP BY 1
// 		ORDER BY 1 DESC;
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var results []output

// 	for rows.Next() {
// 		var result output
// 		if err := rows.Scan(&result.Time, &result.Cnt); err != nil {
// 			log.Fatal(err)
// 		}
// 		results = append(results, result)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	jsonData, err := json.Marshal(results)
// 	if err != nil {
// 		log.Fatal("Error marshaling JSON:", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))

// }

// type output struct {
// 	time string `json:"time"`
// 	cnt  int    `json:"count"`
// }
