package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shuufujita/backlog-tools/common"
)

func init() {
	common.LoadDotEnv()
	err := common.CustomLogger()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	log.Println(fmt.Sprintf("execute migration project."))
	os.Exit(0)
}
