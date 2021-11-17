package main

import (
	"go-jwt-access-refresh-token/databaseConnection"
	"go-jwt-access-refresh-token/environment"
	"go-jwt-access-refresh-token/handler/studentHandler"
	"go-jwt-access-refresh-token/middleware"
	"go-jwt-access-refresh-token/repository/studentRepository"
	"go-jwt-access-refresh-token/service/studentService"
	"go-jwt-access-refresh-token/timeZone"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// var ctx = context.Background()

func main() {

	timeZone.Init()    // ตั้งค่าช่วงเวลาให้เป็น Local
	environment.Init() // เรียกใช้ทรัพยกรที่กำหนดไว้
	db, err := databaseConnection.NewDatabaseConnection().OracleConnection() 
	if err != nil {
		panic(err)
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode) // set release mode using env:   export GIN_MODE=release
	router := gin.Default()  // เรียก function พื้นฐานของ gin-gonic
	router.Use(middleware.NewCorsMiddlewrerAccessControll().CorsMiddlewrerAccessControll()) // Access-Control-Allow

	newStudnetRepo := studentRepository.NewStudentRepository(db)
	newStudnetService := studentService.NewStudentService(newStudnetRepo)
	newStudnetHandler := studentHandler.NewStudentHandler(newStudnetService)

	// Authentication and Authorization
	auth := router.Group("/student/auth")
	{		
		auth.POST("/authentication", newStudnetHandler.Authentication)
		// auth.POST("/authorization")
		// auth.POST("/refresh_authentication")
	}

	router.Run(viper.GetString("survey.port"))

}
