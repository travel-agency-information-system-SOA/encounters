package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HiddenLocationEncounter struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ImageURL         string             `bson:"imageURL" json:"imageURL"`
	ImageLatitude    float64            `bson:"imageLatitude" json:"imageLatitude"`
	ImageLongitude   float64            `bson:"imageLongitude" json:"imageLongitude"`
	DistanceTreshold float64            `bson:"distanceTreshold" json:"distanceTreshold"`
	Encounter        Encounter          `bson:"encounter,omitempty" json:"encounter"`
}

type HiddenLocationEncounters []*HiddenLocationEncounter

func (ens *HiddenLocationEncounters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ens)
}

func (en *HiddenLocationEncounter) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(en)
}

func (en *HiddenLocationEncounter) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(en)
}
