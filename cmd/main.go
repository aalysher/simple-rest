package main

import (
	"log"

	"github.com/aalysher/simple-rest/internal/service"
)

func main() {
	if err := service.New(); err != nil {
		log.Fatalf("init service: %v\n", err)
	}
}
