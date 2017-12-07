package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)

		if err != nil {
			// throw a 500
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		// casting a []byte to a string
		log.Println("Request: %+v", string(reqDump))

		// lets also log it to a file
		next.ServeHTTP(w, r)
	})
}
