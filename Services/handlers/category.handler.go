package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webservice/database"
	"webservice/handlers/logger"
	"webservice/models"
)

// HandleCategory handles the requests for category
func HandleCategory(w http.ResponseWriter, r *http.Request) {
	repo := database.NewCategoryRepo()

	switch r.Method {
	case http.MethodGet:
		category := repo.GetAll()
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		err := enc.Encode(category)
		if err != nil {
			logger.LogErr(err)
		}
		return
	case http.MethodPost:
		var category models.Category
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&category)
		if err != nil {
			logger.LogErr(err)
		}

		fmt.Printf("Decoded category: %v", category)
		repo.Add(category)
		return
	}
}

// HandleProductWithCategory handles the requests for product of specified categoroy
func HandleProductWithCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		repo := database.NewCategoryRepo()
		products := repo.GetCategory(getID(r))
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		err := enc.Encode(products)
		if err != nil {
			logger.LogErr(err)
		}
		return
	}
}
