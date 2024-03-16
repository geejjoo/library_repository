package migrate

import (
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/db/tabler"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

type Migrator struct {
	db           *sqlx.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sqlx.DB, dbConf config.DB) *Migrator {
	out := &Migrator{
		db:           db,
		sqlGenerator: nil,
	}
	switch dbConf.Driver {
	case "postgres":
		out.sqlGenerator = &PostgreSQLGenerator{}
	case "sqlite":
		out.sqlGenerator = &SQLiteGenerator{}
	}
	return out
}

func (m *Migrator) Migrate(tables ...tabler.Tabler) error {
	var errGroup errgroup.Group
	for _, table := range tables {
		createSQL := m.sqlGenerator.CreateTableSQL(table)
		errGroup.Go(func() error {
			_, err := m.db.Exec(createSQL)
			return err
		})
	}
	return errGroup.Wait()
}
