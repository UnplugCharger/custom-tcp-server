package main

import (
	"byron_tcp/src"
	"fmt"
	"log"
)

func main() {
	srv := src.NewServer(":3000")

	go func() {
		for msg := range srv.AllMessages() {
			fmt.Println(string(msg))
		}

	}()
	log.Fatal(srv.Start())
}
