package repository

import "github.com/shuufujita/backlog-tools/domain/model"

// BacklogRepository backlog repository
type BacklogRepository interface {
	GetProject() (model.BacklogProject, error)
	GetIssueType() ([]model.BacklogIssueType, error)
	GetProjectUsers() ([]model.BacklogProjectUser, error)
}
