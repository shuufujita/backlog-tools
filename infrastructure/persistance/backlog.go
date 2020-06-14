package persistance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/shuufujita/backlog-tools/domain/model"
	"github.com/shuufujita/backlog-tools/domain/repository"
)

type backlogPersistance struct{}

// NewBacklogPersistance backlog persistance
func NewBacklogPersistance() repository.BacklogRepository {
	return &backlogPersistance{}
}

func (bpp backlogPersistance) GetProject() (model.BacklogProject, error) {
	projectKeyOrID := os.Getenv("BACKLOG_PROJECT_KEY")
	body, err := backlogGetRequest("/api/v2/projects/" + projectKeyOrID)
	if err != nil {
		return model.BacklogProject{}, err
	}

	project := model.BacklogProject{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return model.BacklogProject{}, err
	}

	return project, nil
}

func (bpp backlogPersistance) GetIssueType() ([]model.BacklogIssueType, error) {
	projectKeyOrID := os.Getenv("BACKLOG_PROJECT_KEY")
	body, err := backlogGetRequest("/api/v2/projects/" + projectKeyOrID + "/issueTypes")
	if err != nil {
		return nil, err
	}

	issueTypes := []model.BacklogIssueType{}
	err = json.Unmarshal(body, &issueTypes)
	if err != nil {
		return nil, err
	}

	return issueTypes, nil
}

func (bpp backlogPersistance) GetProjectUsers() ([]model.BacklogProjectUser, error) {
	projectKeyOrID := os.Getenv("BACKLOG_PROJECT_KEY")
	body, err := backlogGetRequest("/api/v2/projects/" + projectKeyOrID + "/users")
	if err != nil {
		return nil, err
	}

	projectUsers := []model.BacklogProjectUser{}
	err = json.Unmarshal(body, &projectUsers)
	if err != nil {
		return nil, err
	}

	return projectUsers, nil
}

func backlogGetRequest(path string) ([]byte, error) {
	requestURL := os.Getenv("BACKLOG_BASE_URL") + path

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Println(fmt.Sprintf("%v [%v] %v", "error", "backlogGetRequest", err.Error()))
		return nil, err
	}

	params := req.URL.Query()
	params.Add("apiKey", os.Getenv("BACKLOG_API_KEY"))
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "backlogGetRequest", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if resp.StatusCode >= 400 {
		log.Println(fmt.Sprintf("%v: [%v] %v", "warn", resp.StatusCode, string(body)))
		return nil, fmt.Errorf("Error: %s", "not success")
	}
	return body, nil
}
