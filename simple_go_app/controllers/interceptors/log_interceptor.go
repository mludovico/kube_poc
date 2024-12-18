package interceptors

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LogInterceptorWrapper(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %s\n", err)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("%s - %s\n%s\n", r.Method, r.URL.Path, body)
		next(w, r)
	}
}
