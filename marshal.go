package RLBotGo

import (
	schema "github.com/Trey2k/RLBotGo/flat"

	flatbuffers "github.com/google/flatbuffers/go"
)

var boolToByte = map[bool]byte{false: 0, true: 1} // should be in a utils package

func (readyMsg *ReadyMessage) Marshel() []byte {
	builder := flatbuffers.NewBuilder(1024)
	schema.ReadyMessageStart(builder)
	schema.ReadyMessageAddWantsBallPredictions(builder, boolToByte[readyMsg.WantsBallPredictions])
	schema.ReadyMessageAddWantsGameMessages(builder, boolToByte[readyMsg.WantsGameMessages])
	schema.ReadyMessageAddWantsQuickChat(builder, boolToByte[readyMsg.WantsQuickChat])
	msg := schema.ReadyMessageEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}

func (quickChat *QuickChat) Marshel() []byte {
	builder := flatbuffers.NewBuilder(1024)
	schema.QuickChatStart(builder)
	schema.QuickChatAddPlayerIndex(builder, quickChat.PlayerIndex)
	schema.QuickChatAddQuickChatSelection(builder, quickChat.QuickChatSelection)
	schema.QuickChatAddTeamOnly(builder, boolToByte[quickChat.TeamOnly])
	msg := schema.ReadyMessageEnd(builder)
	builder.Finish(msg)
	// Send ready message

	return builder.FinishedBytes()
}
