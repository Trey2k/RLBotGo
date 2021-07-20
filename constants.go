package RLBotGo

const ( // Dropshot tile info const
	DropshotTile_Unkown = iota
	DropshotTile_Filled
	DropshotTile_Damaged
	DropshotTile_Open
)

const (
	PlayerClassType_RLBotPlayer = iota
	PlayerClassType_HumanPlayer
	PlayerClassType_PsyonixBotPlayer
	PlayerClassType_PartyMemberBotPlayer
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
	GameMode_Soccer = iota
	GameMode_Hoops
	GameMode_Dropshot
	GameMode_Hockey
	GameMode_Rumble
	GameMode_Heatseeker
	GameMode_Gridiron
)

const (
	MatchLength_Five_Minutes = iota
	MatchLength_Ten_Minutes
	MatchLength_Twenty_Minutes
	MatchLength_Unlimited
)

const (
	MaxScore_Unlimited = iota
	MaxScore_One_Goal
	MaxScore_Three_Goals
	MaxScore_Five_Goals
)

const (
	OvertimeOption_Unlimited = iota
	OvertimeOption_Five_Max_First_Score
	OvertimeOption_Five_Max_Random_Team
)

const (
	SeriesLengthOption_Unlimited = iota
	SeriesLengthOption_Three_Games
	SeriesLengthOption_Five_Games
	SeriesLengthOption_Seven_Games
)

const (
	GameSpeedOption_Default = iota
	GameSpeedOption_Slo_Mo
	GameSpeedOption_Time_Warp
)

const (
	BallMaxSpeedOption_Default = iota
	BallMaxSpeedOption_Default_Slow
	BallMaxSpeedOption_Default_Fast
	BallMaxSpeedOption_Default_Super_Fast
)

const (
	BallTypeOption_Default = iota
	BallTypeOption_Cube
	BallTypeOption_Puck
	BallTypeOption_Basketball
)

const (
	BallWeightOption_Default = iota
	BallWeightOption_Light
	BallWeightOption_Heavy
	BallWeightOption_Super_Light
)

const (
	BallSizeOption_Default = iota
	BallSizeOption_Small
	BallSizeOption_Large
	BallSizeOption_Gigantic
)

const (
	BallBouncinessOption_Default = iota
	BallBouncinessOption_Low
	BallBouncinessOption_High
	BallBouncinessOption_Super_High
)

const (
	BoostOption_Normal_Boost = iota
	BoostOption_Unlimited_Boost
	BoostOption_Slow_Recharge
	BoostOption_Rapid_Recharge
	BoostOption_No_Boost
)

const (
	RumbleOption_No_Rumble = iota
	RumbleOption_Default
	RumbleOption_Slow
	RumbleOption_Civilized
	RumbleOption_Destruction_Derby
	RumbleOption_Spring_Loaded
	RumbleOption_Spikes_Only
	RumbleOption_Spike_Rush
)

const (
	BoostStrengthOption_One = iota
	BoostStrengthOption_OneAndAHalf
	BoostStrengthOption_Two
	BoostStrengthOption_Ten
)

const (
	GravityOption_Default = iota
	GravityOption_Low
	GravityOption_High
	GravityOption_Super_High
)

const (
	DemolishOption_Default = iota
	DemolishOption_Disabled
	DemolishOption_Friendly_Fire
	DemolishOption_On_Contact
	DemolishOption_On_Contact_FF
)

const (
	RespawnTimeOption_Three_Seconds = iota
	RespawnTimeOption_Two_Seconds
	RespawnTimeOption_One_Seconds
	RespawnTimeOption_Disable_Goal_Reset
)

const (
	/// Restart the match if any match settings differ. This is the default because old RLBot always worked this way.
	ExistingMatchBehavior_Restart_If_Different = iota

	/// Always restart the match, even if config is identical
	ExistingMatchBehavior_Restart

	/// Never restart an existing match, just try to remove or spawn cars to match the configuration.
	/// If we are not in the middle of a match, a match will be started. Handy for LAN matches.
	ExistingMatchBehavior_Continue_And_Spawn
)

const (
	GameMap_DFHStadium = iota
	GameMap_Mannfield
	GameMap_ChampionsField
	GameMap_UrbanCentral
	GameMap_BeckwithPark
	GameMap_UtopiaColiseum
	GameMap_Wasteland
	GameMap_NeoTokyo
	GameMap_AquaDome
	GameMap_StarbaseArc
	GameMap_Farmstead
	GameMap_SaltyShores
	GameMap_DFHStadium_Stormy
	GameMap_DFHStadium_Day
	GameMap_Mannfield_Stormy
	GameMap_Mannfield_Night
	GameMap_ChampionsField_Day
	GameMap_BeckwithPark_Stormy
	GameMap_BeckwithPark_Midnight
	GameMap_UrbanCentral_Night
	GameMap_UrbanCentral_Dawn
	GameMap_UtopiaColiseum_Dusk
	GameMap_DFHStadium_Snowy
	GameMap_Mannfield_Snowy
	GameMap_UtopiaColiseum_Snowy
	GameMap_Badlands
	GameMap_Badlands_Night
	GameMap_TokyoUnderpass
	GameMap_Arctagon
	GameMap_Pillars
	GameMap_Cosmic
	GameMap_DoubleGoal
	GameMap_Octagon
	GameMap_Underpass
	GameMap_UtopiaRetro
	GameMap_Hoops_DunkHouse
	GameMap_DropShot_Core707
	GameMap_ThrowbackStadium
	GameMap_ForbiddenTemple
	GameMap_RivalsArena
	GameMap_Farmstead_Night
	GameMap_SaltyShores_Night
	GameMap_NeonFields
	GameMap_DFHStadium_Circuit
)

const (
	Information_IGotIt = iota
	Information_NeedBoost
	Information_TakeTheShot
	Information_Defending
	Information_GoForIt
	Information_Centering
	Information_AllYours
	Information_InPosition
	Information_Incoming
	Compliments_NiceShot
	Compliments_GreatPass
	Compliments_Thanks
	Compliments_WhatASave
	Compliments_NiceOne
	Compliments_WhatAPlay
	Compliments_GreatClear
	Compliments_NiceBlock
	Reactions_OMG
	Reactions_Noooo
	Reactions_Wow
	Reactions_CloseOne
	Reactions_NoWay
	Reactions_HolyCow
	Reactions_Whew
	Reactions_Siiiick
	Reactions_Calculated
	Reactions_Savage
	Reactions_Okay
	Apologies_Cursing
	Apologies_NoProblem
	Apologies_Whoops
	Apologies_Sorry
	Apologies_MyBad
	Apologies_Oops
	Apologies_MyFault
	PostGame_Gg
	PostGame_WellPlayed
	PostGame_ThatWasFun
	PostGame_Rematch
	PostGame_OneMoreGame
	PostGame_WhatAGame
	PostGame_NiceMoves
	PostGame_EverybodyDance
	/// Custom text chats made by bot makers
	MaxPysonixQuickChatPresets
	/// Waste of CPU cycles
	Custom_Toxic_WasteCPU
	/// Git gud*
	Custom_Toxic_GitGut
	/// De-Allocate Yourself
	Custom_Toxic_DeAlloc
	/// 404: Your skill not found
	Custom_Toxic_404NoSkill
	/// Get a virus
	Custom_Toxic_CatchVirus
	/// Passing!
	Custom_Useful_Passing
	/// Faking!
	Custom_Useful_Faking
	/// Demoing!
	Custom_Useful_Demoing
	/// BOOPING
	Custom_Useful_Bumping
	/// The chances of that was 47525 to 1*
	Custom_Compliments_TinyChances
	/// Who upped your skill level?
	Custom_Compliments_SkillLevel
	/// Your programmer should be proud
	Custom_Compliments_proud
	/// You're the GC of Bots
	Custom_Compliments_GC
	/// Are you <Insert Pro>Bot? *
	Custom_Compliments_Pro
	/// Lag
	Custom_Excuses_Lag
	/// Ghost inputs
	Custom_Excuses_GhostInputs
	/// RIGGED
	Custom_Excuses_Rigged
	/// Mafia plays!
	Custom_Toxic_MafiaPlays
	/// Yeet!
	Custom_Exclamation_Yeet
)
