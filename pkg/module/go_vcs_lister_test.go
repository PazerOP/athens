package module

import (
	"context"
	"testing"
	"time"

	"github.com/gomods/athens/pkg/errors"
	"github.com/spf13/afero"
)

func TestVCSListerInvalidModulePaths(t *testing.T) {
	t.Parallel()
	lister := NewVCSLister("go", nil, afero.NewMemMapFs(), 30*time.Second)

	tests := []struct {
		name string
		mod  string
	}{
		{"bare host", "github.com"},
		{"host with owner only", "github.com/owner"},
		{"empty string", ""},
		{"single element", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, _, err := lister.List(context.Background(), tt.mod)
			if err == nil {
				t.Fatalf("expected error for module path %q, got nil", tt.mod)
			}
			if kind := errors.Kind(err); kind != errors.KindNotFound {
				t.Errorf("expected KindNotFound (%d) for module path %q, got %d", errors.KindNotFound, tt.mod, kind)
			}
		})
	}
}
