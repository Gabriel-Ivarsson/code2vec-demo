// Code generated by "stringer -type=healingMetric -trimprefix=healingMetric erasure-healing.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[healingMetricBucket-0]
	_ = x[healingMetricObject-1]
	_ = x[healingMetricCheckAbandonedParts-2]
}

const _healingMetric_name = "BucketObjectCheckAbandonedParts"

var _healingMetric_index = [...]uint8{0, 6, 12, 31}

func (i healingMetric) String() string {
	if i >= healingMetric(len(_healingMetric_index)-1) {
		return "healingMetric(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _healingMetric_name[_healingMetric_index[i]:_healingMetric_index[i+1]]
}
