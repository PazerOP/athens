package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gomods/athens/pkg/config"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
	"github.com/gomods/athens/pkg/storage/external"
	"github.com/gomods/athens/pkg/storage/fs"
	"github.com/gomods/athens/pkg/storage/mem"
	"github.com/spf13/afero"
)

// StorageFactory creates a storage.Backend from configuration.
type StorageFactory func(storageConfig *config.Storage, timeout time.Duration, client *http.Client) (storage.Backend, error)

// storageRegistry holds registered storage backend factories.
// Backends register themselves via init() in build-tagged files.
var storageRegistry = map[string]StorageFactory{}

// RegisterStorage registers a storage backend factory for a given type name.
func RegisterStorage(name string, factory StorageFactory) {
	storageRegistry[name] = factory
}

// GetStorage returns storage backend based on env configuration.
func GetStorage(storageType string, storageConfig *config.Storage, timeout time.Duration, client *http.Client) (storage.Backend, error) {
	const op errors.Op = "actions.GetStorage"
	switch storageType {
	case "memory":
		return mem.NewStorage()
	case "disk":
		if storageConfig.Disk == nil {
			return nil, errors.E(op, "Invalid Disk Storage Configuration")
		}
		rootLocation := storageConfig.Disk.RootPath
		s, err := fs.NewStorage(rootLocation, afero.NewOsFs())
		if err != nil {
			errStr := fmt.Sprintf("could not create new storage from os fs (%s)", err)
			return nil, errors.E(op, errStr)
		}
		return s, nil
	case "external":
		if storageConfig.External == nil {
			return nil, errors.E(op, "Invalid External Storage Configuration")
		}
		return external.NewClient(storageConfig.External.URL, client), nil
	default:
		// Check the registry for build-tagged backends
		if factory, ok := storageRegistry[storageType]; ok {
			return factory(storageConfig, timeout, client)
		}
		return nil, fmt.Errorf("storage type %q is unknown (you may need to build with -tags %s)", storageType, storageType)
	}
}
