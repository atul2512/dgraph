// automatically generated by the FlatBuffers compiler, do not modify

package task

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Num struct {
	_tab flatbuffers.Table
}

func GetRootAsNum(buf []byte, offset flatbuffers.UOffsetT) *Num {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Num{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Num) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Num) Group() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Num) MutateGroup(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Num) Val() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Num) MutateVal(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Num) Uids(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Num) UidsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func NumStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func NumAddGroup(builder *flatbuffers.Builder, group uint32) {
	builder.PrependUint32Slot(0, group, 0)
}
func NumAddVal(builder *flatbuffers.Builder, val int32) {
	builder.PrependInt32Slot(1, val, 0)
}
func NumAddUids(builder *flatbuffers.Builder, uids flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(uids), 0)
}
func NumStartUidsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func NumEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
