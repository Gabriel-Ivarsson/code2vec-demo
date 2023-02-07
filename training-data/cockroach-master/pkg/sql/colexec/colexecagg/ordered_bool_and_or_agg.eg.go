// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecagg

import (
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
)

// Remove unused warning.
var _ = colexecerror.InternalError

func newBoolAndOrderedAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolAndOrderedAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolAndOrderedAgg struct {
	orderedAggregateFuncBase
	col    coldata.Bools
	curAgg bool
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &boolAndOrderedAgg{}

func (a *boolAndOrderedAgg) SetOutput(vec coldata.Vec) {
	a.orderedAggregateFuncBase.SetOutput(vec)
	a.col = vec.Bool()
}

func (a *boolAndOrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, startIdx, endIdx int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bool(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		// Capture groups and col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		groups := a.groups
		col := col
		if sel == nil {
			_, _ = groups[endIdx-1], groups[startIdx]
			_, _ = col.Get(endIdx-1), col.Get(startIdx)
			if nulls.MaybeHasNulls() {
				for i := startIdx; i < endIdx; i++ {
					//gcassert:bce
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = true
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						//gcassert:bce
						a.curAgg = a.curAgg && col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			} else {
				for i := startIdx; i < endIdx; i++ {
					//gcassert:bce
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = true
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = false
					if !isNull {
						//gcassert:bce
						a.curAgg = a.curAgg && col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			}
		} else {
			sel = sel[startIdx:endIdx]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = true
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						a.curAgg = a.curAgg && col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			} else {
				for _, i := range sel {
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = true
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = false
					if !isNull {
						a.curAgg = a.curAgg && col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsageAfterAllocation(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *boolAndOrderedAgg) Flush(outputIdx int) {
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	col := a.col
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		col[outputIdx] = a.curAgg
	}
}

func (a *boolAndOrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.curAgg = true
	a.foundNonNullForCurrentGroup = false
}

type boolAndOrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []boolAndOrderedAgg
}

var _ aggregateFuncAlloc = &boolAndOrderedAggAlloc{}

const sizeOfBoolAndOrderedAgg = int64(unsafe.Sizeof(boolAndOrderedAgg{}))
const boolAndOrderedAggSliceOverhead = int64(unsafe.Sizeof([]boolAndOrderedAgg{}))

func (a *boolAndOrderedAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(boolAndOrderedAggSliceOverhead + sizeOfBoolAndOrderedAgg*a.allocSize)
		a.aggFuncs = make([]boolAndOrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	f.Reset()
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

func newBoolOrOrderedAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolOrOrderedAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolOrOrderedAgg struct {
	orderedAggregateFuncBase
	col    coldata.Bools
	curAgg bool
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &boolOrOrderedAgg{}

func (a *boolOrOrderedAgg) SetOutput(vec coldata.Vec) {
	a.orderedAggregateFuncBase.SetOutput(vec)
	a.col = vec.Bool()
}

func (a *boolOrOrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, startIdx, endIdx int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bool(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		// Capture groups and col to force bounds check to work. See
		// https://github.com/golang/go/issues/39756
		groups := a.groups
		col := col
		if sel == nil {
			_, _ = groups[endIdx-1], groups[startIdx]
			_, _ = col.Get(endIdx-1), col.Get(startIdx)
			if nulls.MaybeHasNulls() {
				for i := startIdx; i < endIdx; i++ {
					//gcassert:bce
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = false
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						//gcassert:bce
						a.curAgg = a.curAgg || col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			} else {
				for i := startIdx; i < endIdx; i++ {
					//gcassert:bce
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = false
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = false
					if !isNull {
						//gcassert:bce
						a.curAgg = a.curAgg || col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			}
		} else {
			sel = sel[startIdx:endIdx]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = false
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						a.curAgg = a.curAgg || col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			} else {
				for _, i := range sel {
					if groups[i] {
						if !a.isFirstGroup {
							if !a.foundNonNullForCurrentGroup {
								a.nulls.SetNull(a.curIdx)
							} else {
								a.col[a.curIdx] = a.curAgg
							}
							a.curIdx++
							a.curAgg = false
							a.foundNonNullForCurrentGroup = false
						}
						a.isFirstGroup = false
					}

					var isNull bool
					isNull = false
					if !isNull {
						a.curAgg = a.curAgg || col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsageAfterAllocation(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *boolOrOrderedAgg) Flush(outputIdx int) {
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	col := a.col
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		col[outputIdx] = a.curAgg
	}
}

func (a *boolOrOrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.curAgg = false
	a.foundNonNullForCurrentGroup = false
}

type boolOrOrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []boolOrOrderedAgg
}

var _ aggregateFuncAlloc = &boolOrOrderedAggAlloc{}

const sizeOfBoolOrOrderedAgg = int64(unsafe.Sizeof(boolOrOrderedAgg{}))
const boolOrOrderedAggSliceOverhead = int64(unsafe.Sizeof([]boolOrOrderedAgg{}))

func (a *boolOrOrderedAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(boolOrOrderedAggSliceOverhead + sizeOfBoolOrOrderedAgg*a.allocSize)
		a.aggFuncs = make([]boolOrOrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	f.Reset()
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
