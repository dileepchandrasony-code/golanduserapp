package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

var db *gorm.DB
var err error
var users []User
var MyRouter = mux.NewRouter().StrictSlash(true)
var dbDriver = "mysql"
var dbUser = "hidbluevision"
var dbPass = "hidbluevision"
var dbName = "user"
var tcp = "tcp(bluevision-java.cf0qbycayr8x.ap-southeast-2.rds.amazonaws.com:3306)"

func main() {
	db, _ = gorm.Open(dbDriver, dbUser+":"+dbPass+"@"+tcp+"/"+dbName)
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	HandleFunctions()
	defer db.Close()
}
func HandleFunctions() {
	fmt.Println("Run on browser http://127.0.0.1:10000")
	MyRouter.HandleFunc("/users", GetUsers).Methods("GET", "OPTIONS")
	MyRouter.HandleFunc("/user", GetUser).Methods("GET", "OPTIONS")
	MyRouter.HandleFunc("/createusers", CreateUser).Methods("POST", "OPTIONS")
	MyRouter.HandleFunc("/updateuser", UpdateUser).Methods("PUT", "OPTIONS")
	MyRouter.HandleFunc("/deleteuser", DeleteUser).Methods("DELETE", "OPTIONS")
	http.ListenAndServe("127.0.0.1:8000/user/", MyRouter)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	db.Where("id = ?", id).Find(&users)
	json.NewEncoder(w).Encode(users)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{ID: 1, Name: "Shanks", City: "Ajmer", Age: 34}
	db.Create(&user)
	json.NewEncoder(w).Encode(users)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Model(&users).Where("id = ?", 69).Update("name", "nick")
	json.NewEncoder(w).Encode(users)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	db.Where("id = ?", id).Delete(&users)
	json.NewEncoder(w).Encode(users)
}
