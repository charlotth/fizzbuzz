package main

import (
	"log"

	"github.com/tcharlot-datasweet/fizzbuzz/pkg/server"
)

func main() {
	srv := server.New()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
	log.Println("Bye !")
}
