package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lot-tola/project-manager/internal/database"
)

func (apiCfg apiConfig) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ListID      string `json:"list_id"`
		TaskTitle   string `json:"task_title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		AssignedTo  string `json:"assigned_to"`
		Status      string `json:"status"`
		Labels      string `json:"labels"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	list_id, err := uuid.Parse(params.ListID)
	labels := []string{}
	labels = append(labels, params.Labels)
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
		Labels:      labels,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		fmt.Println("Error creating task", err)
	}
	respondWithJSON(w, 200, dbTaskToTask(task))

}
