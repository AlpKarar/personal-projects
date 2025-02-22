package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
		fmt.Fprintf(w, "reflect.TypeOf(x).Kind(): %s", reflect.TypeOf(os.Getenv("REDIS_PORT")).Kind())
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
