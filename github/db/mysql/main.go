package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
	"strings"
)

const data = `{"name":{"first":"hi","last":"world"},"age":47}`

func main() {
	key2field := map[string]string{
		"name": "name",
		"age":  "age",
	}

	field2value := map[string]interface{}{}

	value := gjson.Parse(data)
	value.ForEach(func(key, value gjson.Result) bool {
		keyStr := key.String()
		field, exist := key2field[keyStr]
		if !exist {
			return true
		}
		var v interface{}
		switch value.Type {
		case gjson.JSON:
			v = value.String()
		default:
			v = value.Value()
		}
		field2value[field] = v
		return true
	})

	db, err := sql.Open("mysql", "root:110112@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var fields, wildcard []string
	var values []interface{}
	for k, v := range field2value {
		fields = append(fields, k)
		wildcard = append(wildcard, "?")
		values = append(values, v)
	}

	sql := fmt.Sprintf("INSERT INTO test1(%s)VALUES(%s)", strings.Join(fields, ","), strings.Join(wildcard, ","))
	fmt.Println(sql)

	_, err = db.Exec(sql, values...)
	fmt.Println(err)

	//var ret string
	var ret map[string]interface{}
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 3").Scan(&ret)
	fmt.Println(ret, err)
}
