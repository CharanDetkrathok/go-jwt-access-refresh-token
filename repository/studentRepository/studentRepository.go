package studentRepository

func (r studentRepository) AuthenticateBachelor(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoryFromDB, error) {

	student := StudentAuthenticationRepositoryFromDB{}

	query := "ใส่ query string ตรงนี้"

	err := r.db.Get(&student, query, std_code, birth_date)
	if err != nil {
		return nil, err
	}

	return &student, nil

}

func (r studentRepository) AuthenticateMaster(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoryFromDB, error) {

	student := StudentAuthenticationRepositoryFromDB{}

	query := "ใส่ query string ตรงนี้"


	err := r.db.Get(&student, query, std_code, birth_date)
	if err != nil {
		return nil, err
	}

	return &student, nil

}

func (r studentRepository) AuthenticatePhd(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoryFromDB, error) {

	student := StudentAuthenticationRepositoryFromDB{}

	query := "ใส่ query string ตรงนี้"

	err := r.db.Get(&student, query, std_code, birth_date)
	if err != nil {
		return nil, err
	}

	return &student, nil

}
