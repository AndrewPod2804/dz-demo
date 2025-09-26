package bdjson

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

const filePath = "user.json"

//	func MyJson() {
//		fmt.Println("MyJson")
//	}
func SaveJson(em string, hs string) {
	user := User{
		Email: em,
		Hash:  hs,
	}
	f, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(f))
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write(f)
}

// func ReadJson(user *User) error {
func ReadJson() (User, error) {
	var user User
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	//var user User
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	//fmt.Println(user)
	return user, nil
}
func RemoveJson() {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File successfully deleted")
}
