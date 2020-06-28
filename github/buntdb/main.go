package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"time"
)

func main() {
	db, err := buntdb.Open("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var value string
	err = db.View(func(tx *buntdb.Tx) error {
		v, err := tx.Get("hello")
		if err != nil {
			return err
		}
		value = v
		return nil
	})

	if err != nil {
		fmt.Println("db.View ", err)
	}

	fmt.Println("value", value)

	go func() {
		err = db.Update(func(tx *buntdb.Tx) error {
			time.Sleep(time.Second * 5)
			tx.Set("hello", "cde", nil)
			fmt.Println("set hello abc")
			time.Sleep(time.Second * 5)
			return nil
		})
	}()

	go func() {
		for {
			var value string
			err = db.View(func(tx *buntdb.Tx) error {
				v, err := tx.Get("hello")
				if err != nil {
					return err
				}
				value = v
				return nil
			})

			fmt.Println("value", value)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Hour)

}
