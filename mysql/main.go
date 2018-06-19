package main

import(
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/0990/golearn/mysql/pb"
	sq "github.com/Masterminds/squirrel"
)

func getPb()[]byte{
	contact := test.Contact{}

	person1:=test.Person{
		Id:1,
		Name:"xu",
		PhoneType:test.PhoneType_HOME,
	}

	person2 :=test.Person{
		Id:2,
		Name:"jia",
		PhoneType:test.PhoneType_WORK,
	}

	contact.Person = append(contact.Person, &person1)
	contact.Person = append(contact.Person,&person2)

	data,_ := proto.Marshal(&contact)
	return data
}

func getMapPb()[]byte{
	complexObj:=test.ComplexObj{}
	map1:=test.MapValue{
		Name:"xu",
		Age:11,
	}
	map2:=test.MapValue{
		Name:"jia",
		Age:23,
	}
	complexObj.Map = make(map[string]*test.MapValue)
	complexObj.Map["1"] = &map1;
	complexObj.Map["2"] = &map2;

	data,_ := proto.Marshal(&complexObj)

	return data
}


func main(){
	//db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	//if err!=nil{
	//	panic(err.Error())
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err!=nil{
	//	panic(err.Error())
	//}

	//stmt,err:=db.Prepare(`replace into pb(id,data)values(?,?)`)
	//if err!=nil{
	//	log.Panic(err)
	//}
	//_,err =stmt.Exec(1,getMapPb())
	//if err!=nil{
	//	log.Panic(err)
	//}

	//db.Exec("replace into pb(id,data)values '%s'",)
	//query:=sq.Insert("pb").Columns("id","data").Values(2,[]byte("hi")).RunWith(db)
	//query.QueryRow()
	//rows,err:=db.Query("select id,data from pb where id = 1")
	//defer rows.Close()
	//
	//for rows.Next(){
	//	var id int
	//	var data []byte
	//	err = rows.Scan(&id,&data)
	//	comlexObj:=test.ComplexObj{}
	//	proto.Unmarshal(data,&comlexObj)
	//	fmt.Println(comlexObj)
	//}
	//err = rows.Err()
	//if err!=nil{
	//	panic(err.Error())
	//}

	user:= sq.Select("*").From("pb").Where(sq.Eq{"id":2}).OrderBy("id").Limit(10)
	fmt.Println(user.ToSql())
	rows1,err := user.RunWith(db).Query()
	for rows1.Next(){
		var id int
		var data []byte
		err = rows1.Scan(&id,&data)
		comlexObj:=test.ComplexObj{}
		proto.Unmarshal(data,&comlexObj)
		fmt.Println("rows1")
		fmt.Println(comlexObj)
	}

}
