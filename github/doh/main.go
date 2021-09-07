package main

import (
	"context"
	"fmt"
	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := doh.Use(doh.Quad9Provider)

	resp, err := c.Query(ctx, "www.google.com", dns.TypeA)
	if err != nil {
		panic(err)
	}

	answer := resp.Answer

	for _, a := range answer {
		fmt.Printf("%s->%s\n", a.Name, a.Data)
	}
}
