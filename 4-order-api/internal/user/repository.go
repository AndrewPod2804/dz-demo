package user

import "4-order-api/pkg/db"

type PhoneRepository struct {
	Database *db.Db
}

func NewPhoneRepository(database *db.Db) *PhoneRepository {
	return &PhoneRepository{
		Database: database,
	}
}
func (repo *PhoneRepository) Create(phone *Phone) (*Phone, error) {
	result := repo.Database.DB.Create(phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return phone, nil
}

func (repo *PhoneRepository) FindBySessionId(sessionId string) (*Phone, error) {
	var phone Phone
	result := repo.Database.DB.First(&phone, "session_id = ?", sessionId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &phone, nil
}
