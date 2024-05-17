package entities

import (
	"time"

	"github.com/google/uuid"
)

// Entity represents base for every persistent entity like Account or Item.
type Entity struct {
	ID        uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

// Page is generic struct that represents response made by page request.
type Page[T any] struct {
	TotalPages    int `json:"total_pages"`
	TotalElements int `json:"total_elements"`
	Elements      []T `json:"elements"`
}
