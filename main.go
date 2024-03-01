package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (t templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello world"))
	})
	mux.Handle("/login", &templateHandler{filename: "login.html"})

	server := http.Server{
		Handler: mux,
		Addr:    ":3000",
	}

	log.Println("serve by port :3000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
