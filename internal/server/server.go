package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Config struct {
	Port            string `env:"PORT" envDefault:"8080"`
	HandlerTimeout  int    `env:"PICKUP_LOCATIONS_HANDLER_TIMEOUT" envDefault:"30"`
	ShutdownTimeout int    `env:"PICKUP_LOCATIONS_SHUTDOWN_TIMEOUT" envDefault:"30"`
}

type Server struct {
	config Config
	server *http.Server
	router chi.Router
}

func New(config Config) *Server {
	logger := httplog.NewLogger("pickup-locations", httplog.Options{
		Concise: true,
		JSON:    true,
	})

	router := chi.NewRouter()

	router.Use(middleware.Timeout(
		time.Duration(config.HandlerTimeout) * time.Second),
	)
	router.Use(httplog.RequestLogger(logger))
	router.Use(middleware.Recoverer)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
	))

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	return &Server{
		config: config,
		server: server,
		router: router,
	}
}

func (s *Server) Start() {
	log.Println("Starting pickup-locations...", s.config.Port)
	s.server.ListenAndServe()
}

func (s *Server) Shutdown() {
	shutdownCtx, cancelShutdownCtx := context.WithTimeout(
		context.Background(), time.Duration(s.config.ShutdownTimeout)*time.Second,
	)
	defer cancelShutdownCtx()

	s.server.Shutdown(shutdownCtx)
}
