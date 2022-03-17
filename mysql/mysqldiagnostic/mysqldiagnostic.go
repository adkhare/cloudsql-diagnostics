package mysqldiagnostic

import (
	"context"
	"database/sql"

	"github.com/adkhare/cloudsql-diagnostics/mysql/models"
	"github.com/adkhare/cloudsql-diagnostics/mysql/collector"
	_ "github.com/go-sql-driver/mysql" // nolint
)

type Diagnosis struct {
	DbConfig *models.DbConnConfig
}

//Dianostic method which runs diagnosis on mysql
func (d *Diagnosis) Diagnose(ctx context.Context) error {
	d.DbConfig.MakeInstanceDSN()

	// create connection, but don't open
	db, err := sql.Open("mysql", d.DbConfig.DSN)

	if err != nil {
		return err
	}
	// close db on every test so that we always have to test opening a new connection
	defer db.Close()

	err = collector.GetProcesslist(ctx, db)
	if err != nil {
		return err
	}

	err = collector.GetTransactions(ctx, db)
	if err != nil {
		return err
	}

	err = collector.GetMutexStatus(ctx, db)
	if err != nil {
		return err
	}

	err = collector.GetInnoDbStatus(ctx, db)
	if err != nil {
		return err
	}

	err = collector.GetVariables(ctx, db)
	if err != nil {
		return err
	}
	return nil
}
