package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shuufujita/backlog-tools/usecase"

	"github.com/shuufujita/backlog-tools/common"
	"github.com/shuufujita/backlog-tools/infrastructure/persistance"
)

func init() {
	common.LoadDotEnv()
	err := common.CustomLogger()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	backlogProjectRepository := persistance.NewBacklogProjectPersistance()
	backlogProjectUsecase := usecase.NewBacklogProjectUseCase(backlogProjectRepository)
	log.Println(fmt.Sprintf("%v: [%v] %v", "info", "migrate-project", "execute"))
	err := backlogProjectUsecase.MigrateProject()
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "migrate-project", err.Error()))
		os.Exit(1)
	}
	log.Println(fmt.Sprintf("%v: [%v] %v", "info", "migrate-project", "complete"))
	os.Exit(0)
}
