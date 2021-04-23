package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/routes"
	"gorm.io/gorm"
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

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB)  {

	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
