package instance

import (
	instancemanager "backend/internal/features/instance-manager"
	"backend/internal/templates/disk"
	"backend/internal/templates/provision"
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

	diskTemplate := disk.QemuImg{}
	diskCommnad, dir_output := diskTemplate.CreateDisk(instance.Name, 20)

	manager := instancemanager.InstanceManager{}

	if err := manager.CreateDiskFromTemplate(diskCommnad); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := instance.Init(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	instance.Disk = dir_output

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

func ExecuteProvision(w http.ResponseWriter, r *http.Request) {
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

	template := provision.VirtInstallTemplate{}
	command := template.GenerateTemplate(instance.Name, instance.RAM, instance.Disk, instance.Vcpus, instance.OsVariant, instance.Console, instance.Location, instance.ExtraArgs)
	manager := instancemanager.InstanceManager{}

	if err := manager.ProvisionInstanceFromTemplate(command); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, "Initializing Instance")
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
	instance, err := repository.GetByID(instanceID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	manager := instancemanager.InstanceManager{}
	template := provision.VirtInstallTemplate{}
	command := template.DeleteInstanceFromTemplate(instance.Name)

	if err := manager.DeleteFromTemplate(command); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = repository.Delete(instanceID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
