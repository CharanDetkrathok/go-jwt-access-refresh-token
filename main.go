package main

import (
	"fmt"
	"go-jwt-access-refresh-token/databaseConnection"
	"go-jwt-access-refresh-token/environment"
	"go-jwt-access-refresh-token/middleware"
	"go-jwt-access-refresh-token/repository/studentRepository"
	"go-jwt-access-refresh-token/timeZone"

	"github.com/gin-gonic/gin"
)

// var ctx = context.Background()

func main() {

	// ตั้งค่าช่วงเวลาให้เป็น Local
	timeZone.Init()

	// เรียกใช้ทรัพยกรที่กำหนดไว้
	environment.Init()

	// connect redis cache
	rdb := databaseConnection.NewDatabaseConnection().RedisConnection()
	// connect oracle database
	db, err := databaseConnection.NewDatabaseConnection().OracleConnection()
	if err != nil {
		panic(err)
	}

	newRepo, err:= studentRepository.NewStudentRepository(db).Authenticate("6256000792", "11/1/2534")
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด",err)
		panic(err)
	}

	fmt.Println("ข้อมูล",newRepo)

	// set release mode
	// using env:   export GIN_MODE=release
	gin.SetMode(gin.ReleaseMode)

	// เรียก function พื้นฐานของ gin-gonic
	router := gin.Default()

	// Access-Control-Allow
	router.Use(middleware.NewCorsMiddlewrerAccessControll().CorsMiddlewrerAccessControll())

	// Authentication and Authorization
	auth := router.Group("/student/auth")
	{
		// studentRepo := studentRepository
		auth.POST("/authentication")
		auth.POST("/authorization")
		auth.POST("/refresh_authentication")
	}



	fmt.Println(db.MustBegin().DriverName())
	fmt.Println(rdb)

	

}