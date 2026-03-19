//go:build mongo

package actions

import (
	"net/http"
	"time"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/mongo"
)

func init() {
	RegisterStorage("mongo", func(storageConfig *config.Storage, timeout time.Duration, _ *http.Client) (storage.Backend, error) {
		if storageConfig.Mongo == nil {
			return nil, errors.E("actions.GetStorage", "Invalid Mongo Storage Configuration")
		}
		return mongo.NewStorage(storageConfig.Mongo, timeout)
	})
}
