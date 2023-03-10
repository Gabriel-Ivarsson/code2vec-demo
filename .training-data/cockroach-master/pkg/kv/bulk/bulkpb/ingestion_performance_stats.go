// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package bulkpb

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/bulk"
	"github.com/cockroachdb/cockroach/pkg/util/humanizeutil"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/redact"
	"go.opentelemetry.io/otel/attribute"
)

var _ bulk.TracingAggregatorEvent = (*IngestionPerformanceStats)(nil)

// Identity implements the TracingAggregatorEvent interface.
func (s *IngestionPerformanceStats) Identity() bulk.TracingAggregatorEvent {
	stats := IngestionPerformanceStats{}
	stats.SendWaitByStore = make(map[roachpb.StoreID]time.Duration)
	return &stats
}

// Combine implements the TracingAggregatorEvent interface.
func (s *IngestionPerformanceStats) Combine(other bulk.TracingAggregatorEvent) {
	otherStats, ok := other.(*IngestionPerformanceStats)
	if !ok {
		panic(fmt.Sprintf("`other` is not of type IngestionPerformanceStats: %T", other))
	}

	s.DataSize += otherStats.DataSize
	s.BufferFlushes += otherStats.BufferFlushes
	s.FlushesDueToSize += otherStats.FlushesDueToSize
	s.Batches += otherStats.Batches
	s.BatchesDueToRange += otherStats.BatchesDueToRange
	s.BatchesDueToSize += otherStats.BatchesDueToSize
	s.SplitRetries += otherStats.SplitRetries
	s.Splits += otherStats.Splits
	s.Scatters += otherStats.Scatters
	s.ScatterMoved += otherStats.ScatterMoved
	s.FillWait += otherStats.FillWait
	s.SortWait += otherStats.SortWait
	s.FlushWait += otherStats.FlushWait
	s.BatchWait += otherStats.BatchWait
	s.SendWait += otherStats.SendWait
	s.SplitWait += otherStats.SplitWait
	s.ScatterWait += otherStats.ScatterWait
	s.CommitWait += otherStats.CommitWait
	s.Duration += otherStats.Duration

	for k, v := range otherStats.SendWaitByStore {
		s.SendWaitByStore[k] += v
	}
}

// Tag implements the TracingAggregatorEvent interface.
func (s *IngestionPerformanceStats) Tag() string {
	return "IngestionPerformanceStats"
}

// Render implements the TracingAggregatorEvent interface.
func (s *IngestionPerformanceStats) Render() []attribute.KeyValue {
	const mb = 1 << 20
	tags := make([]attribute.KeyValue, 0)
	if s.Batches > 0 {
		tags = append(tags,
			attribute.KeyValue{
				Key:   "num_batches",
				Value: attribute.Int64Value(s.Batches),
			},
			attribute.KeyValue{
				Key:   "num_batches_due_to_size",
				Value: attribute.Int64Value(s.BatchesDueToSize),
			},
			attribute.KeyValue{
				Key:   "num_batches_due_to_range",
				Value: attribute.Int64Value(s.BatchesDueToRange),
			},
			attribute.KeyValue{
				Key:   "split_retires",
				Value: attribute.Int64Value(s.SplitRetries),
			},
		)
	}

	if s.BufferFlushes > 0 {
		tags = append(tags,
			attribute.KeyValue{
				Key:   "num_flushes",
				Value: attribute.Int64Value(s.BufferFlushes),
			},
			attribute.KeyValue{
				Key:   "num_flushes_due_to_size",
				Value: attribute.Int64Value(s.FlushesDueToSize),
			},
		)
	}

	if s.DataSize > 0 {
		dataSizeMB := float64(s.DataSize) / mb
		tags = append(tags, attribute.KeyValue{
			Key:   "data_size",
			Value: attribute.StringValue(fmt.Sprintf("%.2f MB", dataSizeMB)),
		})

		if s.Duration > 0 {
			throughput := dataSizeMB / s.Duration.Seconds()
			tags = append(tags, attribute.KeyValue{
				Key:   "throughput",
				Value: attribute.StringValue(fmt.Sprintf("%.2f MB/s", throughput)),
			})
		}
	}

	tags = append(tags,
		timeKeyValue("fill_wait", s.FillWait),
		timeKeyValue("sort_wait", s.SortWait),
		timeKeyValue("flush_wait", s.FlushWait),
		timeKeyValue("batch_wait", s.BatchWait),
		timeKeyValue("send_wait", s.SendWait),
		timeKeyValue("split_wait", s.SplitWait),
		attribute.KeyValue{Key: "splits", Value: attribute.Int64Value(s.Splits)},
		timeKeyValue("scatter_wait", s.ScatterWait),
		attribute.KeyValue{Key: "scatters", Value: attribute.Int64Value(s.Scatters)},
		attribute.KeyValue{Key: "scatter_moved", Value: attribute.Int64Value(s.ScatterMoved)},
		timeKeyValue("commit_wait", s.CommitWait),
	)

	// Sort store send wait by IDs before adding them as tags.
	ids := make(roachpb.StoreIDSlice, 0, len(s.SendWaitByStore))
	for i := range s.SendWaitByStore {
		ids = append(ids, i)
	}
	sort.Sort(ids)
	for _, id := range ids {
		tags = append(tags, timeKeyValue(attribute.Key(fmt.Sprintf("store-%d_send_wait", id)), s.SendWaitByStore[id]))
	}

	return tags
}

func timeKeyValue(key attribute.Key, time time.Duration) attribute.KeyValue {
	return attribute.KeyValue{
		Key:   key,
		Value: attribute.StringValue(string(humanizeutil.Duration(time))),
	}
}

// LogTimings logs the timing ingestion stats.
func (s *IngestionPerformanceStats) LogTimings(ctx context.Context, name, action string) {
	log.Infof(ctx,
		"%s adder %s; ingested %s: %s filling; %v sorting; %v / %v flushing; %v sending; %v splitting; %d; %v scattering, %d, %v; %v commit-wait",
		name,
		redact.Safe(action),
		sz(s.DataSize),
		timing(s.FillWait),
		timing(s.SortWait),
		timing(s.FlushWait),
		timing(s.BatchWait),
		timing(s.SendWait),
		timing(s.SplitWait),
		s.Splits,
		timing(s.ScatterWait),
		s.Scatters,
		s.ScatterMoved,
		timing(s.CommitWait),
	)
}

// LogFlushes logs stats about buffering added and SST batcher flushes.
func (s *IngestionPerformanceStats) LogFlushes(
	ctx context.Context, name, action string, bufSize int64, span roachpb.Span,
) {
	log.Infof(ctx,
		"%s adder %s; flushed into %s %d times, %d due to buffer size (%s); flushing chunked into %d files (%d for ranges, %d for sst size) +%d split-retries",
		name,
		redact.Safe(action),
		span,
		s.BufferFlushes,
		s.FlushesDueToSize,
		sz(bufSize),
		s.Batches,
		s.BatchesDueToRange,
		s.BatchesDueToSize,
		s.SplitRetries,
	)
}

// LogPerStoreTimings logs send waits per store.
func (s *IngestionPerformanceStats) LogPerStoreTimings(ctx context.Context, name string) {
	if len(s.SendWaitByStore) == 0 {
		return
	}
	ids := make(roachpb.StoreIDSlice, 0, len(s.SendWaitByStore))
	for i := range s.SendWaitByStore {
		ids = append(ids, i)
	}
	sort.Sort(ids)

	var sb strings.Builder
	for i, id := range ids {
		// Hack: fill the map with placeholder stores if we haven't seen the store
		// with ID below K for all but lowest K, so that next time we print a zero.
		if i > 0 && ids[i-1] != id-1 {
			s.SendWaitByStore[id-1] = 0
			fmt.Fprintf(&sb, "%d: %s;", id-1, timing(0))
		}
		fmt.Fprintf(&sb, "%d: %s;", id, timing(s.SendWaitByStore[id]))

	}
	log.Infof(ctx, "%s waited on sending to: %s", name, redact.Safe(sb.String()))
}

type sz int64

func (b sz) String() string { return string(humanizeutil.IBytes(int64(b))) }
func (b sz) SafeValue()     {}

type timing time.Duration

func (t timing) String() string { return time.Duration(t).Round(time.Second).String() }
func (t timing) SafeValue()     {}

var _ bulk.TracingAggregatorEvent = &IngestionPerformanceStats{}
