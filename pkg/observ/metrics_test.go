package observ

import (
	"context"
	"testing"
	"time"

	"go.opencensus.io/stats/view"
	"github.com/wow-look-at-my/testify/require"
)

func TestCacheLookupMetric(t *testing.T) {
	// Register only the cache view
	require.NoError(t, view.Register(cacheLookupView))

	defer view.Unregister(cacheLookupView)

	ctx := context.Background()

	RecordCacheLookup(ctx, "hit", "info")

	rows, err := view.RetrieveData("cache_lookup_total")
	require.Nil(t, err)

	require.Equal(t, 1, len(rows))

	count := rows[0].Data.(*view.CountData).Value
	require.Equal(t, int64(1), count)

}

func TestUpstreamFetchCounter(t *testing.T) {
	require.NoError(t, view.Register(upstreamFetchView))

	defer view.Unregister(upstreamFetchView)

	ctx := context.Background()

	RecordUpstreamFetch(ctx, "success")

	rows, err := view.RetrieveData("upstream_fetch_total")
	require.Nil(t, err)

	require.Equal(t, 1, len(rows))

	count := rows[0].Data.(*view.CountData).Value
	require.Equal(t, int64(1), count)

}

func TestUpstreamFetchDurationHistogram(t *testing.T) {
	require.NoError(t, view.Register(upstreamFetchLatencyView))

	defer view.Unregister(upstreamFetchLatencyView)

	ctx := context.Background()

	RecordUpstreamFetchDuration(ctx, "success", 2*time.Second)

	rows, err := view.RetrieveData("upstream_fetch_duration_seconds")
	require.Nil(t, err)

	require.Equal(t, 1, len(rows))

	dist := rows[0].Data.(*view.DistributionData)

	require.Equal(t, int64(1), dist.Count)

}
