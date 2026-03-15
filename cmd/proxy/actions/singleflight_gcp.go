//go:build gcp

package actions

import (
	"fmt"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/log"
	"github.com/gomods/athens/pkg/stash"
	"github.com/gomods/athens/pkg/storage"
)

func init() {
	RegisterSingleFlight("gcp", func(_ *log.Logger, c *config.Config, s storage.Backend, _ storage.Checker) (stash.Wrapper, error) {
		if c.StorageType != "gcp" {
			return nil, fmt.Errorf("gcp SingleFlight only works with a gcp storage type and not: %v", c.StorageType)
		}
		return stash.WithGCSLock(c.SingleFlight.GCP.StaleThreshold, s)
	})
}
