package studentRepository

//6256000792 11/1/2534 รูปแบบ วว/ดด/ปปปป
func (r studentRepository) Authenticate(std_code string, birth_date string) (*StudentAuthenticationRepositoryFromDB, error){

	student := StudentAuthenticationRepositoryFromDB{}

	query := "SELECT STD_CODE,PRENAME_NO,PRENAME_THAI,PRENAME_ENG,FIRST_NAME_THAI,FIRST_NAME_ENG,LAST_NAME_THAI,LAST_NAME_ENG,BIRTH_DATE,YEAR_END,FACULTY_NO,FACULTY_NAME_THAI,FACULTY_NAME_ENG,CURR_NO,MAJOR_NO,MAJOR_FLAG,MAJOR_NAME_THAI,MAJOR_NAME_ENG,LEV_ID FROM SCENTER01.VM_STD_GRADUATE WHERE STD_CODE = :param1 AND BIRTH_DATE = TO_DATE(:param2,'DD/MM/YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI')"

	err := r.db.Get(&student,query,std_code,birth_date)
	if err != nil {
		return nil,err
	}

	return &student, nil
	
}
