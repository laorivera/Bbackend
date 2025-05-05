package main

import (
	"encoding/json"
	"math"
	"os"
	"strconv"
)

// EJECUTA CALCULO ALL STATS CURVES

// FIND ARRAY ITEMS BY SLOT TYPE
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
				break
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

func AccessoryBySlotType(slot string) []Item_Accessory {
	var result []Item_Accessory
	for i := 0; i < len(Items.ItemsAccessory); i++ {
		if slot == Items.ItemsAccessory[i].SlotType {
			result = append(result, Items.ItemsAccessory[i])
		}
	}
	return result
}

func AccessoryBySlotType_Json(slot string) []string {
	var result []string
	for i := 0; i < len(Items.ItemsAccessory); i++ {
		if slot == Items.ItemsAccessory[i].SlotType {
			result = append(result, Items.ItemsAccessory[i].Name)
		}
	}
	return result
}

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
	case "10":
		class = "Sorcerer"
	default:

	}
	var result []string
	for i := 0; i < len(Items.ItemsArmor); i++ {
		for _, c := range Items.ItemsArmor[i].Classes {
			if c == class && slot == Items.ItemsArmor[i].SlotType {
				result = append(result, Items.ItemsArmor[i].Name)
				break
			}
		}

	}
	return result
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// FIND SINGLE ITEM BY NAME
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

func ItemsByNameAccesory(name string) Item_Accessory {
	var result Item_Accessory
	for i := 0; i < len(Items.ItemsAccessory); i++ {
		if name == Items.ItemsAccessory[i].Name {
			result = Items.ItemsAccessory[i]
			return result
		}
	}
	return result
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// FIND ARRAY ENCHANTMENTS BY SLOT
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
	allClases_array := [11]Stats{characterStats, classFighter, classBarbarian, classRogue, classWizard, classCleric, classWarlock, classBard, classDruid, classRanger, classSorcerer}
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

// /////////////////////////////////////////////////////////////////////////////////////////////
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

// //////////////////////////////////////////////////////////////////////////////////////////////////
// CALCULATE TOTAL ENCHANTMENT STATS
func setEnchantStats(enchantments []map[string]int) Stats {
	var totalStats Stats
	for _, enchant := range enchantments {
		for key, value := range enchant {
			switch key {
			case "Strength":
				totalStats.Strength += value
			case "Vigor":
				totalStats.Vigor += value
			case "Agility":
				totalStats.Agility += value
			case "Dexterity":
				totalStats.Dexterity += value
			case "Will":
				totalStats.Will += value
			case "Knowledge":
				totalStats.Knowledge += value
			case "Resourcefulness":
				totalStats.Resourcefulness += value
			case "AllAttributes":
				totalStats.Strength += value
				totalStats.Vigor += value
				totalStats.Agility += value
				totalStats.Dexterity += value
				totalStats.Will += value
				totalStats.Knowledge += value
			}
		}
	}
	return totalStats
}

// CALCULATE TOTAL BASE ATTRIBUTE STATS OF ITEMS
func SetItemStats(baseStats Stats, item []Item_Armor, rarity []int /* enchantments []map[string]int */) Stats {
	// Calculate total by sum base stats and item stats
	for i := 0; i < len(item) && i < len(rarity); i++ {
		baseStats = Stats{
			Strength:        baseStats.Strength + item[i].BaseAttribute.Strength[rarity[i]],               // + enchantments[i]["Strength"],
			Vigor:           baseStats.Vigor + item[i].BaseAttribute.Vigor[rarity[i]],                     // + enchantments[i]["Vigor"],
			Agility:         baseStats.Agility + item[i].BaseAttribute.Agility[rarity[i]],                 // + enchantments[i]["Agility"],
			Dexterity:       baseStats.Dexterity + item[i].BaseAttribute.Dexterity[rarity[i]],             // + enchantments[i]["Dexterity"],
			Will:            baseStats.Will + item[i].BaseAttribute.Will[rarity[i]],                       // + enchantments[i]["Will"],
			Knowledge:       baseStats.Knowledge + item[i].BaseAttribute.Knowledge[rarity[i]],             // + enchantments[i]["Knowledge"],
			Resourcefulness: baseStats.Resourcefulness + item[i].BaseAttribute.Resourcefulness[rarity[i]], // + enchantments[i]["Resourcefulness"],
		}
	}
	// Return the total stats struct
	return baseStats
}

func SetItemStatsAccessory(baseStats Stats, item []Item_Accessory, rarity []int /*,enchantments []map[string]int*/) Stats {
	// Calculate total by sum base stats and item stats
	for i := 0; i < len(item) && i < len(rarity); i++ {
		baseStats = Stats{
			Strength:        baseStats.Strength + item[i].BaseAttribute.Strength[rarity[i]],               // + enchantments[i]["Strength"],
			Vigor:           baseStats.Vigor + item[i].BaseAttribute.Vigor[rarity[i]],                     // + enchantments[i]["Vigor"],
			Agility:         baseStats.Agility + item[i].BaseAttribute.Agility[rarity[i]],                 // + enchantments[i]["Agility"],
			Dexterity:       baseStats.Dexterity + item[i].BaseAttribute.Dexterity[rarity[i]],             // + enchantments[i]["Dexterity"],
			Will:            baseStats.Will + item[i].BaseAttribute.Will[rarity[i]],                       // + enchantments[i]["Will"],
			Knowledge:       baseStats.Knowledge + item[i].BaseAttribute.Knowledge[rarity[i]],             // + enchantments[i]["Knowledge"],
			Resourcefulness: baseStats.Resourcefulness + item[i].BaseAttribute.Resourcefulness[rarity[i]], // + enchantments[i]["Resourcefulness"],
		}
	}
	// Return the total stats struct
	return baseStats
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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

		if items[i].SlotType != "Foot" {
			speedrating += items[i].MoveSpeed[1]
		}
		if items[i].SlotType == "Foot" {
			speedrating += items[i].MoveSpeed[rarity[i]]
		}
	}
	return speedrating
}

func BaseItemCalc(items []Item_Armor, rarity []int) Computed_Stats {
	computed := Computed_Stats{}
	for i := 0; i < len(items); i++ {
		computed.Health += items[i].MaxHealthAdd[rarity[i]]
		computed.ProjectileReduction += items[i].ProjectileReduction
		computed.ProjectileReduction += items[i].ProjectileReductionRate[rarity[i]]
		computed.HeadshotReduction += items[i].HeadshotReduction
		computed.MoveSpeedBonus += items[i].MoveSpeedBonus
		computed.ArmorPenetration += items[i].ArmorPenetration
		computed.MagicPenetration += items[i].MagicPenetration
		computed.ActionSpeed += items[i].ActionSpeed[0]
		computed.MagicalDamageReduction += items[i].MagicalDamageReduction
		computed.PhysicalDamageReduction += items[i].PhysicalDamageReduction
		computed.Luck += items[i].Luck
		computed.MagicalHealing += items[i].MagicalHealing[rarity[i]]
		computed.MagicalPower += items[i].MagicalPower[rarity[i]]
		computed.PhysicalPower += items[i].PhysicalPower
		if len(items[i].MagicResistance[rarity[i]]) > 0 {
			computed.MagicRating += items[i].MagicResistance[rarity[i]][0]
		}
		if len(items[i].MaxHealthBonus[rarity[i]]) > 0 {
			computed.MaxHealthBonus += items[i].MaxHealthBonus[rarity[i]][i]
		}
	}
	return computed
}

func ComputedTotal(computedone, computedtwo, computedthree, computedfour Computed_Stats) Computed_Stats {

	stats := []Computed_Stats{computedone, computedtwo, computedthree, computedfour}

	result := Computed_Stats{}

	for _, stat := range stats {
		result.Health += stat.Health
		result.ActionSpeed += stat.ActionSpeed
		result.RegularInteractionSpeed += stat.RegularInteractionSpeed
		result.MoveSpeed += stat.MoveSpeed
		result.MoveSpeedBonus += stat.MoveSpeedBonus
		result.HealthRecovery += stat.HealthRecovery
		result.ManualDexterity += stat.ManualDexterity
		result.EquipSpeed += stat.EquipSpeed
		result.BuffDuration += stat.BuffDuration
		result.DebuffDuration += stat.DebuffDuration
		result.MagicalDamageReduction += stat.MagicalDamageReduction
		result.SpellRecovery += stat.SpellRecovery
		result.SpellCastingSpeed += stat.SpellCastingSpeed
		result.MagicalInteractionSpeed += stat.MagicalInteractionSpeed
		result.Persuasiveness += stat.Persuasiveness
		result.CooldownReduction += stat.CooldownReduction
		result.PhysicalDamageReduction += stat.PhysicalDamageReduction
		result.PhysicalHealing += stat.PhysicalHealing
		result.MagicalHealing += stat.MagicalHealing
		result.Luck += stat.Luck
		result.SpellRecoveryBonus += stat.SpellRecoveryBonus
		result.ArmorPenetration += stat.ArmorPenetration
		result.MagicPenetration += stat.MagicPenetration
		result.HeadshotReduction += stat.HeadshotReduction
		result.ProjectileReduction += stat.ProjectileReduction
		result.FromArmorRating += stat.FromArmorRating
		result.MemoryCapacity += stat.MemoryCapacity
	}

	result.PhysicalPower = computedone.PhysicalPower
	result.MagicalPower = computedone.MagicalPower
	result.PhysicalPowerBonus = computedone.PhysicalPowerBonus
	result.MagicalPowerBonus = computedone.MagicalPowerBonus
	result.MagicRating = computedone.MagicRating
	result.BonusPhysicalDamageReduction = computedtwo.PhysicalDamageReduction
	result.BonusMagicalDamageReduction = computedtwo.MagicalDamageReduction
	result.BonusPhysicalPower = computedtwo.PhysicalPowerBonus
	result.BonusMagicalPower = computedtwo.MagicalPowerBonus

	return result
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

func EnchantValuesCalc(enchantmentvalue string, enchantmenttype map[string][]float32) []float32 {

	switch enchantmentvalue {
	case "ArmorPenetration":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "PhysicalPowerBonus":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MagicPenetration":
		return RangeofEnchanmentValues(enchantmenttype[enchantmentvalue])
	case "MagicalPowerBonus":
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
	case "Strength", "Vigor", "Agility", "Dexterity", "Will", "Knowledge", "Resourcefulness", "AllAttributes":
		return map[string]int{
			enchantmenttype: StringtoInt(enchantmentvalue),
		}
	}
	return map[string]int{}
}

func Enchantother(enchantmenttype string, enchantmentvalue string) map[string]float64 {

	switch enchantmenttype {
	case "Strength", "Vigor", "Agility", "Dexterity", "Will", "Knowledge", "Resourcefulness", "AllAttributes":
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
				result.PhysicalPower += value
			case "PhysicalPowerBonus":
				result.PhysicalPowerBonus += value
			case "MagicalPower":
				result.MagicalPower += value
			case "MagicalPowerBonus":
				result.MagicalPowerBonus += value
			case "MagicPenetration":
				result.MagicPenetration += value
			case "ArmorPenetration":
				result.ArmorPenetration += value
			case "ActionSpeed":
				result.ActionSpeed += value
			case "RegularInteractionSpeed":
				result.RegularInteractionSpeed += value
			case "MagicalHealing":
				result.MagicalHealing += int(value)
			case "Luck":
				result.Luck += int(value)
			case "PhysicalDamageReduction":
				result.PhysicalDamageReduction += value
			case "MagicResistance":
				result.MagicRating += value
			case "MagicalDamageReduction":
				result.MagicalDamageReduction += value
			case "ProjectileReduction":
				result.ProjectileReduction += value
			case "HeadshotReduction":
				result.HeadshotReduction += value
			case "SpellRecoveryBonus":
				result.SpellRecoveryBonus += value
			case "SpellCastingSpeed":
				result.SpellCastingSpeed += value
			case "SpellRecovery":
				result.SpellRecovery += value
			case "MagicalInteractionSpeed":
				result.MagicalInteractionSpeed += value
			case "CooldownReduction":
				result.CooldownReduction += value
			case "PhysicalHealing":
				result.PhysicalHealing += int(value)
			case "ArmorRating":
				result.FromArmorRating += int(value)
			case "MaxHealthAdd":
				result.Health += value
			case "MaxHealthBonus":
				result.MaxHealthBonus += value
			case "BuffDurationBonus":
				result.BuffDuration += value
			case "DebuffDurationBonus":
				result.DebuffDuration += value
			case "MemoryCapacityAdd":
				result.MemoryCapacity += value
			case "MemoryCapacityBonus":
				result.MemoryCapacityBonus += value
			case "MoveSpeedBonus":
				result.MoveSpeedBonus += value
			case "MoveSpeed":
				result.MoveSpeed += value
			}
		}
	}
	return result
}

// ///////////////////////////////////////////////////////////////////////////////////////
// Remove base attributes from the enchantment list
func EnchantBaseAttribExeption(enchantmentlist []string, itemtype Item_Armor) []string {
	// Create a map of attributes to remove
	removeAttrs := make(map[string]bool)
	if itemtype.BaseAttribute.Strength[1] == 1 {
		removeAttrs["Strength"] = true
	}
	if itemtype.BaseAttribute.Vigor[1] == 1 {
		removeAttrs["Vigor"] = true
	}
	if itemtype.BaseAttribute.Agility[1] == 1 {
		removeAttrs["Agility"] = true
	}
	if itemtype.BaseAttribute.Dexterity[1] == 1 {
		removeAttrs["Dexterity"] = true
	}
	if itemtype.BaseAttribute.Will[1] == 1 {
		removeAttrs["Will"] = true
	}
	if itemtype.BaseAttribute.Knowledge[1] == 1 {
		removeAttrs["Knowledge"] = true
	}
	if itemtype.BaseAttribute.Resourcefulness[1] == 1 {
		removeAttrs["Resourcefulness"] = true
	}

	// Build a new list excluding base attributes
	result := make([]string, 0, len(enchantmentlist))
	for _, enchant := range enchantmentlist {
		if !removeAttrs[enchant] {
			result = append(result, enchant)
		}
	}
	return result
}

// Remove a specific enchantment type from the list
func EnchantTypeExeption(enchantmentlist []string, previousSelections []string) []string {
	exclude := make(map[string]bool)
	for _, selection := range previousSelections {
		if selection != "" {
			exclude[selection] = true
		}
	}

	result := make([]string, 0, len(enchantmentlist))
	for _, enchant := range enchantmentlist {
		if !exclude[enchant] {
			result = append(result, enchant)
		}
	}
	return result
}

// For accessories, same logic
func EnchantBaseAttribExeptionAcc(enchantmentlist []string, itemtype Item_Accessory) []string {
	removeAttrs := make(map[string]bool)
	if itemtype.BaseAttribute.Strength[1] == 1 {
		removeAttrs["Strength"] = true
	}
	if itemtype.BaseAttribute.Vigor[1] == 1 {
		removeAttrs["Vigor"] = true
	}
	if itemtype.BaseAttribute.Agility[1] == 1 {
		removeAttrs["Agility"] = true
	}
	if itemtype.BaseAttribute.Dexterity[1] == 1 {
		removeAttrs["Dexterity"] = true
	}
	if itemtype.BaseAttribute.Will[1] == 1 {
		removeAttrs["Will"] = true
	}
	if itemtype.BaseAttribute.Knowledge[1] == 1 {
		removeAttrs["Knowledge"] = true
	}
	if itemtype.BaseAttribute.Resourcefulness[1] == 1 {
		removeAttrs["Resourcefulness"] = true
	}

	result := make([]string, 0, len(enchantmentlist))
	for _, enchant := range enchantmentlist {
		if !removeAttrs[enchant] {
			result = append(result, enchant)
		}
	}
	return result
}

///////////////////////////////////////////--- METHODS ----///////////////////////////////////////////////////////

// add stats method
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

// add enchant method
func (cs Computed_Stats) AddEnchant(others ...Computed_Stats) Computed_Stats {
	result := cs
	// Iterate over the variadic inputs and add them to the result
	for _, other := range others {
		result.Health += other.Health
		result.MaxHealthBonus += other.MaxHealthBonus
		result.ActionSpeed += other.ActionSpeed
		result.RegularInteractionSpeed += other.RegularInteractionSpeed
		result.MoveSpeed += other.MoveSpeed
		result.MoveSpeedBonus += other.MoveSpeedBonus
		result.PhysicalPower += other.PhysicalPower
		result.PhysicalPowerBonus += other.PhysicalPowerBonus
		result.HealthRecovery += other.HealthRecovery
		result.ManualDexterity += other.ManualDexterity
		result.EquipSpeed += other.EquipSpeed
		result.MagicalPower += other.MagicalPower
		result.MagicalPowerBonus += other.MagicalPowerBonus

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
		//result.BonusPhysicalPower += other.BonusPhysicalPower
		//result.BonusMagicalPower += other.BonusMagicalPower

	}

	return result
}

///////////////////////////////////////////////---FROTEND LOGIC--- ////////////////////////////////////////////////////////

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
