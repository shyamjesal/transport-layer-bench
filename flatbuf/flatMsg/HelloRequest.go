// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatMsg

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HelloRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsHelloRequest(buf []byte, offset flatbuffers.UOffsetT) *HelloRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HelloRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsHelloRequest(buf []byte, offset flatbuffers.UOffsetT) *HelloRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &HelloRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *HelloRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HelloRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *HelloRequest) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func HelloRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func HelloRequestAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func HelloRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
