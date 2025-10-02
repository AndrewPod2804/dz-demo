package jwt

import "github.com/golang-jwt/jwt/v5"

type JWTData struct {
	Phone string
}
type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}
func (j *JWT) Create(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
	})
	//t.Claims()
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil

}
func (j *JWT) Parse(tokenString string) (bool, *JWTData) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	phone := token.Claims.(jwt.MapClaims)["phone"]
	return token.Valid, &JWTData{
		Phone: phone.(string),
	}
}
