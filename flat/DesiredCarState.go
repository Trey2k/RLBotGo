// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DesiredCarState struct {
	_tab flatbuffers.Table
}

func GetRootAsDesiredCarState(buf []byte, offset flatbuffers.UOffsetT) *DesiredCarState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DesiredCarState{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DesiredCarState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DesiredCarState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DesiredCarState) Physics(obj *DesiredPhysics) *DesiredPhysics {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DesiredPhysics)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredCarState) BoostAmount(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Float)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredCarState) Jumped(obj *Bool) *Bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Bool)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *DesiredCarState) DoubleJumped(obj *Bool) *Bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Bool)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DesiredCarStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DesiredCarStateAddPhysics(builder *flatbuffers.Builder, physics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(physics), 0)
}
func DesiredCarStateAddBoostAmount(builder *flatbuffers.Builder, boostAmount flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(boostAmount), 0)
}
func DesiredCarStateAddJumped(builder *flatbuffers.Builder, jumped flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(jumped), 0)
}
func DesiredCarStateAddDoubleJumped(builder *flatbuffers.Builder, doubleJumped flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(doubleJumped), 0)
}
func DesiredCarStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}