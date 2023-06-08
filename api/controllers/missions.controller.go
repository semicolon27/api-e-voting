package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/semicolon27/api-e-voting/api/models"
	"github.com/semicolon27/api-e-voting/api/responses"
	"github.com/semicolon27/api-e-voting/api/utils/formaterror"
)

func (server *Server) CreateMission(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	mission := models.Mission{}
	err = json.Unmarshal(body, &mission)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	mission.Prepare()
	err = mission.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	missionCreated, err := mission.SaveMission(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, missionCreated.Id))
	responses.JSON(w, http.StatusCreated, missionCreated)
}

func (server *Server) GetMissions(w http.ResponseWriter, r *http.Request) {

	mission := models.Mission{}

	missions, err := mission.FindAllMissions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, missions)
}

func (server *Server) GetMission(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	mission := models.Mission{}

	missionReceived, err := mission.FindMissionByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, missionReceived)
}

func (server *Server) UpdateMission(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the mission id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it

	// Check if the mission exist
	mission := models.Mission{}
	err = server.DB.Debug().Model(models.Mission{}).Where("id = ?", pid).Take(&mission).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Mission not found"))
		return
	}

	// Read the data missioned
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	missionUpdate := models.Mission{}
	err = json.Unmarshal(body, &missionUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	missionUpdate.Prepare()
	err = missionUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	missionUpdate.Id = mission.Id //this is important to tell the model the mission id to update, the other update field are set above

	missionUpdated, err := missionUpdate.UpdateMission(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, missionUpdated)
}

func (server *Server) DeleteMission(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid mission id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?

	// Check if the mission exist
	mission := models.Mission{}
	err = server.DB.Debug().Model(models.Mission{}).Where("id = ?", pid).Take(&mission).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = mission.DeleteMission(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
