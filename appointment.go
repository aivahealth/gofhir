package fhir

import (
	"encoding/json"
	"fmt"
)

func (c *Connection) GetAppointmentsByPatient(patientId string) (*Appointment, error) {
	b, err := c.Query(fmt.Sprintf("Appointment?patient=%s", patientId))
	if err != nil {
		return nil, err
	}
	data := Appointment{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

type Appointment struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			Identifier      []Identifier  `json:"identifier"`
			Type            Concept       `json:"type"`
			Reason          Concept       `json:"reason"`
			Priority        int           `json:"priority"` // unsignedInt
			Description     string        `json:"description"`
			Start           int           `json:"start"`           // instant
			End             int           `json:"end"`             // instant
			MinutesDuration int           `json:"minutesDuration"` // positiveInt
			Slot            []Thing       `json:"slot"`
			Comment         string        `json:"comment"`
			Participant     []Participant `json:"participant"`
		} `json:"resource"`
	} `json:"entry"`
}
