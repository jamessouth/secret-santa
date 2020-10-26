package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Init starts the server
func Init() {
	// fs := http.FileServer(http.Dir("./dist"))
	// http.Handle("/", fs)
	http.HandleFunc("/", handleConns)

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

func handleConns(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	fmt.Println(r.Form, r.PostForm)

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
