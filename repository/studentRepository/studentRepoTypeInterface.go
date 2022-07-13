package studentRepository

import "github.com/jmoiron/sqlx"

type (
	StudentAuthenticationRepositoryFromDB struct {
		// เพิ่มโครงสร้าง database ตรงนี้
	}

	studentRepository struct {
		db *sqlx.DB
	}

	StudentRepository interface {
		AuthenticateBachelor(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoryFromDB, error)
		AuthenticateMaster(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoryFromDB, error)
		AuthenticatePhd(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoryFromDB, error)
	}
)

func NewStudentRepository(db *sqlx.DB) studentRepository {
	return studentRepository{db: db}
}
