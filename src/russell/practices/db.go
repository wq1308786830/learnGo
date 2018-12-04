package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:92QWWQ0828MM@/russell")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	rows, err := db.Query("SELECT * FROM hotel")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	for rows.Next() {
		var city string
		var name string
		var address string
		var zip string
		if err := rows.Scan(&city, &name, &address, &zip); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%s | %s | %s | %s\n", city, name, address, zip)
		//fmt.Printf("%s is %d\n", name, age)
	}
}
