package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

func main() {
	var st sonyflake.Settings
	st.StartTime = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	st.MachineID = func() (uint16, error) {
		return 101, nil
	}
	s := sonyflake.NewSonyflake(st)
	fmt.Println(s.NextID())
	fmt.Println(s.NextID())
}
