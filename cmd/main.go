package main

import (
	"github.com/go-summer-dev/defaultcase"
	_ "github.com/lib/pq"
	"log"
	"mindstore/internal/router"
	"mindstore/pkg/config"
)

func main() {
	r := router.InitApi()

	log.Fatalln(r.Run(config.GetPort()))
}

func init() {
	defaultcase.SetDefaultCase(defaultcase.Snak_case)
}
