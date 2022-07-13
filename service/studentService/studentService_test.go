package studentService

import (
	"fmt"
	"go-jwt-access-refresh-token/databaseConnection"
	"go-jwt-access-refresh-token/repository/studentRepository"
	"testing"
)

func Test(t *testing.T) {

	db, err := databaseConnection.NewDatabaseConnection().OracleConnection()
	if err != nil {
		panic(err)
	}

	newRepo := studentRepository.NewStudentRepository(db)
	newService := NewStudentService(newRepo)

	resultServiceResponse, err := newService.Authentication("xxx", "xxx","xxx")
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด", err)
		panic(err)
	}

	fmt.Println("ข้อมูล",resultServiceResponse)

}
