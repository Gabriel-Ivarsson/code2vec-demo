// Code generated by "stringer -type=rebalanceMetric -trimprefix=rebalanceMetric erasure-server-pool-rebalance.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[rebalanceMetricRebalanceBuckets-0]
	_ = x[rebalanceMetricRebalanceBucket-1]
	_ = x[rebalanceMetricRebalanceObject-2]
	_ = x[rebalanceMetricRebalanceRemoveObject-3]
	_ = x[rebalanceMetricSaveMetadata-4]
}

const _rebalanceMetric_name = "RebalanceBucketsRebalanceBucketRebalanceObjectRebalanceRemoveObjectSaveMetadata"

var _rebalanceMetric_index = [...]uint8{0, 16, 31, 46, 67, 79}

func (i rebalanceMetric) String() string {
	if i >= rebalanceMetric(len(_rebalanceMetric_index)-1) {
		return "rebalanceMetric(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _rebalanceMetric_name[_rebalanceMetric_index[i]:_rebalanceMetric_index[i+1]]
}