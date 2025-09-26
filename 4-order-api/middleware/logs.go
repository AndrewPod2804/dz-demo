package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func Init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.DebugLevel

}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Loggin")
		//start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		//fmt.Println("After")
		//log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), time.Since(start))
		//log.Printf("code: %v [%s] %q %v\n", wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
		log.WithFields(logrus.Fields{
			"status_code": wrapper.StatusCode,
			"metod":       r.Method,
			"url_path":    r.URL.Path,
		}).Info("Routing request")
	})

}
