package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

// get selected armors
func GetSelectedItems_Armor(c *gin.Context) []Item_Armor {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	values := []Item_Armor{}
	for i := 0; i < len(slots); i++ {
		values = append(values, ItemsByNameArmor(c.Query("item"+slots[i])))
	}
	return values
}

// get selected rarities
func GetSelectedRarities_Armor(c *gin.Context) []int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	values := []int{}
	for i := 0; i < len(slots); i++ {
		values = append(values, StringtoInt(c.Query("rarityselect_"+slots[i])))
	}
	return values
}

// get selected ratings
func GetSelectedRatings_Armor(c *gin.Context) []int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	values := []int{}
	for i := 0; i < len(slots); i++ {
		values = append(values, StringtoInt(c.Query("armorrating_"+slots[i])))
	}
	return values
}

func GetSelectedItems_Accessory(c *gin.Context) []Item_Accessory {
	slots := []string{"necklace", "ring", "ringtwo"}
	values := []Item_Accessory{}
	for i := 0; i < len(slots); i++ {
		values = append(values, ItemsByNameAccesory(c.Query("item"+slots[i])))
	}
	return values
}

func GetSelectedRarities_Accessory(c *gin.Context) []int {
	slots := []string{"necklace", "ring", "ringtwo"}
	values := []int{}
	for i := 0; i < len(slots); i++ {
		values = append(values, StringtoInt(c.Query("rarityselect_"+slots[i])))
	}
	return values
}

///////\\\\\\ --------------> ENCHCANTMENT SELECTION ARMOR <------------------///////////\\\\\\\\\

func GetSelectedEnchantmentsOther_ArmorUncommon(c *gin.Context) []map[string]float64 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type"), c.Query("enchantment_"+slots[i]+"value")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_ArmorRare(c *gin.Context) []map[string]float64 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type2"), c.Query("enchantment_"+slots[i]+"value2")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_ArmorEpic(c *gin.Context) []map[string]float64 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type3"), c.Query("enchantment_"+slots[i]+"value3")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_ArmorLegend(c *gin.Context) []map[string]float64 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type4"), c.Query("enchantment_"+slots[i]+"value4")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_ArmorUnique(c *gin.Context) []map[string]float64 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type5"), c.Query("enchantment_"+slots[i]+"value5")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_ArmorUncommon(c *gin.Context) []map[string]int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type"), c.Query("enchantment_"+slots[i]+"value")))
	}

	return enchantments
}

func GetSelectedEnchantmentsBase_ArmorRare(c *gin.Context) []map[string]int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type2"), c.Query("enchantment_"+slots[i]+"value2")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_ArmorEpic(c *gin.Context) []map[string]int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type3"), c.Query("enchantment_"+slots[i]+"value3")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_ArmorLegend(c *gin.Context) []map[string]int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type4"), c.Query("enchantment_"+slots[i]+"value4")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_ArmorUnique(c *gin.Context) []map[string]int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type5"), c.Query("enchantment_"+slots[i]+"value5")))
	}
	return enchantments
}

///////\\\\\\ --------------> ENCHCANTMENT SELECTION ACCESSORY  <------------------///////////\\\\\\\\\

func GetSelectedEnchantmentsBase_AccessoryUncommon(c *gin.Context) []map[string]int {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type"), c.Query("enchantment_"+slots[i]+"value")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_AccessoryRare(c *gin.Context) []map[string]int {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type2"), c.Query("enchantment_"+slots[i]+"value2")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_AccessoryEpic(c *gin.Context) []map[string]int {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type3"), c.Query("enchantment_"+slots[i]+"value3")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_AccessoryLegend(c *gin.Context) []map[string]int {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type4"), c.Query("enchantment_"+slots[i]+"value4")))
	}
	return enchantments
}

func GetSelectedEnchantmentsBase_AccessoryUnique(c *gin.Context) []map[string]int {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]int{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantattrib(c.Query("enchantment_"+slots[i]+"type5"), c.Query("enchantment_"+slots[i]+"value5")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_AccessoryUncommon(c *gin.Context) []map[string]float64 {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type"), c.Query("enchantment_"+slots[i]+"value")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_AccessoryRare(c *gin.Context) []map[string]float64 {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type2"), c.Query("enchantment_"+slots[i]+"value2")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_AccessoryEpic(c *gin.Context) []map[string]float64 {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type3"), c.Query("enchantment_"+slots[i]+"value3")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_AccessoryLegend(c *gin.Context) []map[string]float64 {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type4"), c.Query("enchantment_"+slots[i]+"value4")))
	}
	return enchantments
}

func GetSelectedEnchantmentsOther_AccessoryUnique(c *gin.Context) []map[string]float64 {
	slots := []string{"necklace", "ring", "ringtwo"}
	enchantments := []map[string]float64{}
	for i := 0; i < len(slots); i++ {
		enchantments = append(enchantments, Enchantother(c.Query("enchantment_"+slots[i]+"type5"), c.Query("enchantment_"+slots[i]+"value5")))
	}
	return enchantments
}

///////\\\\\\ --------------> Race Selection  <------------------///////////\\\\\\\\\

func GetSelectedRace(c *gin.Context) string {
	var raceSelected = c.Query("race")
	return raceSelected
}
