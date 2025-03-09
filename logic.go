package main

import (
	"encoding/json"
	"math"
	"os"
	"strconv"
)

// EJECUTA CALCULO ALL STATS CURVES

// FIND ITEMS BY SLOT TYPE
func ItemsBySlotType_Json(class string, slot string) []string {
	switch class {
	case "1":
		class = "Fighter"
	case "2":
		class = "Barbarian"
	case "3":
		class = "Rogue"
	case "4":
		class = "Wizard"
	case "5":
		class = "Cleric"
	case "6":
		class = "Warlock"
	case "7":
		class = "Bard"
	case "8":
		class = "Druid"
	case "9":
		class = "Ranger"
	default:

	}
	var result []string
	for i := 0; i < len(Items.ItemsArmor); i++ {
		for _, c := range Items.ItemsArmor[i].Classes {
			if c == class && slot == Items.ItemsArmor[i].SlotType {
				result = append(result, Items.ItemsArmor[i].Name)
			}
		}

	}
	return result
}

func ItemsBySlotType(class string, slot string) []Item_Armor {
	switch class {
	case "1":
		class = "Fighter"
	case "2":
		class = "Barbarian"
	case "3":
		class = "Rogue"
	case "4":
		class = "Wizard"
	case "5":
		class = "Cleric"
	case "6":
		class = "Warlock"
	case "7":
		class = "Bard"
	case "8":
		class = "Druid"
	case "9":
		class = "Ranger"
	default:

	}
	var result []Item_Armor
	for i := 0; i < len(Items.ItemsArmor); i++ {
		for _, c := range Items.ItemsArmor[i].Classes {
			if c == class && slot == Items.ItemsArmor[i].SlotType {
				result = append(result, Items.ItemsArmor[i])
			}
		}

	}
	return result
}

func WeaponsBySlotType(class string, slot string) []Item_Weapon {
	switch class {
	case "1":
		class = "Fighter"
	case "2":
		class = "Barbarian"
	case "3":
		class = "Rogue"
	case "4":
		class = "Wizard"
	case "5":
		class = "Cleric"
	case "6":
		class = "Warlock"
	case "7":
		class = "Bard"
	case "8":
		class = "Druid"
	case "9":
		class = "Ranger"
	default:

	}
	var result []Item_Weapon
	for i := 0; i < len(Items.ItemsWeapon); i++ {
		for _, c := range Items.ItemsWeapon[i].Classes {
			if c == class && slot == Items.ItemsWeapon[i].SlotType {
				result = append(result, Items.ItemsWeapon[i])
				break
			}
		}

	}
	return result
}

func AccessoryBySlotType(class string, slot string) []Item_Accessory {
	switch class {
	case "1":
		class = "Fighter"
	case "2":
		class = "Barbarian"
	case "3":
		class = "Rogue"
	case "4":
		class = "Wizard"
	case "5":
		class = "Cleric"
	case "6":
		class = "Warlock"
	case "7":
		class = "Bard"
	case "8":
		class = "Druid"
	case "9":
		class = "Ranger"
	default:

	}
	var result []Item_Accessory
	for i := 0; i < len(Items.ItemsAccessory); i++ {
		for _, c := range Items.ItemsAccessory[i].Classes {
			if c == class && slot == Items.ItemsAccessory[i].SlotType {
				result = append(result, Items.ItemsAccessory[i])
				break
			}
		}

	}
	return result
}

// FIND ITEM BY NAME
func ItemsByNameArmor(name string) Item_Armor {
	var result Item_Armor
	for i := 0; i < len(Items.ItemsArmor); i++ {
		if name == Items.ItemsArmor[i].Name {
			result = Items.ItemsArmor[i]
			return result
		}
	}
	return result
}

func ItemsByNameWeapon(name string) Item_Weapon {
	var result Item_Weapon
	for i := 0; i < len(Items.ItemsWeapon); i++ {
		if name == Items.ItemsWeapon[i].Name {
			result = Items.ItemsWeapon[i]
			return result
		}
	}
	return result
}

func ItemsByNameAccessory(name string) Item_Accessory {
	var result Item_Accessory
	for i := 0; i < len(Items.ItemsAccessory); i++ {
		if name == Items.ItemsAccessory[i].Name {
			result = Items.ItemsAccessory[i]
			return result
		}
	}
	return result
}

// FIND ENCHANTMENT BY TYPE
func EnchamentbySlot(list map[string][]float32) []string {
	keys := []string{}
	for key := range list {
		keys = append(keys, key)
	}
	return keys
}

func RangeofEnchanmentValuesInt(value []float32) []float32 {
	var valueRange []float32
	if len(value) < 2 {
		return []float32{} // Return an empty slice
	}
	for v := value[0]; v <= value[1]; v++ {
		valueRange = append(valueRange, float32(v))
	}
	return valueRange
}

// RANGE OF ENCANTMENT VALUES
func RangeofEnchanmentValues(value []float32) []float32 {
	var valueRange []float32
	if len(value) < 2 {
		return []float32{} // Return an empty slice
	}
	for v := value[0]; v <= value[1]; v += 0.1 {
		v = float32(math.Round(float64(v)*10) / 10)
		valueRange = append(valueRange, v)
	}
	return valueRange
}

// SELECT CLASS STATS AND DEFAULT IF NO CLASS SELECTED
func SelectClass(classselected string) Stats {
	allClases_array := [11]Stats{characterStats, classFighter, classBarbarian, classRogue, classWizard, classCleric, classWarlock, classBard, classDruid, classRanger, classSoccerer}
	int_converted, _ := strconv.Atoi(classselected)
	for i := 0; i <= len(allClases_array); i++ {
		if int_converted == i {
			return allClases_array[i]
		}
	}
	return characterStats
}

// LOAD ITEM FROM JSON
func readJSON(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}

// CREATE AN ITEM
func CreateItemArmor(itemfile string) Item_Armor {
	var item Item_Armor // item to create
	readJSON(itemfile, &item)
	return item
}

func CreateItemWeapon(itemfile string) Item_Weapon {
	var item Item_Weapon // item to create
	readJSON(itemfile, &item)
	return item
}

func CreateItemAccessory(itemfile string) Item_Accessory {
	var item Item_Accessory // item to create
	readJSON(itemfile, &item)
	return item
}

// CALCULATE TOTAL ATTRIBUTE STATS OF ITEMS
func SetItemStats(baseStats Stats, item []Item_Armor, rarity []int, enchantments []map[string]int) Stats {
	// Calculate total by sum base stats and item stats
	for i := 0; i < len(item) && i < len(rarity); i++ {
		baseStats = Stats{
			Strength:        baseStats.Strength + item[i].BaseAttribute.Strength[rarity[i]] + enchantments[i]["Strength"],
			Vigor:           baseStats.Vigor + item[i].BaseAttribute.Vigor[rarity[i]] + enchantments[i]["Vigor"],
			Agility:         baseStats.Agility + item[i].BaseAttribute.Agility[rarity[i]] + enchantments[i]["Agility"],
			Dexterity:       baseStats.Dexterity + item[i].BaseAttribute.Dexterity[rarity[i]] + enchantments[i]["Dexterity"],
			Will:            baseStats.Will + item[i].BaseAttribute.Will[rarity[i]] + enchantments[i]["Will"],
			Knowledge:       baseStats.Knowledge + item[i].BaseAttribute.Knowledge[rarity[i]] + enchantments[i]["Knowledge"],
			Resourcefulness: baseStats.Resourcefulness + item[i].BaseAttribute.Resourcefulness[rarity[i]] + enchantments[i]["Resourcefulness"],
		}
	}
	// Return the total stats struct
	return baseStats
}

func SetItemStatsAccessory(baseStats Stats, item []Item_Accessory, rarity []int, enchantments []map[string]int) Stats {
	// Calculate total by sum base stats and item stats
	for i := 0; i < len(item) && i < len(rarity); i++ {
		baseStats = Stats{
			Strength:        baseStats.Strength + item[i].BaseAttribute.Strength[rarity[i]] + enchantments[i]["Strength"],
			Vigor:           baseStats.Vigor + item[i].BaseAttribute.Vigor[rarity[i]] + enchantments[i]["Vigor"],
			Agility:         baseStats.Agility + item[i].BaseAttribute.Agility[rarity[i]] + enchantments[i]["Agility"],
			Dexterity:       baseStats.Dexterity + item[i].BaseAttribute.Dexterity[rarity[i]] + enchantments[i]["Dexterity"],
			Will:            baseStats.Will + item[i].BaseAttribute.Will[rarity[i]] + enchantments[i]["Will"],
			Knowledge:       baseStats.Knowledge + item[i].BaseAttribute.Knowledge[rarity[i]] + enchantments[i]["Knowledge"],
			Resourcefulness: baseStats.Resourcefulness + item[i].BaseAttribute.Resourcefulness[rarity[i]] + enchantments[i]["Resourcefulness"],
		}
	}
	// Return the total stats struct
	return baseStats
}

// CONVERT STRING INTO CHARACTER CLASS
func InttoClass(convert string) string {
	switch convert {
	case "1":
		convert = "Fighter"
	case "2":
		convert = "Barbarian"
	case "3":
		convert = "Rogue"
	case "4":
		convert = "Wizard"
	case "5":
		convert = "Cleric"
	case "6":
		convert = "Warlock"
	case "7":
		convert = "Bard"
	case "8":
		convert = "Druid"
	case "9":
		convert = "Ranger"
	case "10":
		convert = "Soccerer"
	default:
	}
	return convert
}

// CONVERT STRING INTO INT
func StringtoInt(convert string) int {
	intkey, _ := strconv.Atoi(convert)
	return intkey
}

func StringtoFloat(convert string) float64 {
	floatkey, _ := strconv.ParseFloat(convert, 64)
	return floatkey
}

// CALCULATION TOTAL AMOR RATING
func RatingCalc(ratings []int) int {
	result := 0
	for i := 0; i < len(ratings); i++ {
		result = result + ratings[i]
	}
	return result
}

// CALCULATION TOTAL MOVESPEED
func SpeedCalc(items []Item_Armor, rarity []int) int {
	speedrating := 0
	for i := 0; i < len(items); i++ {
		speedrating += items[i].MoveSpeed[1]
		if items[i].SlotType == "Foot" {
			speedrating += items[i].MoveSpeed[rarity[i]]
		}
	}
	return speedrating
}

func EnchantValuesCalc(enchantmentvalue string, enchantmenttype map[string][]float32) []float32 {

	switch enchantmentvalue {
	case "ArmorPenetration":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "PhysicalDamageBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MagicPenetration":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MagicalDamageBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "PhysicalDamageReduction":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MagicalDamageReduction":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "ProjectileReduction":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "ActionSpeed":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MoveSpeedBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "RegularInteractionSpeed":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MagicalInteractionSpeed":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MaxHealthBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "BuffDurationBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "DebuffDurationBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MemoryCapacityBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "SpellCastingSpeed":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "Luck":
		return RangeofEnchanmentValuesInt(enchantmenttype[enchantmentvalue])
	case "ArmorRating":
		return RangeofEnchanmentValuesInt(enchantmenttype[enchantmentvalue])
	case "MagicResistance":
		return RangeofEnchanmentValuesInt(enchantmenttype[enchantmentvalue])
	}

	return enchantmenttype[enchantmentvalue]
}

func Enchantattrib(enchantmenttype string, enchantmentvalue string) map[string]int {

	switch enchantmenttype {
	case "Strength", "Vigor", "Agility", "Dexterity", "Will", "Knowledge", "Resourcefulness":
		return map[string]int{
			enchantmenttype: StringtoInt(enchantmentvalue),
		}
	}
	return map[string]int{}
}

func Enchantother(enchantmenttype string, enchantmentvalue string) map[string]float64 {

	switch enchantmenttype {
	case "Strength", "Vigor", "Agility", "Dexterity", "Will", "Knowledge", "Resourcefulness":
		return map[string]float64{}
	default:
		return map[string]float64{
			enchantmenttype: StringtoFloat(enchantmentvalue),
		}
	}
}

func EnchantComputedOthers(enchant []map[string]float64) Computed_Stats {
	var result Computed_Stats      // Uses zero-value initialization
	for _, item := range enchant { // Better range loop
		for key, value := range item {
			switch key {
			case "PhysicalPower":
				result.PhysicalPower = value
			case "PhysicalPowerBonus":
				result.PhysicalPowerBonus = value
			case "MagicalPower":
				result.MagicalPower = value
			case "MagicPenetration":
				result.MagicPenetration = value
			case "ArmorPenetration":
				result.ArmorPenetration = value
			case "ActionSpeed":
				result.ActionSpeed = value
			case "RegularInteractionSpeed":
				result.RegularInteractionSpeed = value
			case "MagicalHealing":
				result.MagicalHealing = int(value)
			case "Luck":
				result.Luck = int(value)
			case "PhysicalDamageReduction":
				result.PhysicalDamageReduction = value
			case "MagicResistance":
				result.MagicRating = value
			case "MagicalDamageReduction":
				result.MagicalDamageReduction = value
			case "ProjectileReduction":
				result.ProjectileReduction = value
			case "HeadshotReduction":
				result.HeadshotReduction = value
			case "SpellRecoveryBonus":
				result.SpellRecoveryBonus = value
			case "SpellCastingSpeed":
				result.SpellCastingSpeed = value
			case "SpellRecovery":
				result.SpellRecovery = value
			case "MagicalInteractionSpeed":
				result.MagicalInteractionSpeed = value
			case "CooldownReduction":
				result.CooldownReduction = value
			case "PhysicalHealing":
				result.PhysicalHealing = int(value)
			case "ArmorRating":
				result.FromArmorRating = int(value)
			case "MaxHealthAdd":
				result.Health = value
			case "MaxHealthBonus":
				result.MaxHealthBonus = value
			case "BuffDurationBonus":
				result.BuffDuration = value
			case "DebuffDurationBonus":
				result.DebuffDuration = value
			case "MemoryCapacityAdd":
				result.MemoryCapacity = value
			case "MemoryCapacityBonus":
				result.MemoryCapacityBonus = value
			}
		}
	}
	return result
}

func ComputedTotal(computedone Computed_Stats, computedtwo Computed_Stats) Computed_Stats {
	return Computed_Stats{
		Health:                       computedone.Health + computedtwo.Health,
		ActionSpeed:                  computedone.ActionSpeed + computedtwo.ActionSpeed,
		RegularInteractionSpeed:      computedone.RegularInteractionSpeed + computedtwo.RegularInteractionSpeed,
		MoveSpeed:                    computedone.MoveSpeed + computedtwo.MoveSpeed,
		MoveSpeedCalc:                computedone.MoveSpeedCalc + computedtwo.MoveSpeedCalc,
		PhysicalPower:                computedone.PhysicalPower + computedtwo.PhysicalPower,
		PhysicalPowerBonus:           computedone.PhysicalPowerBonus,
		MagicalPower:                 computedone.MagicalPower + computedtwo.MagicalPower,
		MagicalPowerBonus:            computedone.MagicalPowerBonus,
		HealthRecovery:               computedone.HealthRecovery + computedtwo.HealthRecovery,
		ManualDexterity:              computedone.ManualDexterity + computedtwo.ManualDexterity,
		EquipSpeed:                   computedone.EquipSpeed + computedtwo.EquipSpeed,
		BuffDuration:                 computedone.BuffDuration + computedtwo.BuffDuration,
		DebuffDuration:               computedone.DebuffDuration + computedtwo.DebuffDuration,
		MagicRating:                  computedone.MagicRating,
		MagicalDamageReduction:       computedone.MagicalDamageReduction + computedtwo.MagicalDamageReduction,
		SpellRecovery:                computedone.SpellRecovery + computedtwo.SpellRecovery,
		SpellCastingSpeed:            computedone.SpellCastingSpeed + computedtwo.SpellCastingSpeed,
		MagicalInteractionSpeed:      computedone.MagicalInteractionSpeed + computedtwo.MagicalInteractionSpeed,
		Persuasiveness:               computedone.Persuasiveness + computedtwo.Persuasiveness,
		CooldownReduction:            computedone.CooldownReduction + computedtwo.CooldownReduction,
		PhysicalDamageReduction:      computedone.PhysicalDamageReduction + computedtwo.PhysicalDamageReduction,
		PhysicalHealing:              computedone.PhysicalHealing + computedtwo.PhysicalHealing,
		MagicalHealing:               computedone.MagicalHealing + computedtwo.MagicalHealing,
		Luck:                         computedone.Luck + computedtwo.Luck,
		SpellRecoveryBonus:           computedone.SpellRecoveryBonus + computedtwo.SpellRecoveryBonus,
		ArmorPenetration:             computedone.ArmorPenetration + computedtwo.ArmorPenetration,
		MagicPenetration:             computedone.MagicPenetration + computedtwo.MagicPenetration,
		HeadshotReduction:            computedone.HeadshotReduction + computedtwo.HeadshotReduction,
		ProjectileReduction:          computedone.ProjectileReduction + computedtwo.ProjectileReduction,
		FromArmorRating:              computedone.FromArmorRating + computedtwo.FromArmorRating,
		MemoryCapacity:               computedone.MemoryCapacity + computedtwo.MemoryCapacity,
		BonusPhysicalDamageReduction: computedtwo.PhysicalDamageReduction,
		BonusMagicalDamageReduction:  computedtwo.MagicalDamageReduction,
		BonusPhysicalPower:           computedtwo.PhysicalPower,
		BonusMagicalPower:            computedtwo.MagicalPower,
	}
}

func EchantBaseAttribExeption(enchantmentlist []string, itemtype Item_Armor) []string {
	for i := 0; i < len(enchantmentlist); i++ {
		if itemtype.BaseAttribute.Strength[1] == 1 && enchantmentlist[i] == "Strength" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Vigor[1] == 1 && enchantmentlist[i] == "Vigor" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Agility[1] == 1 && enchantmentlist[i] == "Agility" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Dexterity[1] == 1 && enchantmentlist[i] == "Dexterity" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Will[1] == 1 && enchantmentlist[i] == "Will" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Knowledge[1] == 1 && enchantmentlist[i] == "Knowledge" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Resourcefulness[1] == 1 && enchantmentlist[i] == "Resourcefulness" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
	}
	return enchantmentlist
}

func EnchantTypeExeption(enchantmentlist []string, enchantmenttype string) []string {
	for i := 0; i < len(enchantmentlist); i++ {
		if enchantmentlist[i] == enchantmenttype {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
	}
	return enchantmentlist
}

func EnchantValueExeption(enchantmentlist []string, enchantmenttype string) []string {
	for i := 0; i < len(enchantmentlist); i++ {
		if enchantmentlist[i] == enchantmenttype {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
	}
	return enchantmentlist
}

func (s Stats) AddStats(others ...Stats) Stats {
	result := s
	for _, other := range others {
		result.Strength += other.Strength
		result.Vigor += other.Vigor
		result.Agility += other.Agility
		result.Dexterity += other.Dexterity
		result.Will += other.Will
		result.Knowledge += other.Knowledge
		result.Resourcefulness += other.Resourcefulness
	}
	return result
}

func EchantBaseAttribExeptionAcc(enchantmentlist []string, itemtype Item_Accessory) []string {
	for i := 0; i < len(enchantmentlist); i++ {
		if itemtype.BaseAttribute.Strength[1] == 1 && enchantmentlist[i] == "Strength" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Vigor[1] == 1 && enchantmentlist[i] == "Vigor" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Agility[1] == 1 && enchantmentlist[i] == "Agility" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Dexterity[1] == 1 && enchantmentlist[i] == "Dexterity" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Will[1] == 1 && enchantmentlist[i] == "Will" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Knowledge[1] == 1 && enchantmentlist[i] == "Knowledge" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
		if itemtype.BaseAttribute.Resourcefulness[1] == 1 && enchantmentlist[i] == "Resourcefulness" {
			enchantmentlist = append(enchantmentlist[:i], enchantmentlist[i+1:]...)
			i--
		}
	}
	return enchantmentlist
}

func (cs Computed_Stats) AddEnchant(others ...Computed_Stats) Computed_Stats {
	result := cs // Start with the base `cs`
	// Iterate over the variadic inputs and add them to the result
	for _, other := range others {
		result.Health += other.Health
		result.MaxHealthBonus += other.MaxHealthBonus
		result.ActionSpeed += other.ActionSpeed
		result.RegularInteractionSpeed += other.RegularInteractionSpeed
		result.MoveSpeed += other.MoveSpeed
		result.MoveSpeedCalc += other.MoveSpeedCalc
		result.PhysicalPower += other.PhysicalPower
		result.PhysicalPowerBonus += other.PhysicalPowerBonus
		result.HealthRecovery += other.HealthRecovery
		result.ManualDexterity += other.ManualDexterity
		result.EquipSpeed += other.EquipSpeed
		result.MagicalPower += other.MagicalPower
		result.BuffDuration += other.BuffDuration
		result.MagicRating += other.MagicRating
		result.MagicalDamageReduction += other.MagicalDamageReduction
		result.DebuffDuration += other.DebuffDuration
		result.MemoryCapacity += other.MemoryCapacity
		result.MemoryCapacityBonus += other.MemoryCapacityBonus
		result.SpellRecovery += other.SpellRecovery
		result.SpellCastingSpeed += other.SpellCastingSpeed
		result.MagicalInteractionSpeed += other.MagicalInteractionSpeed
		result.Persuasiveness += other.Persuasiveness
		result.CooldownReduction += other.CooldownReduction
		result.PhysicalDamageReduction += other.PhysicalDamageReduction
		result.SpellRecoveryBonus += other.SpellRecoveryBonus
		result.PhysicalHealing += other.PhysicalHealing
		result.MagicalHealing += other.MagicalHealing
		result.MemorySpellPayload += other.MemorySpellPayload
		result.MemoryMusicPayload += other.MemoryMusicPayload
		result.UtilityEffectiveness += other.UtilityEffectiveness
		result.Luck += other.Luck
		result.ArmorPenetration += other.ArmorPenetration
		result.MagicPenetration += other.MagicPenetration
		result.HeadshotReduction += other.HeadshotReduction
		result.ProjectileReduction += other.ProjectileReduction
		result.FromArmorRating += other.FromArmorRating
		result.BonusPhysicalDamageReduction += other.BonusPhysicalDamageReduction
		result.BonusMagicalDamageReduction += other.BonusMagicalDamageReduction
	}

	return result
}

func RemoveSpaces(s string) string {
	var result []rune
	for _, char := range s {
		if char != ' ' {
			result = append(result, char)
		}
	}
	return string(result)
}

func ImageLocation(itemtype string, itemname []string) []map[string]string {
	result := []map[string]string{}

	for i := 0; i < len(itemname); i++ {
		result = append(result, map[string]string{
			"image": "assets/" + RemoveSpaces(itemname[i]) + ".png",
			"name":  itemname[i],
		})

	}
	return result
}
