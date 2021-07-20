package RLBotGo

import (
	schema "github.com/Trey2k/RLBotGo/flat"

	flatbuffers "github.com/google/flatbuffers/go"
)

func (readyMsg *ReadyMessage) marshal() []byte {
	builder := flatbuffers.NewBuilder(1024)
	schema.ReadyMessageStart(builder)
	schema.ReadyMessageAddWantsBallPredictions(builder, boolToByte[readyMsg.WantsBallPredictions])
	schema.ReadyMessageAddWantsGameMessages(builder, boolToByte[readyMsg.WantsGameMessages])
	schema.ReadyMessageAddWantsQuickChat(builder, boolToByte[readyMsg.WantsQuickChat])
	msg := schema.ReadyMessageEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}

func (playerInput *PlayerInput) marshal() []byte {
	builder := flatbuffers.NewBuilder(1024)
	// Controller State
	schema.ControllerStateStart(builder)
	schema.ControllerStateAddBoost(builder, boolToByte[playerInput.ControllerState.Boost])
	schema.ControllerStateAddHandbrake(builder, boolToByte[playerInput.ControllerState.Handbrake])
	schema.ControllerStateAddJump(builder, boolToByte[playerInput.ControllerState.Jump])
	schema.ControllerStateAddUseItem(builder, boolToByte[playerInput.ControllerState.UseItem])

	schema.ControllerStateAddPitch(builder, playerInput.ControllerState.Pitch)
	schema.ControllerStateAddRoll(builder, playerInput.ControllerState.Roll)
	schema.ControllerStateAddYaw(builder, playerInput.ControllerState.Yaw)
	schema.ControllerStateAddSteer(builder, playerInput.ControllerState.Steer)
	schema.ControllerStateAddThrottle(builder, playerInput.ControllerState.Throttle)
	ControllerState := schema.ControllerStateEnd(builder)

	// Player input
	schema.PlayerInputStart(builder)
	schema.PlayerInputAddControllerState(builder, ControllerState)
	schema.PlayerInputAddPlayerIndex(builder, playerInput.PlayerIndex)
	msg := schema.PlayerInputEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}

func (quickChat *QuickChat) marshal() []byte {
	builder := flatbuffers.NewBuilder(1024)
	schema.QuickChatStart(builder)
	schema.QuickChatAddPlayerIndex(builder, quickChat.PlayerIndex)
	schema.QuickChatAddQuickChatSelection(builder, quickChat.QuickChatSelection)
	schema.QuickChatAddTeamOnly(builder, boolToByte[quickChat.TeamOnly])
	msg := schema.ReadyMessageEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}
