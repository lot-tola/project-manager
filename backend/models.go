package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lot-tola/project-manager/internal/database"
)

type Board struct {
	ID         uuid.UUID `json:"id"`
	BoardTitle string    `json:"board_title"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func dbBoardToBoard(dbBoard database.Board) Board {
	return Board{
		ID:         dbBoard.ID,
		BoardTitle: dbBoard.BoardTitle,
		Created_at: dbBoard.CreatedAt,
		Updated_at: dbBoard.UpdatedAt,
	}
}

func dbBoardsToBoards(dbBoards []database.Board) []Board {
	boards := []Board{}
	for _, board := range dbBoards {
		board := dbBoardToBoard(board)
		boards = append(boards, board)
	}
	return boards

}

type ListRow struct {
	ListID      uuid.UUID      `json:"list_id"`
	ListTitle   string         `json:"list_title"`
	TaskID      uuid.NullUUID  `json:"task_id"`
	TaskTitle   sql.NullString `json:"task_title"`
	Description sql.NullString `json:"description"`
	DueDate     sql.NullTime   `json:"due_date"`
	AssignedTo  sql.NullString `json:"assigned_to"`
	Status      sql.NullString `json:"status"`
	Labels      []string       `json:"labels"`
}

func dbListRowToListRow(dbListRow []database.GetListsRow) []ListRow {
	listRows := []ListRow{}
	for _, row := range dbListRow {
		listRow := ListRow{
			ListID:      row.ListID,
			ListTitle:   row.ListTitle,
			TaskID:      row.TaskID,
			TaskTitle:   row.TaskTitle,
			Description: row.Description,
			DueDate:     row.DueDate,
			AssignedTo:  row.AssignedTo,
			Status:      row.Status,
			Labels:      row.Labels,
		}
		listRows = append(listRows, listRow)
	}
	return listRows
}

type List struct {
	ID        uuid.UUID `json:"id"`
	BoardID   uuid.UUID `json:"board_id"`
	ListTitle string    `json:"list_title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbListToList(dbList database.List) List {
	return List{
		ID:        dbList.ID,
		BoardID:   dbList.BoardID,
		ListTitle: dbList.ListTitle,
		CreatedAt: dbList.CreatedAt,
		UpdatedAt: dbList.UpdatedAt,
	}
}

func dbListsToLists(dbLists []database.List) []List {
	lists := []List{}
	for _, list := range dbLists {
		list := dbListToList(list)
		lists = append(lists, list)
	}
	return lists
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	ListID      uuid.UUID `json:"list_id"`
	TaskTitle   string    `json:"task_title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	AssignedTo  string    `json:"assigned_to"`
	Status      string    `json:"status"`
	Labels      []string  `json:"labels"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func dbTaskToTask(dbTask database.Task) Task {
	return Task{
		ID:          dbTask.ID,
		ListID:      dbTask.ListID,
		TaskTitle:   dbTask.TaskTitle,
		Description: dbTask.Description,
		DueDate:     dbTask.DueDate,
		AssignedTo:  dbTask.AssignedTo,
		Status:      dbTask.Status,
		Labels:      dbTask.Labels,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
}
