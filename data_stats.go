package main

// STRUCT STATS
type Stats struct {
	Strength        int
	Vigor           int
	Agility         int
	Dexterity       int
	Will            int
	Knowledge       int
	Resourcefulness int
}

// STRUCT COMPLETE STATS
type Computed_Stats struct {
	Health                       float64
	MaxHealthBonus               float64
	ActionSpeed                  float64
	RegularInteractionSpeed      float64
	MoveSpeed                    float64
	MoveSpeedCalc                float64
	PhysicalPower                float64
	PhysicalPowerBonus           float64
	HealthRecovery               float64
	ManualDexterity              float64
	EquipSpeed                   float64
	MagicalPower                 float64
	MagicalPowerBonus            float64
	BuffDuration                 float64
	MagicRating                  float64
	MagicalDamageReduction       float64
	DebuffDuration               float64
	MemoryCapacity               float64
	MemoryCapacityBonus          float64
	SpellRecovery                float64
	SpellCastingSpeed            float64
	MagicalInteractionSpeed      float64
	Persuasiveness               float64
	CooldownReduction            float64
	PhysicalDamageReduction      float64
	SpellRecoveryBonus           float64
	PhysicalHealing              int
	MagicalHealing               int
	MemorySpellPayload           int
	MemoryMusicPayload           int
	UtilityEffectiveness         int
	Luck                         int
	ArmorPenetration             float64
	MagicPenetration             float64
	HeadshotReduction            float64
	ProjectileReduction          float64
	FromArmorRating              int
	BonusPhysicalDamageReduction float64
	BonusMagicalDamageReduction  float64
	BonusPhysicalPower           float64
	BonusMagicalPower            float64
}

type Computed_Stats_Weapon struct {
	PrimaryWeapon        WeaponStats
	SecundaryWeapon      WeaponStats
	PrimaryImpactPower   int
	SecondaryImpactPower int
}
