package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"
)

var router *mux.Router
var db *sql.DB


func main() {

	//database.Initialize()
	//db = database.DB

	bootstrap.SetUpDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
