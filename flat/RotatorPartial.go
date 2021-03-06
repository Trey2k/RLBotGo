// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RotatorPartial struct {
	_tab flatbuffers.Table
}

func GetRootAsRotatorPartial(buf []byte, offset flatbuffers.UOffsetT) *RotatorPartial {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RotatorPartial{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RotatorPartial) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RotatorPartial) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RotatorPartial) Pitch(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
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

func (rcv *RotatorPartial) Yaw(obj *Float) *Float {
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

func (rcv *RotatorPartial) Roll(obj *Float) *Float {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
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

func RotatorPartialStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func RotatorPartialAddPitch(builder *flatbuffers.Builder, pitch flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(pitch), 0)
}
func RotatorPartialAddYaw(builder *flatbuffers.Builder, yaw flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(yaw), 0)
}
func RotatorPartialAddRoll(builder *flatbuffers.Builder, roll flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(roll), 0)
}
func RotatorPartialEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
