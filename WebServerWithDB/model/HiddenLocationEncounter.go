package model

import (
	"encoding/json"
	"io"
)

type HiddenLocationEncounter struct {
	Id               int64   `bson:"_id,omitempty" json:"id"`
	ImageURL         string  `bson:"imageURL" json:"imageURL"`
	ImageLatitude    float64 `bson:"imageLatitude" json:"imageLatitude"`
	ImageLongitude   float64 `bson:"imageLongitude" json:"imageLongitude"`
	DistanceTreshold float64 `bson:"distanceTreshold" json:"distanceTreshold"`
	EncounterId      int64   `bson:"encounterId, omitempty" json:"encounterId"`
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
