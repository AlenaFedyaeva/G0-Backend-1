package main

import (
	"homework4/config"
	"homework4/controllers/server"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Error readconfig file")
	}

	confParsed, err := conf.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("Error parse config")
	}

	zerolog.SetGlobalLevel(confParsed.LogLevel)

	server, err := server.NewServer(conf.PGConnInfo, log.Logger)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connect to database")
	}

	defer func() {
		err = server.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error close db conn")
		}
	}()

	rchi := chi.NewRouter()

	// A good base middleware stack
	rchi.Use(middleware.RequestID)
	rchi.Use(middleware.RealIP)
	rchi.Use(middleware.Logger)
	rchi.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	rchi.Use(middleware.Timeout(60 * time.Second))

	rchi.Get("/", server.UserPrintHello)

	s := &http.Server{
		Addr:    conf.ListenAddr,
		Handler: rchi,
	}

	log.Info().Str("Addr", s.Addr).Msg("Listener Server")
	log.Error().Err(s.ListenAndServe()).Msg("Error listener server")
}
