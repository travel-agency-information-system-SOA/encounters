package main

import (
	"context"
	"database-example/handler"
	"database-example/repo"
	"database-example/service"
	"os/signal"

	"log"
	"net/http"
	"os"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//nil - pokazivac ne pokazuje ni na sta

/*
func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer host=database port=5432 sslmode=disable search_path=encounters"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Encounter{}, &model.SocialEncounter{}, &model.HiddenLocationEncounter{}, &model.EncounterExecution{})
	return database
}
*/

/*
func startServer(handlerEnc *handler.EncounterHandler, handlerExec *handler.EncounterExecutionHandler) {
	router := mux.NewRouter().StrictSlash(true) //za rukovanje http zahtevima i definisanje ruta

	//za zahteve iz c# proj ka ovamo
	router.HandleFunc("/encounters/create", handlerEnc.Create).Methods("POST")
	router.HandleFunc("/encounters", handlerEnc.GetAllEncounters).Methods("GET")

		router.HandleFunc("/encounters/createSocialEncounter", handlerEnc.CreateSocialEncounter).Methods("POST")
		router.HandleFunc("/encounters/createHiddenLocationEncounter", handlerEnc.CreateHiddenLocationEncounter).Methods("POST")

		router.HandleFunc("/encounters/update", handlerEnc.Update).Methods("PUT")
		router.HandleFunc("/encounters/updateHiddenLocationEncounter", handlerEnc.UpdateHiddenLocationEncounter).Methods("PUT")
		router.HandleFunc("/encounters/updateSocialEncounter", handlerEnc.UpdateSocialEncounter).Methods("PUT")

		router.HandleFunc("/encounters", handlerEnc.GetAllEncounters).Methods("GET")
		router.HandleFunc("/hiddenLocationEncounters", handlerEnc.GetAllHiddenLocationEncounters).Methods("GET")
		router.HandleFunc("/socialEncounters", handlerEnc.GetAllSocialEncounters).Methods("GET")

		router.HandleFunc("/encounters/getEncounterById/{encounterId}", handlerEnc.GetEncounterById).Methods("GET")

		router.HandleFunc("/encounters/getSocialEncounterId/{baseEncounterId}", handlerEnc.GetSocialEncounterId).Methods("GET")
		router.HandleFunc("/encounters/getHiddenLocationEncounterId/{baseEncounterId}", handlerEnc.GetHiddenLocationEncounterId).Methods("GET")
		router.HandleFunc("/encounters/getHiddenLocationEncounter/{encounterId}", handlerEnc.GetHiddenLocationEncounterByEncounterId).Methods("GET")

		router.HandleFunc("/encounters/deleteEncounter/{baseEncounterId}", handlerEnc.DeleteEncounter).Methods("DELETE")
		router.HandleFunc("/encounters/deleteSocialEncounter/{socialEncounterId}", handlerEnc.DeleteSocialEncounter).Methods("DELETE")
		router.HandleFunc("/encounters/deleteHiddenLocationEncounter/{hiddenLocationEncounterId}", handlerEnc.DeleteHiddenLocationEncounter).Methods("DELETE")

		// Encounter Execution
		router.HandleFunc("/encounterExecution", handlerExec.GetAll).Methods("GET")
		router.HandleFunc("/encounterExecution/getActive/{userId}", handlerExec.GetExecutionByUser).Methods("GET")
		router.HandleFunc("/encounterExecution/completeExecution/{userId}", handlerExec.CompleteEncounter).Methods("GET")
		router.HandleFunc("/encounterExecution/create", handlerExec.Create).Methods("POST")
		router.HandleFunc("/encounterExecution/update/{id}", handlerExec.Update).Methods("PUT")
		router.HandleFunc("/encounterExecution/delete/{id}", handlerExec.Delete).Methods("DELETE")


	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":4000", router))
}
*/

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[encounter-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[encounter-store] ", log.LstdFlags)

	store, err := repo.New(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	store.Ping()

	encounterRepo, err := repo.New(timeoutContext, logger)

	encounterService := service.NewEncounterService(encounterRepo)

	encounterHandler := handler.NewEncounterHandler(encounterService, logger)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/create", encounterHandler.Create).Methods("POST")
	router.HandleFunc("/encounters", encounterHandler.GetAllEncounters).Methods("GET")
	/*
		router.HandleFunc("/encounters/misc", encounterHandler.CreateMiscEncounter).Methods("POST")
		router.HandleFunc("/encounters/social", encounterHandler.CreateSocialEncounter).Methods("POST")
		router.HandleFunc("/encounters/hidden", encounterHandler.CreateHiddenLocationEncounter).Methods("POST")
		router.HandleFunc("/encounters/isInRange/{id}/{long}/{lat}", encounterHandler.IsUserInCompletitionRange).Methods("GET")
		router.HandleFunc("/encounters/{range}/{long}/{lat}", encounterHandler.FindAllInRangeOf).Methods("GET")
		router.HandleFunc("/encounters", encounterHandler.FindAll).Methods("GET")
		router.HandleFunc("/encounters/hidden/{id}", encounterHandler.FindHiddenLocationEncounterById).Methods("GET")
		router.HandleFunc("/encounters/doneByUser/{id}", encounterHandler.FindAllDoneByUser).Methods("GET")
		router.HandleFunc("/encounters/instance/{id}/{encounterId}/encounter", encounterInstanceHandler.FindEncounterInstance).Methods("GET")
		router.HandleFunc("/encounters/touristProgress/{id}", touristProgressHandler.FindTouristProgressByTouristId).Methods("GET")
		router.HandleFunc("/encounters/complete/{userid}/{encounterId}/misc", encounterHandler.CompleteMiscEncounter).Methods("GET")
		router.HandleFunc("/encounters/activate/{id}", encounterHandler.ActivateEncounter).Methods("POST")
		router.HandleFunc("/encounters/complete/{id}", encounterHandler.CompleteHiddenLocationEncounter).Methods("POST")
		router.HandleFunc("/encounters/complete/{encounterId}/social", encounterHandler.CompleteSocialEncounter).Methods("POST")

	*/
	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")

}
