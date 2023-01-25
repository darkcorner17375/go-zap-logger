package main

import (
	"go-zap-logger/log/logger"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log1 := logger.New()
	log1.Error("log1:before submit")
	//新增project name
	log1.Config.SetProjectName("log2")
	//送出設定
	log1.SubmitConfig()
	log1.Error("log2:set project name")
	//將text改變json訊息
	log1.Config.SetJSONFormat(true)
	log1.SubmitConfig()
	log1.Error("log3: change to json")

	//新增gin引擎
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
	}), log1.GinLogger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
