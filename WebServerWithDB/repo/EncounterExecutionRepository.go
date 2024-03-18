package repo

import (
	"database-example/model"
	"gorm.io/gorm"
)

type EncounterExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterExecutionRepository) FindByUserId(userID int) (model.EncounterExecution, error) {
	encounter := model.EncounterExecution{}
	dbResult := repo.DatabaseConnection.Where("user_id = ?", userID).Find(&encounter)
	if dbResult.Error != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}
