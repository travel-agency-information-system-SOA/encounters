package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=encounters" // podesavanje baze
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Encounter{}, &model.SocialEncounter{}, &model.HiddenLocationEncounter{}) // migracije da bismo napravili tabele
	//database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")
	return database
}

func startEncounterServer(handler *handler.EncounterHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	//router.HandleFunc("/students", handler.Create).Methods("POST")

	//za zahteve iz c# proj ka ovamo
	router.HandleFunc("/encounters/create", handler.Create).Methods("POST")
	router.HandleFunc("/encounters/createSocialEncounter", handler.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/createHiddenLocationEncounter", handler.CreateHiddenLocationEncounter).Methods("POST")
	router.HandleFunc("/encounters", handler.GetAllEncounters).Methods("GET")
	router.HandleFunc("/hiddenLocationEncounters", handler.GetAllHiddenLocationEncounters).Methods("GET")
	router.HandleFunc("/socialEncounters", handler.GetAllSocialEncounters).Methods("GET")
	router.HandleFunc("/encounters/update", handler.Update).Methods("PUT")
	router.HandleFunc("/encounters/updateHiddenLocationEncounter", handler.UpdateHiddenLocationEncounter).Methods("PUT")
	router.HandleFunc("/encounters/updateSocialEncounter", handler.UpdateSocialEncounter).Methods("PUT")

	router.HandleFunc("/encounters/getSocialEncounterId/{baseEncounterId}", handler.GetSocialEncounterId).Methods("GET")
	router.HandleFunc("/encounters/getHiddenLocationEncounterId/{baseEncounterId}", handler.GetHiddenLocationEncounterId).Methods("GET")

	router.HandleFunc("/encounters/deleteEncounter/{baseEncounterId}", handler.DeleteEncounter).Methods("DELETE")
	router.HandleFunc("/encounters/deleteSocialEncounter/{socialEncounterId}", handler.DeleteSocialEncounter).Methods("DELETE")
	router.HandleFunc("/encounters/deleteHiddenLocationEncounter/{hiddenLocationEncounterId}", handler.DeleteHiddenLocationEncounter).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.EncounterRepository{DatabaseConnection: database}
	service := &service.EncounterService{EncounterRepo: repo}
	handler := &handler.EncounterHandler{EncounterService: service}

	startEncounterServer(handler)
}