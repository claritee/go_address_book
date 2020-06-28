package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	r.HandleFunc("/person/{id:[0-9]+}", a.getPerson).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (a *App) getPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Person")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p := person{ID: id}

	if err := p.getPerson(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			fmt.Fprintf(w, "Error No Records")
		default:
			fmt.Fprintf(w, "Error "+err.Error())
		}
		return
	}

	fmt.Fprintf(w, p.Name)
}
