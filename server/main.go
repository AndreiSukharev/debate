package main

import (
	"api-test/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)


func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	database.CreateTables()
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/signin", signIn).Methods(http.MethodPost)
	api.HandleFunc("/signup", signUp).Methods(http.MethodPost)
	api.HandleFunc("/tasks", getTasks).Methods(http.MethodGet)
	api.HandleFunc("/tasks", addTasks).Methods(http.MethodPost)
	api.HandleFunc("/tasks", updateTasks).Methods(http.MethodPut)
	api.HandleFunc("/tasks/{id}", deleteTasks).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5010", CORS(r)))

}
