package database

import (
	"context"

	"task-manager/app/usecases"
)

type SQLHandler interface {
	NamedExec(context.Context, string, any) (Result, error)
	NamedQuery(context.Context, string, any) (Rows, error)
	Query(context.Context, string, any) (Rows, error)
	SqlError(error) SqlError
	Transaction() usecases.Transaction
	Close() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	StructScan(any) error
	MapScan(map[string]any) error
	Next() bool
	Close() error
}

type SqlError uint

const (
	SqlOK SqlError = iota
	SqlErrDuplicate
	SqlErrUnknown
)
