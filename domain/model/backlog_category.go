package model

// BacklogCategory backlog category
type BacklogCategory struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	DisplayOrder int64  `json:"diplayOrder"`
}
