package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Ret struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (p *Ret) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	b, _ := src.([]byte)
	return json.Unmarshal(b, p)
}

func (p *Ret) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func main() {
	db, err := sql.Open("mysql", "root:110112@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ret Ret
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 3").Scan(&ret)
	fmt.Println(ret, err)
}
