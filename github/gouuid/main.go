package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	//// panic on error
	//u1 := uuid.Must(uuid.NewV4())
	//fmt.Printf("UUIDv4: %s\n", u1)
	//
	//// or error handling
	//u2, err := uuid.NewV4()
	//if err != nil {
	//	fmt.Printf("Something went wrong: %s", err)
	//	return
	//}
	//fmt.Printf("UUIDv4: %s\n", u2)
	//
	// Parsing UUID from string input
	//u3, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	//if err != nil {
	//	fmt.Printf("Something went wrong: %s", err)
	//	return
	//}
	//fmt.Printf("Successfully parsed: %s \n", u3)

	//for i := 0; i < 100000; i++ {
	u5, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	fmt.Println(u5.Value())
	//u6, err := uuid.FromString(u5.String())
	//u6.

	//}
}
