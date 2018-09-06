package main

import (
	"database/sql"
	"fmt"
	"github.com/0990/golearn/github/sqlboiler/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sgs_nuyan?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	//boil.SetDB(db)

	users, err := models.Accounts().All(db)
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	//fmt.Println(users, err)
}
