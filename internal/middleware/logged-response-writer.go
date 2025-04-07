package middleware

import (
	"net/http"
)

type LoggedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}
