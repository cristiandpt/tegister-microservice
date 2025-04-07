package middleware

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		startTime := time.Now()
		lrw := &LoggedResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next(lrw, r, p)
		duration := time.Since(startTime)
		log.Printf(
			"[%s] %s %s %d %v",
			r.Method,
			r.RequestURI,
			r.Proto,
			lrw.statusCode,
			duration,
		)
	}
}

func (lrw *LoggedResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
