package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"gorm.io/driver/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=encounters"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.EncounterExecution{})
	return database
}

func startServer(handler *handler.EncounterExecutionHandler) {
	router := mux.NewRouter().StrictSlash(true)

	// Encounter Execution
	router.HandleFunc("/encounterExecution/getActive/{userId}", handler.GetExecutionByUser).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	encounterExecutionRepo := &repo.EncounterExecutionRepository{DatabaseConnection: database}
	encounterExecutionService := &service.EncounterExecutionService{EncounterExecutionRepo: encounterExecutionRepo}
	encounterExecutionHandler := &handler.EncounterExecutionHandler{EncounterExecutionService: encounterExecutionService}

	startServer(encounterExecutionHandler)
}
