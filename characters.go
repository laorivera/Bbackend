package main

type WeaponStats struct {
	Attackone   float64
	Attacktwo   float64
	Attackthree float64
}

var characterHolder = Stats{
	Strength:        0,
	Vigor:           0,
	Agility:         0,
	Dexterity:       0,
	Will:            0,
	Knowledge:       0,
	Resourcefulness: 0,
}

var characterStats = Stats{
	Strength:        1,
	Vigor:           1,
	Agility:         1,
	Dexterity:       1,
	Will:            1,
	Knowledge:       1,
	Resourcefulness: 1,
}

var classFighter = Stats{
	Strength:        15,
	Vigor:           15,
	Agility:         15,
	Dexterity:       15,
	Will:            15,
	Knowledge:       15,
	Resourcefulness: 15,
}

var classBarbarian = Stats{
	Strength:        25,
	Vigor:           25,
	Agility:         13,
	Dexterity:       12,
	Will:            18,
	Knowledge:       5,
	Resourcefulness: 7,
}

var classRogue = Stats{
	Strength:        9,
	Vigor:           10,
	Agility:         21,
	Dexterity:       25,
	Will:            10,
	Knowledge:       10,
	Resourcefulness: 20,
}

var classWizard = Stats{
	Strength:        6,
	Vigor:           7,
	Agility:         15,
	Dexterity:       17,
	Will:            20,
	Knowledge:       25,
	Resourcefulness: 15,
}

var classCleric = Stats{
	Strength:        11,
	Vigor:           13,
	Agility:         12,
	Dexterity:       14,
	Will:            23,
	Knowledge:       20,
	Resourcefulness: 12,
}

var classWarlock = Stats{
	Strength:        11,
	Vigor:           14,
	Agility:         14,
	Dexterity:       15,
	Will:            22,
	Knowledge:       15,
	Resourcefulness: 14,
}

var classBard = Stats{
	Strength:        13,
	Vigor:           13,
	Agility:         13,
	Dexterity:       20,
	Will:            11,
	Knowledge:       20,
	Resourcefulness: 15,
}

var classDruid = Stats{
	Strength:        12,
	Vigor:           13,
	Agility:         12,
	Dexterity:       12,
	Will:            18,
	Knowledge:       20,
	Resourcefulness: 18,
}

var classRanger = Stats{
	Strength:        12,
	Vigor:           10,
	Agility:         20,
	Dexterity:       18,
	Will:            10,
	Knowledge:       12,
	Resourcefulness: 23,
}

var classSorcerer = Stats{
	Strength:        10,
	Vigor:           10,
	Agility:         10,
	Dexterity:       18,
	Will:            25,
	Knowledge:       20,
	Resourcefulness: 12,
}

type Character struct {
	EquippedItems []Item_Armor
	BaseAttribute []Stats
}
