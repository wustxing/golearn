package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var server = "192.168.0.166"
var port = 1433
var user = "sa"
var password = "zhengze2016"
var database = "QPAccountsDB"

type User struct {
	UserID int `json:"userID"`
}
type ErrorDes struct {
	Describe string `json:describe`
}

func test(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	body_str := string(body)
	fmt.Println(body_str)
	//var user User
	//err := json.Unmarshal(body, user)r.ParseForm
	//r.ParseForm()

	if value := r.FormValue("userID"); len(value) != 0 {
		userID, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("userID error", err)
		}
		cmd := fmt.Sprintf("select * from dbo.AccountsInfo where UserID=%d", userID)
		fmt.Println(cmd)
		retval, err := getJSON(cmd)
		if err != nil {
			log.Fatal("get invalid json")
		}
		fmt.Fprint(w, retval)
	} else {
		errDes := &ErrorDes{
			Describe: "parameter error",
		}

		retval, err := json.Marshal(errDes)
		if err != nil {
			log.Fatal("marshal error")
		}

		fmt.Fprint(w, string(retval))
	}
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable",
		server, user, password, port, database)
	fmt.Println(connString)
	var err error
	db, err = sql.Open("sqlserver", connString)

	//conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	defer db.Close()

	http.HandleFunc("/", test)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServer", err)
	}
}

//func Exec(db *sql.DB, cmd string) error {
//	rows, err := db.Query(cmd)
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//	cols, err := rows.Columns()
//	if err != nil {
//		return err
//	}
//	if cols == nil {
//		return nil
//	}
//	vals := make([]interface{}, len(cols))
//	for i := 0; i < len(cols); i++ {
//		vals[i] = new(interface{})
//		if i != 0 {
//			fmt.Print("\t")
//		}
//		fmt.Print(cols[i])
//	}
//	fmt.Println()
//	for rows.Next() {
//		err = rows.Scan(vals...)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		for i := 0; i < len(vals); i++ {
//			if i != 0 {
//				fmt.Print("\t")
//			}
//			printValue(vals[i].(*interface{}))
//		}
//		fmt.Println()

//	}
//	if rows.Err() != nil {
//		return rows.Err()
//	}
//	return nil
//}

func getJSON(cmd string) (string, error) {
	if db == nil {
		log.Fatal("db error")
	}
	rows, err := db.Query(cmd)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)

	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
