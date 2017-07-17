package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/geodan/gost/sensorthings/entities"
	"github.com/geodan/gost/sensorthings/models"
	"github.com/geodan/gost/sensorthings/rest/reader"
	"github.com/geodan/gost/sensorthings/rest/writer"
)

// handlePutRequest todo: currently almost same as handlePostRequest, merge if it stays like this
func handlePutRequest(w http.ResponseWriter, e *models.Endpoint, r *http.Request, entity entities.Entity, h *func() (interface{}, []error), indentJSON bool) {
	if !reader.CheckContentType(w, r, indentJSON) {
		return
	}

	byteData, _ := ioutil.ReadAll(r.Body)
	err := entity.ParseEntity(byteData)
	if err != nil {
		writer.SendError(w, []error{err}, indentJSON)
		return
	}

	handle := *h
	data, err2 := handle()
	if err2 != nil {
		writer.SendError(w, err2, indentJSON)
		return
	}
	selfLink := entity.GetSelfLink()
	w.Header().Add("Location", selfLink)
	writer.SendJSONResponse(w, http.StatusOK, data, nil, indentJSON)
}