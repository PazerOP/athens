//go:build minio

package actions

import (
	"net/http"
	"time"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/minio"
)

func init() {
	RegisterStorage("minio", func(storageConfig *config.Storage, timeout time.Duration, _ *http.Client) (storage.Backend, error) {
		if storageConfig.Minio == nil {
			return nil, errors.E("actions.GetStorage", "Invalid Minio Storage Configuration")
		}
		return minio.NewStorage(storageConfig.Minio, timeout)
	})
}
