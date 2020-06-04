package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	db *PostDB
}

func main() {
	db, err := sql.Open("mysql", "webuser:webpassword@tcp(127.0.0.1:31306)/webdb")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to database")
	defer db.Close()

	app := application{
		&PostDB{
			DB: db,
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", app.getPosts)
	fmt.Println("starting server on port 3000")
	http.ListenAndServe(":3000", mux)
}

func (app *application) getPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := app.db.GetAllPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	for _, post := range posts {
		fmt.Fprintf(w, "%v\n", post)
	}
}