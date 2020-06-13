package usecase

import (
	"fmt"
	"log"

	"github.com/shuufujita/backlog-tools/domain/repository"
)

// BacklogProjectUseCase backlog project usecase
type BacklogProjectUseCase interface {
	MigrateProject() error
}

type backlogProjectUseCase struct {
	repository repository.BacklogProjectRepository
}

// NewBacklogProjectUseCase return backlog project usecase instance
func NewBacklogProjectUseCase(bpr repository.BacklogProjectRepository) BacklogProjectUseCase {
	return &backlogProjectUseCase{
		repository: bpr,
	}
}

func (bpu backlogProjectUseCase) MigrateProject() error {
	issues, err := bpu.repository.GetIssueType()
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("issues : %v", issues))
	return nil
}
