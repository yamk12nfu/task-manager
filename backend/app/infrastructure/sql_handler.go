package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"task-manager/app/config"
	"task-manager/app/interface/database"
	"task-manager/app/usecases"
)

var (
	driverName     = "mysql"
	dataSourceName = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/taskmanger-admin?parseTime=true",
		config.Conf.DB.User, config.Conf.DB.Password, config.Conf.DB.Host, config.Conf.DB.Port,
	)
)

func NewSqlHander() database.SQLHandler {
	db := sqlx.MustConnect(driverName, dataSourceName)

	return &sqlHandler{db: db, tx: &transaction{db: db}}
}

type sqlHandler struct {
	db *sqlx.DB
	tx *transaction
}

type transaction struct {
	db *sqlx.DB
}

type txKey struct{}

type executer interface {
	Exec(string, ...any) (sql.Result, error)
	Queryx(string, ...any) (*sqlx.Rows, error)
	Rebind(string) string
}

func (h *sqlHandler) NamedExec(ctx context.Context, query string, arg any) (database.Result, error) {
	ex := h.getExecuter(ctx)

	query, args, err := h.namedIn(ex, query, arg)
	if err != nil {
		return nil, err
	}

	return ex.Exec(query, args...)
}

func (h *sqlHandler) NamedQuery(ctx context.Context, query string, arg any) (database.Rows, error) {
	ex := h.getExecuter(ctx)

	query, args, err := h.namedIn(ex, query, arg)
	if err != nil {
		return nil, err
	}

	return ex.Queryx(query, args...)
}

func (h *sqlHandler) Query(ctx context.Context, query string, arg any) (database.Rows, error) {
	ex := h.getExecuter(ctx)

	query, args, err := h.in(ex, query, arg)
	if err != nil {
		return nil, err
	}

	return ex.Queryx(query, args...)
}

func (h *sqlHandler) getExecuter(ctx context.Context) executer {
	if tx := ctx.Value(txKey{}); tx != nil {
		return tx.(*sqlx.Tx)
	}

	return h.db
}

func (h *sqlHandler) namedIn(ex executer, query string, arg any) (string, []any, error) {
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return query, args, err
	}

	return h.in(ex, query, args...)
}

func (h *sqlHandler) in(ex executer, query string, args ...any) (string, []any, error) {
	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return query, args, err
	}

	query = ex.Rebind(query)

	return query, args, nil
}

func (h *sqlHandler) SqlError(err error) database.SqlError {
	if err == nil {
		return database.SqlOK
	}

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		switch driverErr.Number {
		case mysqlerr.ER_DUP_ENTRY:
			return database.SqlErrDuplicate
		}
	}

	return database.SqlErrUnknown
}

func (h *sqlHandler) Transaction() usecases.Transaction {
	return h.tx
}

func (h *sqlHandler) Close() error {
	return h.db.Close()
}

func (t *transaction) Do(ctx context.Context, fn func(context.Context) error) (err error) {
	tx := t.db.MustBegin()
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			err = tx.Rollback()
		}
	}()

	return
}
