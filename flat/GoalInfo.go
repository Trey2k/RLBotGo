// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GoalInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsGoalInfo(buf []byte, offset flatbuffers.UOffsetT) *GoalInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GoalInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GoalInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GoalInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GoalInfo) TeamNum() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GoalInfo) MutateTeamNum(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *GoalInfo) Location(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vector3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *GoalInfo) Direction(obj *Vector3) *Vector3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vector3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *GoalInfo) Width() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GoalInfo) MutateWidth(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *GoalInfo) Height() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GoalInfo) MutateHeight(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func GoalInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func GoalInfoAddTeamNum(builder *flatbuffers.Builder, teamNum int32) {
	builder.PrependInt32Slot(0, teamNum, 0)
}
func GoalInfoAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(location), 0)
}
func GoalInfoAddDirection(builder *flatbuffers.Builder, direction flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(direction), 0)
}
func GoalInfoAddWidth(builder *flatbuffers.Builder, width float32) {
	builder.PrependFloat32Slot(3, width, 0.0)
}
func GoalInfoAddHeight(builder *flatbuffers.Builder, height float32) {
	builder.PrependFloat32Slot(4, height, 0.0)
}
func GoalInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
