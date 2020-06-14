package repository

import "github.com/shuufujita/backlog-tools/domain/model"

// BacklogMigrationRepository backlog migration repository
type BacklogMigrationRepository interface {
	SaveIssues([]model.BacklogIssue) error
	SaveResolutions(resolutions []model.BacklogResolution, projectID int64) error
	SavePriorities(priorities []model.BacklogPriority, projectID int64) error
	SaveCategory(categories []model.BacklogCategory, projectID int64) error
	SaveVersion([]model.BacklogVersion) error
	SaveProject(model.BacklogProject) error
	SaveIssueType([]model.BacklogIssueType) error
	SaveProjectUsers(projectUsers []model.BacklogProjectUser, projectID int64) error
}
