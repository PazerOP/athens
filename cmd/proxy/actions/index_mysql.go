//go:build mysql

package actions

import (
	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/index"
	"github.com/gomods/athens/pkg/index/mysql"
)

func init() {
	RegisterIndex("mysql", func(c *config.Config) (index.Indexer, error) {
		return mysql.New(c.Index.MySQL)
	})
}
