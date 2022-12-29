package main

import (
	"net/http"
	"os"

	"github.com/davidborzek/tvhgo/api"
	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/db"
	"github.com/davidborzek/tvhgo/health"
	"github.com/davidborzek/tvhgo/repository/session"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/davidborzek/tvhgo/services/channel"
	"github.com/davidborzek/tvhgo/services/epg"
	"github.com/davidborzek/tvhgo/services/picon"
	"github.com/davidborzek/tvhgo/services/recording"
	"github.com/davidborzek/tvhgo/services/streaming"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})

	log.WithField("pid", os.Getpid()).
		Info("tvhgo started")

	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.WithError(err).Fatal("failed to start tvhgo")
	}

	dbConn, err := db.Connect("./tvhgo.db")
	if err != nil {
		log.WithError(err).
			Fatal("failed to create database connection")
	}

	tvhOpts := tvheadend.ClientOpts{
		URL:      cfg.Tvheadend.URL(),
		Username: cfg.Tvheadend.Username,
		Password: cfg.Tvheadend.Password,
	}

	tvhClient := tvheadend.New(tvhOpts)
	tvhStreamingClient := tvheadend.NewStreamingClient(tvhOpts)

	userRepository := user.New(dbConn)

	sessionRepository := session.New(dbConn)

	sessionManager := auth.NewSessionManager(
		sessionRepository,
		cfg.Auth.Session.MaximumInactiveLifetime,
		cfg.Auth.Session.MaximumLifetime,
	)

	passwordAuthenticator := auth.NewLocalPasswordAuthenticator(userRepository)

	channelService := channel.New(tvhClient)
	epgService := epg.New(tvhClient)
	piconService := picon.New(tvhClient)
	recordingService := recording.New(tvhClient)
	streamingService := streaming.New(tvhStreamingClient)

	apiRouter := api.New(
		cfg,
		channelService,
		epgService,
		piconService,
		recordingService,
		streamingService,
		sessionManager,
		userRepository,
		passwordAuthenticator,
	)

	healthRouter := health.New(tvhClient)

	r := chi.NewRouter()
	r.Mount("/api", apiRouter.Handler())
	r.Mount("/health", healthRouter.Handler())

	addr := cfg.Server.Addr()
	log.WithField("addr", addr).
		Info("starting the http server")

	if err := http.ListenAndServe(addr, r); err != nil {
		log.WithError(err).Fatal("failed to start http server")
	}
}
