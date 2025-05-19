package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Car struct { // 注意结构体名大写，公开
	Ind                string  `json:"ind"`
	CarName            string  `json:"carName"`
	DateOfRegistration string  `json:"date_of_registration"`
	CurrentPrice       float64 `json:"CurrentPrice"`
	Mileage            float64 `json:"mileage"`
	SaleType           int     `json:"saleType"`
	Cover              string  `json:"cover"`
}

type ResponseData struct {
	TotalElements int   `json:"totalElements"`
	Pages         int   `json:"pages"`
	Content       []Car `json:"content"`
}

type ApiResponse struct {
	Status  int          `json:"status"`
	Success bool         `json:"success"`
	Msg     string       `json:"msg"`
	Data    ResponseData `json:"data"` // 修改为单个ResponseData
	Time    string       `json:"time"`
}
type Series struct {
	Id      int
	BrandId int
	Series  string
}

type CarDetailInfo struct {
	Ind                  string
	CarName              string
	CarId                string
	SeriesId             int
	BrandId              int
	Brandurl             string
	CarModel             string
	CarStyle             string
	DateOfRegistration   string
	CurrentPrice         float64
	CarColor             string
	GearboxType          string
	FuelType             string
	EmissionStandardType string
	Abbreviation         string
	DriveType            string
	EngineType           string
	ModelOfProduction    string
	Mileage              float64
	Seat                 int
	Displacement         string
	CurrentProvince      string
	CurrentCity          string
	SaleType             int
	CarBrandMaker        string
	Cover                string
	ParameterVo          ParameterVo
	PhotoVo              PhotoVo
	PriceVo              PriceVo
	ReportAssessVo       string
}

type CarDetailInfoCopy struct {
	Ind                  string
	CarName              string
	CarId                string
	SeriesId             int
	BrandId              int
	Brandurl             string
	CarModel             string
	CarStyle             string
	DateOfRegistration   string
	CurrentPrice         float64
	CarColor             string
	GearboxType          string
	FuelType             string
	EmissionStandardType string
	Abbreviation         string
	DriveType            string
	EngineType           string
	ModelOfProduction    string
	Mileage              float64
	Seat                 int
	Displacement         string
	CurrentProvince      string
	CurrentCity          string
	SaleType             int
	CarBrandMaker        string
	Cover                string
	ParameterVo          string
	PhotoVo              string
	PriceVo              string
	ReportAssessVo       string
}

type PhotoVo struct {
	CarId    string
	Carousel string
	Detail   string
}

type ParameterVo struct {
	CarID            string
	CarColor         string
	Engine           string
	GearBox          string
	Fuel             string
	Drive            string
	EmissionStandard string
}
type PriceVo struct {
	CarId            string
	CaHallPrice      float64
	CarDownPay       float64
	MortgageAmount   float64
	MortgageNum      float64
	MonthlyRepayment float64
}

// InsertDataToDB 执行插入函数
func InsertDataToDB(id int) {
	var CarArray []Car
	CarArray, err := GetCarDataFromUrl("https://api.tf2sc.cn/api/tfcar/car/list", id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//fmt.Println(CarArray)

	// 连接数据库准备执行
	db := DBConnect("tengfadata")
	defer db.Close()

	// 写插入数据的语句
	sqlStatement := `
		INSERT IGNORE into carlist (ind,carName,DateOfRegistration,CurrentPric,mileage,saleType,cover)
		VALUES (?,?,?,?,?,?,?,?)
	`

	// 执行插入数据
	for _, car := range CarArray {
		_, err := db.Exec(sqlStatement, car.Ind, car.CarName, car.DateOfRegistration, car.CurrentPrice, car.Mileage, car.SaleType, car.Cover)
		if err != nil {
			fmt.Println("插入数据错误")
			log.Fatal(err)
		}
	}
}

// GetCarDataFromUrl 从URL中获取数据
func GetCarDataFromUrl(url string, id int) ([]Car, error) {
	// 创建一个car[]的实体
	var Cars []Car

	// 循环获取数据
	for i := 1; i <= id; i++ {
		// 构造正确的URL，这里将id作为查询参数
		url := fmt.Sprintf("%s?carModel=%d", url, i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return nil, err
		}
		req.Header.Set("Platformtype", "h5")
		req.Header.Set("User-Agent", "PostmanRuntime/7.39.0")

		// 原生库的http.Get函数无法添加Headers
		//resp, err := http.Get(url)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return nil, err
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}
		respBodyNew := string(respBody)
		var apiResp ApiResponse
		err = json.Unmarshal([]byte(respBodyNew), &apiResp)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}

		// 在空白数组中拼接收到的数据
		for _, car := range apiResp.Data.Content { // 注意字段名首字母大写
			Cars = append(Cars, car)
		}

	}
	//fmt.Println(Cars)
	return Cars, nil
}
