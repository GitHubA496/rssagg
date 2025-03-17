package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GitHubA496/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiconfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found wihtin the env")
	}

	dburl := os.Getenv("DB_URL")
	if dburl == "" {
		log.Fatal("DB_URL is not found wihtin the env")
	}
	// fmt.Println(dburl)
	conn, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	quries := database.New(conn)

	apiCfg := apiconfig{
		DB: quries,
	}

	go startScrapping(quries, 10, time.Second*30)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerError)

	v1router.Post("/user", apiCfg.handleCreateUser)
	v1router.Get("/user", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	v1router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handleGetPosts))
	v1router.Get("/feed", apiCfg.handlerGetFeed)
	v1router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))

	v1router.Post("/feed/follow", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollowers))
	v1router.Get("/feed/follow", apiCfg.middlewareAuth(apiCfg.handleGetFeedFollowers))
	v1router.Delete("/feed/follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollowers))

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server starting on the port %v", port)
	error := srv.ListenAndServe()
	if error != nil {
		log.Fatal(err)
	}

}
