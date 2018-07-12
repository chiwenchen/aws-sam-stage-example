package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
)

// ContentType is the Content-Type header set in responses.
const ContentType = "application/json; charset=utf8"

// Message contains a simple message response.
type Message struct {
	Message string `json:"message"`
}

// Messages used by http.HandlerFunc functions.
var (
	WelcomeMessage = Message{os.Getenv("SomeString")}
	// WelcomeMessage = Message{"I am fixed"}
)

// RootHandler is a http.HandlerFunc for the / path.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessage)
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
