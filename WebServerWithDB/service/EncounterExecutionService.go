package service

import (
	"database-example/model"
	"database-example/repo"
)

type EncounterExecutionService struct {
	EncounterExecutionRepo *repo.EncounterExecutionRepository
}

func (service *EncounterExecutionService) GetExecutionByUser(userID int) (*model.EncounterExecution, error) {
	encounter, err := service.EncounterExecutionRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	if encounter.IsCompleted {
		return nil, nil
	}
	return &encounter, nil
}
