// automatically generated by the FlatBuffers compiler, do not modify

package hello

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Response struct {
	_tab flatbuffers.Table
}

func GetRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) *Response {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Response{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Response) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Response) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Response) Msg() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ResponseAddMsg(builder *flatbuffers.Builder, msg flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(msg), 0)
}
func ResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
