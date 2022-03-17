package collector

import (
	"os"
	"context"
	"database/sql"
	"encoding/json"
)

const (
	// Query.
	variablesQuery = `SHOW VARIABLES`
)

type Variables struct {
	CollectorType string `json:"COLLECTOR_TYPE"`
	VariableName sql.NullString `json:"VARIABLE_NAME"`
	Value sql.NullString `json:"VALUE"`
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func GetVariables(ctx context.Context, db *sql.DB) error {
	var variableslist []Variables
	rows, err := db.QueryContext(ctx, variablesQuery)
	if err != nil {
		return err
	}
	defer rows.Close()


	for rows.Next() {
		var vlist Variables

		if err := rows.Scan(&vlist.VariableName, &vlist.Value); err != nil {
			return err
		}

		vlist.CollectorType = "Variables"

		variableslist = append(variableslist, vlist)
	}

	if variableslist != nil {
		json.NewEncoder(os.Stdout).Encode(&variableslist)
	}
	return nil
}