package model

// BacklogVersion backlog version
type BacklogVersion struct {
	ID             int64  `json:"id"`
	ProjectID      int64  `json:"projectId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	StartDate      string `json:"startDate"`
	ReleaseDueDate string `json:"releaseDueDate"`
	Archived       bool   `json:"archived"`
	DisplayOrder   int64  `json:"diplayOrder"`
}
