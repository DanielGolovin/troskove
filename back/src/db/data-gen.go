package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
)

func genDate() string {
	year := 2021
	randomMonth := rand.Intn(12) + 1
	randomDay := rand.Intn(28) + 1

	return fmt.Sprintf("%04d-%02d-%02d", year, randomMonth, randomDay)
}

func insertData(con *sql.DB) {
	con.Exec("BEGIN TRANSACTION;")

	stmt, _ := con.Prepare("INSERT INTO expenses (id, value, date, type_id) VALUES (?, ?, ?, ?);")

	for i := 0; i < 5000000; i++ {
		stmt.Exec(i, i, genDate(), 1)
		if i%1000 == 0 {
			log.Printf("Inserted: %d", i)
		}
	}

	con.Exec("COMMIT;")
}
