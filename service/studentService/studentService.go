// Bussiness Logic
package studentService

import (
	"fmt"
	"go-jwt-access-refresh-token/middleware"
	"net/http"
)

func (s studentService) Authenticate(std_code string, birth_date string) (*StudentAuthenticationServiceResponse, error) {

	// ส่ง std_code และ birth_date ไปตรวจสอบใน Database
	student, err := s.studentRepository.Authenticate(std_code, birth_date)
	if err != nil {

		// ถ้าไม่พบข้อมูล ทำการเตรียมโครงสร้างข้อมูลแบบว่างเปล่ากลับไป
		studentResponseInfomation := studentInfomation{}
		studentResponse := StudentAuthenticationServiceResponse{
			studentResponseInfomation,
			studentToken{AccessToken: "", RefreshToken: "", Authorized: ""},
			fmt.Sprint(http.StatusNoContent),
			"รหัสนักษึกษาหรือ วัน/เดือน/ปีเกิด ไม่ถูกต้อง!",
		}
		return &studentResponse, err

	}

	generateToken, err := middleware.GenerateToken(student.Lev_id, student.Std_code, fmt.Sprint(" - "+student.First_name_thai+" - "+student.First_name_eng))
	if err != nil {
		return nil, err
	}

	studentResponseGenerateToken := studentToken{
		AccessToken:         generateToken.AccessToken,
		RefreshToken:        generateToken.RefreshToken,
		ExpiresAccessToken:  generateToken.ExpiresAccessToken,
		ExpiresRefreshToken: generateToken.ExpiresRefreshToken,
		AccessTokenUUID:     generateToken.AccessTokenUUID,
		RefreshTokenUUID:    generateToken.RefreshTokenUUID,
	}

	// เก็บข้อมูลที่ได้จากการ Query เพื่อเตรียม Respose
	studentResponseInfomation := studentInfomation{
		Std_code:          student.Std_code,
		Prename_no:        student.Prename_no,
		Prename_thai:      student.Prename_thai,
		Prename_eng:       student.Prename_eng,
		First_name_thai:   student.First_name_thai,
		First_name_eng:    student.First_name_eng,
		Last_name_thai:    student.Last_name_thai,
		Last_name_eng:     student.Last_name_eng,
		Birth_date:        student.Birth_date,
		Year_end:          student.Year_end,
		Faculty_no:        student.Faculty_no,
		Faculty_name_thai: student.Faculty_name_thai,
		Faculty_name_eng:  student.Faculty_name_eng,
		Curr_no:           student.Curr_no,
		Major_no:          student.Major_no,
		Major_flag:        student.Major_flag,
		Major_name_thai:   student.Major_name_thai,
		Major_name_eng:    student.Major_name_eng,
		Lev_id:            student.Lev_id,
	}

	// เตรียม Infomation และ Token สำหรับ Authorization ของนักศึกษา
	studentResponse := StudentAuthenticationServiceResponse{
		studentResponseInfomation,
		studentToken{AccessToken: studentResponseGenerateToken.AccessToken, RefreshToken: studentResponseGenerateToken.RefreshToken,Authorized: studentResponseGenerateToken.Authorized},
		fmt.Sprint(http.StatusCreated),
		"Created tokens successfully",
	}

	return &studentResponse, nil
}
