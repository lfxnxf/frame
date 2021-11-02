package driver

import (
	"context"
	"database/sql/driver"
	"errors"
)

var _ driver.Stmt = (*mockStmt)(nil)

// mockStmt is mock version of sql.Stmt
type mockStmt struct {
	driver.Stmt
	ctx   context.Context
	query string
}

// Close sends a span before closing a statement
func (s *mockStmt) Close() (err error) {
	err = s.Stmt.Close()
	return err
}

// ExecContext is needed to implement the driver.StmtExecContext interface
func (s *mockStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (res driver.Result, err error) {
	if stmtExecContext, ok := s.Stmt.(driver.StmtExecContext); ok {
		res, err = stmtExecContext.ExecContext(ctx, args)
		return res, err
	}
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	res, err = s.Exec(dargs)
	return res, err
}

// QueryContext is needed to implement the driver.StmtQueryContext interface
func (s *mockStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (rows driver.Rows, err error) {
	if stmtQueryContext, ok := s.Stmt.(driver.StmtQueryContext); ok {
		rows, err := stmtQueryContext.QueryContext(ctx, args)
		return rows, err
	}
	dargs, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	rows, err = s.Query(dargs)
	return rows, err
}

// copied from stdlib database/sql package: src/database/sql/ctxutil.go
func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			return nil, errors.New("sql: driver does not support the use of Named Parameters")
		}
		dargs[n] = param.Value
	}
	return dargs, nil
}
