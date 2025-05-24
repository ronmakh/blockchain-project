package api

import (
	"github.com/gin-gonic/gin"
	"blockchain-project/handler"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	  }))

	r.POST("/transaction", handler.HandleTransaction)
	r.GET("/mine", handler.HandleMine)
	r.GET("/blockchain", handler.GetBlockchain)
	r.GET("/wallets", handler.GetWallets)

	return r
}
