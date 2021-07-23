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

func createDesiredPhysics(builder *flatbuffers.Builder, physics Physics) flatbuffers.UOffsetT {
	// Vector 3 partial for BallState Location
	schema.Vector3PartialStart(builder)
	offset := schema.CreateFloat(builder, physics.Location.X)
	schema.Vector3PartialAddX(builder, offset)
	offset = schema.CreateFloat(builder, physics.Location.Y)
	schema.Vector3PartialAddY(builder, offset)
	offset = schema.CreateFloat(builder, physics.Location.Z)
	schema.Vector3PartialAddZ(builder, offset)
	location := schema.Vector3PartialEnd(builder)

	// Vector 3 partial for BallState Velocity
	schema.Vector3PartialStart(builder)
	offset = schema.CreateFloat(builder, physics.Velocity.X)
	schema.Vector3PartialAddX(builder, offset)
	offset = schema.CreateFloat(builder, physics.Velocity.Y)
	schema.Vector3PartialAddY(builder, offset)
	offset = schema.CreateFloat(builder, physics.Velocity.Z)
	schema.Vector3PartialAddZ(builder, offset)
	velocity := schema.Vector3PartialEnd(builder)

	// Rotator partial for BallState Rotation
	schema.RotatorPartialStart(builder)
	offset = schema.CreateFloat(builder, physics.Rotation.Pitch)
	schema.RotatorPartialAddPitch(builder, offset)
	offset = schema.CreateFloat(builder, physics.Rotation.Roll)
	schema.RotatorPartialAddRoll(builder, offset)
	offset = schema.CreateFloat(builder, physics.Rotation.Yaw)
	schema.RotatorPartialAddYaw(builder, offset)
	rotation := schema.RotatorPartialEnd(builder)

	schema.DesiredPhysicsStart(builder)
	schema.DesiredPhysicsAddLocation(builder, location)
	schema.DesiredPhysicsAddVelocity(builder, velocity)
	schema.DesiredPhysicsAddRotation(builder, rotation)
	return schema.DesiredPhysicsEnd(builder)
}

func (desiredGameState *DesiredGameState) marshal() []byte {
	builder := flatbuffers.NewBuilder(1024)

	offset := createDesiredPhysics(builder, desiredGameState.BallState.Physics)
	schema.DesiredBallStateStart(builder)
	schema.DesiredBallStateAddPhysics(builder, offset)
	ballState := schema.DesiredBallStateEnd(builder)

	var boostStates flatbuffers.UOffsetT = 0
	if desiredGameState.BoostStates != nil && len(desiredGameState.BoostStates) > 0 {

		var boostStateOffests []flatbuffers.UOffsetT
		for i := 0; i < len(desiredGameState.BoostStates); i++ {
			schema.DesiredBoostStateStart(builder)
			offset := schema.CreateFloat(builder, desiredGameState.BoostStates[i].RespawnTime)
			schema.DesiredBoostStateAddRespawnTime(builder, offset)
			boostState := schema.DesiredBoostStateEnd(builder)
			boostStateOffests = append(boostStateOffests, boostState)

		}

		schema.DesiredGameStateStartBoostStatesVector(builder, len(desiredGameState.BoostStates))
		for i := 0; i < len(boostStateOffests); i++ {
			builder.PrependUOffsetT(boostStateOffests[i])
		}
		boostStates = builder.EndVector(len(desiredGameState.BoostStates))

	}

	var carStates flatbuffers.UOffsetT = 0
	if desiredGameState.CarStates != nil && len(desiredGameState.CarStates) > 0 {

		var carStateOffsets []flatbuffers.UOffsetT
		for i := 0; i < len(desiredGameState.CarStates); i++ {
			physics := createDesiredPhysics(builder, desiredGameState.CarStates[i].Physics)
			schema.DesiredCarStateStart(builder)

			offset := schema.CreateFloat(builder, desiredGameState.CarStates[i].BoostAmount)
			schema.DesiredCarStateAddBoostAmount(builder, offset)
			offset = schema.CreateBool(builder, boolToByte[desiredGameState.CarStates[i].DoubleJumped])
			schema.DesiredCarStateAddDoubleJumped(builder, offset)
			offset = schema.CreateBool(builder, boolToByte[desiredGameState.CarStates[i].Jumped])
			schema.DesiredCarStateAddJumped(builder, offset)

			schema.DesiredCarStateAddPhysics(builder, physics)
			carState := schema.DesiredCarStateEnd(builder)
			carStateOffsets = append(carStateOffsets, carState)
		}

		schema.DesiredGameStateStartCarStatesVector(builder, len(carStateOffsets))
		for i := 0; i < len(carStateOffsets); i++ {
			builder.PrependUOffsetT(carStateOffsets[i])
		}
		carStates = builder.EndVector(len(carStateOffsets))

	}

	var commands flatbuffers.UOffsetT = 0
	if desiredGameState.ConsoleCommands != nil && len(desiredGameState.ConsoleCommands) > 0 {

		var commandOffsets []flatbuffers.UOffsetT

		for i := 0; i < len(desiredGameState.ConsoleCommands); i++ {
			schema.ConsoleCommandStart(builder)
			offset := builder.CreateString(desiredGameState.ConsoleCommands[i])
			schema.ConsoleCommandAddCommand(builder, offset)
			command := schema.ConsoleCommandEnd(builder)
			commandOffsets = append(commandOffsets, command)
		}

		schema.DesiredGameStateStartConsoleCommandsVector(builder, len(commandOffsets))
		for i := 0; i < len(commandOffsets); i++ {
			builder.PrependUOffsetT(commandOffsets[i])
		}
		commands = builder.EndVector(len(commandOffsets))

	}

	schema.DesiredGameInfoStateStart(builder)
	offset = schema.CreateBool(builder, boolToByte[desiredGameState.GameInfoState.EndMatch])
	schema.DesiredGameInfoStateAddEndMatch(builder, offset)

	offset = schema.CreateBool(builder, boolToByte[desiredGameState.GameInfoState.Paused])
	schema.DesiredGameInfoStateAddPaused(builder, offset)

	offset = schema.CreateFloat(builder, desiredGameState.GameInfoState.GameSpeed)
	schema.DesiredGameInfoStateAddGameSpeed(builder, offset)

	offset = schema.CreateFloat(builder, desiredGameState.GameInfoState.WorldGravityZ)
	schema.DesiredGameInfoStateAddWorldGravityZ(builder, offset)
	infoState := schema.DesiredGameInfoStateEnd(builder)

	schema.DesiredGameStateStart(builder)
	if boostStates != 0 {
		schema.DesiredGameStateAddBoostStates(builder, boostStates)
	}
	if carStates != 0 {
		schema.DesiredGameStateAddCarStates(builder, carStates)
	}
	if commands != 0 {
		schema.DesiredGameStateAddConsoleCommands(builder, commands)
	}

	schema.DesiredGameStateAddBallState(builder, ballState)
	schema.DesiredGameStateAddGameInfoState(builder, infoState)
	desiredState := schema.DesiredGameStateEnd(builder)

	builder.Finish(desiredState)

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

func (renderGroup *RenderGroup) marshal() []byte {
	builder := flatbuffers.NewBuilder(1024)

	var messages []flatbuffers.UOffsetT
	for i := 0; i < len(renderGroup.RenderMessages); i++ {
		text := builder.CreateString(renderGroup.RenderMessages[i].Text)

		schema.ColorStart(builder)
		schema.ColorAddA(builder, renderGroup.RenderMessages[i].Color.A)
		schema.ColorAddR(builder, renderGroup.RenderMessages[i].Color.R)
		schema.ColorAddG(builder, renderGroup.RenderMessages[i].Color.G)
		schema.ColorAddB(builder, renderGroup.RenderMessages[i].Color.B)
		color := schema.ColorEnd(builder)
		schema.RenderMessageStart(builder)
		schema.RenderMessageAddColor(builder, color)
		end := schema.CreateVector3(builder, renderGroup.RenderMessages[i].End.X, renderGroup.RenderMessages[i].End.Y, renderGroup.RenderMessages[i].End.Z)
		schema.RenderMessageAddEnd(builder, end)
		start := schema.CreateVector3(builder, renderGroup.RenderMessages[i].Start.X, renderGroup.RenderMessages[i].Start.Y, renderGroup.RenderMessages[i].Start.Z)
		schema.RenderMessageAddStart(builder, start)
		schema.RenderMessageAddIsFilled(builder, boolToByte[renderGroup.RenderMessages[i].IsFilled])
		schema.RenderMessageAddRenderType(builder, renderGroup.RenderMessages[i].RenderType)
		schema.RenderMessageAddScaleX(builder, renderGroup.RenderMessages[i].ScaleX)
		schema.RenderMessageAddScaleY(builder, renderGroup.RenderMessages[i].ScaleY)

		schema.RenderMessageAddText(builder, text)
		renderMsg := schema.RenderMessageEnd(builder)
		messages = append(messages, renderMsg)
	}
	schema.RenderGroupStartRenderMessagesVector(builder, len(messages))
	for i := 0; i < len(messages); i++ {
		builder.PrependUOffsetT(messages[i])
	}
	renderMessages := builder.EndVector(len(messages))

	schema.RenderGroupStart(builder)
	schema.RenderGroupAddRenderMessages(builder, renderMessages)
	schema.RenderGroupAddId(builder, renderGroup.Id)
	msg := schema.ReadyMessageEnd(builder)
	builder.Finish(msg)

	return builder.FinishedBytes()
}
