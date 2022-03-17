package collector

import (
	"os"
	"context"
	"database/sql"
	"encoding/json"
)

const (
	// Query.
	mutexQuery = `SHOW ENGINE INNODB MUTEX`
)

type MutexStatus struct {
	CollectorType string `json:"COLLECTOR_TYPE"`
	Type sql.NullString `json:"TYPE"`
	Name sql.NullString `json:"NAME"`
	Status sql.NullString `json:"STATUS"`
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func GetMutexStatus(ctx context.Context, db *sql.DB) error {
	var mutexstatuslist []MutexStatus
	rows, err := db.QueryContext(ctx, mutexQuery)
	if err != nil {
		return err
	}
	defer rows.Close()


	for rows.Next() {
		var mstatus MutexStatus

		if err := rows.Scan(&mstatus.Type, &mstatus.Name, &mstatus.Status); err != nil {
			return err
		}

		mstatus.CollectorType = "MutexStatus"

		mutexstatuslist = append(mutexstatuslist, mstatus)
	}

	if mutexstatuslist != nil {
		json.NewEncoder(os.Stdout).Encode(&mutexstatuslist)
	}
	return nil
}