package RLBotGo

import (
	schema "github.com/Trey2k/RLBotGo/flat"

	flatbuffers "github.com/google/flatbuffers/go"
)

func (gameTickPack *GameTickPacket) unmarshal(flatGameTickPack *schema.GameTickPacket) {
	flatBallInfo := &schema.BallInfo{}
	gameTickPack.Ball.unmarshal(flatGameTickPack.Ball(flatBallInfo))

	flatBoostPadStats := &schema.BoostPadState{}
	for i := 0; i < flatGameTickPack.BoostPadStatesLength(); i++ {
		flatGameTickPack.BoostPadStates(flatBoostPadStats, i)
		gameTickPack.BoostPadStates = append(gameTickPack.BoostPadStates, BoostPadState{
			IsActive: flatBoostPadStats.IsActive() == 1,
			Timer:    flatBoostPadStats.Timer(),
		})
	}
	flatGameInfo := &schema.GameInfo{}
	gameTickPack.GameInfo.unmarshal(flatGameTickPack.GameInfo(flatGameInfo))

	flatPlayerInfo := &schema.PlayerInfo{}
	for i := 0; i < flatGameTickPack.PlayersLength(); i++ {
		flatGameTickPack.Players(flatPlayerInfo, i)
		var playerInfo PlayerInfo
		playerInfo.unmarshal(flatPlayerInfo)
		gameTickPack.Players = append(gameTickPack.Players, playerInfo)
	}

	flatTeamInfo := &schema.TeamInfo{}
	for i := 0; i < flatGameTickPack.TeamsLength(); i++ {
		flatGameTickPack.Teams(flatTeamInfo, i)
		gameTickPack.Teams = append(gameTickPack.Teams, TeamInfo{
			TeamIndex: flatTeamInfo.TeamIndex(),
			Score:     flatTeamInfo.Score(),
		})
	}

	flatDropShotTile := &schema.DropshotTile{}
	for i := 0; i < flatGameTickPack.TileInformationLength(); i++ {
		flatGameTickPack.TileInformation(flatDropShotTile, i)
		gameTickPack.TileInformation = append(gameTickPack.TileInformation, DropshotTile{
			TileState: flatDropShotTile.TileState(),
		})
	}
}

func (playerInfo *PlayerInfo) unmarshal(flatPlayerInfo *schema.PlayerInfo) {
	playerInfo.Boost = flatPlayerInfo.Boost()
	playerInfo.DoubleJumped = flatPlayerInfo.DoubleJumped() == 1
	playerInfo.HasWheelContact = flatPlayerInfo.HasWheelContact() == 1
	playerInfo.IsBot = flatPlayerInfo.IsBot() == 1
	playerInfo.IsDemolished = flatPlayerInfo.IsDemolished() == 1
	playerInfo.IsSupersonic = flatPlayerInfo.IsSupersonic() == 1
	playerInfo.Jumped = flatPlayerInfo.Jumped() == 1
	playerInfo.Name = string(flatPlayerInfo.Name())
	playerInfo.SpawnId = flatPlayerInfo.SpawnId()
	playerInfo.Team = flatPlayerInfo.Team()

	flatPhysics := &schema.Physics{}
	playerInfo.Physics.unmarshal(flatPlayerInfo.Physics(flatPhysics))

	flatBoxShape := &schema.BoxShape{}
	playerInfo.Hitbox.marshel(flatPlayerInfo.Hitbox(flatBoxShape))

	flatVector3 := &schema.Vector3{}
	playerInfo.HitboxOffset.unmarshal(flatPlayerInfo.HitboxOffset(flatVector3))

	flatScoreInfo := &schema.ScoreInfo{}
	playerInfo.ScoreInfo.marshel(flatPlayerInfo.ScoreInfo(flatScoreInfo))

}

func (scoreInfo *ScoreInfo) marshel(flatScoreInfo *schema.ScoreInfo) {
	scoreInfo.Assists = flatScoreInfo.Assists()
	scoreInfo.Demolitions = flatScoreInfo.Demolitions()
	scoreInfo.Goals = flatScoreInfo.Goals()
	scoreInfo.OwnGoals = flatScoreInfo.OwnGoals()
	scoreInfo.Saves = flatScoreInfo.Saves()
	scoreInfo.Score = flatScoreInfo.Score()
	scoreInfo.Shots = flatScoreInfo.Shots()
}

func (boxShape *BoxShape) marshel(flatBoxShape *schema.BoxShape) {
	boxShape.Height = flatBoxShape.Height()
	boxShape.Length = flatBoxShape.Length()
	boxShape.Width = flatBoxShape.Width()
}

func (gameInfo *GameInfo) unmarshal(flatGameInfo *schema.GameInfo) {
	gameInfo.FrameNum = flatGameInfo.FrameNum()
	gameInfo.GameSpeed = flatGameInfo.GameSpeed()
	gameInfo.GameTimeRemaining = flatGameInfo.GameTimeRemaining()
	gameInfo.IsKickoffPause = flatGameInfo.IsKickoffPause() == 1
	gameInfo.IsMatchEnded = flatGameInfo.IsMatchEnded() == 1
	gameInfo.IsOvertime = flatGameInfo.IsOvertime() == 1
	gameInfo.IsRoundActive = flatGameInfo.IsRoundActive() == 1
	gameInfo.IsUnlimitedTime = flatGameInfo.IsUnlimitedTime() == 1
	gameInfo.SecondsElapsed = flatGameInfo.SecondsElapsed()
	gameInfo.WorldGravityZ = flatGameInfo.WorldGravityZ()
}

func (ballInfo *BallInfo) unmarshal(flatBallInfo *schema.BallInfo) {
	flatPhysics := &schema.Physics{}
	ballInfo.Physics.unmarshal(flatBallInfo.Physics(flatPhysics))

	flatDropShotInfo := &schema.DropShotBallInfo{}
	ballInfo.DropShotInfo.marshel(flatBallInfo.DropShotInfo(flatDropShotInfo))

	flatTouch := &schema.Touch{}
	if flatBallInfo.LatestTouch(flatTouch) != nil {
		ballInfo.LatestTouch.marshel(flatTouch)
	}

	flatShape := &flatbuffers.Table{}
	flatBallInfo.Shape(flatShape)
	ballInfo.Shape.Diameter = flatShape.GetFloat32(flatbuffers.UOffsetT(flatBallInfo.ShapeType())) // This may not work
}

func (physics *Physics) unmarshal(flatPhysics *schema.Physics) {
	flatVector3 := &schema.Vector3{}
	physics.Location.unmarshal(flatPhysics.Location(flatVector3))
	physics.AngularVelocity.unmarshal(flatPhysics.AngularVelocity(flatVector3))
	physics.Velocity.unmarshal(flatPhysics.Velocity(flatVector3))

	flatRotator := &schema.Rotator{}
	physics.Rotation.unmarshal(flatPhysics.Rotation(flatRotator))
}

func (touch *Touch) marshel(flatTouch *schema.Touch) {
	touch.GameSeconds = flatTouch.GameSeconds()
	touch.PlayerIndex = flatTouch.PlayerIndex()
	touch.PlayerName = string(flatTouch.PlayerName())
	touch.Team = flatTouch.Team()

	flatVector3 := &schema.Vector3{}
	touch.Location.unmarshal(flatTouch.Location(flatVector3))
	touch.Normal.unmarshal(flatTouch.Normal(flatVector3))
}

func (dropShotInfo *DropShotBallInfo) marshel(flatDropShotInfo *schema.DropShotBallInfo) {
	dropShotInfo.AbsorbedForce = flatDropShotInfo.AbsorbedForce()
	dropShotInfo.DamageIndex = flatDropShotInfo.DamageIndex()
	dropShotInfo.ForceAccumRecent = flatDropShotInfo.ForceAccumRecent()
}

func (vector3 *Vector3) unmarshal(flatVector3 *schema.Vector3) {
	vector3.X = flatVector3.X()
	vector3.Y = flatVector3.Y()
	vector3.Z = flatVector3.Z()
}

func (rotator *Rotator) unmarshal(flatRotator *schema.Rotator) {
	rotator.Pitch = flatRotator.Pitch()
	rotator.Yaw = flatRotator.Yaw()
	rotator.Roll = flatRotator.Roll()
}
