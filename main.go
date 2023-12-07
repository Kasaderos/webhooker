package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var port = flag.String("port", "8080", "port number")

func CatchWebhooks(w http.ResponseWriter, r *http.Request) {
	msg, err := io.ReadAll(r.Body)
	headers := make([]string, 0)
	for k, vs := range r.Header {
		headers = append(headers, fmt.Sprintf("\t%s: %s", k, strings.Join(vs, ",")))
	}
	log.Printf("\nURL:\n\t%s\nHEADERS:\n%s\nMSG:\n\t%s\nERR:\n\t%v\n\n",
		r.URL.String(),
		strings.Join(headers, "\n"),
		string(msg),
		err,
	)
}

func main() {
	http.HandleFunc("/", CatchWebhooks)

	fmt.Println("starting server at", ":"+*port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
