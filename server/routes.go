package main

import (
	"api-test/database"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type userData struct {
	Login    string
	Password string
}

func handlePostData(r *http.Request) *userData {
	var data userData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	return &data
}

func getPassword(login string) string {
	sqlString := "SELECT password FROM users WHERE login=$1;"
	//values := []string {login}
	row := database.GetSpecificFromDB(sqlString, login)
	return row
}

func hashPassword(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))
	sum := hex.EncodeToString(hash.Sum(nil))
	return sum
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	//sqlString := "SELECT * FROM tasks;"
	//values := []string {""}
	//res := database.InsertToDB(sqlString, values)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "ok"}`))
}

func signIn(w http.ResponseWriter, r *http.Request) {
	data := handlePostData(r)
	password := getPassword(data.Login)
	if password != "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "error"}`))
	}
	hashedPassword := hashPassword(password)
	if hashedPassword != password {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "error"}`))
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "ok"}`))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	data := handlePostData(r)
	password := getPassword(data.Login)
	if password != "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "error"}`))
	}
	hashedPassword := hashPassword(password)
	sqlString := "INSERT INTO users(login, password) VALUES($1, $2) returning id;"
	values := []string {data.Login,hashedPassword}
	database.InsertToDB(sqlString, values)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "ok"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
