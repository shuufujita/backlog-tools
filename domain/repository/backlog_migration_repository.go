package repository

import "github.com/shuufujita/backlog-tools/domain/model"

// BacklogMigrationRepository backlog migration repository
type BacklogMigrationRepository interface {
	SaveProject(model.BacklogProject) error
	SaveIssueType([]model.BacklogIssueType) error
	SaveProjectUsers(projectUsers []model.BacklogProjectUser, projectID int64) error
}
