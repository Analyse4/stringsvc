// StringService provides operations on strings.
package main

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringservice{}
	svc = loggingMiddleware{logger, svc}

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

	http.ListenAndServe(":9001", nil)
}
