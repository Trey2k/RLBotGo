// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Expresses the rotation state of an object.
/// Learn about quaternions here: https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation
/// You can tinker with them here to build an intuition: https://quaternions.online/
type Quaternion struct {
	_tab flatbuffers.Struct
}

func (rcv *Quaternion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Quaternion) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Quaternion) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Quaternion) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Quaternion) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Quaternion) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Quaternion) Z() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Quaternion) MutateZ(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func (rcv *Quaternion) W() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(12))
}
func (rcv *Quaternion) MutateW(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(12), n)
}

func CreateQuaternion(builder *flatbuffers.Builder, x float32, y float32, z float32, w float32) flatbuffers.UOffsetT {
	builder.Prep(4, 16)
	builder.PrependFloat32(w)
	builder.PrependFloat32(z)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
