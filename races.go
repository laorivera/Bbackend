package main

//var races = []string{"Skeleton", "Elite Skeleton", "Nightmare Skeleton", "Skeleton Champion", "Lizardmen", "FrostWalker", "Zombie", "DarkElf", "Mummy", "Panther", "Felidian", "Lycan", "Orc", "Elf"}

var raceStats = map[string]Stats{
	"Orc": {
		Strength: 1,
		Agility:  -1,
	},
	"Elf": {
		Agility:  1,
		Strength: -1,
	},
	"Dark Elf": {
		Will:      1,
		Dexterity: -1,
	},
	"Mummy": {
		Vigor:   1,
		Agility: -1,
	},
	"Panther": {
		Agility:  2,
		Strength: -1,
		Vigor:    -1,
	},
	"Felidian": {
		Dexterity: 2,
		Strength:  -1,
		Vigor:     -1,
	},
	"Lycan": {
		Vigor: 4,
	},
}

var raceComputed = map[string]Computed_Stats{
	"Frost Walker": {
		Health:      3,
		ActionSpeed: -1.5,
	},
	"Zombie": {
		BuffDuration:   5,
		DebuffDuration: -5,
	},
	"Skeleton": {
		MagicRating:     10,
		FromArmorRating: -10,
	},
	"Elite Skeleton": {
		MagicRating:     10,
		FromArmorRating: -10,
	},
	"Nightmare Skeleton": {
		FromArmorRating: 10,
		MagicRating:     -10,
	},
	"Skeleton Champion": {
		FromArmorRating: 5,
		MagicRating:     -5,
	},
	"Lizardmen": {
		FromArmorRating:   20,
		HeadshotReduction: 10,
	},
	"Lycan": {
		HeadshotReduction: 10,
	},
}
