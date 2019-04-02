package main

import (
	"log"
	"os"

	_ "github.com/button-tech/utils-node-tool/btc/docs"
	"github.com/button-tech/utils-node-tool/btc/handlers"
	"github.com/button-tech/utils-node-tool/btc/handlers/storage"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prazd/round-robin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {

	// @title Swagger BUTTON Node API
	// @version 1.0
	// @description This is BUTTON Node API responseModels documentation

	// @contact.name API Support
	// @contact.email nk
	// ap@buttonwallet.com

	// @license.name MIT
	// @license.url https://opensource.org/licenses/MIT

	// @BasePath /

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	gin.SetMode(gin.ReleaseMode)

	r.GET("/btc/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/btc/balance/:address", handlers.GetBalance)

	r.GET("/btc/utxo/:address", handlers.GetUTXO)

	r.GET("/btc/transactionFee", handlers.GetTxFee)

	r.GET("/btc/bestTransactionFee", handlers.GetBextTxFee)

	r.POST("/btc/balances", handlers.GetBalances)

	if err := r.Run(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
