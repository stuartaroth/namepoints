package namepoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/invopop/jsonschema"
	"net/http"
)

type Namepoint interface {
	Name() string
	UpdatesData() bool
	Request() *jsonschema.Schema
	Response() *jsonschema.Schema
	GetResponse(r *http.Request) (interface{}, error)
}

func NewHttpHandler(points []Namepoint, headers map[string]string) (http.Handler, error) {
	pointsMap := make(map[string]Namepoint)
	schemas := make(map[string]interface{})

	for i := 0; i < len(points); i++ {
		handler := points[i]
		name := handler.Name()

		if name == "" {
			return nil, errors.New("Namepoint Name() should not be empty string")
		}

		_, exists := pointsMap[name]
		if exists {
			return nil, errors.New(fmt.Sprintf("Namepoint Name() should be unique. Found duplicate: %s", name))
		}

		pointsMap[name] = handler
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

	genericErrorMessage := npError{"Something went wrong"}

	genericErrorMessageBits, err := json.Marshal(genericErrorMessage)
	if err != nil {
		return nil, err
	}

	return namepointsHandler{
		pointsMap:               pointsMap,
		schemas:                 schemas,
		schemaBits:              schemaBits,
		genericErrorMessageBits: genericErrorMessageBits,
		headers:                 headers,
	}, nil
}

type npError struct {
	Error string `json:"error"`
}

type namepointsHandler struct {
	pointsMap               map[string]Namepoint
	schemas                 map[string]interface{}
	schemaBits              []byte
	genericErrorMessageBits []byte
	headers                 map[string]string
}

func (npHandler namepointsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for key, value := range npHandler.headers {
		w.Header().Set(key, value)
	}

	if r.Method != "POST" {
		writeError(w, errors.New("requests must be POST"), npHandler)
		return
	}

	name := r.URL.Query().Get("name")

	nameHandler, exists := npHandler.pointsMap[name]
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

	errorMessage := npError{err.Error()}

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
