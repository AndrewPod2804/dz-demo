package main
import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type RandomHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	h := &RandomHandler{}
	router.HandleFunc("/random", h.random())
}
func (handler *RandomHandler) random() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		b := strconv.Itoa(rand.Intn(6) + 1)
		_, err := w.Write([]byte(b))
		if err != nil {
			fmt.Println(err)
		}
	}
}
func main() {
	fmt.Println("2-random-api")
	router := http.NewServeMux()
	NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()

}
