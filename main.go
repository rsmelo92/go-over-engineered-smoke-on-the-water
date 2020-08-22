package main

// TODO: Fetch from this api https://pipl.ir/v1/getPerson 

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func booksRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main()  {
	r := mux.NewRouter()
	
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)
	r.HandleFunc("/person/{title}/page/{page}", booksRoute)

	port := ":8080"
	fmt.Printf("Running server on http://localhost%s\n", port)
	
	err := http.ListenAndServe(port, nil)
  if err != nil {
    log.Fatal(err)
  }
}
