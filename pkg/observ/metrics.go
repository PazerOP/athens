package observ

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	cacheResult = tag.MustNewKey("cache_result")
	cacheType   = tag.MustNewKey("cache_type")
	fetchResult = tag.MustNewKey("fetch_result")
)

var (
	cacheStats                 = stats.Int64("cache_lookup_total", "Count of cache lookup results", stats.UnitDimensionless)
	upstreamFetchStats         = stats.Int64("upstream_fetch_total", "Count of upstream fetch attempts", stats.UnitDimensionless)
	upstreamFetchDurationStats = stats.Float64("upstream_fetch_duration_seconds", "Distribution of upstream fetch latency in seconds", stats.UnitSeconds)
	cacheHitsStats             = stats.Int64("athens_cache_hits_total", "Module served from disk storage", stats.UnitDimensionless)
	cacheMissesStats           = stats.Int64("athens_cache_misses_total", "Module fetched from upstream", stats.UnitDimensionless)
	bytesServedStats           = stats.Int64("athens_bytes_served_total", "Bytes served to clients", stats.UnitBytes)
	bytesFetchedStats          = stats.Int64("athens_bytes_fetched_total", "Bytes fetched from upstream", stats.UnitBytes)
)

var upstreamExponentialBuckets = []float64{0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 30}

var (
	cacheLookupView = &view.View{
		Name:        "cache_lookup_total",
		Measure:     cacheStats,
		Description: "Count of cache lookup results",
		TagKeys:     []tag.Key{cacheResult, cacheType},
		Aggregation: view.Count(),
	}
	upstreamFetchView = &view.View{
		Name:        "upstream_fetch_total",
		Measure:     upstreamFetchStats,
		Description: "Count of upstream fetch attempts",
		TagKeys:     []tag.Key{fetchResult},
		Aggregation: view.Count(),
	}
	upstreamFetchLatencyView = &view.View{
		Name:        "upstream_fetch_duration_seconds",
		Measure:     upstreamFetchDurationStats,
		Description: "Distribution of upstream fetch latency in seconds",
		TagKeys:     []tag.Key{fetchResult},
		Aggregation: view.Distribution(upstreamExponentialBuckets...),
	}
	cacheHitsView = &view.View{
		Name:        "athens_cache_hits_total",
		Measure:     cacheHitsStats,
		Description: "Module served from disk storage",
		TagKeys:     []tag.Key{cacheType},
		Aggregation: view.Count(),
	}
	cacheMissesView = &view.View{
		Name:        "athens_cache_misses_total",
		Measure:     cacheMissesStats,
		Description: "Module fetched from upstream",
		TagKeys:     []tag.Key{cacheType},
		Aggregation: view.Count(),
	}
	bytesServedView = &view.View{
		Name:        "athens_bytes_served_total",
		Measure:     bytesServedStats,
		Description: "Bytes served to clients",
		TagKeys:     []tag.Key{cacheType},
		Aggregation: view.Sum(),
	}
	bytesFetchedView = &view.View{
		Name:        "athens_bytes_fetched_total",
		Measure:     bytesFetchedStats,
		Description: "Bytes fetched from upstream",
		TagKeys:     []tag.Key{cacheType},
		Aggregation: view.Sum(),
	}
)

func customViews() []*view.View {
	return []*view.View{
		cacheLookupView, upstreamFetchView, upstreamFetchLatencyView,
		cacheHitsView, cacheMissesView, bytesServedView, bytesFetchedView,
	}
}

func RecordCacheLookup(ctx context.Context, result, typ string) {
	ctx, _ = tag.New(ctx,
		tag.Insert(cacheResult, result),
		tag.Insert(cacheType, typ),
	)
	stats.Record(ctx, cacheStats.M(1))
}

func RecordUpstreamFetch(ctx context.Context, result string) {
	ctx, _ = tag.New(ctx,
		tag.Insert(fetchResult, result),
	)
	stats.Record(ctx, upstreamFetchStats.M(1))
}

func RecordUpstreamFetchDuration(ctx context.Context, result string, duration time.Duration) {
	ctx, _ = tag.New(ctx,
		tag.Insert(fetchResult, result),
	)
	stats.Record(ctx, upstreamFetchDurationStats.M(duration.Seconds()))
}

func RecordCacheHit(ctx context.Context, typ string) {
	ctx, _ = tag.New(ctx,
		tag.Insert(cacheType, typ),
	)
	stats.Record(ctx, cacheHitsStats.M(1))
}

func RecordCacheMiss(ctx context.Context, typ string) {
	ctx, _ = tag.New(ctx,
		tag.Insert(cacheType, typ),
	)
	stats.Record(ctx, cacheMissesStats.M(1))
}

func RecordBytesServed(ctx context.Context, typ string, n int64) {
	ctx, _ = tag.New(ctx,
		tag.Insert(cacheType, typ),
	)
	stats.Record(ctx, bytesServedStats.M(n))
}

func RecordBytesFetched(ctx context.Context, typ string, n int64) {
	ctx, _ = tag.New(ctx,
		tag.Insert(cacheType, typ),
	)
	stats.Record(ctx, bytesFetchedStats.M(n))
}
