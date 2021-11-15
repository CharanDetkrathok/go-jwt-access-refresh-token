package studentRepository

type Student struct {
	Std_code          string `db:"STD_CODE"`
	Prename_no        string `db:"PRENAME_NO"`
	Prename_thai      string `db:"PRENAME_THAI"`
	Prename_eng       string `db:"PRENAME_ENG"`
	First_name_thai   string `db:"FIRST_NAME_THAI"`
	First_name_eng    string `db:"FIRST_NAME_ENG"`
	Last_name_thai    string `db:"LAST_NAME_THAI"`
	Last_name_eng     string `db:"LAST_NAME_ENG"`
	Birth_date        string `db:"BIRTH_DATE"`
	Year_end          string `db:"YEAR_END"`
	Faculty_no        string `db:"FACULTY_NO"`
	Faculty_name_thai string `db:"FACULTY_NAME_THAI"`
	Faculty_name_eng  string `db:"FACULTY_NAME_ENG"`
	Curr_no           string `db:"CURR_NO"`
	Major_no          string `db:"MAJOR_NO"`
	Major_flag        string `db:"MAJOR_FLAG"`
	Major_name_thai   string `db:"MAJOR_NAME_THAI"`
	Major_name_eng    string `db:"MAJOR_NAME_ENG"`
	Lev_id            string `db:"LEV_ID"`
}

type StudentRepository interface {
	Authenticate(std_code string, birth_date string) (*Student, error)
}