package main

const (
	// --- Window ---
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "Dino Game"
	TargetFPS    = 60

	FightDemoWindowTitle = "Dino Fight Demo"

	// --- Dino (shared) ---
	DinoWidth  = 50
	DinoHeight = 100

	// --- Enemy dino generation ---
	EnemyMinHealth   = 50
	EnemyHealthRange = 50 // health = rand(EnemyHealthRange) + EnemyMinHealth
	EnemyMinDamage   = 5
	EnemyDamageRange = 5 // damage = rand(EnemyDamageRange) + EnemyMinDamage

	// --- Dino movement ---
	DefaultDinoSpeed = 2.0

	// --- Helmet generation ---
	HelmetMinCapacity   = 15
	HelmetCapacityRange = 16 // capacity = rand(HelmetCapacityRange) + HelmetMinCapacity
	HelmetMinCost       = 10
	HelmetCostRange     = 41 // cost = rand(HelmetCostRange) + HelmetMinCost
	HelmetMinNumHits    = 20
	HelmetNumHitsRange  = 31 // numHits = rand(HelmetNumHitsRange) + HelmetMinNumHits
	HelmetMinBlock      = 1
	HelmetBlockRange    = 5 // baseBlock = rand(HelmetBlockRange) + HelmetMinBlock
	HelmetDamageRange   = 4 // baseDamage = rand(HelmetDamageRange), min is 0

	// --- Modifier generation ---
	ModifierMinCost     = 1
	ModifierCostRange   = 20 // cost = rand(ModifierCostRange) + ModifierMinCost
	ModifierMinSize     = 1
	ModifierSizeRange   = 5  // size = rand(ModifierSizeRange) + ModifierMinSize
	ModifierDamageRange = 10 // dmg = rand(ModifierDamageRange), min is 0
	ModifierBlockRange  = 10 // block = rand(ModifierBlockRange), min is 0

	// --- Player defaults ---
	PlayerStartGold      = 100
	PlayerStartHealth    = 75
	PlayerStartMaxHealth = 75
	PlayerStartDamage    = 8
	PlayerStartBlock     = 0

	// --- Shop ---
	ShopCapacity      = 5
	ShopStartingItems = 3

	// --- Battle timing (in frames) ---
	KnockbackDuration = 60
	HitCooldown       = 30

	// --- Ground ---
	GroundHeight     = 120
	GroundDirtY      = 28 // y offset where dirt color starts within ground rect
	GroundDirtHeight = 92 // height of dirt portion

	// --- Health bar (on dino sprite) ---
	DinoHealthBarYOffset = -20 // y offset above dino
	DinoHealthBarHeight  = 10

	// --- Battle screen ---
	PlayerDinoStartXDivisor  = 4
	EnemyDinoStartXNumerator = 3
	EnemyDinoStartXDivisor   = 4
	BattleResultY            = 150
	BattleResultFontSize     = 36

	// --- Stat panel ---
	StatPanelWidth         = 220
	StatPanelHeight        = 130
	StatPanelPadding       = 12
	StatPanelFontSizeLarge = 18
	StatPanelFontSizeSmall = 16
	StatPanelMargin        = 24
	StatPanelTitleY        = 10
	StatPanelHealthY       = 38
	StatPanelHealthHeight  = 14
	StatPanelHPTextY       = 60
	StatPanelStatTextY     = 82
	StatPanelSpeedTextY    = 104
	StatPanelBlockX        = 110

	// --- Shop UI ---
	ShopBGPadding = 40

	// --- Button ---
	ButtonFontSize    = 20
	ButtonBorderWidth = 2.0

	// --- Fight screen buttons ---
	// X values are offsets from screen center
	FightButtonY    = 260
	StartButtonXOff = -70
	StartButtonW    = 140
	StartButtonH    = 48
	ShopButtonXOff  = -150
	ShopButtonW     = 120
	ShopButtonH     = 48
	FightAgainXOff  = 30
	FightAgainW     = 150
	FightAgainH     = 48

	// --- Death screen ---
	DeathScreenFontSize = 60
	DeathText           = "YOU DIED"
)
