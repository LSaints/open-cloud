package instance

import (
	"backend/pkg/database"
	"backend/pkg/http/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var instance Instance
	if err = json.Unmarshal(request, &instance); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := instance.Init(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := NewInstanceRepository(db)
	instance.ID, err = repository.Create(instance)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, instance)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	queryParams := strings.ToLower(r.URL.Query().Get("instance"))

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := NewInstanceRepository(db)
	instance, err := repository.GetAll(queryParams)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, instance)
}

func Get(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	instanceID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := NewInstanceRepository(db)
	instance, err := repository.GetByID(instanceID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, instance)
}

func Update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	instanceID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var instance Instance
	if err = json.Unmarshal(requestBody, &instance); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}
	if err = instance.Init(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	repository := NewInstanceRepository(db)
	if err = repository.Update(instanceID, instance); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	instanceID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := NewInstanceRepository(db)
	if err = repository.Delete(instanceID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
