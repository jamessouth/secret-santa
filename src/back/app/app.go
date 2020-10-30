package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session

// Init starts the server
func Init(ses *session.Session) {
	// fs := http.FileServer(http.Dir("./dist"))
	// http.Handle("/", fs)
	sess = ses

	// http.HandleFunc("/", handleConns)
	http.HandleFunc("/create", handleCreate)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	log.Println("server running on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleCreate(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	fmt.Println("pf", r.PostForm)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	// fmt.Println("sv", svc)

	svc.Handlers.Send.PushFront(func(r *request.Request) {
		r.HTTPRequest.Header.Set("CustomHeader", fmt.Sprintf("%d", 10))
	})

	a, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("errr", err)
	}
	fmt.Println("ddd", a)

	val, ok := r.PostForm["email"]
	if !ok {
		w.Write([]byte("not ok"))
	}
	_, err = fmt.Fprintf(w, val[0])

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4444")
	// pl, err := json.Marshal(finalPlayer)

}

func handleConns(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	fmt.Println("conns", r.URL)

	if err != nil {
		w.Write([]byte(err.Error()))
	}
	_, err = fmt.Fprintf(w, "Hello")

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4444")
	// pl, err := json.Marshal(finalPlayer)

}
