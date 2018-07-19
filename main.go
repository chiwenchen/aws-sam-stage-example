package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// ContentType is the Content-Type header set in responses.
const ContentType = "application/json; charset=utf8"

// Message contains a simple message response.
type Message struct {
	Message string `json:"message"`
}

// Messages used by http.HandlerFunc functions.
var (
	RootMessage = Message{os.Getenv("HELLO")}
	TestMessage = Message{os.Getenv("DBConnection")}
	awsRegion   = "us-east-1"
)

func generateSsmClient() *ssm.SSM {
	session, _ := session.NewSession(&aws.Config{Region: aws.String(awsRegion)})
	return ssm.New(session)
}

// RootHandler is a http.HandlerFunc for the / path.
func RootHandler(w http.ResponseWriter, r *http.Request) {

	ssmClient := generateSsmClient()

	key := "first-param"
	keyPointer := &key
	gpo, _ := ssmClient.GetParameter(&ssm.GetParameterInput{Name: keyPointer})

	value := gpo.Parameter.Value

	json.NewEncoder(w).Encode(Message{*value})
	// json.Unmarshal(data, w)
}

// HelloHandler is a http.HandlerFunc for the /hello path.
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	json.NewEncoder(w).Encode(RootMessage)
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
	http.HandleFunc("/hello", HelloHandler)
}

// h wraps a http.HandlerFunc and adds common headers.
func h(hf http.HandlerFunc) http.Handler {
	generalHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		hf.ServeHTTP(w, r)
	}
	return http.HandlerFunc(generalHandler)
}

type hotdog int
type food string
type water string

func (h hotdog) eat(f food, w water) {
	fmt.Println(f)
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))

}
