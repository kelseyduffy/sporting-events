package models

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Team struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	FoundedYear   string         `json:"founded_year"`
	DissolvedYear sql.NullString `json:"dissolved_year"`
	Sport         string         `json:"sport"`
}
type TeamList struct {
	Teams []Team `json:"teams"`
}

func (i *Team) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	if i.FoundedYear == "" {
		return fmt.Errorf("founded year is a required field")
	}
	if i.Sport == "" {
		return fmt.Errorf("sport is a required field")
	}
	return nil
}
func (*TeamList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Team) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
