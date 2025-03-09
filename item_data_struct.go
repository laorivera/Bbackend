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
	Name                string          `json:"name"`
	ArmorRatings        map[int][]int   `json:"armorRatings"`
	MaxHealthAdd        map[int]int     `json:"maxHealthAdd"`
	MagicResistance     map[int][]int   `json:"magicResistance"`
	ProjectileReduction float64         `json:"projectileReduction"` // %
	HeadshotReduction   float64         `json:"headshotReduction"`   // %
	BaseAttribute       BaseAttribute   `json:"BaseAttribute"`       // BaseStats 1 to 7
	MoveSpeed           map[int]int     `json:"moveSpeed"`           // %
	MoveSpeedBonus      map[int]float64 `json:"moveSpeedbonus"`
	ArmorPenetration    float64         `json:"armorpPenetration"`
	MagicPenetration    float64         `json:"magicPenetration"`
	MagicalPower        int             `json:"magicalPower"`
	RegularInteraction  int             `json:"regularInteraction"`
	MagicalHealing      int             `json:"magicalHealing"`
	Luck                map[int]int     `json:"luck"`
	TruePhysicalDamage  map[int]int     `json:"truePhysicaldamage"`
	TrueMagicalDamage   map[int]int     `json:"trueMagicaldamage"`
	ArmorType           string          `json:"armorType"`    // Armor type (e.g., "Plate")
	Rarities            []int           `json:"rarities"`     // List of rarities
	SlotType            string          `json:"slotType"`     // Slot type (e.g., "Head")
	InvWidth            int             `json:"invWidth"`     // Inventory width
	InvHeight           int             `json:"invHeight"`    // Inventory height
	FlavorText          string          `json:"flavorText"`   // Description text
	MaxAmmoCount        int             `json:"maxAmmoCount"` // Max ammo count
	MaxCount            int             `json:"maxCount"`     // Max count
	AP                  map[int]int     `json:"ap"`           // Action points for levels 3 to 7
	XP                  map[int]int     `json:"xp"`           // XP for levels 1 to 7
	GearScore           map[int]int     `json:"gearScore"`    // Gear score for levels 0 to 7
	Classes             []string        `json:"classes"`      // List of classes (e.g., "Fighter")
	SellPrices          map[int]int     `json:"sellPrices"`   // Sell prices for levels 1 to 7
	NumEnchants         map[int]int     `json:"numEnchants"`  // Number of enchants for levels 3 to 7
}

type Item_Weapon struct {
	Name          string
	Classes       []string      `json:"classes"` // List of classes (e.g., "Fighter", "Barbarian")
	DamageRatings map[int][]int `json:"damageRatings"`
	MoveSpeed     int           `json:"moveSpeed"`
	DamageType    string        `json:"damageType"`
	ImpactZones   []int         `json:"impactZones"`
	ImpactPower   int           `json:"impactPower"`
	WeaponType    string        `json:"type"`        // Weapon type (e.g., "Axe")
	Rarities      []int         `json:"rarities"`    // List of rarities
	HandType      string        `json:"handType"`    // Hand type (e.g., "Two Handed")
	SlotType      string        `json:"slotType"`    // Slot type (e.g., "Main Hand")
	GearScore     map[int]int   `json:"gearScore"`   // Gear score for levels 0 to 7
	InvWidth      int           `json:"invWidth"`    // Inventory width
	InvHeight     int           `json:"invHeight"`   // Inventory height
	FlavorText    string        `json:"flavorText"`  // Description text
	AP            map[int]int   `json:"ap"`          // Action points for levels 3 to 7
	XP            map[int]int   `json:"xp"`          // XP for levels 1 to 7
	SellPrices    map[int]int   `json:"sellPrices"`  // Sell prices for levels 1 to 7
	NumEnchants   map[int]int   `json:"numEnchants"` // Number of enchants for levels 3 to 7

}

type Item_Accessory struct {
	Name          string        `json:"name"`
	Classes       []string      `json:"classes"`       // List of classes (e.g., "Fighter", "Barbarian")
	BaseAttribute BaseAttribute `json:"BaseAttribute"` // BaseStats 1 to 7
	Luck          map[int]int   `json:"luck"`
	MoveSpeed     int           `json:"moveSpeed"`
	Type          string        `json:"type"`         // Accessory type
	Rarities      []int         `json:"rarities"`     // List of rarities
	SlotType      string        `json:"slotType"`     // Slot type (e.g., "Necklace")
	InvWidth      int           `json:"invWidth"`     // Inventory width
	InvHeight     int           `json:"invHeight"`    // Inventory height
	FlavorText    string        `json:"flavorText"`   // Description text
	MaxAmmoCount  int           `json:"maxAmmoCount"` // Max ammo count
	MaxCount      int           `json:"maxCount"`     // Max count
	AP            map[int]int   `json:"ap"`           // Action points for levels 3 to 7
	XP            map[int]int   `json:"xp"`           // XP for levels 1 to 7
	GearScore     map[int]int   `json:"gearScore"`    // Gear score for levels 0 to 7
	SellPrices    map[int]int   `json:"sellPrices"`   // Sell prices for levels 1 to 7
	NumEnchants   map[int]int   `json:"numEnchants"`  // Number of enchants for levels 3 to 7

}
