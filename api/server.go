package api

import (
	"encoding/json"
	"fsa-food-api/client"
	"fsa-food-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct{}

func (s Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api", s.getAuthorities)
	router.HandleFunc("/api/{authorityID}", s.getAuthority)
	// router.HandleFunc("/api/localAuthorities/{localAuthorityID}/ratingPercentages",s.)

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
	authorityRating := []model.AuthorityRating{
		{Name: "5-star", Value: 22.41},
		{Name: "4-star", Value: 43.13},
		{Name: "3-star", Value: 12.97},
		{Name: "2-star", Value: 1.54},
		{Name: "1-star", Value: 17.84},
		{Name: "Exempt", Value: 2.11},
	}

	data, _ := json.Marshal(authorityRating)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s Server) getRatingPercentages(w http.ResponseWriter, r *http.Request) {

}
