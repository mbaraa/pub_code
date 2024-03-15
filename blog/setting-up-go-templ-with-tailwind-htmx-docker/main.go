package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"spendings/components"
	"spendings/db"
	"spendings/handlers"
	"spendings/services"
)

//go:embed static/*
var static embed.FS

//go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m

func main() {
	ctx := context.Background()

	balanceStore := db.NewBalanceStoreJson()
	spendingsStore := db.NewSpendingsStoreJson()
	balanceService := services.NewBalanceService(balanceStore)
	spendingsService := services.NewSpendingService(spendingsStore)

	pagesHandler := http.NewServeMux()
	pagesHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spendings, err := spendingsService.ListItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		components.Index(balanceService.GetBalance(), spendings).Render(ctx, w)
	})
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	spendingsHandler := handlers.NewSpendingHandler(*spendingsService)
	restHandler := http.NewServeMux()
	restHandler.HandleFunc("POST /spending", spendingsHandler.HandleAddSpendingItem)
	restHandler.HandleFunc("PUT /spending", spendingsHandler.HandleUpdateSpendingItem)
	restHandler.HandleFunc("DELETE /spending", spendingsHandler.HandleRemoveSpendingItem)

	applicationHandler := http.NewServeMux()
	applicationHandler.Handle("/", pagesHandler)
	applicationHandler.Handle("/api/", http.StripPrefix("/api", restHandler))

	log.Println("starting server on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", applicationHandler))
}
