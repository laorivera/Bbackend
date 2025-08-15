package main

type Enchanmentcat struct {
	Attributes map[string][]float32
	Others     map[string][]float32
}

type Enchantment struct {
	Helmet          map[string][]float32
	Chest           map[string][]float32
	Gloves          map[string][]float32
	Pants           map[string][]float32
	Boots           map[string][]float32
	Cloak           map[string][]float32
	Necklace        map[string][]float32
	Ring            map[string][]float32
	WeaponOneHand   map[string][]float32
	ShieldOneHand   map[string][]float32
	ShieldTwoHand   map[string][]float32
	PhysicalTwoHand map[string][]float32
	HybridTwoHand   map[string][]float32
	MagicalTwoHand  map[string][]float32
}

var Enchantments = Enchantment{
	// Armor
	Helmet: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.7, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3},
		"Luck":                {15, 30},
	},
	Chest: map[string][]float32{
		"Strength": {1, 3}, "Vigor": {1, 3}, "Agility": {1, 3}, "Dexterity": {1, 3}, "Will": {1, 3}, "Knowledge": {1, 3}, "Resourcefulness": {1, 3},
		"ArmorPenetration": {0.9, 1.8}, "PhysicalPowerBonus": {0.9, 1.8}, "MagicalPowerBonus": {0.7, 1.9}, "MagicPenetration": {1, 1.9},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {0.5, 1}, "MagicalDamageReduction": {0.5, 1}, "ProjectileReduction": {1, 2},
		"ActionSpeed": {0.9, 1.8}, "SpellCastingSpeed": {1, 2}, "RegularInteractionSpeed": {1.5, 3}, "MagicalInteractionSpeed": {1.5, 3},
		"BuffDurationBonus": {1, 2}, "DebuffDurationBonus": {1, 2}, "CooldownReduction": {1, 2},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {2, 4.5},
		"Luck": {20, 45},
	},
	Gloves: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"Luck": {15, 30},
	},
	Pants: map[string][]float32{
		"Strength": {1, 3}, "Vigor": {1, 3}, "Agility": {1, 3}, "Dexterity": {1, 3}, "Will": {1, 3}, "Knowledge": {1, 3}, "Resourcefulness": {1, 3},
		"ArmorPenetration": {0.9, 1.8}, "PhysicalPowerBonus": {0.9, 1.8}, "MagicalPowerBonus": {0.7, 1.8}, "MagicPenetration": {1, 1.8},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {0.5, 1}, "MagicalDamageReduction": {0.5, 1}, "ProjectileReduction": {1, 2},
		"ActionSpeed": {0.9, 1.8}, "SpellCastingSpeed": {1, 2}, "RegularInteractionSpeed": {1.5, 3}, "MagicalInteractionSpeed": {1.5, 3},
		"BuffDurationBonus": {1, 2}, "DebuffDurationBonus": {1, 2}, "CooldownReduction": {1, 2},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {2, 4.5},
		"Luck": {20, 45},
	},
	Boots: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"Luck": {15, 30},
	},
	Cloak: map[string][]float32{
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"MagicalHealing": {1, 2}, "MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3}, "PhysicalHealing": {1, 2},
		"Luck": {15, 30},
	},
	// Accessories
	Necklace: map[string][]float32{
		"AllAttributes":    {1},
		"ArmorPenetration": {0.9, 1.8}, "PhysicalPowerBonus": {0.9, 1.8}, "MagicalPowerBonus": {0.7, 1.8}, "MagicPenetration": {1, 1.8},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {0.5, 1}, "MagicalDamageReduction": {0.5, 1}, "ProjectileReduction": {1, 2},
		"ActionSpeed": {0.9, 1.8}, "SpellCastingSpeed": {1, 2}, "RegularInteractionSpeed": {1.5, 3}, "MagicalInteractionSpeed": {1.5, 3},
		"BuffDurationBonus": {1, 2}, "DebuffDurationBonus": {1, 2}, "CooldownReduction": {1, 2},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {2, 4.5},
		"MagicalHealing": {1, 2}, "MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3}, "PhysicalHealing": {1, 2},
		"Luck": {15, 30},
	},
	Ring: map[string][]float32{
		"AllAttributes":    {1},
		"ArmorPenetration": {0.6, 1.2}, "PhysicalPowerBonus": {0.6, 1.2}, "MagicalPowerBonus": {0.6, 1.2}, "MagicPenetration": {0.7, 1.2},
		"ArmorRating": {3, 6}, "MagicResistance": {3, 6}, "PhysicalDamageReduction": {0.3, 0.6}, "MagicalDamageReduction": {0.2, 0.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.2}, "SpellCastingSpeed": {0.7, 1.5}, "RegularInteractionSpeed": {1, 2}, "MagicalInteractionSpeed": {1, 2},
		"BuffDurationBonus": {1, 1.5}, "DebuffDurationBonus": {0.7, 1.5}, "CooldownReduction": {0.7, 1.4},
		"MemoryCapacityBonus": {1.5, 3}, "MemoryCapacityAdd": {1, 2},
		"MagicalHealing": {1, 2}, "MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3}, "PhysicalHealing": {1, 2},
		"Luck": {15, 30},
	},
	// Weapons
	WeaponOneHand: map[string][]float32{
		"PhysicalPower": {1, 2}, "ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {1, 2}, "AdditionalWeaponDamage": {1},
		"MagicalPower": {1, 2}, "MagicPenetration": {2, 3}, "MagicalPowerBonus": {1, 2},
		"CooldownReduction": {1, 2},
		"ActionSpeed":       {1, 2}, "RegularIntractionSpeed": {2, 3}, "MagicalInteractionSpeed": {2, 3}, "SpellCastingSpeed": {1, 2},
		"MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3},
		"MagicalHealing":    {1, 2},
		"BuffDurationBonus": {2, 3}, "DebuffDurationBonus": {2, 3},
		"MemoryCapacityAdd": {1, 3}, "MemoryCapacityBonus": {3, 6},
		"MoveSpeed": {1, 3}, "MoveSpeedBonus": {0.5, 1},
	},
	ShieldOneHand: map[string][]float32{
		"PhysicalPowerBonus": {1, 2}, "AdditionalWeaponDamage": {1},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {1, 2}, "MagicalDamageReduction": {1, 2}, "ProjectileReduction": {1, 2}, "CooldownReduction": {1, 2},
		"ActionSpeed": {1, 2}, "RegularInteractionSpeed": {2, 3}, "MagicalInteractionSpeed": {2, 3}, "SpellCastingSpeed": {1, 2},
		"MaxHealthBonus": {1, 2}, "MaxHealthAdd": {1, 3},
		"MagicalHealing":    {1, 2},
		"BuffDurationBonus": {2, 3}, "DebuffDurationBonus": {2, 3},
		"MoveSpeed": {1, 3}, "MoveSpeedBonus": {0.5, 1},
	},
	ShieldTwoHand: map[string][]float32{
		"ArmorRating": {5, 15}, "MagicResistance": {10, 20}, "PhysicalDamageReduction": {2, 3}, "MagicalDamageReduction": {2, 3}, "ProjectileReduction": {2, 3}, "CooldownReduction": {1, 2},
		"ActionSpeed": {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
	PhysicalTwoHand: map[string][]float32{
		"PhysicalPower": {1, 3}, "ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {2, 3}, "AdditionalWeaponDamage": {1, 2},
		"CooldownReduction": {2, 3},
		"ActionSpeed":       {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
	HybridTwoHand: map[string][]float32{
		"PhysicalPower": {1, 3}, "ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {2, 3}, "AdditionalWeaponDamage": {1, 2},
		"MagicalPower": {1, 3}, "MagicPenetration": {3, 4}, "MagicalPowerBonus": {2, 3}, "CooldownReduction": {2, 3},
		"ActionSpeed": {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4}, "SpellCastingSpeed": {2, 3},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {4, 8},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
	MagicalTwoHand: map[string][]float32{
		"MaigcalPower": {1, 3}, "MagicPenetration": {3, 4}, "MagicalPowerBonus": {2, 3},
		"CooldownReduction": {2, 3},
		"ActionSpeed":       {2, 3}, "RegularInteractionSpeed": {3, 4}, "MagicalInteractionSpeed": {3, 4}, "SpellCastingSpeed": {2, 3},
		"MaxHealthBonus": {2, 3}, "MaxHealthAdd": {1, 4},
		"MagicalHealing":    {1, 3},
		"BuffDurationBonus": {3, 4}, "DebuffDurationBonus": {3, 4},
		"MemoryCapacityAdd": {1, 4}, "MemoryCapacityBonus": {4, 8},
		"MoveSpeed": {1, 4}, "MoveSpeedBonus": {1, 2},
	},
}

/*
type List_Enchantment struct {
	Enchantments []map[string][]float32
}

var Listx = List_Enchantment{
	Enchantments: []map[string][]float32{
		Enchantments.Helmet,
	},
}
*/
