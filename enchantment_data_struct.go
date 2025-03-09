package main

type Enchanmentcat struct {
	Attributes map[string][]float32
	Others     map[string][]float32
}

type Enchantment struct {
	Helmet   map[string][]float32
	Chest    map[string][]float32
	Gloves   map[string][]float32
	Pants    map[string][]float32
	Boots    map[string][]float32
	Cloak    map[string][]float32
	Necklace map[string][]float32
}

var Enchantments = Enchantment{

	Helmet: map[string][]float32{
		"Strength": {1, 2}, "Vigor": {1, 2}, "Agility": {1, 2}, "Dexterity": {1, 2}, "Will": {1, 2}, "Knowledge": {1, 2}, "Resourcefulness": {1, 2},
		"ArmorPenetration": {2, 3}, "PhysicalPowerBonus": {1, 2}, "PhysicalPower": {1, 2}, "MagicalPower": {1, 2}, "MagicPowerBonus": {1, 2}, "TrueMagicalDamage": {1, 1}, "AddiontalMagicalDamage": {1, 1}, "MagicPenetration": {2, 3},
		"ArmorRating": {3, 5}, "MagicResistance": {3, 5}, "PhysicalDamageReduction": {1, 1.5}, "MagicalDamageReduction": {1, 1.5}, "ProjectileReduction": {1, 1.5},
		"ActionSpeed": {1, 1.5}, "SpellCastingSpeed": {1, 1.5},
		"MaxHealthBonus": {1, 1.5}, "MaxHealthAdd": {2, 3},
		"BuffDurationBonus": {3, 5}, "DebuffDurationBonus": {1, 3},
		"MemoryCapacityAdd": {1, 2}, "MemoryCapacityBonus": {2, 5},
		"Luck": {15, 25},
	},
	Necklace: map[string][]float32{
		"AllAttributes":    {1},
		"ArmorPenetration": {2.5, 3.5}, "TruePhysicalDamage": {1}, "AddiontalPhysicalDamage": {1}, "PhysicalPower": {1, 3},
		"MagicPenetration": {2.5, 3.5}, "TrueMagicalDamage": {1}, "AddiontalMagicalDamage": {1}, "MagicalPower": {1, 3},
		"ArmorRating": {5, 10}, "MagicResistance": {5, 10}, "PhysicalDamageReduction": {1.5, 2}, "MagicalDamageReduction": {1.5, 2}, "ProjectileReduction": {1.5, 2},
		"ActionSpeed": {1.5, 2}, "SpellCastingSpeed": {1.5, 2},
		"MaxHealthBonus": {1.5, 2}, "MaxHealthAdd": {3, 5},
		"PhysicalHealing": {1, 3}, "MagicalHealing": {1, 3},
		"BuffDurationBonus": {5, 10}, "DebuffDurationBonus": {2, 5},
		"MemoryCapacityAdd": {2, 3}, "MemoryCapacityBonus": {5, 10},
		"Luck": {25, 50},
	},
}

type List_Enchantment struct {
	Enchantments []map[string][]float32
}

var Listx = List_Enchantment{
	Enchantments: []map[string][]float32{
		Enchantments.Helmet,
	},
}
