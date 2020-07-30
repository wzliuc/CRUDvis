package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webservice/database"
	"webservice/handlers/logger"
	"webservice/models"
)

// HandleProducts handles the requests for product
func HandleProducts(w http.ResponseWriter, r *http.Request) {
	logger.LogInfo(fmt.Sprintf("Handle products %s", r.Method))
	repo := database.NewProductRepo()
	switch r.Method {
	case http.MethodGet:
		product := repo.GetAll()
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		err := enc.Encode(product)
		if err != nil {
			logger.LogErr(err)
		}
		return
	case http.MethodPost:
		var product models.Product
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&product)
		if err != nil {
			logger.LogErr(err)
		}
		repo.Add(product)
		return
	}
}

// HandleProduct handles the requests for product
func HandleProduct(w http.ResponseWriter, r *http.Request) {
	logger.LogInfo(fmt.Sprintf("Handle product %s", r.Method))
	repo := database.NewProductRepo()
	switch r.Method {
	case http.MethodGet:
		product := repo.Get(getID(r))
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		err := enc.Encode(product)
		if err != nil {
			logger.LogErr(err)
		}
		return
	case http.MethodPut:
		var product models.Product
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&product)
		if err != nil {
			logger.LogErr(err)
		}
		repo.Update(product)
		return
	case http.MethodDelete:
		pID := getID(r)
		repo.Delete(pID)
		return
	}
}
