package gocloudfunction

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

const message = "Hello, %s!"

type Name struct {
	// value from request
	Value string `json:"name"`
	// default value
	defaultValue string
}

func (n *Name) Get() string {
	if n.Value != "" {
		return n.Value
	}

	return n.defaultValue
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	name := Name{defaultValue: "World"}
	if r.Method == http.MethodGet {
		if err := r.ParseForm(); err != nil {
			log.Printf("r.ParseForm: %v", err)
		} else {
			name.Value = r.Form.Get("name")
		}
		_, _ = writeResponse(w, message, name.Get())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&name); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = writeResponse(w, "Bad Request: %s", err.Error())
		return
	}

	_, _ = writeResponse(w, message, name.Get())
}

// writeResponse writes a response with the given template and name and log error if exists
func writeResponse(w http.ResponseWriter, template, name string) (n int, err error) {
	if n, err = fmt.Fprintf(w, template, html.EscapeString(name)); err != nil {
		log.Printf("fmt.Fprintf: %v", err)
	}

	return
}
