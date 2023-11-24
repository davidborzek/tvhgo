package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	_ "github.com/davidborzek/tvhgo/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type router struct {
	cfg                   *config.Config
	channels              core.ChannelService
	epg                   core.EpgService
	picons                core.PiconService
	recordings            core.RecordingService
	streaming             core.StreamingService
	sessionManager        core.SessionManager
	sessions              core.SessionRepository
	users                 core.UserRepository
	passwordAuthenticator core.PasswordAuthenticator
	tokens                core.TokenRepository
	tokenService          core.TokenService
	twoFactorService      core.TwoFactorAuthService
}

var corsOpts = cors.Options{
	AllowedOrigins:   []string{"*"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
	AllowCredentials: true,
	MaxAge:           300,
}

func New(
	cfg *config.Config,
	channels core.ChannelService,
	epg core.EpgService,
	picons core.PiconService,
	recordings core.RecordingService,
	streaming core.StreamingService,
	sessionManager core.SessionManager,
	sessions core.SessionRepository,
	users core.UserRepository,
	passwordAuthenticator core.PasswordAuthenticator,
	tokens core.TokenRepository,
	tokenService core.TokenService,
	twoFactorService core.TwoFactorAuthService,
) *router {
	return &router{
		cfg:                   cfg,
		channels:              channels,
		epg:                   epg,
		picons:                picons,
		recordings:            recordings,
		sessionManager:        sessionManager,
		sessions:              sessions,
		users:                 users,
		streaming:             streaming,
		passwordAuthenticator: passwordAuthenticator,
		tokens:                tokens,
		tokenService:          tokenService,
		twoFactorService:      twoFactorService,
	}
}

func (s *router) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(cors.Handler(corsOpts))

	r.Post("/login", s.Login)

	authenticated := r.With(s.HandleAuthentication)

	authenticated.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/swagger/index.html", http.StatusMovedPermanently)
	})

	authenticated.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/api/swagger/doc.json"),
	))

	authenticated.Post("/logout", s.Logout)

	authenticated.Get("/user", s.GetUser)
	authenticated.Patch("/user", s.UpdateUser)
	authenticated.Patch("/user/password", s.UpdateUserPassword)

	authenticated.Get("/two-factor-auth", s.GetTwoFactorAuthSettings)
	authenticated.Put("/two-factor-auth/setup", s.SetupTwoFactorAuth)
	authenticated.Put("/two-factor-auth/activate", s.ActivateTwoFactorAuth)
	authenticated.Put("/two-factor-auth/deactivate", s.DeactivateTwoFactorAuth)

	authenticated.Get("/sessions", s.GetSessions)
	authenticated.Delete("/sessions/{id}", s.DeleteSession)

	authenticated.Get("/tokens", s.GetTokens)
	authenticated.Post("/tokens", s.CreateToken)
	authenticated.Delete("/tokens/{id}", s.DeleteToken)

	authenticated.Get("/epg", s.GetEpg)
	authenticated.Get("/epg/events", s.GetEpgEvents)
	authenticated.Get("/epg/events/{id}", s.GetEpgEvent)
	authenticated.Get("/epg/events/{id}/related", s.GetRelatedEpgEvents)
	authenticated.Get("/epg/content-types", s.GetEpgContentTypes)

	authenticated.Get("/channels", s.GetChannels)
	authenticated.Get("/channels/{id}", s.GetChannel)
	authenticated.Get("/channels/{number}/stream", s.StreamChannel)

	authenticated.Get("/picon/{id}", s.GetPicon)

	authenticated.Get("/recordings", s.GetRecordings)
	authenticated.Post("/recordings", s.CreateRecording)

	authenticated.Delete("/recordings", s.BatchRemoveRecordings)
	authenticated.Put("/recordings/stop", s.BatchStopRecordings)
	authenticated.Put("/recordings/cancel", s.BatchCancelRecordings)

	authenticated.Post("/recordings/event", s.CreateRecordingByEvent)

	authenticated.Get("/recordings/{id}", s.GetRecording)
	authenticated.Delete("/recordings/{id}", s.RemoveRecording)
	authenticated.Patch("/recordings/{id}", s.UpdateRecording)
	authenticated.Put("/recordings/{id}/stop", s.StopRecording)
	authenticated.Put("/recordings/{id}/cancel", s.CancelRecording)
	authenticated.Put("/recordings/{id}/move/{dest}", s.MoveRecording)
	authenticated.Get("/recordings/{id}/stream", s.StreamRecording)

	return r
}
