package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sgs_nuyan?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Exec("INSERT INTO load_balance (name,v,log_time) VALUES (?,?,?)", "game", 1, time.Now())
	if err != nil {
		fmt.Println(result, err)
	}
	err = db.QueryRow("SELECT userid FROM auth WHERE userid=? LIMIT 1;", 1).Scan()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("over")
}
