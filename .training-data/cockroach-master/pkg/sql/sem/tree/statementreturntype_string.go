// Code generated by "stringer"; DO NOT EDIT.

package tree

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ack-0]
	_ = x[DDL-1]
	_ = x[RowsAffected-2]
	_ = x[Rows-3]
	_ = x[CopyIn-4]
	_ = x[Unknown-5]
}

const _StatementReturnType_name = "AckDDLRowsAffectedRowsCopyInUnknown"

var _StatementReturnType_index = [...]uint8{0, 3, 6, 18, 22, 28, 35}

func (i StatementReturnType) String() string {
	if i < 0 || i >= StatementReturnType(len(_StatementReturnType_index)-1) {
		return "StatementReturnType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatementReturnType_name[_StatementReturnType_index[i]:_StatementReturnType_index[i+1]]
}