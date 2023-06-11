package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/semicolon27/api-e-voting/api/auth"
	"github.com/semicolon27/api-e-voting/api/models"
	"github.com/semicolon27/api-e-voting/api/responses"
	"github.com/semicolon27/api-e-voting/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

// ADMIN SECTION
func (server *Server) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	admin := models.Admin{}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	admin.PrepareAdmin()
	err = admin.ValidateAdmin("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignInAdmin(admin.Username, admin.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnauthorized, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignInAdmin(username string, password string) (string, error) {

	var err error

	admin := models.Admin{}

	err = server.DB.Debug().Model(models.Admin{}).Where("username = ?", username).Take(&admin).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPasswordAdmin(admin.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateTokenAdmin(uint32(admin.Id), admin.Name, admin.Username)
}

// PARTICIPANT SECTION
func (server *Server) LoginParticipant(w http.ResponseWriter, r *http.Request) {
	log.Print("Login participant")
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

	participant.PrepareParticipant()
	err = participant.ValidateParticipant("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignInParticipant(participant.RegNumber, participant.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignInParticipant(RegNumber, password string) (string, error) {

	var err error

	participant := models.Participant{}

	err = server.DB.Debug().Model(models.Participant{}).Where("reg_number = ?", RegNumber).Take(&participant).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPasswordParticipant(participant.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateTokenParticipant(uint32(participant.Id), participant.RegNumber)
}
