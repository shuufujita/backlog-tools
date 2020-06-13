package model

// BacklogIssueType backlog issue type
type BacklogIssueType struct {
	ID                  int64  `json:"id"`
	ProjectID           int64  `json:"projectId"`
	Name                string `json:"name"`
	Color               string `json:"color"`
	DisplayOrder        int64  `json:"diplayOrder"`
	TemplateSummary     string `json:"templateSummary"`
	TemplateDescription string `json:"templateDescription"`
}
