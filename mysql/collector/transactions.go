package collector

import (
	"os"
	"context"
	"database/sql"
	"encoding/json"
)

const (
	// Query.
	transactionsQuery = `SELECT trx_id,
								trx_state,
								trx_started,
								trx_requested_lock_id,
								trx_wait_started,
								trx_weight,
								trx_mysql_thread_id,
								trx_query,
								trx_operation_state,
								trx_tables_in_use,
								trx_tables_locked,
								trx_lock_structs,
								trx_lock_memory_bytes,
								trx_rows_locked,
								trx_rows_modified,
								trx_concurrency_tickets,
								trx_isolation_level,
								trx_unique_checks,
								trx_foreign_key_checks,
								trx_last_foreign_key_error,
								trx_adaptive_hash_latched,
								trx_adaptive_hash_timeout,
								trx_is_read_only,
								trx_autocommit_non_locking
						FROM INFORMATION_SCHEMA.INNODB_TRX`
)

type Transactions struct {
	CollectorType string `json:"COLLECTOR_TYPE"`
	TrxID 	sql.NullInt64 `json:"TRX_ID"`
	TrxState sql.NullString `json:"TRX_STATE"`
	TrxStarted sql.NullString `json:"TRX_STARTED"`
	TrxRequestedLockID sql.NullString `json:"TRX_REQUESTED_LOCK_ID"`
	TrxWaitStarted sql.NullString `json:"TRX_WAIT_STARTED"`
	TrxWeight 	sql.NullInt64 `json:"TRX_WEIGHT"`
	TrxMysqlThreadID 	sql.NullInt64 `json:"TRX_MYSQL_THREAD_ID"`
	TrxQuery sql.NullString `json:"TRX_QUERY"`
	TrxOperationState 	sql.NullString `json:"TRX_OPERATION_STATE"`
	TrxTablesInUse 	sql.NullInt64 `json:"TRX_TABLES_IN_USE"`
	TrxTablesLocked 	sql.NullInt64 `json:"TRX_TABLES_LOCKED"`
	TrxLockStructs 	sql.NullInt64 `json:"TRX_LOCK_STRUCTS"`
	TrxLockMemoryBytes 	sql.NullInt64 `json:"TRX_LOCK_MEMORY_BYTES"`
	TrxRowsLocked 	sql.NullInt64 `json:"TRX_ROWS_LOCKED"`
	TrxRowsModified 	sql.NullInt64 `json:"TRX_ROWS_MODIFIED"`
	TrxConcurrencyTickets 	sql.NullInt64 `json:"TRX_CONCURRENCY_TICKETS"`
	TrxIsolationLevel 	sql.NullString `json:"TRX_ISOLATION_LEVEL"`
	TrxUniqueChecks 	sql.NullInt64 `json:"TRX_UNIQUE_CHECKS"`
	TrxForeignKeyChecks 	sql.NullInt64 `json:"TRX_FOREIGN_KEY_CHECKS"`
	TrxLastForeignKeyError 	sql.NullString `json:"TRX_LAST_FOREIGN_KEY_ERROR"`
	TrxAdaptiveHashLatched 	sql.NullInt64 `json:"TRX_ADAPTIVE_HASH_LATCHED"`
	TrxAdaptiveHashTimeout 	sql.NullInt64 `json:"TRX_ADAPTIVE_HASH_TIMEOUT"`
	TrxIsReadOnly 	sql.NullInt64 `json:"TRX_IS_READ_ONLY"`
	TrxAutoCommitNonLocking 	sql.NullInt64 `json:"TRX_AUTO_COMMIT_NON_LOCKING"`
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func GetTransactions(ctx context.Context, db *sql.DB) error {
	var transactions []Transactions
	rows, err := db.QueryContext(ctx, transactionsQuery)
	if err != nil {
		return err
	}
	defer rows.Close()


	for rows.Next() {
		var transaction Transactions

		if err := rows.Scan(
				&transaction.TrxID,
				&transaction.TrxState,
				&transaction.TrxStarted,
				&transaction.TrxRequestedLockID,
				&transaction.TrxWaitStarted,
				&transaction.TrxWeight,
				&transaction.TrxMysqlThreadID,
				&transaction.TrxQuery,
				&transaction.TrxOperationState,
				&transaction.TrxTablesInUse,
				&transaction.TrxTablesLocked,
				&transaction.TrxLockStructs,
				&transaction.TrxLockMemoryBytes,
				&transaction.TrxRowsLocked,
				&transaction.TrxRowsModified,
				&transaction.TrxConcurrencyTickets,
				&transaction.TrxIsolationLevel,
				&transaction.TrxUniqueChecks,
				&transaction.TrxForeignKeyChecks,
				&transaction.TrxLastForeignKeyError,
				&transaction.TrxAdaptiveHashLatched,
				&transaction.TrxAdaptiveHashTimeout,
				&transaction.TrxIsReadOnly,
				&transaction.TrxAutoCommitNonLocking,
			); err != nil {
			return err
		}

		transaction.CollectorType = "Transactions"

		transactions = append(transactions, transaction)
	}

	if transactions != nil {
		json.NewEncoder(os.Stdout).Encode(&transactions)
	}
	return nil
}