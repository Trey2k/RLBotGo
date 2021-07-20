package RLBotGo

// CONSTS

const ( // Dropshot tile info const
	Unkown  = 0
	Filled  = 1
	Damaged = 2
	Open    = 3
)

const (
	DataType_TickPacket = iota + 1
	DataType_FieldInfo
	DataType_MatchSettings
	DataType_PlayerInput
	// Depercated!!!
	DataType_ActorMapping
	// Depercated!!!
	DataType_ComputerId
	DataType_DesiredGameState
	DataType_RenderGroup
	DataType_QuickChat
	DataType_BallPrediction
	DataType_ReadyMessage
	DataType_MessagePacket
)

const (
	Information_IGotIt      = 0
	Information_NeedBoost   = 1
	Information_TakeTheShot = 2
	Information_Defending   = 3
	Information_GoForIt     = 4
	Information_Centering   = 5
	Information_AllYours    = 6
	Information_InPosition  = 7
	Information_Incoming    = 8
	Compliments_NiceShot    = 9
	Compliments_GreatPass   = 10
	Compliments_Thanks      = 11
	Compliments_WhatASave   = 12
	Compliments_NiceOne     = 13
	Compliments_WhatAPlay   = 14
	Compliments_GreatClear  = 15
	Compliments_NiceBlock   = 16
	Reactions_OMG           = 17
	Reactions_Noooo         = 18
	Reactions_Wow           = 19
	Reactions_CloseOne      = 20
	Reactions_NoWay         = 21
	Reactions_HolyCow       = 22
	Reactions_Whew          = 23
	Reactions_Siiiick       = 24
	Reactions_Calculated    = 25
	Reactions_Savage        = 26
	Reactions_Okay          = 27
	Apologies_Cursing       = 28
	Apologies_NoProblem     = 29
	Apologies_Whoops        = 30
	Apologies_Sorry         = 31
	Apologies_MyBad         = 32
	Apologies_Oops          = 33
	Apologies_MyFault       = 34
	PostGame_Gg             = 35
	PostGame_WellPlayed     = 36
	PostGame_ThatWasFun     = 37
	PostGame_Rematch        = 38
	PostGame_OneMoreGame    = 39
	PostGame_WhatAGame      = 40
	PostGame_NiceMoves      = 41
	PostGame_EverybodyDance = 42
	/// Custom text chats made by bot makers
	MaxPysonixQuickChatPresets = 43
	/// Waste of CPU cycles
	Custom_Toxic_WasteCPU = 44
	/// Git gud*
	Custom_Toxic_GitGut = 45
	/// De-Allocate Yourself
	Custom_Toxic_DeAlloc = 46
	/// 404: Your skill not found
	Custom_Toxic_404NoSkill = 47
	/// Get a virus
	Custom_Toxic_CatchVirus = 48
	/// Passing!
	Custom_Useful_Passing = 49
	/// Faking!
	Custom_Useful_Faking = 50
	/// Demoing!
	Custom_Useful_Demoing = 51
	/// BOOPING
	Custom_Useful_Bumping = 52
	/// The chances of that was 47525 to 1*
	Custom_Compliments_TinyChances = 53
	/// Who upped your skill level?
	Custom_Compliments_SkillLevel = 54
	/// Your programmer should be proud
	Custom_Compliments_proud = 55
	/// You're the GC of Bots
	Custom_Compliments_GC = 56
	/// Are you <Insert Pro>Bot? *
	Custom_Compliments_Pro = 57
	/// Lag
	Custom_Excuses_Lag = 58
	/// Ghost inputs
	Custom_Excuses_GhostInputs = 59
	/// RIGGED
	Custom_Excuses_Rigged = 60
	/// Mafia plays!
	Custom_Toxic_MafiaPlays = 61
	/// Yeet!
	Custom_Exclamation_Yeet = 62
)

type ControllerState struct {
	/// -1 for full reverse, 1 for full forward
	Throttle float32

	/// -1 for full left, 1 for full right
	Steer float32

	/// -1 for nose down, 1 for nose up
	Pitch float32

	/// -1 for full left, 1 for full right
	Yaw float32

	/// -1 for roll left, 1 for roll right
	Roll float32

	/// true if you want to press the jump button
	Jump bool

	/// true if you want to press the boost button
	Boost bool

	/// true if you want to press the handbrake button
	Handbrake bool

	/// true if you want to press the 'use item' button, used in rumble etc.
	UseItem bool
}

type PlayerInput struct {
	PlayerIndex     int32
	ControllerState ControllerState
}

type PlayerInputChange struct {
	PlayerIndex     int32
	ControllerState ControllerState

	// These are provided by Rocket League, and I'm passing them through. Theoretically they could be
	// inferred by jump + pitch + roll, but nice to have clarity.
	DodgeForward float32
	DodgeRight   float32
}

type ReadyMessage struct {
	// If this is set, RLBot will send BallPrediction data back to the client when available.
	WantsBallPredictions bool
	// If this is set, RLBot will send QuickChatMessages to the client when available.
	WantsQuickChat bool
	// If this is set, RLBot will send MessagePacket data back to the client when available.
	WantsGameMessages bool
}

type QuickChat struct {
	QuickChatSelection int8

	/// The index of the player that sent the quick chat
	PlayerIndex int32

	/// True if the chat is team only false if everyone can see it.
	TeamOnly bool

	MessageIndex int32

	TimeStamp float32
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Rotator struct {
	Pitch float32
	Yaw   float32
	Roll  float32
}

type BoxShape struct {
	Length float32
	Width  float32
	Height float32
}

type SphereShape struct {
	Diameter float32
}

type Physics struct {
	Location        Vector3
	Rotation        Rotator
	Velocity        Vector3
	AngularVelocity Vector3
}

type ScoreInfo struct {
	Score       int32
	Goals       int32
	OwnGoals    int32
	Assists     int32
	Saves       int32
	Shots       int32
	Demolitions int32
}

type BoostPadState struct {
	/// True if the boost can be picked up
	IsActive bool

	/// The number of seconds since the boost has been picked up, or 0.0 if the boost is active.
	Timer float32
}

type Touch struct {
	/// The name of the player involved with the touch.
	PlayerName string

	/// Seconds that had elapsed in the game when the touch occurred.
	GameSeconds float32

	/// The point32 of contact for the touch.
	Location Vector3

	/// The direction of the touch.
	Normal Vector3

	/// The Team which the touch belongs to, 0 for blue 1 for orange.
	Team int32

	/// The index of the player involved with the touch.
	PlayerIndex int32
}

type DropShotBallInfo struct {
	AbsorbedForce    float32
	DamageIndex      int32
	ForceAccumRecent float32
}

type DropshotTile struct {
	/// The amount of damage the tile has sustained.
	TileState int8
}

type BallInfo struct {
	Physics      Physics
	LatestTouch  Touch
	DropShotInfo DropShotBallInfo
	Shape        SphereShape
}

type PlayerInfo struct {
	Physics      Physics
	ScoreInfo    ScoreInfo
	IsDemolished bool
	/// True if your wheels are on the ground, the wall, or the ceiling. False if you're midair or turtling.
	HasWheelContact bool
	IsSupersonic    bool
	IsBot           bool
	/// True if the player has jumped. Falling off the ceiling / driving off the goal post does not count.
	Jumped bool
	///  True if player has double jumped. False does not mean you have a jump remaining, because the
	///  aerial timer can run out, and that doesn't affect this flag.
	DoubleJumped bool
	Name         string
	Team         int32
	Boost        int32
	Hitbox       BoxShape
	HitboxOffset Vector3
	/// In the case where the requested player index is not available, spawnId will help
	/// the framework figure out what index was actually assigned to this player instead.
	SpawnId int32
}

type GameInfo struct {
	SecondsElapsed    float32
	GameTimeRemaining float32
	IsOvertime        bool
	IsUnlimitedTime   bool
	/// True when cars are allowed to move, and during the pause menu. False during replays.
	IsRoundActive bool
	/// True when the clock is paused due to kickoff, but false during kickoff countdown. In other words, it is true
	/// while cars can move during kickoff. Note that if both players sit still, game clock start and this will become false.
	IsKickoffPause bool
	/// Turns true after final replay, the moment the 'winner' screen appears. Remains true during next match
	/// countdown. Turns false again the moment the 'choose team' screen appears.
	IsMatchEnded  bool
	WorldGravityZ float32
	/// Game speed multiplier, 1.0 is regular game speed.
	GameSpeed float32
	/// Tracks the number of physics frames the game has computed.
	/// May increase by more than one across consecutive packets.
	/// Data type will roll over after 207 days at 120Hz.
	FrameNum int32
}

type TeamInfo struct {
	TeamIndex int32
	/// number of goals scored.
	Score int32
}

type GameTickPacket struct {
	Players         []PlayerInfo
	BoostPadStates  []BoostPadState
	Ball            BallInfo
	GameInfo        GameInfo
	TileInformation []DropshotTile
	Teams           []TeamInfo
}
