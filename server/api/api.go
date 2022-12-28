package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type router struct {
	cfg                   *config.Config
	channels              core.ChannelService
	epg                   core.EpgService
	picons                core.PiconService
	recordings            core.RecordingService
	streaming             core.StreamingService
	sessions              core.SessionManager
	users                 core.UserRepository
	passwordAuthenticator core.PasswordAuthenticator
}

func New(
	cfg *config.Config,
	channels core.ChannelService,
	epg core.EpgService,
	picons core.PiconService,
	recordings core.RecordingService,
	streaming core.StreamingService,
	sessions core.SessionManager,
	users core.UserRepository,
	passwordAuthenticator core.PasswordAuthenticator,
) *router {
	return &router{
		cfg:                   cfg,
		channels:              channels,
		epg:                   epg,
		picons:                picons,
		recordings:            recordings,
		sessions:              sessions,
		users:                 users,
		streaming:             streaming,
		passwordAuthenticator: passwordAuthenticator,
	}
}

func (s *router) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	// TODO: cors

	r.Post("/login", s.Login)
	r.With(s.HandleAuthentication).Post("/logout", s.Logout)

	r.With(s.HandleAuthentication).Get("/user", s.GetUser)

	r.With(s.HandleAuthentication).Get("/epg/events", s.GetEpg)
	r.Get("/epg/events/{id}", s.GetEpgEvent)
	r.Get("/epg/content-types", s.GetEpgContentTypes)

	r.Get("/channels", s.GetChannels)
	r.Get("/channels/{id}", s.GetChannel)
	r.Get("/channels/{number}/stream", s.StreamChannel)

	r.Get("/picon/{id}", s.GetPicon)

	r.Get("/recordings", s.GetRecordings)
	r.Get("/recordings/{id}", s.GetRecording)
	r.Post("/recordings", s.CreateRecording)
	r.Post("/recordings/event", s.CreateRecordingByEvent)
	r.Put("/recordings/{id}/stop", s.StopRecording)
	r.Put("/recordings/{id}/cancel", s.CancelRecording)
	r.Put("/recordings/{id}/move/{dest}", s.MoveRecording)
	r.Delete("/recordings/{id}", s.RemoveRecording)
	r.Put("/recordings/{id}", s.UpdateRecording)

	return r
}
