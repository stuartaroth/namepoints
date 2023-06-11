package main

import (
	"github.com/invopop/jsonschema"
	"net/http"
)

type getAllBooksHandler struct {
	myEnv *env
}

type getAllBooksResponse struct {
	Books []book `json:"books"`
}

func (g getAllBooksHandler) Name() string {
	return "getAllBooks"
}

func (g getAllBooksHandler) UpdatesData() bool {
	return false
}

func (g getAllBooksHandler) Request() *jsonschema.Schema {
	return nil
}

func (g getAllBooksHandler) Response() *jsonschema.Schema {
	return jsonschema.Reflect(&getAllBooksResponse{})
}

func (g getAllBooksHandler) GetResponse(r *http.Request) (interface{}, error) {
	return getAllBooksResponse{g.myEnv.Books}, nil
}
