package financial_management_infrastructure

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

func GetDBConnection() *sql.DB {
	dbPath := getDBPath()
	dbPathDSN := getPathDSN(dbPath)

	db, err := sql.Open("sqlite", dbPathDSN)

	if err != nil {
		log.Println("Database path: ", dbPath)
		log.Fatalln(err)
	}

	return db
}

func SetupDB() {
	ensureDBDir(getDBPath())
	db := GetDBConnection()
	defer db.Close()

	err := InitTransactionCategoryTable(db)

	if err != nil {
		log.Fatalln(err)
	}

	err = InitTransactionTable(db)

	if err != nil {
		log.Fatalln(err)
	}
}

func ensureDBDir(dbPath string) {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create directory for database: %v", err)
	}
}

func getPathDSN(dbPath string) string {
	return "file:" + dbPath + "?_foreign_keys=on"
}

func getDBPath() string {
	path := os.Getenv("SQLITE_DB_PATH")

	if path == "" {
		log.Fatalln("SQLITE_DB_PATH is not set")
	}

	absolutePath, err := filepath.Abs(path)

	if err != nil {
		log.Fatalln(err)
	}

	return absolutePath
}
