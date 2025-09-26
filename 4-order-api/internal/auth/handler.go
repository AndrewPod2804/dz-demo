package auth

import (
	"4-order-api/configs"
	"4-order-api/pkg/jwt"
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}
type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/phone/", handler.AuthByPhone())
	router.HandleFunc("POST /auth/code/", handler.VerifyCode())

	//io.ReadAll()
}

func (handler *AuthHandler) AuthByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[PhoneRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		//data := PhoneResponse{
		//	SessionId: "wuerwiruyweriu",
		//}
		data, err := handler.AuthService.Register(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		res.Json(w, data, 200)
	}
}
func (handler *AuthHandler) VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CodeRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		phone, err := handler.Verify(body.SessionId, body.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := CodeResponse{
			Token: token,
		}
		res.Json(w, data, 200)
	}
}
