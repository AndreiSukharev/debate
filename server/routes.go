package main

import (
	"api-test/database"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	sqlString := "SELECT * FROM tasks;"
	values := []string {""}
	res := database.InsertToDB(sqlString, values)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "ok"}`))
}

func signIn(w http.ResponseWriter, r *http.Request) {
	sqlString := "SELECT password FROM users WHERE id=$1;"
	values := []string {"1"}
	res := database.InsertToDB(sqlString, values)
	// todo: check password
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "ok"}`))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	sqlString := "INSERT INTO users(login, password) VALUES($1, $2) returning id;"
	values := []string {"log", "pass"}
	// todo: check user exists hash password
	database.InsertToDB(sqlString, values)
	w.WriteHeader(http.StatusAccepted)
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

