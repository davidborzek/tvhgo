package api

import (
	"net/http"
	"strconv"

	"github.com/davidborzek/tvhgo/api/response"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

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
