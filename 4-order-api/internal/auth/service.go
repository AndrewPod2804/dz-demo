package auth

import (
	"4-order-api/internal/user"
	"crypto/sha1"
	"encoding/base64"
	"errors"
)

const SMS = 3245

type AuthService struct {
	PhoneRepository *user.PhoneRepository
}

func NewAuthService(phoneRep *user.PhoneRepository) *AuthService {
	return &AuthService{
		PhoneRepository: phoneRep,
	}
}

func (service *AuthService) Register(phone string) (string, error) {
	hash := EmToHashSt(phone)
	p := &user.Phone{
		Phone:     phone,
		SessionId: hash,
		Code:      SMS,
	}
	_, err := service.PhoneRepository.Create(p)
	if err != nil {
		return "", err
	}
	return hash, nil
}
func (service *AuthService) Verify(sessionId string, code uint) (string, error) {
	ph, err := service.PhoneRepository.FindBySessionId(sessionId)
	if err != nil {
		return "", err
	}
	if (ph.SessionId == sessionId) && (ph.Code == code) {
		return ph.Phone, nil
	} else {
		return "", errors.New("invalid code")
	}
}
func EmToHashSt(st string) string {
	data := []byte(st)
	hasher := sha1.New()
	hasher.Write(data)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
