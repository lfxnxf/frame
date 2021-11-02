package driver

import (
	"context"
	"database/sql/driver"
)

var _ driver.Conn = (*mockConn)(nil)

type mockConn struct {
	driver.Conn
}

func (tc *mockConn) BeginTx(ctx context.Context, opts driver.TxOptions) (tx driver.Tx, err error) {
	if connBeginTx, ok := tc.Conn.(driver.ConnBeginTx); ok {
		tx, err = connBeginTx.BeginTx(ctx, opts)
		if err != nil {
			return nil, err
		}
		return &mockTx{tx, ctx}, nil
	}
	tx, err = tc.Conn.Begin()
	if err != nil {
		return nil, err
	}
	return &mockTx{tx, ctx}, nil
}

func (tc *mockConn) PrepareContext(ctx context.Context, query string) (stmt driver.Stmt, err error) {
	if connPrepareCtx, ok := tc.Conn.(driver.ConnPrepareContext); ok {
		stmt, err := connPrepareCtx.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}
		return &mockStmt{stmt, ctx, query}, nil
	}
	stmt, err = tc.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &mockStmt{stmt, ctx, query}, nil
}

func (tc *mockConn) Exec(query string, args []driver.Value) (res driver.Result, err error) {
	if execer, ok := tc.Conn.(driver.Execer); ok {
		return execer.Exec(query, args)
	}
	return nil, driver.ErrSkip
}

func (tc *mockConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (r driver.Result, err error) {
	if execContext, ok := tc.Conn.(driver.ExecerContext); ok {
		r, err := execContext.ExecContext(ctx, query, args)
		return r, err
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
	r, err = tc.Exec(query, dargs)
	return r, err
}

// mockConn has a Ping method in order to implement the pinger interface
func (tc *mockConn) Ping(ctx context.Context) (err error) {
	if pinger, ok := tc.Conn.(driver.Pinger); ok {
		err = pinger.Ping(ctx)
	}
	return err
}

func (tc *mockConn) Query(query string, args []driver.Value) (row driver.Rows, err error) {
	if queryer, ok := tc.Conn.(driver.Queryer); ok {
		return queryer.Query(query, args)
	}
	return nil, driver.ErrSkip
}

func (tc *mockConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	if queryerContext, ok := tc.Conn.(driver.QueryerContext); ok {
		rows, err = queryerContext.QueryContext(ctx, query, args)
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
	rows, err = tc.Query(query, dargs)
	return rows, err
}
