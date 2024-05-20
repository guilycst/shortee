package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/guilycst/shortee/internal/services"
)

type Short struct {
	Url string `json:"url"`
}

func main() {

	host := "localhost:8080"
	db, err := sql.Open("sqlite3", "db/shortee.db")
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

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		url, err := short.Resolve(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if url == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})

	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		b := r.Body
		v := Short{}
		json.NewDecoder(b).Decode(&v)

		id, err := short.Shorten(v.Url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		respUrl, err := url.JoinPath("https://", host, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(respUrl))
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /health", healthCheck(db))
	http.ListenAndServe(":8080", mux)
}

func healthCheck(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	}
}
