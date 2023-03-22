package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/core"
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
//	@Failure	404	{object}	response.ErrorResponse
//	@Failure	500	{object}	response.ErrorResponse
//	@Security	JWT
//	@Router		/picon/{id} [get]
func (s *router) GetPicon(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	picon, err := s.picons.Get(r.Context(), id)
	if err != nil {
		if err == core.ErrPiconNotFound {
			response.NotFound(w, err)
			return
		}

		log.WithError(err).
			Error("failed to get picon")
		response.InternalErrorCommon(w)
		return
	}

	io.Copy(w, picon)
}
