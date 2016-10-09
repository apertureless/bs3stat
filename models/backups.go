package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Backup struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Started  string `json:"started"`
	Finished string `json:"finished"`
	Duration string `json:"duration"`
	Status   string `json:"status"`
}

type BackupCollection struct {
	Backups []Backup `json:"items"`
}

func GetBackups(db *sql.DB) BackupCollection {
	sql := "SELECT * FROM backups"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}
	// Cleanup if exit
	defer rows.Close()

	result := BackupCollection{}

	for rows.Next() {
		backup := Backup{}
		err2 := rows.Scan(&backup.ID, &backup.Name)

		if err2 != nil {
			panic(err2)
		}

		result.Backups = append(result.Backups, backup)
	}
	return result
}

func PutBackup(db *sql.DB, name string, starting string, finished string, duration string, status string) (int64, error) {
	sql := "INSERT INTO backups(name, starting, finished, duration, status) VALUES(?,?,?,?,?)"

	// Create prepared sql statement
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name)

	if err2 != nil {
		panic(err)
	}

	return result.LastInsertId()

}
