package main

import (
	"api-test/database"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func signIn(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "ok"}`))
}

func signUp(w http.ResponseWriter, r *http.Request) {
	sqlString := "INSERT INTO users(login, password) VALUES($1, $2) returning id;"
	values := []string {"log", "pass"}
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

