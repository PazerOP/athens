//go:build s3

package actions

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/s3"
)

func init() {
	RegisterStorage("s3", func(storageConfig *config.Storage, timeout time.Duration, client *http.Client) (storage.Backend, error) {
		if storageConfig.S3 == nil {
			return nil, errors.E("actions.GetStorage", "Invalid S3 Storage Configuration")
		}
		return s3.New(storageConfig.S3, timeout, func(config *aws.Config) {
			config.HTTPClient = client
		})
	})
}
