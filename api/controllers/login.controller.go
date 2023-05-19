package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/semicolon27/api-e-voting/api/auth"
	"github.com/semicolon27/api-e-voting/api/models"
	"github.com/semicolon27/api-e-voting/api/responses"
	"github.com/semicolon27/api-e-voting/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
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
	token, err := server.SignIn(admin.Username, admin.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, password string) (string, error) {

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
	return auth.CreateToken(uint32(admin.Id))
}
