package server

import (
	"net/http"
	"os"

	"github.com/davidborzek/tvhgo/api"
	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/db"
	"github.com/davidborzek/tvhgo/health"
	"github.com/davidborzek/tvhgo/metrics"
	"github.com/davidborzek/tvhgo/repository/session"
	"github.com/davidborzek/tvhgo/repository/token"
	twofactorsettings "github.com/davidborzek/tvhgo/repository/two_factor_settings"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/davidborzek/tvhgo/services/channel"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/davidborzek/tvhgo/services/epg"
	"github.com/davidborzek/tvhgo/services/picon"
	"github.com/davidborzek/tvhgo/services/recording"
	"github.com/davidborzek/tvhgo/services/streaming"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/davidborzek/tvhgo/ui"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:   "server",
	Usage:  "Starts the tvhgo server",
	Action: start,
}

func start(ctx *cli.Context) error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})

	log.WithField("pid", os.Getpid()).
		Info("tvhgo started")

	cfg, err := config.Load()
	if err != nil {
		log.WithError(err).Fatal("failed to start tvhgo")
	}

	dbConn, err := db.Connect(cfg.Database.Path)
	if err != nil {
		log.WithError(err).
			WithField("db", cfg.Database.Path).
			Fatal("failed to create database connection")
	}

	log.WithField("db", cfg.Database.Path).
		Info("database connection established")

	clock := clock.NewClock()

	tvhOpts := tvheadend.ClientOpts{
		URL:      cfg.Tvheadend.URL(),
		Username: cfg.Tvheadend.Username,
		Password: cfg.Tvheadend.Password,
	}

	tvhClient := tvheadend.New(tvhOpts)
	tvhStreamingClient := tvheadend.NewStreamingClient(tvhOpts)

	userRepository := user.New(dbConn, clock)
	sessionRepository := session.New(dbConn, clock)

	// TODO clock
	tokenRepository := token.New(dbConn)
	twoFactorSettingsRepository := twofactorsettings.New(dbConn)

	sessionManager := auth.NewSessionManager(
		sessionRepository,
		clock,
		cfg.Auth.Session.MaximumInactiveLifetime,
		cfg.Auth.Session.MaximumLifetime,
		cfg.Auth.Session.TokenRotationInterval,
	)

	tokenService := auth.NewTokenService(tokenRepository)

	twoFactorService := auth.NewTwoFactorAuthService(
		twoFactorSettingsRepository,
		userRepository,
		&cfg.Auth.TOTP,
	)
	passwordAuthenticator := auth.NewLocalPasswordAuthenticator(userRepository, twoFactorService)

	channelService := channel.New(tvhClient)
	epgService := epg.New(tvhClient)
	piconService := picon.New(tvhClient)
	recordingService := recording.New(tvhClient)
	streamingService := streaming.New(tvhStreamingClient)

	sessionCleaner := auth.NewSessionCleaner(
		sessionRepository,
		clock,
		cfg.Auth.Session.CleanupInterval,
		cfg.Auth.Session.MaximumInactiveLifetime,
		cfg.Auth.Session.MaximumLifetime,
	)
	sessionCleaner.Start()

	apiRouter := api.New(
		cfg,
		channelService,
		epgService,
		piconService,
		recordingService,
		streamingService,
		sessionManager,
		sessionRepository,
		userRepository,
		passwordAuthenticator,
		tokenRepository,
		tokenService,
		twoFactorService,
	)

	healthRouter := health.New(tvhClient, dbConn)

	uiRouter, err := ui.NewRouter()
	if err != nil {
		log.WithError(err).Fatal("failed to create embedded ui router")
	}

	r := chi.NewRouter()

	r.Mount("/api", apiRouter.Handler())
	r.Mount("/health", healthRouter.Handler())

	r.HandleFunc(cfg.Metrics.Path, metrics.Handler(
		cfg.Metrics,
		metrics.NewTvheadendCollector(tvhClient),
	))

	r.Mount("/", uiRouter)

	addr := cfg.Server.Addr()
	log.WithField("addr", addr).
		Info("starting the http server")

	if err := http.ListenAndServe(addr, r); err != nil {
		log.WithError(err).Fatal("failed to start http server")
	}

	return nil
}
