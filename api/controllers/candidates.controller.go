package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/semicolon27/api-e-voting/api/auth"
	"github.com/semicolon27/api-e-voting/api/models"
	"github.com/semicolon27/api-e-voting/api/responses"
	"github.com/semicolon27/api-e-voting/api/utils/formaterror"
)

func (server *Server) CreateCandidate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	candidate := models.Candidate{}
	err = json.Unmarshal(body, &candidate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	candidate.Prepare()
	err = candidate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = auth.TokenAdminValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	candidateCreated, err := candidate.SaveCandidate(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, candidateCreated.Id))
	responses.JSON(w, http.StatusCreated, candidateCreated)
}

func (server *Server) GetCandidates(w http.ResponseWriter, r *http.Request) {

	candidate := models.Candidate{}

	candidates, err := candidate.FindAllCandidates(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, candidates)
}

func (server *Server) GetCandidate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	candidate := models.Candidate{}

	candidateReceived, err := candidate.FindCandidateByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, candidateReceived)
}

func (server *Server) UpdateCandidate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the candidate id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	err = auth.TokenAdminValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the candidate exist
	candidate := models.Candidate{}
	err = server.DB.Debug().Model(models.Candidate{}).Where("id = ?", pid).Take(&candidate).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Candidate not found"))
		return
	}

	// Read the data candidateed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	candidateUpdate := models.Candidate{}
	err = json.Unmarshal(body, &candidateUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	candidateUpdate.Prepare()
	err = candidateUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	candidateUpdate.Id = candidate.Id //this is important to tell the model the candidate id to update, the other update field are set above

	candidateUpdated, err := candidateUpdate.UpdateCandidate(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, candidateUpdated)
}

func (server *Server) DeleteCandidate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid candidate id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	err = auth.TokenAdminValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the candidate exist
	candidate := models.Candidate{}
	err = server.DB.Debug().Model(models.Candidate{}).Where("id = ?", pid).Take(&candidate).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = candidate.DeleteCandidate(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
