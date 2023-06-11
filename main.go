package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"authservice/config"
	"authservice/controller"
	"authservice/logger"
	"authservice/repository"
)

func main() {
	cfg := config.NewConfig()
	logger.InitLogger(cfg.LogLevel)
	repo := repository.FactoryRepository(cfg.Store, cfg.DSN)
	handlers := controller.NewAuthServiceHandler(repo)
	mux := http.NewServeMux()

	mux.HandleFunc("/registration", handlers.Registration)
	mux.HandleFunc("/verification", handlers.Verification)
	mux.HandleFunc("/login", handlers.Login)

	addr := cfg.Host + ":" + cfg.Port
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Info().Str("addr", addr).Msg("starting server")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
