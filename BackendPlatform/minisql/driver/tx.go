package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Tx = (*mockTx)(nil)

// mockTx is a mock version of sql.Tx
type mockTx struct {
	driver.Tx
	ctx context.Context
}

// Commit sends a span at the end of the transaction
func (t *mockTx) Commit() (err error) {
	err = t.Tx.Commit()
	return err
}

// Rollback sends a span if the connection is aborted
func (t *mockTx) Rollback() (err error) {
	err = t.Tx.Rollback()
	return err
}
