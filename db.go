package walletcore

import (
	"fmt"

	"github.com/nutsdb/nutsdb"
)

type db struct {
	*nutsdb.DB
}

func newDB(dbPath string) (*db, error) {
	dbConn, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(dbPath),
	)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return &db{
		dbConn,
	}, nil
}

func (d *db) close() error {
	if d.DB != nil {
		return d.DB.Close()
	}

	return nil
}
