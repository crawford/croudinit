package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/crawford/crowdconfig/validator/report"
	"github.com/crawford/crowdconfig/validator/rules"

	"github.com/gorilla/mux"
)

var (
	flags = struct {
		port    int
		address string
	}{}
	templates = struct {
		validate *template.Template
	}{}
)

type Redirector string

func (s Redirector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, string(s), http.StatusTemporaryRedirect)
}

type Handler func(http.ResponseWriter, *http.Request)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	h(w, r)
}

func init() {
	flag.StringVar(&flags.address, "address", "0.0.0.0", "address to listen on")
	flag.IntVar(&flags.port, "port", 80, "port to bind on")

	templates.validate = template.Must(template.New("main").ParseFiles("views/validate.html"))
}

func main() {
	flag.Parse()

	router := mux.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", flags.address, flags.port),
		Handler: router,
	}

	router.NotFoundHandler = Redirector("/validate")
	router.HandleFunc("/validate", getValidate).Methods("GET")
	router.HandleFunc("/validate", postValidate).Methods("POST")

	log.Fatalln(server.ListenAndServe())
}

func getValidate(w http.ResponseWriter, r *http.Request) {
	if err := templates.validate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func postValidate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	context := struct {
		Report *report.Report
		Config string
	}{
		Report: &report.Report{},
	}
	if c, ok := r.Form["config"]; ok {
		if len(c) > 0 {
			context.Config = strings.Replace(c[0], "\r", "", -1)
		}
	}

	for _, r := range rules.Rules {
		r([]byte(context.Config), context.Report)
	}
	if err := templates.validate.Execute(w, context); err != nil {
		panic(err)
	}
}
