package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func main() {
	user := sq.Select("*").From("pb").Where(sq.And{sq.Eq{"id": 1}, sq.Or{sq.Eq{"isdel": false}, sq.Eq{"emailtype": 1}, sq.Eq{"emailtype": 2}}})
	sql, args, err := user.ToSql()
	fmt.Println(sql, args, err)

	//sql,args,err = sq.Insert("test").Columns("id","data").Values("2",[]byte("hello")).Values("larry",sq.Expr("?+5",12)).ToSql()
	fmt.Println(sql, args, err)

	cmd, args, err := sq.Update("user").Set("name", "name+1").Where(sq.Eq{"id": 1}).ToSql()
	fmt.Println(cmd, args, err)
}
