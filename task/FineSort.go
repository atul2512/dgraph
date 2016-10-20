// automatically generated by the FlatBuffers compiler, do not modify

package task

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FineSort struct {
	_tab flatbuffers.Table
}

func GetRootAsFineSort(buf []byte, offset flatbuffers.UOffsetT) *FineSort {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FineSort{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FineSort) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FineSort) Uid(obj *UidList) *UidList {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(UidList)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FineSortStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FineSortAddUid(builder *flatbuffers.Builder, uid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(uid), 0)
}
func FineSortEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}