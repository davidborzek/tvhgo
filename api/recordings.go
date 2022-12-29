package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func (s *router) GetRecordings(w http.ResponseWriter, r *http.Request) {
	var q core.PaginationSortQueryParams
	if err := request.BindQuery(r, &q); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := q.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	recordings, err := s.recordings.GetAll(r.Context(), q)
	if err != nil {
		log.WithError(err).
			Error("failed to get recordings")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, recordings, 200)
}

func (s *router) GetRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	recordings, err := s.recordings.Get(r.Context(), id)
	if err != nil {
		if err == core.ErrRecordingNotFound {
			response.NotFound(w, err)
			return
		}

		log.WithError(err).
			WithField("id", id).
			Error("failed to get recording")

		response.InternalErrorCommon(w)
		return
	}

	response.JSON(w, recordings, 200)
}

func (s *router) CreateRecording(w http.ResponseWriter, r *http.Request) {
	var in core.CreateRecording
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := in.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	err := s.recordings.Create(r.Context(), in)
	if err != nil {
		log.WithError(err).
			Error("failed to create recording")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(201)
}

func (s *router) CreateRecordingByEvent(w http.ResponseWriter, r *http.Request) {
	var in core.CreateRecordingByEvent
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := in.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	err := s.recordings.CreateByEvent(r.Context(), in)
	if err != nil {
		log.WithError(err).
			Error("failed to create recording by event")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(201)
}

func (s *router) StopRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := s.recordings.Stop(r.Context(), id)
	if err != nil {
		log.WithError(err).
			WithField("id", id).
			Error("failed to stop recording")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

func (s *router) CancelRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := s.recordings.Cancel(r.Context(), id)
	if err != nil {
		log.WithError(err).
			WithField("id", id).
			Error("failed to cancel recording")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

func (s *router) RemoveRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := s.recordings.Remove(r.Context(), id)
	if err != nil {
		log.WithError(err).
			WithField("id", id).
			Error("failed to remove recording")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

func (s *router) MoveRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dest := chi.URLParam(r, "dest")

	var err error
	switch dest {
	case "finished":
		err = s.recordings.MoveFinished(r.Context(), id)
	case "failed":
		err = s.recordings.MoveFailed(r.Context(), id)
	default:
		response.BadRequest(w, err)
		return
	}

	if err != nil {
		log.WithError(err).
			WithField("id", id).
			WithField("destination", dest).
			Error("failed to move recording")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

func (s *router) UpdateRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var in core.UpdateRecording
	if err := request.BindJSON(r, &in); err != nil {
		response.BadRequest(w, err)
		return
	}

	if err := in.Validate(); err != nil {
		response.BadRequest(w, err)
		return
	}

	err := s.recordings.UpdateRecording(r.Context(), id, in)
	if err != nil {
		log.WithError(err).
			WithField("id", id).
			Error("failed to update recording")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(201)
}
