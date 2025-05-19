package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/url"
	"strconv"
)

func SendCarData(c *gin.Context) {
	allQueryParams := c.Request.URL.Query()
	allQueryParams.Get("page") // 另一种获取单个参数的方式

	// 捕获参数
	queryParam := ""
	for key, values := range allQueryParams {
		// values 是一个字符串切片
		if len(values) > 0 {
			queryParam += fmt.Sprintf("%s=%s&", key, values[0])
		}
	}

	db := DBConnect("tengfadata")
	carList, err := QueryData(db, queryParam)
	if err != nil {
		c.JSON(200, gin.H{
			"Golang_study": 500,
			"msg":          "查询失败",
		})
	}

	c.JSON(200, gin.H{
		"carList": carList,
	})
}
func SendSeriesData(c *gin.Context) {
	fmt.Println("-------------------------------------------")
	brandIdStr := c.Query("brandId") // Use c.Query instead of c.GetQuery for simplicity
	if brandIdStr == "" {
		c.JSON(400, gin.H{
			"Golang_study": 400,
			"msg":          "Missing 'brandId' query parameter",
		})
		return
	}

	brandId, err := strconv.Atoi(brandIdStr)
	if err != nil {
		c.JSON(400, gin.H{
			"Golang_study": 400,
			"msg":          "Invalid 'brandId', must be an integer",
		})
		return
	}

	db := DBConnect("tengfadata")
	seriesList, err := QuerySeries(db, brandId) // Ensure brandId is converted back to string if needed
	if err != nil {
		c.JSON(500, gin.H{
			"Golang_study": 500,
			"msg":          "Failed to fetch series data",
			"err":          err.Error(),
		})
		return
	}
	fmt.Println(seriesList)

	c.JSON(200, gin.H{
		"seriesList": seriesList,
	})
}
func SendPredictPrice(c *gin.Context) {
	//Params := c.Request.URL.Query()
	// 10到100  范围内生成随机数
	predictPrice := rand.Intn(90) + 10
	c.JSON(200, gin.H{
		"predictPrice": predictPrice,
	})
}
func SendDetailInfo(c *gin.Context) {
	var carDetail CarDetailInfo
	carId := c.Query("carId")
	db := DBConnect("tengfadata")
	defer db.Close()
	carDetail, err := QueryDetail(db, carId)
	if err != nil {
		c.JSON(500, gin.H{})
	}

	c.JSON(200, gin.H{
		"carDetail": carDetail,
	})
}

func SendSearchData(c *gin.Context) {
	QueryParams := c.Query("info")
	// 对QueryParams进行url解码
	QueryParams, err := url.QueryUnescape(QueryParams)
	fmt.Println(QueryParams)
	if err != nil {
		c.JSON(500, gin.H{})
	}
	db := DBConnect("tengfadata")
	defer db.Close()
	carList, err := SearchData(db, QueryParams)
	if err != nil {
		c.JSON(500, gin.H{})
	}
	c.JSON(200, gin.H{
		"carList": carList,
	})
}
