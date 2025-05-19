package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//   Exec 用于执行那些不返回结果集的SQL语句，如INSERT、UPDATE、DELETE等

// DBConnect 连接数据库
func DBConnect(DBName string) *sql.DB {

	dsn := "root:rootROOT1234.@tcp(127.0.0.1:3306)/" + DBName //"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	maxOpenConnections := 100
	db.SetMaxOpenConns(maxOpenConnections) // 正在连接的数量
	db.SetMaxIdleConns(maxOpenConnections) // 空闲的连接数量

	// 验证连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	return db
}

// QueryData 默认查询100条数据
func QueryData(db *sql.DB, queryParam string) ([]Car, error) {
	var Cars []Car      // 返回对象
	var pageId int      // 查询的数据范围
	var sqlParam string // 查询的参数
	var sqlQuery string // 查询语句
	// 如果param为空，直接查
	if len(queryParam) != 0 {
		// 去除末尾的&
		queryParam, _ = strings.CutSuffix(queryParam, "&")
		// 使用正则表达式对page进行处理
		if strings.Contains(queryParam, "page=") {
			// 编译一个正则表达式，匹配 "page=" 后面跟着一个或多个数字（\d+）
			re := regexp.MustCompile(`page=(\d+)`)
			matches := re.FindStringSubmatch(queryParam)
			// 使用正则表达式的 FindStringSubmatch 方法查找匹配项
			pageId, _ = strconv.Atoi(matches[1])
			fmt.Println("pageId:", pageId)
			// 使用正则表达式的 ReplaceAllString 方法移除匹配到的子串，第二个参数为空字符串表示移除匹配到的内容
			queryParam = re.ReplaceAllString(queryParam, "")
			queryParam, _ = strings.CutSuffix(queryParam, "&")
			queryParam, _ = strings.CutPrefix(queryParam, "&")
			sqlParam = strings.ReplaceAll(queryParam, "&", " and ")
			sqlQuery = fmt.Sprintf("select ind,carName,DateOfRegistration,currentPrice,mileage,saleType,cover from car_info_data where %s limit %d OFFSET %d;", sqlParam, 100*pageId, 100*(pageId-1))
		} else {
			//queryParam, _ = strings.CutSuffix(queryParam, "&")
			queryParam, _ = strings.CutPrefix(queryParam, "&")
			sqlParam = strings.ReplaceAll(queryParam, "&", " and ")
			sqlQuery = fmt.Sprintf("select ind,carName,DateOfRegistration,currentPrice,mileage,saleType,cover from car_info_data where %s limit 100", sqlParam)
		}
	} else {
		sqlQuery = "select ind,carName,DateOfRegistration,currentPrice,mileage,saleType,cover from car_info_data limit 100"
	}

	result, err := db.Query(sqlQuery)
	fmt.Println(sqlQuery)
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}
	defer result.Close()

	for result.Next() {
		var car Car
		err := result.Scan(&car.Ind, &car.CarName, &car.DateOfRegistration, &car.CurrentPrice, &car.Mileage, &car.SaleType, &car.Cover)
		if err != nil {
			log.Fatal("绑定数据失败:", err)
		}
		Cars = append(Cars, car)
	}
	fmt.Println("\n", sqlQuery, "\n")
	return Cars, nil
}

func QuerySeries(db *sql.DB, queryParam int) ([]Series, error) {
	var series []Series
	sqlQuery := fmt.Sprintf("select  DISTINCT seriesId,brandId,series from car_info where BrandID = %d;", queryParam)
	fmt.Println(sqlQuery)
	result, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}
	defer result.Close()
	for result.Next() {
		var s Series
		err := result.Scan(&s.Id, &s.BrandId, &s.Series)
		if err != nil {
			log.Fatal("绑定数据失败:", err)
		}
		series = append(series, s)

	}
	return series, nil
}

func QueryDetail(db *sql.DB, queryParam string) (CarDetailInfo, error) {
	var car CarDetailInfoCopy
	var ReallyCar CarDetailInfo
	sqlQuery := fmt.Sprintf("select * from car_info_data where ind = '%s' limit 1;", queryParam)
	fmt.Println(sqlQuery)
	result, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&car.Ind, &car.CarId, &car.CarName, &car.BrandId, &car.SeriesId, &car.CarModel, &car.CarStyle, &car.DateOfRegistration, &car.CurrentPrice,
			&car.CarColor, &car.GearboxType, &car.FuelType, &car.EmissionStandardType, &car.DriveType, &car.EngineType, &car.ModelOfProduction, &car.Mileage, &car.Seat,
			&car.Displacement, &car.CurrentProvince, &car.CurrentCity, &car.SaleType, &car.CarBrandMaker, &car.Cover, &car.ParameterVo, &car.PhotoVo, &car.PriceVo, &car.ReportAssessVo)

		if err != nil {
			log.Fatal("绑定数据失败:", err)
		}
	}
	//car = strings.ReplaceAll(car,"'",""")
	var photoVo PhotoVo
	var priceVo PriceVo
	var parameterVo ParameterVo
	fmt.Println(car.PhotoVo)
	fmt.Println(car.PriceVo)
	fmt.Println(car.PriceVo)
	if err := json.Unmarshal([]byte(car.PhotoVo), &photoVo); err != nil {
		log.Printf("反序列化失败: %v", err)
	}
	if err := json.Unmarshal([]byte(car.PriceVo), &priceVo); err != nil {
		log.Printf("反序列化失败: %v", err)
	}
	if err := json.Unmarshal([]byte(car.ParameterVo), &parameterVo); err != nil {
		log.Printf("反序列化失败: %v", err)
	}
	fmt.Println(photoVo)
	fmt.Println(priceVo)
	fmt.Println(parameterVo)
	ReallyCar.Ind = car.Ind
	ReallyCar.CarId = car.CarId
	ReallyCar.CarName = car.CarName
	ReallyCar.BrandId = car.BrandId
	ReallyCar.SeriesId = car.SeriesId
	ReallyCar.CarModel = car.CarModel
	ReallyCar.CarStyle = car.CarStyle
	ReallyCar.DateOfRegistration = car.DateOfRegistration
	ReallyCar.CurrentPrice = car.CurrentPrice
	ReallyCar.CarColor = car.CarColor
	ReallyCar.GearboxType = car.GearboxType
	ReallyCar.FuelType = car.FuelType
	ReallyCar.EmissionStandardType = car.EmissionStandardType
	ReallyCar.DriveType = car.DriveType
	ReallyCar.EngineType = car.EngineType
	ReallyCar.ModelOfProduction = car.ModelOfProduction
	ReallyCar.Mileage = car.Mileage
	ReallyCar.Seat = car.Seat
	ReallyCar.Displacement = car.Displacement
	ReallyCar.CurrentProvince = car.CurrentProvince
	ReallyCar.CurrentCity = car.CurrentCity
	ReallyCar.SaleType = car.SaleType
	ReallyCar.CarBrandMaker = car.CarBrandMaker
	ReallyCar.Cover = car.Cover
	ReallyCar.PhotoVo = photoVo
	ReallyCar.PriceVo = priceVo
	ReallyCar.ParameterVo = parameterVo
	ReallyCar.ReportAssessVo = car.ReportAssessVo
	ReallyCar.Abbreviation = car.Abbreviation
	ReallyCar.Brandurl = car.Brandurl

	return ReallyCar, nil
}

func SearchData(db *sql.DB, queryParam string) ([]Car, error) {
	var Cars []Car
	// 根据queryParam进行模糊查询
	sqlQuery := fmt.Sprintf("select ind,carName,DateOfRegistration,currentPrice,mileage,saleType,cover from car_info_data where carName like '%%%s%%' limit 100;", queryParam)
	result, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var car Car
		err := result.Scan(&car.Ind, &car.CarName, &car.DateOfRegistration, &car.CurrentPrice, &car.Mileage, &car.SaleType, &car.Cover)
		if err != nil {
			return nil, err
		}
		Cars = append(Cars, car)
	}

	return Cars, nil
}
