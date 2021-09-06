package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getenv(key, default_key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return default_key
	}
	return value
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Hostname: ", hostname)
}

func main() {
	http.HandleFunc("/", handler)
	port := getenv("SIMPLE_APP_PORT", "8888")
	log.Fatal(http.ListenAndServe(":"+port, nil))
	fmt.Println(os.Hostname())
}
