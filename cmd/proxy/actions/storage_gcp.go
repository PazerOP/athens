//go:build gcp

package actions

import (
	"context"
	"net/http"
	"time"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/gcp"
)

func init() {
	RegisterStorage("gcp", func(storageConfig *config.Storage, timeout time.Duration, _ *http.Client) (storage.Backend, error) {
		if storageConfig.GCP == nil {
			return nil, errors.E("actions.GetStorage", "Invalid GCP Storage Configuration")
		}
		return gcp.New(context.Background(), storageConfig.GCP, timeout)
	})
}
