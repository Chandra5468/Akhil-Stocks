package main

import (
	"log"
	"net/http"

	"github.com/Chandra5468/Akhil-Stocks/router"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := router.Router()
	r.Use(corsMiddleware) // You can also wrap this around any particular route.
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	/*Command to run

	APP_ENV=local go run main.go

	To Build
	APP_ENV=local go build -o myapp
	*/
}
