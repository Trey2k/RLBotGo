package RLBotGo

type GameState struct {
	GameTick       *GameTickPacket
	BallPrediction *BallPrediction

	MatchSettigns   *MatchSettings
	MatchSettingsOK bool
	FieldInfo       *FieldInfo
	FieldInfoOK     bool
}

type PredictionSlice struct {
	// The moment in game time that this prediction corresponds to.
	// This corresponds to 'secondsElapsed' in the GameInfo table.
	GameSeconds float32

	// The predicted location and motion of the object.
	Physics Physics
}

type BallPrediction struct {
	// A list of places the ball will be at specific times in the future.
	// It is guaranteed to sorted so that time increases with each slice.
	// It is NOT guaranteed to have a consistent amount of time between slices.
	Slices []PredictionSlice
}

type ControllerState struct {
	// -1 for full reverse, 1 for full forward
	Throttle float32

	// -1 for full left, 1 for full right
	Steer float32

	// -1 for nose down, 1 for nose up
	Pitch float32

	// -1 for full left, 1 for full right
	Yaw float32

	// -1 for roll left, 1 for roll right
	Roll float32

	// true if you want to press the jump button
	Jump bool

	// true if you want to press the boost button
	Boost bool

	// true if you want to press the handbrake button
	Handbrake bool

	// true if you want to press the 'use item' button, used in rumble etc.
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

type PlayerClass struct {
	Type     int
	BotSkill float64
}

type PlayerLoadout struct {
	TeamColorId     int32
	CustomColorId   int32
	CarId           int32
	DecalId         int32
	WheelsId        int32
	BoostId         int32
	AntennaId       int32
	HatId           int32
	PaintFinishId   int32
	CustomFinishId  int32
	EngineAudioId   int32
	TrailsId        int32
	GoalExplosionId int32
	LoadoutPaint    LoadoutPaint
	// Sets the primary color of the car to the swatch that most closely matches the provided
	// RGB color value. If set, this overrides teamColorId.
	PrimaryColorLookup Color
	// Sets the secondary color of the car to the swatch that most closely matches the provided
	// RGB color value. If set, this overrides customColorId.
	SecondaryColorLookup Color
}

type LoadoutPaint struct {
	CarPaintId           int32
	DecalPaintId         int32
	WheelsPaintId        int32
	BoostPaintId         int32
	AntennaPaintId       int32
	HatPaintId           int32
	TrailsPaintId        int32
	GoalExplosionPaintId int32
}

type Color struct {
	A uint8
	R uint8
	G uint8
	B uint8
}

type PlayerConfiguration struct {
	// Cannot be named 'class' because that breaks Java.
	// Cannot be named 'playerClass' because that breaks C#.
	Variety PlayerClass
	Name    string
	Team    int32
	Loadout PlayerLoadout
	// In the case where the requested player index is not available, spawnId will help
	// the framework figure out what index was actually assigned to this player instead.
	SpawnId int32
}

type MutatorSettings struct {
	MatchLength          int8
	MaxScore             int8
	OvertimeOption       int8
	SeriesLengthOption   int8
	GameSpeedOption      int8
	BallMaxSpeedOption   int8
	BallTypeOption       int8
	BallWeightOption     int8
	BallSizeOption       int8
	BallBouncinessOption int8
	BoostOption          int8
	RumbleOption         int8
	BoostStrengthOption  int8
	GravityOption        int8
	DemolishOption       int8
	RespawnTimeOption    int8
}

type MatchSettings struct {
	PlayerConfigurations  []PlayerConfiguration
	GameMode              int8
	GameMap               int8
	SkipReplays           bool
	InstantStart          bool
	MutatorSettings       MutatorSettings
	ExistingMatchBehavior int8
	EnableLockstep        bool
	EnableRendering       bool
	EnableStateSetting    bool
	AutoSaveReplay        bool
	// The name of a upk file, like UtopiaStadium_P, which should be loaded.
	// If specified, this overrides gameMap. On Steam version of Rocket League,
	// this can be used to load custom map files, but on Epic version it only
	// works on the Psyonix maps. Still useful because maintaining the gameMap
	// enum as new Psyonix maps are added is annoying.
	GameMapUpk string
}

type GoalInfo struct {
	TeamNum   int32
	Location  Vector3
	Direction Vector3
	Width     float32
	Height    float32
}

type FieldInfo struct {
	BoostPads []BoostPad // These will be sorted according to (y * 100 + x), and BoostInfo will be provided in the same order.
	Goals     []GoalInfo
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

	// The index of the player that sent the quick chat
	PlayerIndex int32

	// True if the chat is team only false if everyone can see it.
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
	// True if the boost can be picked up
	IsActive bool

	// The number of seconds since the boost has been picked up, or 0.0 if the boost is active.
	Timer float32
}

type BoostPad struct {
	Location    Vector3
	IsFullBoost bool
}

type Touch struct {
	// The name of the player involved with the touch.
	PlayerName string

	// Seconds that had elapsed in the game when the touch occurred.
	GameSeconds float32

	// The point32 of contact for the touch.
	Location Vector3

	// The direction of the touch.
	Normal Vector3

	// The Team which the touch belongs to, 0 for blue 1 for orange.
	Team int32

	// The index of the player involved with the touch.
	PlayerIndex int32
}

type DropShotBallInfo struct {
	AbsorbedForce    float32
	DamageIndex      int32
	ForceAccumRecent float32
}

type DropshotTile struct {
	// The amount of damage the tile has sustained.
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
	// True if your wheels are on the ground, the wall, or the ceiling. False if you're midair or turtling.
	HasWheelContact bool
	IsSupersonic    bool
	IsBot           bool
	// True if the player has jumped. Falling off the ceiling / driving off the goal post does not count.
	Jumped bool
	//  True if player has double jumped. False does not mean you have a jump remaining, because the
	//  aerial timer can run out, and that doesn't affect this flag.
	DoubleJumped bool
	Name         string
	Team         int32
	Boost        int32
	Hitbox       BoxShape
	HitboxOffset Vector3
	// In the case where the requested player index is not available, spawnId will help
	// the framework figure out what index was actually assigned to this player instead.
	SpawnId int32
}

type GameInfo struct {
	SecondsElapsed    float32
	GameTimeRemaining float32
	IsOvertime        bool
	IsUnlimitedTime   bool
	// True when cars are allowed to move, and during the pause menu. False during replays.
	IsRoundActive bool
	// True when the clock is paused due to kickoff, but false during kickoff countdown. In other words, it is true
	// while cars can move during kickoff. Note that if both players sit still, game clock start and this will become false.
	IsKickoffPause bool
	// Turns true after final replay, the moment the 'winner' screen appears. Remains true during next match
	// countdown. Turns false again the moment the 'choose team' screen appears.
	IsMatchEnded  bool
	WorldGravityZ float32
	// Game speed multiplier, 1.0 is regular game speed.
	GameSpeed float32
	// Tracks the number of physics frames the game has computed.
	// May increase by more than one across consecutive packets.
	// Data type will roll over after 207 days at 120Hz.
	FrameNum int32
}

type TeamInfo struct {
	TeamIndex int32
	// number of goals scored.
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
