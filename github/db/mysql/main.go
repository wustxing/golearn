package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"sanguosha.com/games/sgs/framework/util"
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

	//for i := 0; i < 10000; i++ {
	//	v := i
	//	go func() {
	//		//_, err := db.Exec("Insert into hello(nickname) VALUES (?)", fmt.Sprintf("hi%d", v))
	//		err := db.QueryRow("SELECT userid FROM auth WHERE userid=? LIMIT 1;", v).Scan()
	//		if err != nil {
	//			logrus.WithError(err).Error("insert db error")
	//		}
	//	}()
	//}
	//fmt.Println("over")
	_, err = db.Exec("INSERT INTO auth (userid, account,showid, account_type,login_id,created_ip,invite_code) VALUES (?,?,?,?,?,?,?)",
		100000110210, "test_xujialong602", 11011258, 1, 1000000002, "sfdsfds", "sfdsf")
	if err != nil {

		// 已存在
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
			fmt.Println("errNum", err.Number)
		}
		fmt.Println("err", err)
		newerr := util.ParseOriginMySQLError(err)
		if newerr == util.SqlDuplicateErr {
			fmt.Println("not primary", err)
		} else if newerr == util.SqlPrimaryKeyDuplicateErr {
			fmt.Println("primary", err)
		}

	}
}
