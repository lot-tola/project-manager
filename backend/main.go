package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
	"github.com/lot-tola/project-manager/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Try to load .env file, but don't fail if it doesn't exist
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "5000" // Default port
		fmt.Println("Using default port 5000")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		// Use a default database URL for development
		dbURL = "postgres://postgres:password@localhost/project_manager?sslmode=disable"
		fmt.Println("Using default database URL. Please set DB_URL environment variable for production.")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Cannot connect to database, Err: ", err)
		os.Exit(1)
	}

	// Test the database connection
	if err := conn.Ping(); err != nil {
		fmt.Println("Database connection failed:", err)
		fmt.Println("Please make sure PostgreSQL is running and the database exists.")
		fmt.Println("You can create the database with: createdb project_manager")
		os.Exit(1)
	}

	fmt.Println("Database connection successful!")
	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
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
	v1Router.Post("/tasks", apiCfg.CreateTaskHandler)

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
