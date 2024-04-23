package postgresql

import (
	"cars_catalog/internal/repository/transaction"

	"github.com/jmoiron/sqlx"
)

func SqlxTx(ts transaction.Session) *sqlx.Tx {
	tx := ts.Tx()
	return tx.(*sqlx.Tx)
}
