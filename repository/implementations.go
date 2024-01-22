package repository

func (r *Repository) SaveUser(entity *Users) (err error) {
	err = r.DB.Save(entity).Error
	if err != nil {
		return
	}
	return
}

func (r *Repository) FindUserByPhoneNumber(phoneNumber string) (entity *Users, err error) {
	err = r.DB.Where("phone_number = ?", phoneNumber).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *Repository) FindUserById(userID string) (entity *Users, err error) {
	err = r.DB.Where("id = ?", userID).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *Repository) FindUserByName(name string) (entity *Users, err error) {
	err = r.DB.Where("id = ?", name).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return
}
