package studentRepository

import (
	"fmt"
	"go-jwt-access-refresh-token/databaseConnection"
	"testing"
)

func Test(t *testing.T) {
	
	db, err := databaseConnection.NewDatabaseConnection().OracleConnection()
	if err != nil {
		panic(err)
	}

	newRepo, err:= NewStudentRepository(db).AuthenticateBachelor("xxx", "xxx","xxx")
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด",err)
		panic(err)
	}

	fmt.Println("ข้อมูล",newRepo)

}
