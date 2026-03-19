//go:build azureblob

package actions

import (
	"fmt"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/log"
	"github.com/gomods/athens/pkg/stash"
	"github.com/gomods/athens/pkg/storage"
)

func init() {
	RegisterSingleFlight("azureblob", func(_ *log.Logger, c *config.Config, _ storage.Backend, checker storage.Checker) (stash.Wrapper, error) {
		if c.StorageType != "azureblob" {
			return nil, fmt.Errorf("azureblob SingleFlight only works with a azureblob storage type and not: %v", c.StorageType)
		}
		return stash.WithAzureBlobLock(c.Storage.AzureBlob, c.TimeoutDuration(), checker)
	})
}
