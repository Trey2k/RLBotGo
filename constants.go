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
	// Restart the match if any match settings differ. This is the default because old RLBot always worked this way.
	ExistingMatchBehavior_Restart_If_Different = iota

	// Always restart the match, even if config is identical
	ExistingMatchBehavior_Restart

	// Never restart an existing match, just try to remove or spawn cars to match the configuration.
	// If we are not in the middle of a match, a match will be started. Handy for LAN matches.
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
	QuickChat_Information_IGotIt = iota
	QuickChat_Information_NeedBoost
	QuickChat_Information_TakeTheShot
	QuickChat_Information_Defending
	QuickChat_Information_GoForIt
	QuickChat_Information_Centering
	QuickChat_Information_AllYours
	QuickChat_Information_InPosition
	QuickChat_Information_Incoming
	QuickChat_Compliments_NiceShot
	QuickChat_Compliments_GreatPass
	QuickChat_Compliments_Thanks
	QuickChat_Compliments_WhatASave
	QuickChat_Compliments_NiceOne
	QuickChat_Compliments_WhatAPlay
	QuickChat_Compliments_GreatClear
	QuickChat_Compliments_NiceBlock
	QuickChat_Reactions_OMG
	QuickChat_Reactions_Noooo
	QuickChat_Reactions_Wow
	QuickChat_Reactions_CloseOne
	QuickChat_Reactions_NoWay
	QuickChat_Reactions_HolyCow
	QuickChat_Reactions_Whew
	QuickChat_Reactions_Siiiick
	QuickChat_Reactions_Calculated
	QuickChat_Reactions_Savage
	QuickChat_Reactions_Okay
	QuickChat_Apologies_Cursing
	QuickChat_Apologies_NoProblem
	QuickChat_Apologies_Whoops
	QuickChat_Apologies_Sorry
	QuickChat_Apologies_MyBad
	QuickChat_Apologies_Oops
	QuickChat_Apologies_MyFault
	QuickChat_PostGame_Gg
	QuickChat_PostGame_WellPlayed
	QuickChat_PostGame_ThatWasFun
	QuickChat_PostGame_Rematch
	QuickChat_PostGame_OneMoreGame
	QuickChat_PostGame_WhatAGame
	QuickChat_PostGame_NiceMoves
	QuickChat_PostGame_EverybodyDance
	// Custom text chats made by bot makers
	QuickChat_MaxPysonixQuickChatPresets
	// Waste of CPU cycles
	QuickChat_Custom_Toxic_WasteCPU
	// Git gud*
	QuickChat_Custom_Toxic_GitGut
	// De-Allocate Yourself
	QuickChat_Custom_Toxic_DeAlloc
	// 404: Your skill not found
	QuickChat_Custom_Toxic_404NoSkill
	// Get a virus
	QuickChat_Custom_Toxic_CatchVirus
	// Passing!
	QuickChat_Custom_Useful_Passing
	// Faking!
	QuickChat_Custom_Useful_Faking
	// Demoing!
	QuickChat_Custom_Useful_Demoing
	// BOOPING
	QuickChat_Custom_Useful_Bumping
	// The chances of that was 47525 to 1*
	QuickChat_Custom_Compliments_TinyChances
	// Who upped your skill level?
	QuickChat_Custom_Compliments_SkillLevel
	// Your programmer should be proud
	QuickChat_Custom_Compliments_proud
	// You're the GC of Bots
	QuickChat_Custom_Compliments_GC
	// Are you <Insert Pro>Bot? *
	QuickChat_Custom_Compliments_Pro
	// Lag
	QuickChat_Custom_Excuses_Lag
	// Ghost inputs
	QuickChat_Custom_Excuses_GhostInputs
	// RIGGED
	QuickChat_Custom_Excuses_Rigged
	// Mafia plays!
	QuickChat_Custom_Toxic_MafiaPlays
	// Yeet!
	QuickChat_Custom_Exclamation_Yeet
)

const (
	RenderType_DrawLine2D = iota + 1
	RenderType_DrawLine3D
	RenderType_DrawLine2D_3D
	RenderType_DrawRect2D
	RenderType_DrawRect3D
	RenderType_DrawString2D
	RenderType_DrawString3D
	RenderType_DrawCenteredRect3D
)

const (
	GamesMessageType_PlayerStatEvent = iota
	GamesMessageType_PlayerSpectate
	GamesMessageType_PlayerInputChange
)
