// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A normal human player
type HumanPlayer struct {
	_tab flatbuffers.Table
}

func GetRootAsHumanPlayer(buf []byte, offset flatbuffers.UOffsetT) *HumanPlayer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HumanPlayer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *HumanPlayer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HumanPlayer) Table() flatbuffers.Table {
	return rcv._tab
}

func HumanPlayerStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func HumanPlayerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
