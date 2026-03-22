package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Shenith404/go-ecom/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
type application struct {
	config config
}

// run -> start the server
// mount ->
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.RequestID) // for rate limiting
	r.Use(middleware.RealIP) // important for logging and rate limiting  
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from craches

	r.Use(middleware.Timeout(60*time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	productService := products.NewService()
	productsHandler := products.NewHandler(productService)
	r.Get("/products",productsHandler.ListProducts)

	return r

}

func(app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at addr %s",app.config.addr)
	return srv.ListenAndServe()

}


type config struct {
	addr string // server address eg: ":4000"
	db   dbConfig
}

type dbConfig struct {
	dsn string // data source name
}