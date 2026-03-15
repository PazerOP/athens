//go:build azureblob

package actions

import (
	"net/http"
	"time"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/azureblob"
)

func init() {
	RegisterStorage("azureblob", func(storageConfig *config.Storage, timeout time.Duration, _ *http.Client) (storage.Backend, error) {
		if storageConfig.AzureBlob == nil {
			return nil, errors.E("actions.GetStorage", "Invalid AzureBlob Storage Configuration")
		}
		return azureblob.New(storageConfig.AzureBlob, timeout)
	})
}
