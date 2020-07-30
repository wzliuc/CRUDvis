package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"webservice/handlers/logger"
	"webservice/handlers/middlewares"
)

// RegisterHandlers registers handlers for the default ServerMux
func RegisterHandlers() {
	logger.LogInfo("Registering handlers...")
	CategoryHandler := http.HandlerFunc(HandleCategory)
	CategoryProductHandler := http.HandlerFunc(HandleProductWithCategory)
	ProductsHandler := http.HandlerFunc(HandleProducts)
	ProductHandler := http.HandlerFunc(HandleProduct)

	http.Handle("/categories", middlewares.CorsPolicyMiddleware(CategoryHandler))
	http.Handle("/categories/", middlewares.CorsPolicyMiddleware(CategoryProductHandler))
	http.Handle("/products", middlewares.CorsPolicyMiddleware(ProductsHandler))
	http.Handle("/products/", middlewares.CorsPolicyMiddleware(ProductHandler))
	logger.LogInfo("Handlers registration complete.")
}

// getID retrives last element from routing path
func getID(r *http.Request) int {
	pathSlice := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSlice[len(pathSlice)-1])
	if err != nil {
		logger.LogErr("Error when mapping route parameter")
	}

	return id
}

