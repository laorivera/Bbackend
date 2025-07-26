package main

type BaseAttribute struct {
	Strength        map[int]int `json:"strength"`
	Vigor           map[int]int `json:"vigor"`
	Agility         map[int]int `json:"agility"`
	Dexterity       map[int]int `json:"dexterity"`
	Will            map[int]int `json:"will"`
	Knowledge       map[int]int `json:"knowledge"`
	Resourcefulness map[int]int `json:"resourcefulness"`
}

type Item_Armor struct {
	Name                    string            `json:"name"`
	ArmorRatings            map[int][]int     `json:"armorRatings"`
	MaxHealthAdd            map[int]float64   `json:"maxHealthAdd"`
	MaxHealthBonus          map[int][]float64 `json:"maxHealthBonus"` // exception
	MagicResistance         map[int][]float64 `json:"magicalResistance"`
	ProjectileReduction     float64           `json:"projectileReduction"` //
	ProjectileReductionRate map[int]float64   `json:"projectileReductionRate"`
	HeadshotReduction       float64           `json:"headshotReduction"` //
	BaseAttribute           BaseAttribute     `json:"BaseAttribute"`     //
	PhysicalPower           float64           `json:"physicalPower"`
	MoveSpeed               map[int]int       `json:"moveSpeed"` //
	MoveSpeedBonus          float64           `json:"moveSpeedbonus"`
	ArmorPenetration        float64           `json:"armorpPenetration"`
	MagicPenetration        float64           `json:"magicPenetration"`
	BuffDuration            float64           `json:"buffDuration"`
	MagicalPower            map[int]float64   `json:"magicalPower"`
	MagicalPowerBonus       map[int]float64   `json:"magicalPowerBonus"`
	RegularInteractionSpeed map[int]float64   `json:"regularInteraction"`
	MagicalHealing          map[int]int       `json:"magicalHealing"`
	ActionSpeed             map[int]float64   `json:"actionSpeed"`
	MagicalDamageReduction  float64           `json:"magicalDamageReduction"`
	PhysicalDamageReduction float64           `json:"physicalDamageReduction"`
	Luck                    int               `json:"luck"`
	ArmorType               string            `json:"armorType"`   //
	Rarities                []int             `json:"rarities"`    //
	SlotType                string            `json:"slotType"`    //
	InvWidth                int               `json:"invWidth"`    //
	InvHeight               int               `json:"invHeight"`   //
	FlavorText              string            `json:"flavorText"`  //
	AP                      map[int]int       `json:"ap"`          //
	XP                      map[int]int       `json:"xp"`          //
	GearScore               map[int]int       `json:"gearScore"`   //
	Classes                 []string          `json:"classes"`     //
	NumEnchants             map[int]int       `json:"numEnchants"` //
}

type Item_Weapon struct {
	Name                    string            `json:"name"`
	Classes                 []string          `json:"classes"` // List of classes (e.g., "Fighter", "Barbarian")
	DamageRatings           map[int][]int     `json:"damageRatings"`
	MoveSpeed               map[int]int       `json:"moveSpeed"`
	MoveSpeedBonus          float64           `json:"moveSpeedbonus"`
	DamageTypes             []string          `json:"damageTypes"`
	ImpactZones             []int             `json:"impactZones"`
	ImpactPower             int               `json:"impactPower"`
	ComboDamage             []int             `json:"comboDamage"`
	WeaponType              []string          `json:"type"`        // Weapon type (e.g., "Axe")
	Rarities                []int             `json:"rarities"`    // List of rarities
	HandType                string            `json:"handType"`    // Hand type (e.g., "Two Handed")
	SlotType                string            `json:"slotType"`    // Slot type (e.g., "Main Hand")
	GearScore               map[int]int       `json:"gearScore"`   // Gear score for levels 0 to 7
	InvWidth                int               `json:"invWidth"`    // Inventory width
	InvHeight               int               `json:"invHeight"`   // Inventory height
	FlavorText              string            `json:"flavorText"`  // Description text
	AP                      map[int]int       `json:"ap"`          // Action points for levels 3 to 7
	XP                      map[int]int       `json:"xp"`          // XP for levels 1 to 7
	NumEnchants             map[int]int       `json:"numEnchants"` // Number of enchants for levels 3 to 7
	MaxHealthAdd            map[int]float64   `json:"maxHealthAdd"`
	MaxHealthBonus          map[int][]float64 `json:"maxHealthBonus"`
	MagicResistance         map[int][]float64 `json:"magicalResistance"`
	ProjectileReduction     float64           `json:"projectileReduction"`
	ProjectileReductionRate map[int]float64   `json:"projectileReductionRate"`
	HeadshotReduction       float64           `json:"headshotReduction"`
	BaseAttribute           BaseAttribute     `json:"BaseAttribute"`
	PhysicalPower           float64           `json:"physicalPower"`
	ArmorPenetration        float64           `json:"armorpPenetration"`
	MagicPenetration        float64           `json:"magicPenetration"`
	BuffDuration            float64           `json:"buffDuration"`
	MagicalPower            map[int]float64   `json:"magicalPower"`
	MagicalPowerBonus       map[int]float64   `json:"magicalPowerBonus"`
	RegularInteractionSpeed map[int]float64   `json:"regularInteraction"`
	MagicalHealing          map[int]int       `json:"magicalHealing"`
	ActionSpeed             map[int]float64   `json:"actionSpeed"`
	MagicalDamageReduction  float64           `json:"magicalDamageReduction"`
	PhysicalDamageReduction float64           `json:"physicalDamageReduction"`
	Luck                    int               `json:"luck"`
}

type Item_Accessory struct {
	Name           string          `json:"name"`
	Classes        []string        `json:"classes"` // List of classes (e.g., "Fighter", "Barbarian")
	MaxHealthAdd   map[int]float64 `json:"maxHealthAdd"`
	MoveSpeed      map[int]int     `json:"moveSpeed"` //
	MoveSpeedBonus float64         `json:"moveSpeedbonus"`
	BaseAttribute  BaseAttribute   `json:"BaseAttribute"` // BaseStats 1 to 7
	Luck           map[int]int     `json:"luck"`
	Type           string          `json:"type"`         // Accessory type
	Rarities       []int           `json:"rarities"`     // List of rarities
	SlotType       string          `json:"slotType"`     // Slot type (e.g., "Necklace")
	InvWidth       int             `json:"invWidth"`     // Inventory width
	InvHeight      int             `json:"invHeight"`    // Inventory height
	FlavorText     string          `json:"flavorText"`   // Description text
	MaxAmmoCount   int             `json:"maxAmmoCount"` // Max ammo count
	MaxCount       int             `json:"maxCount"`     // Max count
	AP             map[int]int     `json:"ap"`           // Action points for levels 3 to 7
	XP             map[int]int     `json:"xp"`           // XP for levels 1 to 7
	GearScore      map[int]int     `json:"gearScore"`    // Gear score for levels 0 to 7
	NumEnchants    map[int]int     `json:"numEnchants"`  // Number of enchants for levels 3 to 7

}
