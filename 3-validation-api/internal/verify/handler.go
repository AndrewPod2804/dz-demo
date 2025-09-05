package verify

import (
	"HomeVal3/configs"
	"HomeVal3/internal/bdjson"
	"HomeVal3/pkg/hash"
	"HomeVal3/pkg/req"
	"HomeVal3/pkg/res"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/http"
	"net/smtp"
	"net/textproto"
)

type VerifyHandlerDeps struct {
	*configs.Config
}
type VerifyHandler struct {
	*configs.Config
}

var EmailVer string

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.VerifyEmail())
}
func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body.Email)
		res.Json(w, "Ok", 200)
		handler.sendEmail(body.Email)
		EmailVer = body.Email
	}
}
func (handler *VerifyHandler) VerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hs := r.PathValue("hash")
		user, err := bdjson.ReadJson()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("user=", user)
		if user.Email == EmailVer {
			fmt.Println("Email is Ok")
		}
		if user.Hash == hs {
			fmt.Println("Hash  is Ok")
		} else {
			fmt.Println("Hash  is Not Ok")
		}
		bdjson.RemoveJson()
	}
}

func (handler *VerifyHandler) sendEmail(em string) {
	emF := handler.Config.Ver.Email
	password := handler.Config.Ver.Password
	address := handler.Config.Ver.Address
	hs := hash.EmToHashSt(em)
	st := `<a href="http://localhost:8081/verify/` + hs + `">Верификация email!</a>`
	e := &email.Email{
		To:      []string{em},
		From:    "Andrew <" + emF + ">",
		Subject: "Andrew ",
		HTML:    []byte(st),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(address, smtp.PlainAuth("", emF, password, "smtp.yandex.ru"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	bdjson.SaveJson(em, hs)
	fmt.Println("End")
}
