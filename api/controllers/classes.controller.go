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

func (server *Server) CreateClass(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	class := models.Class{}
	err = json.Unmarshal(body, &class)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	class.Prepare()
	err = class.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = auth.TokenValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	classCreated, err := class.SaveClass(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, classCreated.Id))
	responses.JSON(w, http.StatusCreated, classCreated)
}

func (server *Server) GetClasses(w http.ResponseWriter, r *http.Request) {

	class := models.Class{}

	classs, err := class.FindAllClasses(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, classs)
}

func (server *Server) GetClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	class := models.Class{}

	classReceived, err := class.FindClassByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, classReceived)
}

func (server *Server) UpdateClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the class id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	err = auth.TokenValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the class exist
	class := models.Class{}
	err = server.DB.Debug().Model(models.Class{}).Where("id = ?", pid).Take(&class).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Class not found"))
		return
	}

	// Read the data classed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	classUpdate := models.Class{}
	err = json.Unmarshal(body, &classUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	classUpdate.Prepare()
	err = classUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	classUpdate.Id = class.Id //this is important to tell the model the class id to update, the other update field are set above

	classUpdated, err := classUpdate.UpdateClass(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, classUpdated)
}

func (server *Server) DeleteClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid class id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	err = auth.TokenValid(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the class exist
	class := models.Class{}
	err = server.DB.Debug().Model(models.Class{}).Where("id = ?", pid).Take(&class).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = class.DeleteClass(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
