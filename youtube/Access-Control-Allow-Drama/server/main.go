package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	allowedOrigins    string
	allowedOriginsMap = map[string]bool{}
)

func init() {
	allowedOrigins = os.Getenv("ALLOWED_ORIGINS")
	parseAllowedOringins()
}

func parseAllowedOringins() {
	_allowedOrigins := strings.Split(
		regexp.MustCompile(`\s*,\s*`).ReplaceAllString(allowedOrigins, ","),
		",",
	)

	if len(_allowedOrigins) == 0 {
		allowedOriginsMap["null"] = true
	}

	for _, allowedOrigin := range _allowedOrigins {
		allowedOriginsMap[allowedOrigin] = true
	}
}

func main() {
	http.HandleFunc("/", someHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func allowCorsHandler(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		// the linear substring check method
		// if !strings.Contains(allowedOrigins, origin) {
		// 	return
		// }
		// w.Header().Set("Access-Control-Allow-Origin", origin)

		// the super fast hash map method
		if allowedOriginsMap[origin] || allowedOriginsMap["*"] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		handler(w, r)
	})
}

func someHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ay Wassup!")
}
