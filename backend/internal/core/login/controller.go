package login

import (
	"backend/internal/core/user"
	auth "backend/pkg/Auth"
	"backend/pkg/database"
	"backend/pkg/http/response"
	"backend/pkg/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var userBody user.User
	if err = json.Unmarshal(request, &userBody); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := user.NewUserRepository(db)
	userInDatabase, err := repository.GetByEmail(userBody.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPassword(userInDatabase.Password, userBody.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.GenerateToken(userInDatabase.ID)
	w.Write([]byte("token: " + token))
}
