package models

import "time"

// TypeOfResource use to set status and type of resources, like LinkCategory
type TypeOfResource int

const (
	// ActivePublic active and available to all users
	ActivePublic TypeOfResource = iota + 1
	// ActivePrivate active, but only for owner
	ActivePrivate
	// DeletedPublic deleted
	DeletedPublic
	// DeletedPrivate deleted
	DeletedPrivate
)

// Response is base JSON response
type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// User represents user who is able to use VsLinks system
type User struct {
	UserID   string         `json:"user_id"`
	StatusID TypeOfResource `json:"status_id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
}

// Link is model for create/transfer links
type Link struct {
	LinkID         string    `json:"link_id"`
	UserID         string    `json:"user_id"`
	TypeID         int       `json:"type_id"`
	LinkCategoryID string    `json:"link_category_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	URL            string    `json:"url"`
	Rate           int       `json:"rate"`
	CreatedAt      time.Time `json:"created_at"`
}

// LinkCategory to store information about categories
type LinkCategory struct {
	LinkCategoryID string         `json:"link_category_id"`
	StatusID       TypeOfResource `json:"status_id"`
	UserID         string         `json:"user_id"`
	Name           string         `json:"name"`
	CreatedAt      time.Time      `json:"created_at"`
}
