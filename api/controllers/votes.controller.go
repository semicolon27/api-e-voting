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

func (server *Server) CreateVote(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vote := models.Vote{}
	err = json.Unmarshal(body, &vote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vote.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = auth.TokenValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	voteCreated, err := vote.SaveVote(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, voteCreated.Id))
	responses.JSON(w, http.StatusCreated, voteCreated)
}

func (server *Server) GetVotes(w http.ResponseWriter, r *http.Request) {

	vote := models.Vote{}

	votes, err := vote.FindAllVotes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, votes)
}

func (server *Server) GetVote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	vote := models.Vote{}

	voteReceived, err := vote.FindVoteByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, voteReceived)
}
