package service

import (
	"database-example/model"
	"database-example/repo"
	"time"
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

func (service *EncounterExecutionService) CompleteEncounter(userID int) (*model.EncounterExecution, error) {
	encounter, err := service.EncounterExecutionRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	
	encounter.CompletionTime = time.Now()
	encounter.IsCompleted = true

	// Neki update XP za korisnika

	err = service.EncounterExecutionRepo.Update(&encounter)
	if err != nil {
		return nil, err
	}
	return &encounter, nil
}
