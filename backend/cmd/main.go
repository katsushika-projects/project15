package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moto340/project15/backend/internal/db"
	"github.com/moto340/project15/backend/internal/routes"
)

func main() {
	// データベースの初期化
	database := db.InitDB()

	// Ginのルーターを作成
	r := gin.Default()

	// サインアップルートを登録
	routes.RegisterRoutes(r, database)
	routes.AuthRoutes(r, database)
	routes.AdminRoutes(r, database)

	r.Run(":8080") // サーバーを8080ポートで起動
}
