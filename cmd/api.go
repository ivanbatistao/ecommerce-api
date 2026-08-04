package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	repo "github.com/ivanbatistao/ecommerce-api/internal/adapters/postgresql/sqlc"
	"github.com/ivanbatistao/ecommerce-api/internal/orders"
	"github.com/ivanbatistao/ecommerce-api/internal/products"
	"github.com/jackc/pgx/v5"
)

// mount
func (app *Application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// Middlewares
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	})

	productService := products.NewService(repo.New(app.db))
	productsHandler := products.NewHandler(productService)
	r.Get("/products", productsHandler.ListProducts)
	r.Get("/products/{id}", productsHandler.FindProductById)

	ordersService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(ordersService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

// run
func (app *Application) run(h http.Handler) error {
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

type Application struct {
	config Config
	// logger
	db *pgx.Conn
}

type Config struct {
	addr string
	db   DbConfig
}

type DbConfig struct {
	dsn string
}
