package middleware

import "net/http"

type CorsMiddleware struct {
	Handler http.Handler
}

func NewCorsMiddleware(handler http.Handler) *CorsMiddleware {
	return &CorsMiddleware{Handler: handler}
}

func (middleware *CorsMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	middleware.Handler.ServeHTTP(writer, request)
}
