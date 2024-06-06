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

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log.SetOutput(os.Stderr)

	//TRACING

	//shutdown - pravilno zatvaranje traceova
	shutdown, err := initTracer() //inicjalizacija traceera
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	//DEFER
	// - odlozeno zatvaranje tracer-a
	// - odlaže izvršavanje funkcije do trenutka kada se funkcija u kojoj se defer nalazi završi
	// - u ovom slučaju, funkcija za zatvaranje tracer-a će se izvršiti na kraju glavne funkcije (main ili bilo koje druge funkcije u kojoj je ovaj kod)

	//

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

func (s Server) CreateSocialEncounter(ctx context.Context, socialEncounterDto *encounter.SocialEnc) (*encounter.SocialEnc, error) {

	println("usao je na ENCOUNTERS GO")
	id := primitive.NewObjectID()

	tracer := otel.Tracer("service") //vraca tracer objekat, service - ime tracer-a
	//tracer koristimo za stvaranje novih span-ova
	//span - vremenski interval u kome se prati izvrsavanje odredjenog dela koda
	ctx, span := tracer.Start(ctx, "CreateSocialEncounter") //zapocinje novi span, ime span-a opisuje operaciju koja se prati
	defer span.End()
	//osigurava da će span.End() biti pozvan kada funkcija u kojoj se ovaj kod nalazi završi svoje izvršenje, bez obzira na način završetka - normalno ili usred greske

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

	//TRACING
	tracer := otel.Tracer("service")
	ctx, span := tracer.Start(ctx, "CreateHiddenLocationEncounter")
	defer span.End()
	//

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

func (s *Server) GetAllEncounters(ctx context.Context, request *encounter.PageRequest) (*encounter.ListEnc, error) {
	println("usao je na ENCOUNTERS GO GET ALL ENC")

	tracer := otel.Tracer("service")
	ctx, span := tracer.Start(ctx, "GetAllEncounters")
	defer span.End()

	encounterService := service.NewEncounterService(s.EncounterRepo)
	encounters, err := encounterService.GetAllEncounters()
	if err != nil {
		return nil, err
	}

	var encDtos []*encounter.Enc
	for _, encDto := range encounters {
		encDto := &encounter.Enc{
			Id:               encDto.Id.Hex(),
			Name:             encDto.Name,
			Description:      encDto.Description,
			XpPoints:         int32(encDto.XpPoints),
			Type:             encDto.Type,
			Latitude:         encDto.Latitude,
			Longitude:        encDto.Longitude,
			ShouldBeApproved: encDto.ShouldBeApproved,
		}
		encDtos = append(encDtos, encDto)
	}

	response := &encounter.ListEnc{
		Results:    encDtos,
		TotalCount: int32(len(encounters)),
	}

	return response, nil
}

//TRACING

func initTracer() (func(context.Context) error, error) {

	jaegerExporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint("jaeger:4318"), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", "encounters"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(jaegerExporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}
