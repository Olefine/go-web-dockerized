package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/olefine/go-dockerized-web/dao"
)

// Persons is handler returns persons in JSON format
func Persons(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(dao.GetAll())

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
