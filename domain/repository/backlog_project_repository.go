package repository

import "github.com/shuufujita/backlog-tools/domain/model"

// BacklogProjectRepository backlog project repository
type BacklogProjectRepository interface {
	GetIssueType() ([]model.BacklogIssueType, error)
}
