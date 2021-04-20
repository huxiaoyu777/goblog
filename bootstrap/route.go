package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/routes"
	"time"
)

// 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}

func SetUpDB()  {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
