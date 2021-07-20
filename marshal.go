package RLBotGo

import (
	schema "github.com/Trey2k/RLBotGo/flat"

	flatbuffers "github.com/google/flatbuffers/go"
)

var boolToByte = map[bool]byte{false: 0, true: 1} // should be in a utils file

type rlData interface {
	marshel() []byte
}

func (readyMsg *ReadyMessage) marshel() []byte {
	builder := flatbuffers.NewBuilder(1024)
	schema.ReadyMessageStart(builder)
	schema.ReadyMessageAddWantsBallPredictions(builder, boolToByte[readyMsg.WantsBallPredictions])
	schema.ReadyMessageAddWantsGameMessages(builder, boolToByte[readyMsg.WantsGameMessages])
	schema.ReadyMessageAddWantsQuickChat(builder, boolToByte[readyMsg.WantsQuickChat])
	msg := schema.ReadyMessageEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}

func (playerInput *PlayerInput) marshel() []byte {
	builder := flatbuffers.NewBuilder(1024)
	//Controller State
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

	//Player input
	schema.PlayerInputStart(builder)
	schema.PlayerInputAddControllerState(builder, ControllerState)
	schema.PlayerInputAddPlayerIndex(builder, playerInput.PlayerIndex)
	msg := schema.ControllerStateEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}

func (playerInputChange *PlayerInputChange) marshel() []byte {
	builder := flatbuffers.NewBuilder(1024)
	//Controller State
	schema.ControllerStateStart(builder)
	schema.ControllerStateAddBoost(builder, boolToByte[playerInputChange.ControllerState.Boost])
	schema.ControllerStateAddHandbrake(builder, boolToByte[playerInputChange.ControllerState.Handbrake])
	schema.ControllerStateAddJump(builder, boolToByte[playerInputChange.ControllerState.Jump])
	schema.ControllerStateAddUseItem(builder, boolToByte[playerInputChange.ControllerState.UseItem])

	schema.ControllerStateAddPitch(builder, playerInputChange.ControllerState.Pitch)
	schema.ControllerStateAddRoll(builder, playerInputChange.ControllerState.Roll)
	schema.ControllerStateAddYaw(builder, playerInputChange.ControllerState.Yaw)
	schema.ControllerStateAddSteer(builder, playerInputChange.ControllerState.Steer)
	schema.ControllerStateAddThrottle(builder, playerInputChange.ControllerState.Throttle)
	ControllerState := schema.ControllerStateEnd(builder)
	builder.Finish(ControllerState)

	schema.PlayerInputChangeStart(builder)
	schema.PlayerInputChangeAddControllerState(builder, ControllerState)
	schema.PlayerInputChangeAddPlayerIndex(builder, playerInputChange.PlayerIndex)
	schema.PlayerInputChangeAddDodgeRight(builder, playerInputChange.DodgeRight)
	schema.PlayerInputChangeAddDodgeForward(builder, playerInputChange.DodgeForward)
	msg := schema.PlayerInputEnd(builder)
	builder.Finish(msg)

	schema.MessagePacketStart(builder)
	schema.MessagePacketAddMessages(builder, msg)
	msgPack := schema.MessagePacketEnd(builder)

	builder.Finish(msgPack)

	return builder.FinishedBytes()
}

func (quickChat *QuickChat) marshel() []byte {
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
