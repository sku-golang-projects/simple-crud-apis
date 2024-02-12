package main

/*
// API naming rule
//https://medium.com/@nadinCodeHat/rest-api-naming-conventions-and-best-practices-1c4e781eb6a5
*/
import (
	"log"
	"net/http"

	api "SIMPLE_CRUD_APIS/src/api"
	model "SIMPLE_CRUD_APIS/src/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func simulateDeviceData() {
	//Add item
	model.Devices = append(model.Devices, model.DeviceList{DeviceId: uuid.New().String(),
		DeviceName: "device01",
		Applications: []model.ApplicationInfo{{ApplicationId: uuid.New().String(),
			ApplicationName: "app-01"}},
		AvailableMemory: "0",
		TotalMemory:     "0",
		Status:          "Active",
	})

	//Add item
	model.Devices = append(model.Devices, model.DeviceList{DeviceId: uuid.New().String(),
		DeviceName: "device02",
		Applications: []model.ApplicationInfo{{ApplicationId: uuid.New().String(),
			ApplicationName: "app-01"}},
		AvailableMemory: "0",
		TotalMemory:     "0",
		Status:          "Active",
	})
}

func main() {
	log.Println("Starting application...")

	simulateDeviceData()

	r := mux.NewRouter()

	pathV1 := r.PathPrefix("/api/v1").Subrouter()
	pathV1.HandleFunc("/devices", api.GetAllDevices).Methods("GET")
	pathV1.HandleFunc("/devices/{id}", api.GetDevice).Methods("GET")
	pathV1.HandleFunc("/devices", api.AddDevice).Methods("POST")
	pathV1.HandleFunc("/devices/{id}", api.UpdateDevice).Methods("PUT")
	pathV1.HandleFunc("/devices/{id}", api.DeleteDevice).Methods("DELETE")

	//Handling Future API's
	pathV2 := r.PathPrefix("/api/v2").Subrouter()
	pathV2.HandleFunc("/devices/{id}", api.DeleteAndListRemainingDevice).Methods("DELETE")

	log.Println("Running the server into the port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
