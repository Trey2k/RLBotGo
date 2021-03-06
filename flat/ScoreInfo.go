// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ScoreInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsScoreInfo(buf []byte, offset flatbuffers.UOffsetT) *ScoreInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ScoreInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ScoreInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScoreInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ScoreInfo) Score() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateScore(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *ScoreInfo) Goals() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateGoals(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *ScoreInfo) OwnGoals() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateOwnGoals(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *ScoreInfo) Assists() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateAssists(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *ScoreInfo) Saves() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateSaves(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func (rcv *ScoreInfo) Shots() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateShots(n int32) bool {
	return rcv._tab.MutateInt32Slot(14, n)
}

func (rcv *ScoreInfo) Demolitions() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScoreInfo) MutateDemolitions(n int32) bool {
	return rcv._tab.MutateInt32Slot(16, n)
}

func ScoreInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func ScoreInfoAddScore(builder *flatbuffers.Builder, score int32) {
	builder.PrependInt32Slot(0, score, 0)
}
func ScoreInfoAddGoals(builder *flatbuffers.Builder, goals int32) {
	builder.PrependInt32Slot(1, goals, 0)
}
func ScoreInfoAddOwnGoals(builder *flatbuffers.Builder, ownGoals int32) {
	builder.PrependInt32Slot(2, ownGoals, 0)
}
func ScoreInfoAddAssists(builder *flatbuffers.Builder, assists int32) {
	builder.PrependInt32Slot(3, assists, 0)
}
func ScoreInfoAddSaves(builder *flatbuffers.Builder, saves int32) {
	builder.PrependInt32Slot(4, saves, 0)
}
func ScoreInfoAddShots(builder *flatbuffers.Builder, shots int32) {
	builder.PrependInt32Slot(5, shots, 0)
}
func ScoreInfoAddDemolitions(builder *flatbuffers.Builder, demolitions int32) {
	builder.PrependInt32Slot(6, demolitions, 0)
}
func ScoreInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
