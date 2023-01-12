package main

import (
	"flag"
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
		/* TODO: Need to beautify the input... */

		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			log.Fatal(err)
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
