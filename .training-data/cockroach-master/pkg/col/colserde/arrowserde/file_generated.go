// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package arrowserde

import flatbuffers "github.com/google/flatbuffers/go"

// / ----------------------------------------------------------------------
// / Arrow File metadata
// /
type Footer struct {
	_tab flatbuffers.Table
}

func GetRootAsFooter(buf []byte, offset flatbuffers.UOffsetT) *Footer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Footer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Footer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Footer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Footer) Version() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Footer) MutateVersion(n int16) bool {
	return rcv._tab.MutateInt16Slot(4, n)
}

func (rcv *Footer) Schema(obj *Schema) *Schema {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Schema)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Footer) Dictionaries(obj *Block, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 24
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Footer) DictionariesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Footer) RecordBatches(obj *Block, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 24
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Footer) RecordBatchesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FooterStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FooterAddVersion(builder *flatbuffers.Builder, version int16) {
	builder.PrependInt16Slot(0, version, 0)
}
func FooterAddSchema(builder *flatbuffers.Builder, schema flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(schema), 0)
}
func FooterAddDictionaries(builder *flatbuffers.Builder, dictionaries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dictionaries), 0)
}
func FooterStartDictionariesVector(
	builder *flatbuffers.Builder, numElems int,
) flatbuffers.UOffsetT {
	return builder.StartVector(24, numElems, 8)
}
func FooterAddRecordBatches(builder *flatbuffers.Builder, recordBatches flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(recordBatches), 0)
}
func FooterStartRecordBatchesVector(
	builder *flatbuffers.Builder, numElems int,
) flatbuffers.UOffsetT {
	return builder.StartVector(24, numElems, 8)
}
func FooterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type Block struct {
	_tab flatbuffers.Struct
}

func (rcv *Block) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Block) Table() flatbuffers.Table {
	return rcv._tab.Table
}

// / Index to the start of the RecordBlock (note this is past the Message header)
func (rcv *Block) Offset() int64 {
	return rcv._tab.GetInt64(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}

// / Index to the start of the RecordBlock (note this is past the Message header)
func (rcv *Block) MutateOffset(n int64) bool {
	return rcv._tab.MutateInt64(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

// / Length of the metadata
func (rcv *Block) MetaDataLength() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}

// / Length of the metadata
func (rcv *Block) MutateMetaDataLength(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

// / Length of the data (this is aligned so there can be a gap between this and
// / the metadata).
func (rcv *Block) BodyLength() int64 {
	return rcv._tab.GetInt64(rcv._tab.Pos + flatbuffers.UOffsetT(16))
}

// / Length of the data (this is aligned so there can be a gap between this and
// / the metadata).
func (rcv *Block) MutateBodyLength(n int64) bool {
	return rcv._tab.MutateInt64(rcv._tab.Pos+flatbuffers.UOffsetT(16), n)
}

func CreateBlock(
	builder *flatbuffers.Builder, offset int64, metaDataLength int32, bodyLength int64,
) flatbuffers.UOffsetT {
	builder.Prep(8, 24)
	builder.PrependInt64(bodyLength)
	builder.Pad(4)
	builder.PrependInt32(metaDataLength)
	builder.PrependInt64(offset)
	return builder.Offset()
}
