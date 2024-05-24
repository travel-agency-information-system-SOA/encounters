package main

import (
	"context"
	"database-example/model"
	"database-example/proto/encounter"
	"database-example/repo"
	"database-example/service"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

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

	encounterRepo := repo.NewEncounterRepository(store)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}

	//lis, err := net.Listen("tcp", "localhost:81")
	lis, err := net.Listen("tcp", "encounters:4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Pass a pointer to Server here
	encounter.RegisterEncounterServer(grpcServer, &Server{EncounterRepo: encounterRepo})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

}

type Server struct {
	encounter.UnimplementedEncounterServer
	EncounterRepo *repo.EncounterRepository
}

func CreateId() int64 {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	return currentTimestamp + int64(uniqueID)
}

/*
func (s *Server) GetAllEncounters(ctx context.Context, req *encounter.GetAllEncountersRequest) (*encounter.GetAllEncountersResponse, error) {
	// Retrieve all encounters from the repository
	encounters, err := s.EncounterRepo.GetAllEncounters(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// Convert to proto response
	var response encounter.GetAllEncountersResponse
	for _, e := range encounters {
		response.Encounters = append(response.Encounters, &encounter.EncounterMongoDto{
			Id:               e.Id,
			Name:             e.Name,
			Description:      e.Description,
			XpPoints:         e.XpPoints,
			Status:           e.Status,
			Type:             e.Type,
			Latitude:         e.Latitude,
			Longitude:        e.Longitude,
			ShouldBeApproved: e.ShouldBeApproved,
		})
	}
	response.TotalItems = int32(len(encounters))

	return &response, nil
}
*/

/*
func (s Server) CreateWholeHiddenLocationEncounter(ctx context.Context, req *encounter.WholeHiddenLocationEncounterRequest) (*encounter.WholeHiddenLocationEncounterResponse, error) {
	// Pravimo ID za novi HiddenLocationEncounter
	// Kreiramo novi HiddenLocationEncounter na osnovu podataka iz zahteva
	newHiddenLocationEncounter := &encounter.HiddenLocationEncounterMongoDto{
		Encounter: &encounter.EncounterMongoDto{
			Name:             req.Name,
			Description:      req.Description,
			XpPoints:         req.XpPoints,
			Status:           req.Status,
			Type:             req.Type,
			Latitude:         req.Latitude,
			Longitude:        req.Longitude,
			ShouldBeApproved: req.ShouldBeApproved,
		},
		ImageUrl:         req.ImageUrl,
		ImageLatitude:    req.ImageLatitude,
		ImageLongitude:   req.ImageLongitude,
		DistanceTreshold: req.DistanceTreshold,
	}

	// Pozivamo odgovarajuću metodu za kreiranje HiddenLocationEncountera
	if err := s.EncounterRepo.CreateHiddenLocationEncounter(newHiddenLocationEncounter); err != nil {
		return nil, err
	}

	// Vraćamo odgovor sa podacima o novom HiddenLocationEncounter-u
	return &encounter.WholeHiddenLocationEncounterResponse{
		Name:             req.Name,
		Description:      req.Description,
		XpPoints:         req.XpPoints,
		Status:           req.Status,
		Type:             req.Type,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
		ShouldBeApproved: req.ShouldBeApproved,
		ImageLatitude:    req.ImageLatitude,
		ImageLongitude:   req.ImageLongitude,
		ImageUrl:         req.ImageUrl,
		DistanceTreshold: req.DistanceTreshold,
	}, nil
}
*/

func (s Server) CreateSocialEncounter(ctx context.Context, socialEncounterDto *encounter.SocialEnc) (*encounter.SocialEnc, error) {

	println("usao je na ENCOUNTERS GO")
	id := primitive.NewObjectID()

	touristIDs := make([]int, len(socialEncounterDto.TouristIDs))
	for i, v := range socialEncounterDto.TouristIDs {
		touristIDs[i] = int(v)
	}

	println("Parametri metode:")
	println("socialEncounterDto.Name:", socialEncounterDto.Name)
	println("socialEncounterDto.Description:", socialEncounterDto.Description)
	println("socialEncounterDto.XpPoints:", socialEncounterDto.XpPoints)

	// Kreiranje novog SocialEncounter objekta prema novom modelu
	newSocialEncounter := model.SocialEncounter{
		Id: id,
		Encounter: model.Encounter{
			Id:               id,
			Name:             socialEncounterDto.Name,
			Description:      socialEncounterDto.Description,
			XpPoints:         int(socialEncounterDto.XpPoints),
			Status:           socialEncounterDto.Status,
			Type:             socialEncounterDto.Type,
			Longitude:        socialEncounterDto.Longitude,
			Latitude:         socialEncounterDto.Latitude,
			ShouldBeApproved: socialEncounterDto.ShouldBeApproved,
		},
		TouristsRequiredForCompletion: int(socialEncounterDto.TouristsRequiredForCompletion),
		DistanceTreshold:              socialEncounterDto.DistanceTreshold,
		TouristIDs:                    touristIDs, // Koristimo konvertovanu listu
	}

	encounterService := service.NewEncounterService(s.EncounterRepo)

	err := encounterService.CreateSocialEncounter(&newSocialEncounter)
	if err != nil {
		println("Error while creating a new social encounter")
		return nil, err // Vraćamo error umesto nil
	}

	return &encounter.SocialEnc{
		Id:                            id.Hex(),
		Name:                          socialEncounterDto.Name,
		Description:                   socialEncounterDto.Description,
		XpPoints:                      socialEncounterDto.XpPoints,
		Status:                        socialEncounterDto.Status,
		Type:                          socialEncounterDto.Type,
		Longitude:                     socialEncounterDto.Longitude,
		Latitude:                      socialEncounterDto.Latitude,
		ShouldBeApproved:              socialEncounterDto.ShouldBeApproved,
		TouristsRequiredForCompletion: int32(newSocialEncounter.TouristsRequiredForCompletion),
		DistanceTreshold:              newSocialEncounter.DistanceTreshold,
		TouristIDs:                    socialEncounterDto.TouristIDs,
	}, nil
}

func (s Server) CreateHiddenLocationEncounter(ctx context.Context, hiddenLocationEncounter *encounter.HiddenLocationEnc) (*encounter.HiddenLocationEnc, error) {
	println("usao je na ENCOUNTERS GO HIDDEN LOCATION ENC")
	println(hiddenLocationEncounter)

	id := primitive.NewObjectID()

	newHiddenLocationEncounter := model.HiddenLocationEncounter{
		Id:               id,
		ImageURL:         hiddenLocationEncounter.ImageURL,
		ImageLatitude:    float64(hiddenLocationEncounter.ImageLatitude),
		ImageLongitude:   float64(hiddenLocationEncounter.ImageLongitude),
		DistanceTreshold: float64(hiddenLocationEncounter.DistanceTreshold),
		Encounter: model.Encounter{
			Id:               id,
			Name:             hiddenLocationEncounter.Name,
			Description:      hiddenLocationEncounter.Description,
			XpPoints:         int(hiddenLocationEncounter.XpPoints),
			Status:           hiddenLocationEncounter.Status,
			Type:             hiddenLocationEncounter.Type,
			Longitude:        hiddenLocationEncounter.Longitude,
			Latitude:         hiddenLocationEncounter.Latitude,
			ShouldBeApproved: hiddenLocationEncounter.ShouldBeApproved,
		},
	}

	encounterService := service.NewEncounterService(s.EncounterRepo)

	err := encounterService.CreateHiddenLocationEncounter(&newHiddenLocationEncounter)
	if err != nil {
		println("Error while creating a new social encounter")
		return nil, err // Vraćamo error umesto nil
	}

	return &encounter.HiddenLocationEnc{
		Id:               id.Hex(),
		Name:             hiddenLocationEncounter.Name,
		Description:      hiddenLocationEncounter.Description,
		XpPoints:         hiddenLocationEncounter.XpPoints,
		Status:           hiddenLocationEncounter.Status,
		Type:             hiddenLocationEncounter.Type,
		Latitude:         hiddenLocationEncounter.Latitude,
		Longitude:        hiddenLocationEncounter.Longitude,
		EncounterId:      newHiddenLocationEncounter.Encounter.Id.Hex(),
		ImageURL:         hiddenLocationEncounter.ImageURL,
		ImageLatitude:    hiddenLocationEncounter.ImageLatitude,
		ImageLongitude:   hiddenLocationEncounter.ImageLongitude,
		DistanceTreshold: hiddenLocationEncounter.DistanceTreshold,
		ShouldBeApproved: hiddenLocationEncounter.ShouldBeApproved,
	}, nil
}
