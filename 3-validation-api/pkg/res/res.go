package res

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	//json.NewEncoder(w).Encode(res)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println("error:", err)
	}
}

//func Json() {
//
//}
