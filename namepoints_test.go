package namepoints

import (
	"github.com/invopop/jsonschema"
	"log"
	"net/http"
	"testing"
)

type testNameHandler struct {
	LocalName string
}

func (t testNameHandler) Name() string {
	return t.LocalName
}

func (t testNameHandler) UpdatesData() bool {
	return false
}

func (t testNameHandler) Request() *jsonschema.Schema {
	return nil
}

func (t testNameHandler) Response() *jsonschema.Schema {
	return nil
}

func (t testNameHandler) GetResponse(r *http.Request) (interface{}, error) {
	return nil, nil
}

func TestHappyPath(t *testing.T) {
	nameHandlers := []Namepoint{testNameHandler{"test"}}
	_, err := NewHttpHandler(nameHandlers, map[string]string{})
	if err != nil {
		log.Fatalf("Error on HappyPath %s", err.Error())
	}
}

func TestNameAsEmptyString(t *testing.T) {
	nameHandlers := []Namepoint{testNameHandler{""}}
	handler, err := NewHttpHandler(nameHandlers, map[string]string{})

	if handler != nil {
		log.Fatalf("Handler should be nil")
	}

	if err == nil {
		log.Fatalf("Error should not be nil")
	}

	if err.Error() != "Namepoint Name() should not be empty string" {
		log.Fatalf("Error string is not expected")
	}
}

func TestDuplicateName(t *testing.T) {
	nameHandlers := []Namepoint{
		testNameHandler{"test"},
		testNameHandler{"test"},
	}
	handler, err := NewHttpHandler(nameHandlers, map[string]string{})

	if handler != nil {
		log.Fatalf("Handler should be nil")
	}

	if err == nil {
		log.Fatalf("Error should not be nil")
	}

	if err.Error() != "Namepoint Name() should be unique. Found duplicate: test" {
		log.Fatalf("Error string is not expected")
	}
}
