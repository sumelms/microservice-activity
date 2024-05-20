package domain

import (
	"time"

	"github.com/google/uuid"
)

// Activity struct.
type Activity struct {
	UUID        uuid.UUID  `json:"uuid"`
	ContentUUID uuid.UUID  `db:"content_uuid"  json:"content_uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ContentType string     `db:"content_type" json:"content_type"`
	Taxonomy    string     `json:"taxonomy"`
	CreatedAt   time.Time  `db:"created_at"    json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"    json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"    json:"deleted_at"`
}
