package RLBotGo

import (
	"fmt"

	schema "github.com/Trey2k/RLBotGo/flat"

	flatbuffers "github.com/google/flatbuffers/go"
)

func (gameTickPack *GameTickPacket) unmarshal(flatGameTickPack *schema.GameTickPacket) {
	flatBallInfo := &schema.BallInfo{}
	gameTickPack.Ball.unmarshal(flatGameTickPack.Ball(flatBallInfo))

	flatBoostPadStats := &schema.BoostPadState{}
	for i := 0; i < flatGameTickPack.BoostPadStatesLength(); i++ {
		ok := flatGameTickPack.BoostPadStates(flatBoostPadStats, i)
		if !ok {
			continue
		}
		gameTickPack.BoostPadStates = append(gameTickPack.BoostPadStates, BoostPadState{
			IsActive: flatBoostPadStats.IsActive() == 1,
			Timer:    flatBoostPadStats.Timer(),
		})
	}
	flatGameInfo := &schema.GameInfo{}
	gameTickPack.GameInfo.unmarshal(flatGameTickPack.GameInfo(flatGameInfo))

	flatPlayerInfo := &schema.PlayerInfo{}
	for i := 0; i < flatGameTickPack.PlayersLength(); i++ {
		ok := flatGameTickPack.Players(flatPlayerInfo, i)
		if !ok {
			continue
		}
		var playerInfo PlayerInfo
		playerInfo.unmarshal(flatPlayerInfo)
		gameTickPack.Players = append(gameTickPack.Players, playerInfo)
	}

	flatTeamInfo := &schema.TeamInfo{}
	for i := 0; i < flatGameTickPack.TeamsLength(); i++ {
		ok := flatGameTickPack.Teams(flatTeamInfo, i)
		if !ok {
			continue
		}
		gameTickPack.Teams = append(gameTickPack.Teams, TeamInfo{
			TeamIndex: flatTeamInfo.TeamIndex(),
			Score:     flatTeamInfo.Score(),
		})
	}

	flatDropShotTile := &schema.DropshotTile{}
	for i := 0; i < flatGameTickPack.TileInformationLength(); i++ {
		ok := flatGameTickPack.TileInformation(flatDropShotTile, i)
		if !ok {
			continue
		}
		gameTickPack.TileInformation = append(gameTickPack.TileInformation, DropshotTile{
			TileState: flatDropShotTile.TileState(),
		})
	}
}

func (fieldInfo *FieldInfo) unmarshal(flatFieldInfo *schema.FieldInfo) {
	flatBoostPad := &schema.BoostPad{}
	for i := 0; i < flatFieldInfo.BoostPadsLength(); i++ {
		ok := flatFieldInfo.BoostPads(flatBoostPad, i)
		if !ok {
			continue
		}
		var vec3 Vector3
		flatVec3 := &schema.Vector3{}
		vec3.unmarshal(flatBoostPad.Location(flatVec3))
		if fieldInfo == nil {
			fmt.Println("FUCCKKKKK")
		}
		fieldInfo.BoostPads = append(fieldInfo.BoostPads, BoostPad{
			Location:    vec3,
			IsFullBoost: flatBoostPad.IsFullBoost() == 1,
		})
	}

	flatGoal := &schema.GoalInfo{}
	for i := 0; i < flatFieldInfo.GoalsLength(); i++ {
		ok := flatFieldInfo.Goals(flatGoal, i)
		if !ok {
			continue
		}
		var location Vector3
		var direction Vector3
		flatVec3 := &schema.Vector3{}
		location.unmarshal(flatGoal.Location(flatVec3))
		direction.unmarshal(flatGoal.Direction(flatVec3))
		fieldInfo.Goals = append(fieldInfo.Goals, GoalInfo{
			TeamNum:   flatGoal.TeamNum(),
			Location:  location,
			Direction: direction,
			Width:     flatGoal.Width(),
			Height:    flatGoal.Height(),
		})
	}
}

func (ballPrediction *BallPrediction) unmarshal(flatBallPrediction *schema.BallPrediction) {
	flatPredictionSlice := &schema.PredictionSlice{}
	for i := 0; i < flatBallPrediction.SlicesLength(); i++ {
		ok := flatBallPrediction.Slices(flatPredictionSlice, i)
		if !ok {
			continue
		}
		var physics Physics
		flatPhysics := &schema.Physics{}
		physics.unmarshal(flatPredictionSlice.Physics(flatPhysics))
		ballPrediction.Slices = append(ballPrediction.Slices, PredictionSlice{
			GameSeconds: flatPredictionSlice.GameSeconds(),
			Physics:     physics,
		})
	}
}

func (matchSettings *MatchSettings) unmarshal(flatMatchSettings *schema.MatchSettings) {

	matchSettings.GameMode = flatMatchSettings.GameMode()
	matchSettings.GameMap = flatMatchSettings.GameMap()
	matchSettings.ExistingMatchBehavior = flatMatchSettings.ExistingMatchBehavior()

	matchSettings.SkipReplays = flatMatchSettings.SkipReplays() == 1
	matchSettings.InstantStart = flatMatchSettings.InstantStart() == 1
	matchSettings.EnableLockstep = flatMatchSettings.EnableLockstep() == 1
	matchSettings.EnableRendering = flatMatchSettings.EnableRendering() == 1
	matchSettings.EnableStateSetting = flatMatchSettings.EnableStateSetting() == 1
	matchSettings.AutoSaveReplay = flatMatchSettings.AutoSaveReplay() == 1

	matchSettings.GameMapUpk = string(flatMatchSettings.GameMapUpk())

	matchSettings.MutatorSettings.unmarshal(flatMatchSettings.MutatorSettings(&schema.MutatorSettings{}))

	flatPlayerConfig := &schema.PlayerConfiguration{}
	for i := 0; i < flatMatchSettings.PlayerConfigurationsLength(); i++ {
		ok := flatMatchSettings.PlayerConfigurations(flatPlayerConfig, i)
		if !ok {
			continue
		}
		temp := PlayerConfiguration{}
		temp.unmarshal(flatPlayerConfig)
		matchSettings.PlayerConfigurations = append(matchSettings.PlayerConfigurations, temp)
	}
}

func (mutatorSettings *MutatorSettings) unmarshal(flatMutatorSettings *schema.MutatorSettings) {
	mutatorSettings.MatchLength = flatMutatorSettings.MatchLength()
	mutatorSettings.MaxScore = flatMutatorSettings.MaxScore()
	mutatorSettings.OvertimeOption = flatMutatorSettings.OvertimeOption()
	mutatorSettings.SeriesLengthOption = flatMutatorSettings.SeriesLengthOption()
	mutatorSettings.GameSpeedOption = flatMutatorSettings.GameSpeedOption()
	mutatorSettings.BallMaxSpeedOption = flatMutatorSettings.BallMaxSpeedOption()
	mutatorSettings.BallTypeOption = flatMutatorSettings.BallTypeOption()
	mutatorSettings.BallWeightOption = flatMutatorSettings.BallWeightOption()
	mutatorSettings.BallSizeOption = flatMutatorSettings.BallSizeOption()
	mutatorSettings.BallBouncinessOption = flatMutatorSettings.BallBouncinessOption()
	mutatorSettings.BoostOption = flatMutatorSettings.BoostOption()
	mutatorSettings.RumbleOption = flatMutatorSettings.RumbleOption()
	mutatorSettings.BoostStrengthOption = flatMutatorSettings.BoostStrengthOption()
	mutatorSettings.GravityOption = flatMutatorSettings.GravityOption()
	mutatorSettings.DemolishOption = flatMutatorSettings.DemolishOption()
	mutatorSettings.RespawnTimeOption = flatMutatorSettings.RespawnTimeOption()

}

func (playerConfig *PlayerConfiguration) unmarshal(flatPlayerConfig *schema.PlayerConfiguration) {
	playerConfig.Name = string(flatPlayerConfig.Name())
	playerConfig.Team = flatPlayerConfig.Team()
	playerConfig.SpawnId = flatPlayerConfig.SpawnId()
	playerConfig.Loadout.unmarshal(flatPlayerConfig.Loadout(&schema.PlayerLoadout{}))

}

func (playerLoadout *PlayerLoadout) unmarshal(flatPlayerLoadout *schema.PlayerLoadout) {
	playerLoadout.TeamColorId = flatPlayerLoadout.TeamColorId()
	playerLoadout.CustomColorId = flatPlayerLoadout.CustomColorId()
	playerLoadout.CarId = flatPlayerLoadout.CarId()
	playerLoadout.DecalId = flatPlayerLoadout.DecalId()
	playerLoadout.WheelsId = flatPlayerLoadout.WheelsId()
	playerLoadout.BoostId = flatPlayerLoadout.BoostId()
	playerLoadout.AntennaId = flatPlayerLoadout.AntennaId()
	playerLoadout.HatId = flatPlayerLoadout.HatId()
	playerLoadout.PaintFinishId = flatPlayerLoadout.PaintFinishId()
	playerLoadout.CustomFinishId = flatPlayerLoadout.CustomFinishId()
	playerLoadout.EngineAudioId = flatPlayerLoadout.EngineAudioId()
	playerLoadout.TrailsId = flatPlayerLoadout.TrailsId()
	playerLoadout.GoalExplosionId = flatPlayerLoadout.GoalExplosionId()
	playerLoadout.LoadoutPaint.unmarshal(flatPlayerLoadout.LoadoutPaint(&schema.LoadoutPaint{}))
	// playerLoadout.PrimaryColorLookup.unmarshal(flatPlayerLoadout.PrimaryColorLookup(&schema.Color{}))
	// playerLoadout.SecondaryColorLookup.unmarshal(flatPlayerLoadout.SecondaryColorLookup(&schema.Color{}))
	// TODO: For some reason these 2 fail
}

func (loadoutPaint *LoadoutPaint) unmarshal(flatLoadoutPaint *schema.LoadoutPaint) {
	loadoutPaint.CarPaintId = flatLoadoutPaint.CarPaintId()
	loadoutPaint.DecalPaintId = flatLoadoutPaint.DecalPaintId()
	loadoutPaint.WheelsPaintId = flatLoadoutPaint.WheelsPaintId()
	loadoutPaint.BoostPaintId = flatLoadoutPaint.BoostPaintId()
	loadoutPaint.AntennaPaintId = flatLoadoutPaint.AntennaPaintId()
	loadoutPaint.HatPaintId = flatLoadoutPaint.HatPaintId()
	loadoutPaint.TrailsPaintId = flatLoadoutPaint.TrailsPaintId()
	loadoutPaint.GoalExplosionPaintId = flatLoadoutPaint.GoalExplosionPaintId()
}

func (color *Color) unmarshal(flatColor *schema.Color) {
	color.A = flatColor.A()
	color.R = flatColor.R()
	color.G = flatColor.G()
	color.B = flatColor.B()
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
	playerInfo.Hitbox.unmarshal(flatPlayerInfo.Hitbox(flatBoxShape))

	flatVector3 := &schema.Vector3{}
	playerInfo.HitboxOffset.unmarshal(flatPlayerInfo.HitboxOffset(flatVector3))

	flatScoreInfo := &schema.ScoreInfo{}
	playerInfo.ScoreInfo.unmarshal(flatPlayerInfo.ScoreInfo(flatScoreInfo))

}

func (scoreInfo *ScoreInfo) unmarshal(flatScoreInfo *schema.ScoreInfo) {
	scoreInfo.Assists = flatScoreInfo.Assists()
	scoreInfo.Demolitions = flatScoreInfo.Demolitions()
	scoreInfo.Goals = flatScoreInfo.Goals()
	scoreInfo.OwnGoals = flatScoreInfo.OwnGoals()
	scoreInfo.Saves = flatScoreInfo.Saves()
	scoreInfo.Score = flatScoreInfo.Score()
	scoreInfo.Shots = flatScoreInfo.Shots()
}

func (boxShape *BoxShape) unmarshal(flatBoxShape *schema.BoxShape) {
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
	ballInfo.DropShotInfo.unmarshal(flatBallInfo.DropShotInfo(flatDropShotInfo))

	flatTouch := &schema.Touch{}
	if flatBallInfo.LatestTouch(flatTouch) != nil {
		ballInfo.LatestTouch.unmarshal(flatTouch)
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

	//flatRotator := &schema.Rotator{}
	//physics.Rotation.unmarshal(flatPhysics.Rotation(flatRotator))
}

func (touch *Touch) unmarshal(flatTouch *schema.Touch) {
	touch.GameSeconds = flatTouch.GameSeconds()
	touch.PlayerIndex = flatTouch.PlayerIndex()
	touch.PlayerName = string(flatTouch.PlayerName())
	touch.Team = flatTouch.Team()

	flatVector3 := &schema.Vector3{}
	touch.Location.unmarshal(flatTouch.Location(flatVector3))
	touch.Normal.unmarshal(flatTouch.Normal(flatVector3))
}

func (dropShotInfo *DropShotBallInfo) unmarshal(flatDropShotInfo *schema.DropShotBallInfo) {
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
