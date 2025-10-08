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

func (apiCfg apiConfig) GetOneBoardHandler(w http.ResponseWriter, r *http.Request) {

	boardIDStr := chi.URLParam(r, "id")
	boardID, err := uuid.Parse(boardIDStr)
	if err != nil {
		http.Error(w, "Error parsing board ID", 400)
		return
	}
	board, err := apiCfg.DB.GetOneBoard(r.Context(), boardID)
	if err != nil {
		fmt.Println("Error fetching one board", err)
	}
	respondWithJSON(w, 200, dbBoardToBoard(board))

}

func (apiCfg apiConfig) GetboardHandler(w http.ResponseWriter, r *http.Request) {
	boards, err := apiCfg.DB.GetBoards(r.Context())
	if err != nil {
		fmt.Println(err)
	}
	respondWithJSON(w, 200, dbBoardsToBoards(boards))

}

func (apiCfg apiConfig) CreateBoardHandler(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Title string `json:"title"`
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		//TODO: implement respondWithError
		fmt.Println(err)
	}
	board, err := apiCfg.DB.CreateBoard(r.Context(), database.CreateBoardParams{
		ID:         uuid.New(),
		BoardTitle: params.Title,
	})
	if err != nil {
		fmt.Println(err)
	}
	respondWithJSON(w, 200, board)

}

// UpdateBoardHandler handles updating a board
func (apiCfg apiConfig) UpdateBoardHandler(w http.ResponseWriter, r *http.Request) {
	boardIDStr := chi.URLParam(r, "id")
	boardID, err := uuid.Parse(boardIDStr)
	if err != nil {
		http.Error(w, "Error parsing board ID", 400)
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

	board, err := apiCfg.DB.UpdateBoard(r.Context(), database.UpdateBoardParams{
		BoardTitle: params.Title,
		UpdatedAt:  time.Now().UTC(),
		ID:         boardID,
	})
	if err != nil {
		http.Error(w, "Error updating board", 500)
		return
	}
	respondWithJSON(w, 200, board)
}

// DeleteBoardHandler handles deleting a board
func (apiCfg apiConfig) DeleteBoardHandler(w http.ResponseWriter, r *http.Request) {
	boardIDStr := chi.URLParam(r, "id")
	boardID, err := uuid.Parse(boardIDStr)
	if err != nil {
		http.Error(w, "Error parsing board ID", 400)
		return
	}

	err = apiCfg.DB.DeleteBoard(r.Context(), boardID)
	if err != nil {
		http.Error(w, "Error deleting board", 500)
		return
	}

	respondWithJSON(w, 200, map[string]string{"message": "Board deleted successfully"})
}
