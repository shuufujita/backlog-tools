package usecase

import (
	"fmt"
	"log"

	"github.com/shuufujita/backlog-tools/domain/repository"
)

// BacklogProjectUseCase backlog project usecase
type BacklogProjectUseCase interface {
	LoadProject() error
}

type backlogUseCase struct {
	repository          repository.BacklogRepository
	migrationRepository repository.BacklogMigrationRepository
}

// NewBacklogUseCase return backlog usecase instance
func NewBacklogUseCase(br repository.BacklogRepository, bmr repository.BacklogMigrationRepository) BacklogProjectUseCase {
	return &backlogUseCase{
		repository:          br,
		migrationRepository: bmr,
	}
}

func (bu backlogUseCase) LoadProject() error {
	project, err := bu.repository.GetProject()
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("project : %v", project))

	err = bu.migrationRepository.SaveProject(project)
	if err != nil {
		return err
	}

	projectUsers, err := bu.repository.GetProjectUsers()
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("project_users : %v", projectUsers))

	err = bu.migrationRepository.SaveProjectUsers(projectUsers, project.ID)
	if err != nil {
		return err
	}

	issueTypes, err := bu.repository.GetIssueType()
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("issue_types : %v", issueTypes))

	err = bu.migrationRepository.SaveIssueType(issueTypes)
	if err != nil {
		return err
	}

	return nil
}
