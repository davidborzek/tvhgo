package api

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/request"
	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// GetRecordings godoc
//
//	@Summary	Get list of recordings
//	@Tags		recordings
//
//	@Param		limit		query	int		false	"Limit"
//	@Param		offset		query	int		false	"Offset"
//	@Param		sort_key	query	string	false	"Sort key"
//	@Param		sort_dir	query	string	false	"Sort direction"
//	@Param		status		query	string	false	"Recording status"
//
//	@Produce	json
//	@Success	200	{array}		core.Recording
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings [get]
func (s *router) GetRecordings(w http.ResponseWriter, r *http.Request) {
	var q core.GetRecordingsParams
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

// GetRecording godoc
//
//	@Summary	Get a recording by id
//	@Tags		recordings
//
//	@Param		id	path	string	true	"Recording id"
//
//	@Produce	json
//	@Success	200	{object}	core.Recording
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id} [get]
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

// CreateRecording godoc
//
//	@Summary	Create a recording
//	@Tags		recordings
//	@Accept		json
//	@Param		body	body	core.CreateRecording	true	"Body"
//	@Produce	json
//	@Success	201
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings [post]
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

// CreateRecordingByEvent godoc
//
//	@Summary	Create a recording by a event
//	@Tags		recordings
//	@Accept		json
//	@Param		body	body	core.CreateRecordingByEvent	true	"Body"
//	@Produce	json
//	@Success	201
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/event [post]
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

// StopRecording godoc
//
//	@Summary	Stops a recording
//	@Tags		recordings
//	@Param		id	path	string	true	"Recording id"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id}/stop [put]
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

// BatchStopRecordings godoc
//
//	@Summary	Stop multiple recordings
//	@Tags		recordings
//	@Param		ids	query	[]string	true	"recording ids"	collectionFormat(multi)
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/stop [put]
func (s *router) BatchStopRecordings(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids := r.Form["ids"]

	if len(ids) < 1 {
		response.BadRequestf(w, "ids are required")
		return
	}

	err := s.recordings.BatchStop(r.Context(), ids)
	if err != nil {
		log.WithError(err).
			WithField("ids", ids).
			Error("failed to stop recordings")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

// CancelRecording godoc
//
//	@Summary	Cancels a recording
//	@Tags		recordings
//	@Param		id	path	string	true	"Recording id"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id}/cancel [put]
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

// BatchCancelRecordings godoc
//
//	@Summary	Cancel multiple recordings
//	@Tags		recordings
//	@Param		ids	query	[]string	true	"recording ids"	collectionFormat(multi)
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/cancel [put]
func (s *router) BatchCancelRecordings(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids := r.Form["ids"]

	if len(ids) < 1 {
		response.BadRequestf(w, "ids are required")
		return
	}

	err := s.recordings.BatchCancel(r.Context(), ids)
	if err != nil {
		log.WithError(err).
			WithField("ids", ids).
			Error("failed to cancel recordings")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

// RemoveRecording godoc
//
//	@Summary	Removes a recording
//	@Tags		recordings
//	@Param		id	path	string	true	"Recording id"
//	@Produce	json
//	@Success	204
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id} [delete]
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

// BatchRemoveRecordings godoc
//
//	@Summary	Remove multiple recordings
//	@Tags		recordings
//	@Param		ids	query	[]string	true	"recording ids"	collectionFormat(multi)
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings [delete]
func (s *router) BatchRemoveRecordings(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids := r.Form["ids"]

	if len(ids) < 1 {
		response.BadRequestf(w, "ids are required")
		return
	}

	err := s.recordings.BatchRemove(r.Context(), ids)
	if err != nil {
		log.WithError(err).
			WithField("ids", ids).
			Error("failed to remove recordings")

		response.InternalErrorCommon(w)
		return
	}

	w.WriteHeader(204)
}

// MoveRecording godoc
//
//	@Summary	Moves a recording
//	@Tags		recordings
//	@Param		id		path	string	true	"Recording id"
//	@Param		dest	path	string	true	"Recording id"
//	@Produce	json
//	@Success	204
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id}/move/{dest} [put]
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

// UpdateRecording godoc
//
//	@Summary	Updates a recording
//	@Tags		recordings
//	@Accept		json
//	@Param		body	body	core.UpdateRecording	true	"Body"
//	@Produce	json
//	@Success	201
//	@Failure	400	{object}	response.ErrorResponse
//	@Failure	401	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/recordings/{id} [patch]
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
