package main

import (
	"github.com/go-summer-dev/defaultcase"
	_ "github.com/lib/pq"
	"log"
	"mindstore/internal/router"
)

func main() {
	r := router.InitApi()

	log.Fatalln(r.Run(":8765"))
}

func init() {
	defaultcase.SetDefaultCase(defaultcase.Snak_case)
}
