package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	//	"strings"
)

func main() {
	uuids := []string{"8211ea1c-3fb4-11e9-8066-00ff4bc475e8", "c39f39bc-3fb4-11e9-8066-00ff4bc475e8", "ba14cbfc-3fb6-11e9-8066-00ff4bc475e8", "ba14cbfc-3fb6-11e9-8066-00ff4bc475e8", "0cf22cfc-3fb7-11e9-93be-00ff4bc475e8"}
	//uuids := []string{"8211ea1c-3fb4-11e9-8066-00ff4bc475e8", "c39f39bc-3fb4-11e9-8066-00ff4bc475e8"}
	//uuids := []string{"8211ea1c-3fb4-11e9-8066-00ff4bc475e8"}
	selectInParam := strings.Join(uuids, ",")
	args := make([]interface{}, len(uuids))
	for i, id := range uuids {
		args[i] = id
	}
	//sqlStr := fmt.Sprintf("SELECT uuid,season_id, start_time, over_time, over_type, start_info, over_info FROM %s WHERE uuid in (?);",
	//	"2566_mode5", strings.Repeat(",?", len(uuids)-1)+`)`)
	sqlStr := "SELECT uuid,season_id, start_time, over_time, over_type, start_info, over_info FROM " + "2566_mode5" + " WHERE uuid in (?" + strings.Repeat(",?", len(uuids)-1) + `);`
	fmt.Println(sqlStr)
	db, err := sql.Open("mysql", "nuyan_api:KOgNpKmlhPtk5UsB@tcp(10.225.136.159:3306)/sgs_nuyan_game_info?parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(selectInParam)
	rows, e := db.Query(sqlStr, args...)

	if e != nil || rows == nil {
		panic(e.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println("result")
	}
}
