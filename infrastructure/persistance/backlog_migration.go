package persistance

import (
	"github.com/shuufujita/backlog-tools/domain/model"
	"github.com/shuufujita/backlog-tools/domain/repository"
)

type backlogMigrationPersistance struct{}

// NewBacklogMigrationPersistance backlog migration persistance
func NewBacklogMigrationPersistance() repository.BacklogMigrationRepository {
	return &backlogMigrationPersistance{}
}

func (bmp backlogMigrationPersistance) SaveProject(project model.BacklogProject) error {
	tx, err := dbCon.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO projects (`project_id`, `project_key`) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		project.ID,
		project.ProjectKey,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (bmp backlogMigrationPersistance) SaveProjectUsers(projectUsers []model.BacklogProjectUser, projectID int64) error {
	tx, err := dbCon.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO project_users (`project_user_id`, `project_id`, `user_id`, `name`, `mail_address`, `lang`, `role_type`) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	for i := 0; i < len(projectUsers); i++ {
		_, err = stmt.Exec(
			projectUsers[i].ID,
			projectID,
			projectUsers[i].UserID,
			projectUsers[i].Name,
			projectUsers[i].MailAddress,
			projectUsers[i].Lang,
			projectUsers[i].RoleType,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (bmp backlogMigrationPersistance) SaveIssueType(issueTypes []model.BacklogIssueType) error {
	tx, err := dbCon.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO issue_types (`issue_type_id`, `project_id`, `name`, `color`, `display_order`) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	for i := 0; i < len(issueTypes); i++ {
		_, err = stmt.Exec(
			issueTypes[i].ID,
			issueTypes[i].ProjectID,
			issueTypes[i].Name,
			issueTypes[i].Color,
			issueTypes[i].DisplayOrder,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
