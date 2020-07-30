package main

import (
	"net/http"
	"webservice/database"
	"webservice/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.SetupDataBase()
	handlers.RegisterHandlers()
	http.ListenAndServe(":7070", nil)
}
