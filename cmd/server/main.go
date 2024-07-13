package main

import (
	"fmt"
	"log"
	"net/http"

	"quackbin/internal/api"
	"quackbin/internal/database"
	"quackbin/internal/ratelimit"
)

func main() {
	db, err := database.InitDB("quackbin.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createLimiter := ratelimit.NewRateLimiter(5.0/60.0, 10)
	accessLimiter := ratelimit.NewRateLimiter(10.0/60.0, 20)

	http.HandleFunc("/", api.EnableCORS(api.ServeFile("web/create.html")))
	http.HandleFunc("/read", api.EnableCORS(api.ServeFile("web/read.html")))

	http.HandleFunc("/api/paste", api.EnableCORS(api.RateLimitMiddleware(createLimiter, api.HandlePaste(db))))
	http.HandleFunc("/api/paste/", api.EnableCORS(api.RateLimitMiddleware(accessLimiter, api.GetPaste(db))))

	fmt.Println("QuackBin starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
