package api_handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/querybuilder/model"
	"github.com/querybuilder/service"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/fetchPlayers", PostHandler).Methods("POST")
	http.Handle("/", r)

	fmt.Println("Server is starting on port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var data model.QueryRequest

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Received: %+v\n", data)

		response := service.Serve(data)

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, "Failed to write the response", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}
