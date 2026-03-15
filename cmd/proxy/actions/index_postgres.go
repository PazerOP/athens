//go:build postgres

package actions

import (
	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/index"
	"github.com/gomods/athens/pkg/index/postgres"
)

func init() {
	RegisterIndex("postgres", func(c *config.Config) (index.Indexer, error) {
		return postgres.New(c.Index.Postgres)
	})
}
