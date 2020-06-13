package common

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/comail/colog"
)

// CustomLogger custom logger settings
func CustomLogger() error {
	file, err := os.OpenFile(getLogFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "CustomLogger", err.Error()))
		return err
	}
	colog.Register()
	colog.SetOutput(io.MultiWriter(file, os.Stdout))
	colog.SetFormatter(&colog.StdFormatter{
		Flag: log.Ldate | log.Ltime | log.Lshortfile,
	})
	return nil
}

func getLogFilePath() string {
	return os.Getenv("LOG_DIR_PATH") + "/" + time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("20060102") + ".log"
}
