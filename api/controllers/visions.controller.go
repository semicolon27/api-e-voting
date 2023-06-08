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

func (server *Server) CreateVision(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vision := models.Vision{}
	err = json.Unmarshal(body, &vision)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vision.Prepare()
	err = vision.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	visionCreated, err := vision.SaveVision(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, visionCreated.Id))
	responses.JSON(w, http.StatusCreated, visionCreated)
}

func (server *Server) GetVisions(w http.ResponseWriter, r *http.Request) {

	vision := models.Vision{}

	visions, err := vision.FindAllVisions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, visions)
}

func (server *Server) GetVision(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	vision := models.Vision{}

	visionReceived, err := vision.FindVisionByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, visionReceived)
}

func (server *Server) UpdateVision(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the vision id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it

	// Check if the vision exist
	vision := models.Vision{}
	err = server.DB.Debug().Model(models.Vision{}).Where("id = ?", pid).Take(&vision).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Vision not found"))
		return
	}

	// Read the data visioned
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	visionUpdate := models.Vision{}
	err = json.Unmarshal(body, &visionUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	visionUpdate.Prepare()
	err = visionUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	visionUpdate.Id = vision.Id //this is important to tell the model the vision id to update, the other update field are set above

	visionUpdated, err := visionUpdate.UpdateVision(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, visionUpdated)
}

func (server *Server) DeleteVision(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid vision id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?

	// Check if the vision exist
	vision := models.Vision{}
	err = server.DB.Debug().Model(models.Vision{}).Where("id = ?", pid).Take(&vision).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = vision.DeleteVision(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
