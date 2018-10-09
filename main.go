// StringService provides operations on strings.
package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {
	svc := stringservice{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":9001", nil))
}
