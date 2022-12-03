package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Collect our quotes.
	file, err := os.ReadFile("quotes.txt")
	if err != nil {
		log.Fatalln("reading quotes.txt:", err)
	}
	quotes := strings.Split(string(file), "\n\n")

	// Setup our random numbers.
	rand.Seed(time.Now().UnixNano())

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, quotes[rand.Intn(len(quotes))])
	}))

	http.ListenAndServe(":8080", nil)
}
