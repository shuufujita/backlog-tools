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

type backlogProjectPersistance struct{}

// NewBacklogProjectPersistance backlog project persistance
func NewBacklogProjectPersistance() repository.BacklogProjectRepository {
	return &backlogProjectPersistance{}
}

func (bpp backlogProjectPersistance) GetIssueType() ([]model.BacklogIssueType, error) {
	projectKeyOrID := os.Getenv("BACKLOG_PROJECT_KEY")
	body, err := backlogGetRequest("/api/v2/projects/" + projectKeyOrID + "/issueTypes")
	if err != nil {
		return nil, err
	}

	issues := []model.BacklogIssueType{}
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
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
