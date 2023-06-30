package main

import (
	"github.com/stuartaroth/namepoints"
	"log"
	"net/http"
)

type book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type env struct {
	Books []book
}

func main() {

	books := []book{
		book{"It", "Stephen King", "Horror"},
		book{"Ringworld", "Larry Niven", "Science Fiction"},
		book{"A Game of Thrones", "George R.R. Martin", "Fantasy"},
		book{"Murder on the Orient Express", "Agatha Christie", "Mystery"},
		book{"In Cold Blood", "Truman Capote", "True Crime"},
	}

	myEnv := env{books}

	var nameHandlers []namepoints.Namepoint
	nameHandlers = append(nameHandlers, getAllBooksHandler{&myEnv})
	nameHandlers = append(nameHandlers, addBookHandler{&myEnv})
	httpHandler, err := namepoints.NewHttpHandler(nameHandlers)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
