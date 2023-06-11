package main

import (
	"encoding/json"
	"errors"
	"github.com/invopop/jsonschema"
	"net/http"
)

type addBookRequest struct {
	Book book `json:"book"`
}

type addBookResponse struct {
	Book book `json:"book"`
}

type addBookHandler struct {
	myEnv *env
}

func (a addBookHandler) Name() string {
	return "addBook"
}

func (a addBookHandler) UpdatesData() bool {
	return true
}

func (a addBookHandler) Request() *jsonschema.Schema {
	return jsonschema.Reflect(&addBookRequest{})
}

func (a addBookHandler) Response() *jsonschema.Schema {
	return jsonschema.Reflect(&addBookResponse{})
}

func (a addBookHandler) GetResponse(r *http.Request) (interface{}, error) {
	var addBook addBookRequest
	err := json.NewDecoder(r.Body).Decode(&addBook)
	if err != nil {
		return nil, err
	}

	b := addBook.Book
	if b.Title == "" || b.Author == "" {
		return nil, errors.New("`title` and `author` cannot be empty strings")
	}

	a.myEnv.Books = append(a.myEnv.Books, b)

	return addBookResponse{b}, nil
}
