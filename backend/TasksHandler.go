package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lot-tola/project-manager/internal/database"
)
type parameters struct {
	ListID      string `json:"list_id"`
	TaskTitle   string `json:"task_title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	AssignedTo  string `json:"assigned_to"`
	Status      string `json:"status"`
	Labels      string `json:"labels"`
}

func (apiCfg apiConfig) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	list_id, err := uuid.Parse(params.ListID)
	dueDate, err := time.Parse("2006-01-02", params.DueDate)
	// formattedTime := dueDate.Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Println("Error parsing time layout")
	}
	if err != nil {
		fmt.Println("error parsing listID to uuid")
	}
	if err != nil {
		fmt.Println("Error getting params while creating task", err)
	}
	task, err := apiCfg.DB.CreateTask(r.Context(), database.CreateTaskParams{
		ID:          uuid.New(),
		ListID:      list_id,
		TaskTitle:   params.TaskTitle,
		Description: params.Description,
		DueDate:     dueDate,
		AssignedTo:  params.AssignedTo,
		Status:      params.Status,
		Labels:      []string{params.Labels},
	})
	if err != nil {
		fmt.Println("Error creating task", err)
	}
	respondWithJSON(w, 200, dbTaskToTask(task))

}


func (apiCfg apiConfig) DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	taskIdStr := chi.URLParam(r, "id")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		http.Error(w, "Error parsing task ID", 400)
		return
	}
	err = apiCfg.DB.DeleteTask(r.Context(), taskId)
	if err != nil {
		fmt.Println("Error deleting task. ", err)
	}
	respondWithJSON(w, 200, map[string]string{"message": "Task deleted successfully"})
}
func (apiCfg apiConfig) GetTaskHandler(w http.ResponseWriter, r *http.Request){
	taskIdStr := chi.URLParam(r, "id")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		http.Error(w, "Error parsing task ID", 400)
		return
	}
	task, err := apiCfg.DB.GetTask(r.Context(), taskId)
	if err != nil {
		fmt.Println("Error getting task. ", err)
	}
	respondWithJSON(w, 200, dbTaskToTask(task))
}

func (apiCfg apiConfig) UpdateTaskHandler(w http.ResponseWriter, r *http.Request){
	taskIdStr := chi.URLParam(r, "id")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		http.Error(w, "Error parsing task ID", 400)
		return
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		fmt.Println("Error decoding params", err)
	}
	dueDate, err := time.Parse("2006-01-02", params.DueDate)
	if err != nil {
		fmt.Println("Error parsing time", err)
	}
	err = apiCfg.DB.UpdateTask(r.Context(), database.UpdateTaskParams{
		ID: taskId,
		TaskTitle: params.TaskTitle,
		Description: params.Description,
		DueDate: dueDate,
		AssignedTo: params.AssignedTo,
		Status: params.Status,
		Labels: []string{params.Labels},
	})
	if err != nil {
		fmt.Println("Error while updating task", err)
	}
}
