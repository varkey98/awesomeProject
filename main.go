package main

import (
	"log"
	"time"
)

func main() {
	config := InitConfig()

	lib, err := New(config)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.Delete()

	if err := lib.Start(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(30 * time.Second)
}
