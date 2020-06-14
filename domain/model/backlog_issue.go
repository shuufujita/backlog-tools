package model

// BacklogIssue backlog issue
type BacklogIssue struct {
	ID             int64              `json:"id"`
	ProjectID      int64              `json:"projectId"`
	IssueKey       string             `json:"issueKey"`
	KeyID          int64              `json:"keyId"`
	IssueType      BacklogIssueType   `json:"issueType"`
	Summary        string             `json:"summary"`
	Description    string             `json:"description"`
	Resolutions    BacklogResolution  `json:"resolutions"`
	Priority       BacklogPriority    `json:"priority"`
	Status         BacklogStatus      `json:"status"`
	Assignee       BacklogAssignee    `json:"assignee"`
	Category       []BacklogCategory  `json:"category"`
	Versions       []BacklogVersion   `json:"versions"`
	Milestone      []BacklogVersion   `json:"milestone"`
	StartDate      string             `json:"startDate"`
	DueDate        string             `json:"dueDate"`
	EstimatedHours int64              `json:"estimatedHours"`
	ActualHours    int64              `json:"actualHours"`
	ParentIssueID  int64              `json:"parentIssueId"`
	Createduser    BacklogProjectUser `json:"createdUser"`
	Created        string             `json:"created"`
	UpdatedUser    BacklogProjectUser `json:"updatedUser"`
	Updated        string             `json:"updated"`
	// CustomFields   []string           `json:"customFields"`
	// Attachments    []string           `json:"attachments"`
}

// BacklogStatus backlog status
type BacklogStatus struct {
	ID           int64  `json:"id"`
	ProjectID    int64  `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int64  `json:"displayOrder"`
}

// BacklogAssignee backlog assignee
type BacklogAssignee struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	RoleType    int64  `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}
