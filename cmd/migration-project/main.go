package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shuufujita/backlog-tools/usecase"

	_ "github.com/go-sql-driver/mysql"
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
	log.Println(fmt.Sprintf("%v: [%v] %v", "info", "migrate-project", "execute"))

	err := persistance.MySQLInit()
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "migrate-project", err.Error()))
		os.Exit(1)
	}
	defer persistance.CloseMySQL()

	backlogRepository := persistance.NewBacklogPersistance()
	backlogMigrationRepository := persistance.NewBacklogMigrationPersistance()
	backlogUsecase := usecase.NewBacklogUseCase(backlogRepository, backlogMigrationRepository)
	err = backlogUsecase.LoadProject()
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "migrate-project", err.Error()))
		os.Exit(1)
	}

	log.Println(fmt.Sprintf("%v: [%v] %v", "info", "migrate-project", "complete"))
	os.Exit(0)
}
