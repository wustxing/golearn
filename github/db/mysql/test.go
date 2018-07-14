package main

import(
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func main(){
	user:= sq.Select("*").From("pb").Where(sq.Eq{"id":2}).OrderBy("id").Limit(10)
	fmt.Println(user.ToSql())
}

