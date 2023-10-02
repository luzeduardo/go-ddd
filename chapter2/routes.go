package chapter2

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func router(u UserHandler) {
	m := mux.NewRouter()
	//exposing a simple endpoint that allows another team
	//to extract info within our module as a Open Host Service example
	m.HandleFunc("/user/{userID}/subscription/active", func(w http.ResponseWriter, r *http.Request) {
		//TODO check auth

		uID := mux.Vars(r)["userID"]
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		isActive := u.IsUserSubscriptionActive(r.Context(), uID)
		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(b)
	}).Methods(http.MethodGet)
}
