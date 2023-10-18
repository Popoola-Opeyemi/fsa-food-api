package api

import (
	"encoding/json"
	"fsa-food-api/client"
	"fsa-food-api/helpers"
	"fsa-food-api/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
}

func (s Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api", s.getAuthorities)
	router.HandleFunc("/api/{id}", s.getAuthority)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(srv.ListenAndServe(), router)
}

func (s Server) getAuthorities(w http.ResponseWriter, r *http.Request) {

	var authorities []model.Authority
	fsaAuthorities, err := client.GetAuthorities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, authority := range fsaAuthorities.Authorities {
		authorities = append(authorities, model.Authority{
			ID:   authority.ID,
			Name: authority.Name,
		})
	}

	data, _ := json.Marshal(authorities)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s Server) getAuthority(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid authorityID", http.StatusBadRequest)
		return
	}

	resp, err := client.GetLocalAuthority(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d, count, err := helpers.ProcessData(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	percentage, err := helpers.GetPercentages(d, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(percentage)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
