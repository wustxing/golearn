package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sgs_nuyan?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < 10000; i++ {
		v := i
		go func() {
			//_, err := db.Exec("Insert into hello(nickname) VALUES (?)", fmt.Sprintf("hi%d", v))
			err := db.QueryRow("SELECT userid FROM auth WHERE userid=? LIMIT 1;", v).Scan()
			if err != nil {
				logrus.WithError(err).Error("insert db error")
			}
		}()
	}
	fmt.Println("over")
	time.Sleep(time.Minute)
}
