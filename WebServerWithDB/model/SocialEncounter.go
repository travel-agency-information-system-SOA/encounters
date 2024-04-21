package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialEncounter struct {
	Id                            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Encounter                     Encounter          `bson:"encounter,omitempty" json:"encounter"`
	TouristsRequiredForCompletion int                `bson:"touristsRequiredForCompletion" json:"touristsRequiredForCompletion"`
	DistanceTreshold              float64            `bson:"distanceTreshold" json:"distanceTreshold"`
	TouristIDs                    []int              `bson:"touristIDs,omitempty" json:"touristIDs"`
}

type SocialEncounters []*SocialEncounter

func (ens *SocialEncounters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ens)
}

func (en *SocialEncounter) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(en)
}

func (en *SocialEncounter) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(en)
}
