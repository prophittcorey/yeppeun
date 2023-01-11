package main

import (
	"log"

	"github.com/prophittcorey/yeppeun/internal/web"
)

func main() {
	if err := web.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
