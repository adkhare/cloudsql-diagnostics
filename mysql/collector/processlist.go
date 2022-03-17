package collector

import (
	"os"
	"context"
	"database/sql"
	"encoding/json"
)

const (
	// Query.
	processlistQuery = `SELECT ID, USER, DB, COMMAND, TIME, STATE, INFO FROM INFORMATION_SCHEMA.PROCESSLIST`
)

type Processlist struct {
	CollectorType string `json:"COLLECTOR_TYPE"`
	ID 	sql.NullInt64 `json:"ID"`
	User sql.NullString `json:"USER"`
	Db sql.NullString `json:"DB"`
	Command sql.NullString `json:"COMMAND"`
	Time sql.NullInt64 `json:"TIME"`
	State sql.NullString `json:"STATE"`
	Info sql.NullString `json:"INFO"`
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func GetProcesslist(ctx context.Context, db *sql.DB) error {
	var processlists []Processlist
	rows, err := db.QueryContext(ctx, processlistQuery)
	if err != nil {
		return err
	}
	defer rows.Close()


	for rows.Next() {
		var plist Processlist

		if err := rows.Scan(&plist.ID, &plist.User, &plist.Db, &plist.Command, &plist.Time, &plist.State, &plist.Info); err != nil {
			return err
		}

		plist.CollectorType = "Processlist"

		processlists = append(processlists, plist)
	}

	if processlists != nil {
		json.NewEncoder(os.Stdout).Encode(&processlists)
	}
	return nil
}