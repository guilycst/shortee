package main

import (
	"database/sql"
	"os"

	"github.com/guilycst/shortee/internal/services"
)

func main() {
	host := "localhost"
	db, err := sql.Open("sqlite3", "shortee.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	gen, err := services.NewSQLiteBigIntGenerator(db)
	if err != nil {
		panic(err)
	}

	short, err := services.NewShortener(db, gen)
	if err != nil {
		panic(err)
	}

	// this is a CLI app, usage is shortee -s <url>
	// we will generate a short link for the given URL
	// and print it to stdout
	// also can be used as shortee -r short_id to resolve a short link to the original URL
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	if os.Args[1] == "-s" {
		s, err := short.Shorten(os.Args[2])
		if err != nil {
			panic(err)
		}
		println("https://" + host + "/" + s)
	} else if os.Args[1] == "-r" {
		r, err := short.Resolve(os.Args[2])
		if err != nil {
			panic(err)
		}
		println(r)
	} else {
		println("usage:")
		os.Exit(1)
	}
}

func printUsage() {
	println("usage:")
	println("Shorten: shortee -s <url>")
	println("Resolve: shortee -r <url>")
}
