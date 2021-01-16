package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:xu2009@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
	}

	err = engine.Sync(new(User))
	fmt.Println(err)
	//
	//_,err = engine.Insert(&User{
	//	Id:      100,
	//	Name:    "xujialong",
	//	Salt:    "",
	//	Age:     0,
	//	Passwd:  "",
	//	Created: time.Now(),
	//	Updated: time.Time{},
	//})
	user := User{
		Name: "xujialong",
	}
	engine.Get(&user)
	fmt.Println(user)
}
