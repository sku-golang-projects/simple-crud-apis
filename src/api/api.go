package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	model "SIMPLE_CRUD_APIS/src/model"
)

func GetAllDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if model.Devices != nil {
		err := json.NewEncoder(w).Encode(&model.Devices)

		if err != nil {
			log.Printf("Error is:%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"Response\":\"Error in parsing JSON\"}")
			return
		}
	} else {
		log.Println("No data is available.")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{\"Response\":\"No data available.\"}"))
	}
}

func GetDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if len(params["id"]) < 36 {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte("Invalid input provided, does not seems to be valid ID."))
		return
	}

	log.Println("Requested device information for id:", params["id"])

	found := false
	// C++ Style
	// for index := 0; index < len(model.Devices); index++ {
	// 	log.Println(model.Devices[index].DeviceId)
	// 	if model.Devices[index].DeviceId == params["id"] {
	// 		found = true
	// 		json.NewEncoder(w).Encode(&model.Devices[index])
	// 		break
	// 	}
	// }

	for _, item := range model.Devices {
		if item.DeviceId == params["id"] {
			found = true
			json.NewEncoder(w).Encode(item)
			break
		}
	}

	if !found {
		log.Println("Unable to found the requested device id.")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{\"Response\":\"Unable to found the requested device id.\"}"))
	}
}

func AddDevice(w http.ResponseWriter, r *http.Request) {
	data := r.Body
	fmt.Println("Received data is:", data)
	var device model.DeviceRegistration
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte("Invalid input provided, please provide the valid data."))
		return
	}
	fmt.Println("Decoded data is:", device.DeviceName)

	deviceData := model.DeviceList{DeviceId: uuid.New().String(),
		DeviceName:      device.DeviceName,
		Applications:    []model.ApplicationInfo{{}},
		AvailableMemory: "0",
		TotalMemory:     "0",
		Status:          "Onboarding",
	}

	model.Devices = append(model.Devices, deviceData)
	w.Write([]byte("{\"Response\":\"Requested device is added successfully.\"}"))
}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {

}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	log.Println("Request to delete ID:", params["id"])

	if len(params["id"]) < 36 {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte("Invalid input provided, does't seems to be valid ID."))
		return
	}

	found := false
	for index, item := range model.Devices {
		if item.DeviceId == params["id"] {
			log.Println("Item found, will be deleted")
			// Reference
			// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
			model.Devices = append(model.Devices[:index], model.Devices[index+1:]...)
			w.Write([]byte("{\"Response\":\"Requested device id is deleted successfully.\"}"))
			return
		}
	}

	if !found {
		log.Println("Unable to found the requested device id.")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{\"Response\":\"Unable to found the requested device id.\"}"))
	}
}

func DeleteAndListRemainingDevice(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	log.Println("Request to delete ID:", params["id"])

	if len(params["id"]) < 36 {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte("Invalid input provided, does't seems to be valid ID."))
		return
	}

	found := false
	for index, item := range model.Devices {
		if item.DeviceId == params["id"] {
			log.Println("Item found, will be deleted")
			model.Devices = append(model.Devices[:index], model.Devices[index+1:]...)
			GetAllDevices(w, r)
			return
		}
	}

	if !found {
		log.Println("Unable to found the requested device id.")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{\"Response\":\"Unable to found the requested device id.\"}"))
	}
}
