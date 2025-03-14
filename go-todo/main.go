package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-todo/db"
	"go-todo/repositories"
)

func main() {
	db.InitDB() // เชื่อมต่อ Database

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasksHandler).Methods("GET")

	fmt.Println("🚀 Server เริ่มที่ http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := repositories.GetAllTasks()
	if err != nil {
		http.Error(w, "ไม่สามารถดึงข้อมูลได้", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
