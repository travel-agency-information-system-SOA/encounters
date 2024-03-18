package model

import "time"

type EncounterExecution struct {
	ID             int        `json:"id" gorm:"column:Id;primaryKey;autoIncrement"`
	UserID         int        `json:"userId"`
	EncounterID    int        `json:"encounterID"`
	CompletionTime *time.Time `json:"completionTime"`
	IsCompleted    bool       `json:"isCompleted"`
}

//func NewEncounterExecution(userID, encounterID int64, completionTime *time.Time, isCompleted bool) *EncounterExecution {
//	return &EncounterExecution{
//		UserID:         userID,
//		EncounterID:    encounterID,
//		CompletionTime: completionTime,
//		IsCompleted:    isCompleted,
//	}
//}
