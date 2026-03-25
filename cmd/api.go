package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Shenith404/go-ecom/internal/adapters/postgre/sqlc"
	"github.com/Shenith404/go-ecom/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

// run -> start the server
// mount ->
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.RequestID) // for rate limiting
	r.Use(middleware.RealIP)    // important for logging and rate limiting
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from craches

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	productService := products.NewService(repo.New(app.db))
	productsHandler := products.NewHandler(productService)
	r.Get("/products", productsHandler.ListProducts)

	orderHandler := orders.NewHandler(nil)
	r.Post("/orders",orderHandler.PlaceOrder)

	return r

}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()

}

type config struct {
	addr string // server address eg: ":4000"
	db   dbConfig
}

type dbConfig struct {
	dsn string // data source name
}
