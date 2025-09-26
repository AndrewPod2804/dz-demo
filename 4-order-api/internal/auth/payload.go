package auth

type PhoneRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type PhoneResponse struct {
	SessionId string `json:"sessionId"`
}

type CodeRequest struct {
	SessionId string `json:"sessionid" validate:"required"`
	Code      uint   `json:"code" validate:"required"`
}
type CodeResponse struct {
	Token string `json:"token"`
}
