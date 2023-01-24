package api

import (
	"net/http"
	"strconv"

	"github.com/davidborzek/tvhgo/api/response"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// GetPicon godoc
//
//	@Summary	Get channel picon
//	@Tags		picon
//	@Param		id	path	string	true	"Picon id"
//	@Produce	image/*
//	@Produce	json
//	@Success	200
//	@Failure	401	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/picon/{id} [get]
func (s *router) GetPicon(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	res, err := s.picons.Get(r.Context(), id)
	if err != nil {
		log.WithError(err).
			Error("failed to get picon")
		response.InternalErrorCommon(w)
		return
	}

	response.CopyResponse(w, res)
}
