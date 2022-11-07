package models

import "time"

type DateTime struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (dt *DateTime) Omit() {
	dt.CreatedAt = nil
	dt.UpdatedAt = nil
}

func (dt *DateTime) OmitCreatedAt() {
	dt.CreatedAt = nil
}

func (dt *DateTime) OmitUpdatedAt() {
	dt.UpdatedAt = nil
}
