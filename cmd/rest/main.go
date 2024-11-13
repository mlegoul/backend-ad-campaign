package main

import (
	"backend-ad-campaign/internal/adapters/api"
	"backend-ad-campaign/internal/adapters/repository"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	connStr := "user=postgres dbname=backend_db password=example host=db port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	handler := api.NewCampaignHandler(repo)

	router := mux.NewRouter()
	router.HandleFunc("/campaigns", handler.HandleCreateCampaign).Methods("POST")

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
