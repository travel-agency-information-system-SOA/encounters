package repo

import (
	"database-example/model"
	"gorm.io/gorm"
	"errors"
)

type EncounterRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterRepository) CreateEncounter(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}


func (repo *EncounterRepository) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (r *EncounterRepository) GetAllEncounters() ([]*model.Encounter, error) {
    // Ovde bi trebalo da izvršimo upit ka bazi podataka ili drugom skladištu podataka da dobijemo sve susrete
    // Na primer, koristeći ORM poput GORM-a, možemo uraditi nešto poput sledećeg:
    var encounters []*model.Encounter
    if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
        // Ukoliko dođe do greške pri izvršavanju upita, vraćamo grešku
        return nil, err
    }

    return encounters, nil
}

func (r *EncounterRepository) GetAllHiddenLocationEncounters() ([]*model.HiddenLocationEncounter, error) {
    // Ovde bi trebalo da izvršimo upit ka bazi podataka ili drugom skladištu podataka da dobijemo sve susrete
    // Na primer, koristeći ORM poput GORM-a, možemo uraditi nešto poput sledećeg:
    var encounters []*model.HiddenLocationEncounter
    if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
        // Ukoliko dođe do greške pri izvršavanju upita, vraćamo grešku
        return nil, err
    }

    return encounters, nil
}

func (r *EncounterRepository) GetAllSocialEncounters() ([]*model.SocialEncounter, error) {
    // Ovde bi trebalo da izvršimo upit ka bazi podataka ili drugom skladištu podataka da dobijemo sve susrete
    // Na primer, koristeći ORM poput GORM-a, možemo uraditi nešto poput sledećeg:
    var encounters []*model.SocialEncounter
    if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
        // Ukoliko dođe do greške pri izvršavanju upita, vraćamo grešku
        return nil, err
    }

    return encounters, nil
}

func (repo *EncounterRepository) Update(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Model(&model.Encounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
		"name":              encounter.Name,
		"description":       encounter.Description,
		"xp_points":         encounter.XpPoints,
		"status":            encounter.Status,
		"type":              encounter.Type,
		"longitude":         encounter.Longitude,
		"latitude":          encounter.Latitude,
		"should_be_approved": encounter.ShouldBeApproved,
	})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("Encounter not found")
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) UpdateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
    dbResult := repo.DatabaseConnection.Model(&model.HiddenLocationEncounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
        "image_url":            encounter.ImageURL,
        "image_latitude":       encounter.ImageLatitude,
        "image_longitude":      encounter.ImageLongitude,
        "distance_treshold":    encounter.DistanceTreshold,
        "encounter_id":         encounter.EncounterId,
        // Dodajte ostale polja ovde prema potrebi
    })
    if dbResult.Error != nil {
        return dbResult.Error
    }
    if dbResult.RowsAffected == 0 {
        return errors.New("HiddenLocationEncounter not found")
    }
    println("Rows affected: ", dbResult.RowsAffected)
    return nil
}

func (repo *EncounterRepository) UpdateSocialEncounter(encounter *model.SocialEncounter) error {
    dbResult := repo.DatabaseConnection.Model(&model.SocialEncounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
        "tourists_required_for_completion": encounter.TouristsRequiredForCompletion,
        "distance_treshold":                encounter.DistanceTreshold,
        "tourist_ids":                      encounter.TouristIDs,
        // Dodajte ostale polja ovde prema potrebi
    })
    if dbResult.Error != nil {
        return dbResult.Error
    }
    if dbResult.RowsAffected == 0 {
        return errors.New("SocialEncounter not found")
    }
    println("Rows affected: ", dbResult.RowsAffected)
    return nil
}

func (r *EncounterRepository) GetSocialEncounterId(baseEncounterID int) (int, error) {
    var socialEncounterID int

    // Izvršavanje upita za dobavljanje ID-a društvenog susreta
    result := r.DatabaseConnection.Model(&model.SocialEncounter{}).Select("Id").Where("encounter_id = ?", baseEncounterID).First(&socialEncounterID)
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
    return socialEncounterID, nil
}

func (r *EncounterRepository) GetHiddenLocationEncounterId(baseEncounterID int) (int, error) {
    var hiddenLocationEncounterID int

    // Izvršavanje upita za dobavljanje ID-a društvenog susreta
    result := r.DatabaseConnection.Model(&model.HiddenLocationEncounter{}).Select("Id").Where("encounter_id = ?", baseEncounterID).First(&hiddenLocationEncounterID)
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
    return hiddenLocationEncounterID, nil
}

func (r *EncounterRepository) DeleteSocialEncounter(socialEncounterID int) error {
    // Izvršavanje SQL upita za brisanje socijalnog susreta na osnovu njegovog ID-ja
    result := r.DatabaseConnection.Exec("DELETE FROM social_encounters WHERE id = ?", socialEncounterID)
    if result.Error != nil {
        // Ukoliko dođe do greške prilikom izvršavanja SQL upita, vraćamo je kao rezultat
        return result.Error
    }
    return nil
}

func (r *EncounterRepository) DeleteHiddenLocationEncounter(hiddenLocationEncounterID int) error {
    // Izvršavanje SQL upita za brisanje skrivenog susreta na osnovu njegovog ID-ja
    result := r.DatabaseConnection.Exec("DELETE FROM hidden_location_encounters WHERE id = ?", hiddenLocationEncounterID)
    if result.Error != nil {
        // Ukoliko dođe do greške prilikom izvršavanja SQL upita, vraćamo je kao rezultat
        return result.Error
    }
    return nil
}

func (r *EncounterRepository) DeleteEncounter(baseEncounterID int) error {
    // Izvršavanje upita za brisanje susreta na osnovu ID-ja
    result := r.DatabaseConnection.Where("id = ?", baseEncounterID).Delete(&model.Encounter{})
    if result.Error != nil {
        // Provera greške prilikom brisanja
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            // Ako ne postoji susret sa datim ID-jem, ne vraćamo grešku, već samo ne obavljamo nikakvo brisanje
            return nil
        }
        // Ako postoji druga greška, vraćamo je
        return result.Error
    }
    return nil
}