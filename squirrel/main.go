package main
import(
	sq "github.com/Masterminds/squirrel"
	"fmt"
)
func main(){
	user:= sq.Select("*").From("pb").Where(sq.Eq{"id":1})
	sql,args,err :=user.ToSql()
	fmt.Println(sql,args,err)

	sql,args,err = sq.Insert("test").Columns("id","data").Values("2",[]byte("hello")).Values("larry",sq.Expr("?+5",12)).ToSql()
	fmt.Println(sql,args,err)
}
