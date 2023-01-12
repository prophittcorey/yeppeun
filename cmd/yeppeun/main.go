package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/prophittcorey/yeppeun/internal/web"
)

func setenv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		log.Printf("yeppeun: failed to set a default environment variable; %s", err)
	}
}

func hasBytesToRead() bool {
	file := os.Stdin

	fi, err := file.Stat()

	if err != nil {
		return false
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

func main() {
	var (
		help   bool
		port   string
		host   string
		domain string
	)

	flag.BoolVar(&help, "h", false, "Displays the program's usage.")
	flag.StringVar(&port, "port", "3000", "The port to run the server on (default: 3000).")
	flag.StringVar(&host, "host", "127.0.0.1", "The host to run the server on (default: 127.0.0.1).")
	flag.StringVar(&domain, "domain", "localhost", "The domain name for the server (default: localhost).")

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if hasBytesToRead() {
		if bs, err := io.ReadAll(os.Stdin); err == nil {
			var data any

			if err = json.Unmarshal(bs, &data); err == nil {
				if bs, err = json.MarshalIndent(data, "", "  "); err == nil {
					fmt.Println(string(bs))
				}
			}
		}

		return
	}

	setenv("HOST", host)
	setenv("PORT", port)
	setenv("DOMAIN", domain)

	if err := web.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
