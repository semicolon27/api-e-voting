package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/semicolon27/api-e-voting/api/auth"
	"github.com/semicolon27/api-e-voting/api/models"
	"github.com/semicolon27/api-e-voting/api/responses"
	"github.com/semicolon27/api-e-voting/api/utils/formaterror"
)

func (server *Server) CreateParticipant(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	participant := models.Participant{}
	err = json.Unmarshal(body, &participant)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	participant.PrepareParticipant()
	err = participant.ValidateParticipant("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	participantCreated, err := participant.SaveParticipant(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, participantCreated.Id))
	responses.JSON(w, http.StatusCreated, participantCreated)
}

func (server *Server) GetParticipants(w http.ResponseWriter, r *http.Request) {

	participant := models.Participant{}

	participants, err := participant.FindAllParticipants(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, participants)
}

func (server *Server) GetParticipant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := fmt.Sprintf(vars["id"])
	participant := models.Participant{}

	participantGotten, err := participant.FindParticipantByID(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, participantGotten)
}

func (server *Server) UpdateParticipant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := fmt.Sprintf(vars["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	participant := models.Participant{}
	err = json.Unmarshal(body, &participant)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = auth.TokenAdminValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	participant.PrepareParticipant()
	err = participant.ValidateParticipant("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedParticipant, err := participant.UpdateParticipant(server.DB, uid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedParticipant)
}

func (server *Server) DeleteParticipant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	participant := models.Participant{}

	uid := fmt.Sprintf(vars["id"])
	err := auth.TokenAdminValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = participant.DeleteParticipant(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
