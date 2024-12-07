package main

import (
	"github.com/gin-gonic/gin"
	"my-gin-app/internal/db"
	"my-gin-app/internal/routes"
)

func main() {
	// データベースの初期化
	database := db.InitDB()

	// Ginのルーターを作成
	r := gin.Default()

	// サインアップルートを登録
	routes.AuthRoutes(r, database)

	r.Run(":8080") // サーバーを8080ポートで起動
}
