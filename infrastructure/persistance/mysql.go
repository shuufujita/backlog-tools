package persistance

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var dbCon *sql.DB
var once sync.Once

// MySQLInit MySQL initialization
func MySQLInit() error {
	connectionParams := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DBNAME")
	db, err := sql.Open("mysql", connectionParams)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "MySQLInit", err.Error()))
		return err
	}

	dbCon = db
	dbCon.SetMaxIdleConns(50)
	dbCon.SetMaxOpenConns(50)
	dbCon.SetConnMaxLifetime(120 * time.Second)

	return nil
}

// CloseMySQL close MySQL
func CloseMySQL() error {
	if dbCon != nil {
		return dbCon.Close()
	}
	return nil
}
