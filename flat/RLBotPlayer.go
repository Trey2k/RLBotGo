// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A bot controlled by the RLBot framework
type RLBotPlayer struct {
	_tab flatbuffers.Table
}

func GetRootAsRLBotPlayer(buf []byte, offset flatbuffers.UOffsetT) *RLBotPlayer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RLBotPlayer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RLBotPlayer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RLBotPlayer) Table() flatbuffers.Table {
	return rcv._tab
}

func RLBotPlayerStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func RLBotPlayerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
