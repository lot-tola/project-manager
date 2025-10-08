package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lot-tola/project-manager/internal/database"
)

// func (apiCfg apiConfig) GetListsHandler(w http.ResponseWriter, r *http.Request) {
// 	lists, err := apiCfg.DB.GetLists(context.Background())
// 	if err != nil {
// 		// TODO: respondWithError
// 	}
// 	respondWithJSON(w, 200, lists)
//
// }

func (apiCfg apiConfig) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	boardIDStr := chi.URLParam(r, "id")
	boardID, err := uuid.Parse(boardIDStr)
	if err != nil {
		fmt.Println("Error parsing uuid", err)
	}

	lists, err := apiCfg.DB.GetLists(context.Background(), boardID)
	if err != nil {
		fmt.Println("errors getting lists")
	}
	respondWithJSON(w, 200, dbListRowToListRow(lists))
}

func (apiCfg apiConfig) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type Parameters struct {
		BoardID uuid.UUID `json:"board_id"`
		Title   string    `json:"title"`
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, "Invalid body", 400)
		return
	}

	list, err := apiCfg.DB.CreateList(context.Background(), database.CreateListParams{
		ID:        uuid.New(),
		BoardID:   params.BoardID,
		ListTitle: params.Title,
	})
	if err != nil {
		http.Error(w, "Error creating list", 500)
		return
	}

	respondWithJSON(w, 200, dbListToList(list))
}

// UpdateListHandler handles updating a list
func (apiCfg apiConfig) UpdateListHandler(w http.ResponseWriter, r *http.Request) {
	listIDStr := chi.URLParam(r, "id")
	listID, err := uuid.Parse(listIDStr)
	if err != nil {
		http.Error(w, "Error parsing list ID", 400)
		return
	}

	type Parameters struct {
		Title string `json:"title"`
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		http.Error(w, "Error parsing request body", 400)
		return
	}

	list, err := apiCfg.DB.UpdateList(r.Context(), database.UpdateListParams{
		ListTitle: params.Title,
		UpdatedAt: time.Now().UTC(),
		ID:        listID,
	})
	if err != nil {
		http.Error(w, "Error updating list", 500)
		return
	}
	respondWithJSON(w, 200, list)
}

// DeleteListHandler handles deleting a list
func (apiCfg apiConfig) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	listIDStr := chi.URLParam(r, "id")
	listID, err := uuid.Parse(listIDStr)
	if err != nil {
		http.Error(w, "Error parsing list ID", 400)
		return
	}

	err = apiCfg.DB.DeleteList(r.Context(), listID)
	if err != nil {
		http.Error(w, "Error deleting list", 500)
		return
	}

	respondWithJSON(w, 200, map[string]string{"message": "List deleted successfully"})
}

func (apiCfg apiConfig) GetAllListsHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := apiCfg.DB.GetAllLists(r.Context())
	if err != nil {
		fmt.Println("Error getting all lists")
	}
	respondWithJSON(w, 200, dbListsToLists(lists))
}
