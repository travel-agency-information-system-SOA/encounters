package repo

import (
	"context"
	"database-example/model"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
)

/*
type EncounterRepository struct {
	//DatabaseConnection *gorm.DB //za konekciju sa bazom podataka
	cli    *mongo.Client
	logger *log.Logger
}
*/

type EncounterRepository struct {
	store *Repository
}

func NewEncounterRepository(r *Repository) *EncounterRepository {
	return &EncounterRepository{r}
}

func (er *EncounterRepository) getCollection() *mongo.Collection {
	encounterDatabase := er.store.cli.Database("encounters")
	encoutersCollection := encounterDatabase.Collection("encounters_collection")
	return encoutersCollection
}

func (er *EncounterRepository) getSocialEncountersCollection() *mongo.Collection {
	encounterDatabase := er.store.cli.Database("encounters") //ista baza druga kolekcija
	socialEncoutersCollection := encounterDatabase.Collection("socialEncounters_collection")
	return socialEncoutersCollection
}

func (er *EncounterRepository) getHiddenLocationEncountersCollection() *mongo.Collection {
	encounterDatabase := er.store.cli.Database("encounters") //ista baza druga kolekcija
	hiddenLocationEncoutersCollection := encounterDatabase.Collection("hiddenLocationEncouters_collection")
	return hiddenLocationEncoutersCollection
}

func (repo *EncounterRepository) CreateEncounter(encounter *model.Encounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	encountersCollection := repo.getCollection()

	result, err := encountersCollection.InsertOne(ctx, &encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

/*
OVO VALJA AL DA PROBAM JOS NESTO
func (repo *EncounterRepository) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	socialEncountersCollection := repo.getSocialEncountersCollection()

	result, err := socialEncountersCollection.InsertOne(ctx, &encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}
*/

func (repo *EncounterRepository) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	baseEncounter := &model.Encounter{
		Id:               encounter.Encounter.Id,
		Name:             encounter.Encounter.Name,
		Description:      encounter.Encounter.Description,
		XpPoints:         encounter.Encounter.XpPoints,
		Status:           encounter.Encounter.Status,
		Type:             encounter.Encounter.Type,
		Latitude:         encounter.Encounter.Latitude,
		Longitude:        encounter.Encounter.Longitude,
		ShouldBeApproved: encounter.Encounter.ShouldBeApproved,
	}

	if err := repo.CreateEncounter(baseEncounter); err != nil {
		repo.store.logger.Println(err)
		return err
	}

	//encounter.EncounterId = baseEncounter.ID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hiddenLocationEncounterCollection := repo.getHiddenLocationEncountersCollection()

	//tracing
	tracer := otel.Tracer("repository")
	ctx, span := tracer.Start(ctx, "CreateHiddenLocationEncounter")
	defer span.End()

	result, err := hiddenLocationEncounterCollection.InsertOne(ctx, encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *EncounterRepository) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	baseEncounter := &model.Encounter{
		Id:               encounter.Encounter.Id,
		Name:             encounter.Encounter.Name,
		Description:      encounter.Encounter.Description,
		XpPoints:         encounter.Encounter.XpPoints,
		Status:           encounter.Encounter.Status,
		Type:             encounter.Encounter.Type,
		Latitude:         encounter.Encounter.Latitude,
		Longitude:        encounter.Encounter.Longitude,
		ShouldBeApproved: encounter.Encounter.ShouldBeApproved,
	}

	if err := repo.CreateEncounter(baseEncounter); err != nil {
		repo.store.logger.Println(err)
		return err
	}

	//encounter.EncounterId = baseEncounter.ID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	socialEncountersCollection := repo.getSocialEncountersCollection()

	//tracing
	tracer := otel.Tracer("repository")
	ctx, span := tracer.Start(ctx, "CreateSocialEncounter")
	defer span.End()

	result, err := socialEncountersCollection.InsertOne(ctx, encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

/*
func (repo *EncounterRepository) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hiddenLocationEncounterCollection := repo.getHiddenLocationEncountersCollection()

	result, err := hiddenLocationEncounterCollection.InsertOne(ctx, &encounter)
	if err != nil {
		repo.store.logger.Println(err)
		return err
	}
	repo.store.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}
*/

func (r *EncounterRepository) GetAllEncounters() (model.Encounters, error) {
	// Initialise context (after 5 seconds timeout, abort operation)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//tracing
	tracer := otel.Tracer("repository")
	ctx, span := tracer.Start(ctx, "GetAllEncounters")
	defer span.End()

	enocuntersCollection := r.getCollection()

	var encounters model.Encounters
	encountersCursor, err := enocuntersCollection.Find(ctx, bson.M{})
	if err != nil {
		r.store.logger.Println(err)
		return nil, err
	}
	if err = encountersCursor.All(ctx, &encounters); err != nil {
		r.store.logger.Println(err)
		return nil, err
	}
	return encounters, nil
}

func (r *EncounterRepository) GetAllHiddenLocationEncounters() (model.HiddenLocationEncounters, error) {
	// Ovde bi trebalo da izvršimo upit ka bazi podataka ili drugom skladištu podataka da dobijemo sve susrete
	// Na primer, koristeći ORM poput GORM-a, možemo uraditi nešto poput sledećeg:
	/*var encounters []*model.HiddenLocationEncounter
	if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
		// Ukoliko dođe do greške pri izvršavanju upita, vraćamo grešku
		return nil, err
	}

	return encounters, nil*/

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	enocuntersCollection := r.getHiddenLocationEncountersCollection()

	var encounters model.HiddenLocationEncounters
	encountersCursor, err := enocuntersCollection.Find(ctx, bson.M{})
	if err != nil {
		r.store.logger.Println(err)
		return nil, err
	}
	if err = encountersCursor.All(ctx, &encounters); err != nil {
		r.store.logger.Println(err)
		return nil, err
	}
	return encounters, nil
}

func (r *EncounterRepository) GetAllSocialEncounters() (model.SocialEncounters, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	enocuntersCollection := r.getSocialEncountersCollection()

	var encounters model.SocialEncounters
	encountersCursor, err := enocuntersCollection.Find(ctx, bson.M{})
	if err != nil {
		r.store.logger.Println(err)
		return nil, err
	}
	if err = encountersCursor.All(ctx, &encounters); err != nil {
		r.store.logger.Println(err)
		return nil, err
	}
	return encounters, nil
}

func (repo *EncounterRepository) Update(encounter *model.Encounter) error {
	/*dbResult := repo.DatabaseConnection.Model(&model.Encounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
		"name":               encounter.Name,
		"description":        encounter.Description,
		"xp_points":          encounter.XpPoints,
		"status":             encounter.Status,
		"type":               encounter.Type,
		"longitude":          encounter.Longitude,
		"latitude":           encounter.Latitude,
		"should_be_approved": encounter.ShouldBeApproved,
	})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("Encounter not found")
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getCollection()
	filter := bson.M{"id": encounter.Id}

	update := bson.M{
		"$set": bson.M{
			"name":               encounter.Name,
			"description":        encounter.Description,
			"xp_points":          encounter.XpPoints,
			"status":             encounter.Status,
			"type":               encounter.Type,
			"longitude":          encounter.Longitude,
			"latitude":           encounter.Latitude,
			"should_be_approved": encounter.ShouldBeApproved,
		},
	}

	result, err := encountersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("Encounter not found")
	}

	repo.store.logger.Println("Rows affected: ", result.ModifiedCount)
	return nil
}

func (repo *EncounterRepository) UpdateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getHiddenLocationEncountersCollection()
	filter := bson.M{"id": encounter.Id}

	update := bson.M{
		"$set": bson.M{
			"image_url":         encounter.ImageURL,
			"image_latitude":    encounter.ImageLatitude,
			"image_longitude":   encounter.ImageLongitude,
			"distance_treshold": encounter.DistanceTreshold,
			"encounter_id":      encounter.Encounter,
		},
	}

	result, err := encountersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("HiddenLocationEncounter not found")
	}

	repo.store.logger.Println("Rows affected: ", result.ModifiedCount)
	return nil
}

func (repo *EncounterRepository) UpdateSocialEncounter(encounter *model.SocialEncounter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := repo.getSocialEncountersCollection()
	filter := bson.M{"id": encounter.Id}

	update := bson.M{
		"$set": bson.M{
			"tourists_required_for_completion": encounter.TouristsRequiredForCompletion,
			"distance_treshold":                encounter.DistanceTreshold,
			"tourist_ids":                      encounter.TouristIDs,
		},
	}

	result, err := encountersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("SocialEncounter not found")
	}

	repo.store.logger.Println("Rows affected: ", result.ModifiedCount)
	return nil
}

func (r *EncounterRepository) GetSocialEncounterId(baseEncounterID string) (string, error) {
	/*var socialEncounterID int

	// Izvršavanje upita za dobavljanje ID-a društvenog susreta
	result := r.DatabaseConnection.Model(&model.SocialEncounter{}).Select("id").Where("encounter_id = ?", baseEncounterID).First(&socialEncounterID)
	if result.Error != nil {
		// Provera greške prilikom izvršavanja upita
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Ako red ne postoji, vraćamo -1 kao ID društvenog susreta
			return -1, nil
		}
		// Ako postoji druga greška, vraćamo grešku
		return 0, result.Error
	}

	// Ako nema greške, vraćamo ID društvenog susreta
	return socialEncounterID, nil*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := r.getSocialEncountersCollection()
	filter := bson.M{"encounter.id": baseEncounterID}

	var socialEncounter model.SocialEncounter
	err := encountersCollection.FindOne(ctx, filter).Decode(&socialEncounter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		return "", err
	}

	return socialEncounter.Id.Hex(), nil
}

func (r *EncounterRepository) GetHiddenLocationEncounterId(baseEncounterID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := r.getHiddenLocationEncountersCollection()
	filter := bson.M{"encounter.id": baseEncounterID}

	var hiddenEncounter model.SocialEncounter
	err := encountersCollection.FindOne(ctx, filter).Decode(&hiddenEncounter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		return "", err
	}

	return hiddenEncounter.Id.Hex(), nil
}

func (r *EncounterRepository) DeleteSocialEncounter(socialEncounterID string) error {
	// Izvršavanje SQL upita za brisanje socijalnog susreta na osnovu njegovog ID-ja
	/*result := r.DatabaseConnection.Exec("DELETE FROM social_encounters WHERE id = ?", socialEncounterID)
	if result.Error != nil {
		// Ukoliko dođe do greške prilikom izvršavanja SQL upita, vraćamo je kao rezultat
		return result.Error
	}
	return nil*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := r.getSocialEncountersCollection()

	objID, err := primitive.ObjectIDFromHex(socialEncounterID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := encountersCollection.DeleteOne(ctx, filter)
	if err != nil {
		r.store.logger.Println(err)
		return err
	}

	if result.DeletedCount == 0 {
		return nil
	}

	r.store.logger.Printf("Deleted document ID: %v\n", objID)
	return nil
}

func (r *EncounterRepository) DeleteHiddenLocationEncounter(hiddenLocationEncounterID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := r.getHiddenLocationEncountersCollection()

	objID, err := primitive.ObjectIDFromHex(hiddenLocationEncounterID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := encountersCollection.DeleteOne(ctx, filter)
	if err != nil {
		r.store.logger.Println(err)
		return err
	}

	if result.DeletedCount == 0 {
		return nil
	}

	r.store.logger.Printf("Deleted document ID: %v\n", objID)
	return nil
}

func (r *EncounterRepository) DeleteEncounter(baseEncounterID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	encountersCollection := r.getCollection()

	objID, err := primitive.ObjectIDFromHex(baseEncounterID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := encountersCollection.DeleteOne(ctx, filter)
	if err != nil {
		r.store.logger.Println(err)
		return err
	}

	if result.DeletedCount == 0 {
		return nil
	}

	r.store.logger.Printf("Deleted document ID: %v\n", objID)
	return nil
}

/*
func (r *EncounterRepository) GetHiddenLocationEncounterByEncounterId(baseEncounterID int) (*model.HiddenLocationEncounter, error) {
	var hiddenLocationEncounter model.HiddenLocationEncounter

	// Execute the query to fetch the hidden location encounter by baseEncounterID
	result := r.DatabaseConnection.Model(&model.HiddenLocationEncounter{}).Where("encounter_id = ?", baseEncounterID).First(&hiddenLocationEncounter)
	if result.Error != nil {
		// Check for error while executing the query
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// If the record does not exist, return nil and nil error
			return nil, nil
		}
		// If there's another error, return nil and the error
		return nil, result.Error
	}

	// If no error, return the hidden location encounter object
	return &hiddenLocationEncounter, nil
}

func (r *EncounterRepository) GetEncounterById(encounterId int) (*model.Encounter, error) {
	var encounter model.Encounter

	//upit nad bazom podataka
	//First - ocekuje pokazivac na objekat u koji ce upisati podatke
	result := r.DatabaseConnection.Model(&model.Encounter{}).Where("id = ?", encounterId).First(&encounter)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	//pokazivac na objekat
	return &encounter, nil
}
*/
