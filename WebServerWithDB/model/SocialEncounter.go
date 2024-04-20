package model

import (
	"encoding/json"
	"io"
)

type SocialEncounter struct {
	Id                            int64   `bson:"_id,omitempty" json:"id"`
	EncounterId                   int64   `bson:"encounterId, omitempty" json:"encounterId"`
	TouristsRequiredForCompletion int     `bson:"touristsRequiredForCompletion" json:"touristsRequiredForCompletion"`
	DistanceTreshold              float64 `bson:"distanceTreshold" json:"distanceTreshold"`
	TouristIDs                    []int   `bson:"touristIDs,omitempty" json:"touristIDs"`
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
