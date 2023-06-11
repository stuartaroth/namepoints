package namepoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/invopop/jsonschema"
	"net/http"
)

type NameHandler interface {
	Name() string
	UpdatesData() bool
	Request() *jsonschema.Schema
	Response() *jsonschema.Schema
	GetResponse(r *http.Request) (interface{}, error)
}

func NewHandler(handlers []NameHandler) (http.Handler, error) {
	nameHandlers := make(map[string]NameHandler)
	schemas := make(map[string]interface{})

	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]
		name := handler.Name()

		if name == "" {
			return nil, errors.New("NameHandler Name() should not be empty string")
		}

		_, exists := nameHandlers[name]
		if exists {
			return nil, errors.New(fmt.Sprintf("NameHandler Name() should be unique. Found duplicate: %s", name))
		}

		nameHandlers[name] = handler
		schemas[name] = map[string]interface{}{
			"updatesData":    handler.UpdatesData(),
			"requestSchema":  handler.Request(),
			"responseSchema": handler.Response(),
		}
	}

	schemaBits, err := json.Marshal(schemas)
	if err != nil {
		return nil, err
	}

	genericErrorMessage := map[string]interface{}{
		"error": "Error",
	}

	genericErrorMessageBits, err := json.Marshal(genericErrorMessage)
	if err != nil {
		return nil, err
	}

	return namepointsHandler{
		nameHandlers:            nameHandlers,
		schemas:                 schemas,
		schemaBits:              schemaBits,
		genericErrorMessageBits: genericErrorMessageBits,
	}, nil
}

type namepointsHandler struct {
	nameHandlers            map[string]NameHandler
	schemas                 map[string]interface{}
	schemaBits              []byte
	genericErrorMessageBits []byte
}

func (npHandler namepointsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		writeError(w, errors.New("requests must be POST"), npHandler)
		return
	}

	name := r.URL.Query().Get("name")

	nameHandler, exists := npHandler.nameHandlers[name]
	if exists {
		response, err := nameHandler.GetResponse(r)
		if err != nil {
			writeError(w, err, npHandler)
			return
		}

		bits, err := json.Marshal(response)
		if err != nil {
			writeError(w, err, npHandler)
			return
		}

		_, err = w.Write(bits)
		if err != nil {
			writeError(w, err, npHandler)
			return
		}

	} else {
		_, err := w.Write(npHandler.schemaBits)
		if err != nil {
			writeError(w, err, npHandler)
			return
		}
	}
}

func writeError(w http.ResponseWriter, err error, npHandler namepointsHandler) {
	w.WriteHeader(500)

	errorMessage := map[string]interface{}{
		"error": err.Error(),
	}

	bits, err := json.Marshal(errorMessage)
	if err != nil {
		_, err = w.Write(npHandler.genericErrorMessageBits)
		if err != nil {
			return
		}
	}

	_, err = w.Write(bits)
	if err != nil {
		return
	}
}
