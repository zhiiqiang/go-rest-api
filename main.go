package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Task struct {
	ID     string
	Title  string
	Status string
}

var tasks []Task

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var task Task
		json.Unmarshal(body, &task)
		tasks = append(tasks, task)
		json.NewEncoder(w).Encode(tasks)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/tasks/")
		for _, item := range tasks {
			if item.ID == id {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		http.Error(w, "Task not found", http.StatusNotFound)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	tasks = append(tasks, Task{ID: "1", Title: "Task one", Status: "Pending"})
	http.HandleFunc("/tasks/", taskHandler)
	http.HandleFunc("/tasks", tasksHandler)
	fmt.Println("Server listening to port 8000")
	http.ListenAndServe(":8000", nil)
}
