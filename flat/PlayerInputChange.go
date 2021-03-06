// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Rocket League is notifying us that some player has moved their controller. This is an *output*
type PlayerInputChange struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerInputChange(buf []byte, offset flatbuffers.UOffsetT) *PlayerInputChange {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerInputChange{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PlayerInputChange) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerInputChange) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PlayerInputChange) PlayerIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PlayerInputChange) MutatePlayerIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *PlayerInputChange) ControllerState(obj *ControllerState) *ControllerState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ControllerState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PlayerInputChange) DodgeForward() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *PlayerInputChange) MutateDodgeForward(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *PlayerInputChange) DodgeRight() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *PlayerInputChange) MutateDodgeRight(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func PlayerInputChangeStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func PlayerInputChangeAddPlayerIndex(builder *flatbuffers.Builder, playerIndex int32) {
	builder.PrependInt32Slot(0, playerIndex, 0)
}
func PlayerInputChangeAddControllerState(builder *flatbuffers.Builder, controllerState flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(controllerState), 0)
}
func PlayerInputChangeAddDodgeForward(builder *flatbuffers.Builder, dodgeForward float32) {
	builder.PrependFloat32Slot(2, dodgeForward, 0.0)
}
func PlayerInputChangeAddDodgeRight(builder *flatbuffers.Builder, dodgeRight float32) {
	builder.PrependFloat32Slot(3, dodgeRight, 0.0)
}
func PlayerInputChangeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
