package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	router := gin.Default()

	// 设置 CORS 中间件
	corsMiddleware := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://118.178.120.11:5173"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	// 将 corsMiddleware 包装成 gin.HandlerFunc
	router.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
	})
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.GET("/query", SendCarData)
	router.GET("/series", SendSeriesData)
	router.GET("/predictPrice", SendPredictPrice)
	router.GET("/detailInformation", SendDetailInfo)
	router.GET("/search", SendSearchData)

	router.Run(":8080")
}
