package main

import (
	"database/sql"
	"log"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
	"github.com/lot-tola/project-manager/internal/database"
	"github.com/pressly/goose/v3"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "5000" 
		fmt.Println("Using default port 5000")
	}

	dbURL := os.Getenv("DATABASE_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Cannot connect to database, Err: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	if err := goose.SetDialect("postgres"); err != nil {
			log.Fatal("Failed to set dialect:", err)
	}
	if err := goose.Up(conn, "../backend/sql/schema"); err != nil {  // Path relative to your project root
			log.Fatal("Migration failed:", err)
	}
	log.Println("Migrations completed successfully")

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://taskcosmos-pro.netlify.app"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handleTest)

	// Board routes
	v1Router.Get("/boards", apiCfg.GetboardHandler)
	v1Router.Post("/boards", apiCfg.CreateBoardHandler)
	v1Router.Get("/boards/{id}", apiCfg.GetOneBoardHandler)
	v1Router.Put("/boards/{id}", apiCfg.UpdateBoardHandler)
	v1Router.Delete("/boards/{id}", apiCfg.DeleteBoardHandler)

	// List routes
	v1Router.Get("/lists/{id}", apiCfg.GetListsHandler)
	v1Router.Get("/lists", apiCfg.GetAllListsHandler)
	v1Router.Post("/lists", apiCfg.CreateListHandler)
	v1Router.Put("/lists/{id}", apiCfg.UpdateListHandler)
	v1Router.Delete("/lists/{id}", apiCfg.DeleteListHandler)

	// Task routes
	v1Router.Get("/task/{id}", apiCfg.GetTaskHandler)
	v1Router.Post("/tasks", apiCfg.CreateTaskHandler)
	v1Router.Delete("/tasks/{id}", apiCfg.DeleteTaskHandler)
	v1Router.Put("/tasks/edit/{id}", apiCfg.UpdateTaskHandler)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	fmt.Printf("Sever is starting on port %s", portString)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
