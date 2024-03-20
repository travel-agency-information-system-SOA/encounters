package model

import "time"

type EncounterExecution struct {
	ID             int       `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	UserID         int       `json:"userId"`
	EncounterID    int       `json:"encounterID"`
	CompletionTime time.Time `json:"completionTime"`
	IsCompleted    bool      `json:"isCompleted"`
}
