package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EncounterStatus int

const (
	Draft EncounterStatus = iota
	Archived
	Active
)

type EncounterType int

const (
	Social EncounterType = iota
	Location
	Misc
)

type Encounter struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name             string             `bson:"name" json:"name"`
	Description      string             `bson:"description,omitempty" json:"description"`
	XpPoints         int                `bson:"xpPoints" json:"xpPoints"`
	Status           string             `bson:"status" json:"status"`
	Type             string             `bson:"type" json:"type"`
	Latitude         float64            `bson:"latitude" json:"latitude"`
	Longitude        float64            `bson:"longitude" json:"longitude"`
	ShouldBeApproved bool               `bson:"shouldBeApproved" json:"shouldBeApproved"`
}

type Encounters []*Encounter

func (ens *Encounters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ens)
}

func (en *Encounter) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(en)
}

func (en *Encounter) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(en)
}
