package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

type statusLogger struct {
	status int
	http.ResponseWriter
}

func (sl *statusLogger) WriteHeader(code int) {
	sl.status = code
	sl.ResponseWriter.WriteHeader(code)
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sl := &statusLogger{
			ResponseWriter: w,
		}

		// call the handler function
		next.ServeHTTP(sl, r)

		// check if file exists, if not, create it and append to it - file permission 600
		file, err := os.OpenFile("logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

		if err != nil {
			log.Fatalf(errors.Wrap(err, "error opening file\n").Error())
		}

		defer file.Close()

		// yes, you have to use this specific time if you want to use Format()
		line := fmt.Sprintf("%v %s %d %s\n", time.Now().Format("2006-01-02 15:04:05.999"), r.RequestURI, sl.status, r.Method)

		if _, err := file.Write([]byte(line)); err != nil {
			log.Printf(errors.Wrap(err, "error writing logs to file\n").Error())
		}
	})
}
