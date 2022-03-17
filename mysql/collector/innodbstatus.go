package collector

import (
	"os"
	"context"
	"database/sql"
	"encoding/json"
)

const (
	// Query.
	innodbStatusQuery = `SHOW ENGINE INNODB STATUS`
)

type InnodbStatus struct {
	CollectorType string `json:"COLLECTOR_TYPE"`
	Type sql.NullString `json:"TYPE"`
	Name sql.NullString `json:"NAME"`
	Status sql.NullString `json:"STATUS"`
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func GetInnoDbStatus(ctx context.Context, db *sql.DB) error {
	var innodbstatuslist []InnodbStatus
	rows, err := db.QueryContext(ctx, innodbStatusQuery)
	if err != nil {
		return err
	}
	defer rows.Close()


	for rows.Next() {
		var idbstatus InnodbStatus

		if err := rows.Scan(&idbstatus.Type, &idbstatus.Name, &idbstatus.Status); err != nil {
			return err
		}

		idbstatus.CollectorType = "InnoDbStatus"

		innodbstatuslist = append(innodbstatuslist, idbstatus)
	}

	if innodbstatuslist != nil {
		json.NewEncoder(os.Stdout).Encode(&innodbstatuslist)
	}
	return nil
}