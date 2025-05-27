package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

// get item lists

func GetItemLists_Armor_Json(c *gin.Context, class string) map[string][]string {
	slots := []string{"Head", "Chest", "Hands", "Legs", "Foot", "Back"}
	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		lists[slots[i]] = ItemsBySlotType_Json(class, slots[i])
	}
	return lists
}

func GetItemLists_Accesory_Json(c *gin.Context) map[string][]string {
	slots := []string{"Necklace", "Ring"}
	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		lists[slots[i]] = AccessoryBySlotType_Json(slots[i])
	}
	return lists
}

// get rating lists
func GetRatingLists_Armor(c *gin.Context) map[string][]int {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	ratings := map[string][]int{}
	for i := 0; i < len(slots); i++ {
		ratings[slots[i]] = CompleteArrayInt(ItemsByNameArmor(c.Query("item" + slots[i])).ArmorRatings[StringtoInt(c.Query("rarityselect_"+slots[i]))])
	}
	return ratings
}

func GetItemLists_Weapon_Json(c *gin.Context, class string) map[string][]string {
	slots := []string{"Main Hand", "Off Hand"}
	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		lists[slots[i]] = WeaponsBySlotType_Json(class, slots[i])
	}
	return lists
}

func GetRatingLists_Weapon(c *gin.Context) map[string][]int {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	ratings := map[string][]int{}
	for i := 0; i < len(slots); i++ {
		ratings[slots[i]] = CompleteArrayInt(ItemsByNameWeapon(c.Query("item" + slots[i])).DamageRatings[StringtoInt(c.Query("rarityselect_"+slots[i]))])
	}
	return ratings
}

// /////\\\\\\ --------------> ENCHCANTMENT LISTS ARMOR <------------------///////////\\\\\\\\\
func GetEnchatmentLists_Armor_ExceptionBase(c *gin.Context) map[string][]string {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantBaseAttribExeption(EnchamentbySlot(Enchantments.Helmet), ItemsByNameArmor(c.Query("item"+slots[i])))
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantBaseAttribExeption(EnchamentbySlot(Enchantments.Chest), ItemsByNameArmor(c.Query("item"+slots[i])))
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantBaseAttribExeption(EnchamentbySlot(Enchantments.Gloves), ItemsByNameArmor(c.Query("item"+slots[i])))
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantBaseAttribExeption(EnchamentbySlot(Enchantments.Pants), ItemsByNameArmor(c.Query("item"+slots[i])))
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantBaseAttribExeption(EnchamentbySlot(Enchantments.Boots), ItemsByNameArmor(c.Query("item"+slots[i])))
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantBaseAttribExeption(EnchamentbySlot(Enchantments.Cloak), ItemsByNameArmor(c.Query("item"+slots[i])))
		}

	}
	return lists

}

func GetEnchatmentLists_Armor_ValuesUncommon(c *gin.Context) map[string][]float32 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := map[string][]float32{}
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type")
		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Cloak)
		}
	}
	return lists
}

func GetEnchatmentLists_Armor_TypeRare(c *gin.Context) map[string][]string {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"), // Uncommon
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}

func GetEnchatmentLists_Armor_ValuesRare(c *gin.Context) map[string][]float32 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type2")
		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Cloak)
		}
	}
	return lists
}

func GetEnchatmentLists_Armor_TypeEpic(c *gin.Context) map[string][]string {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}

func GetEnchatmentLists_Armor_ValuesEpic(c *gin.Context) map[string][]float32 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type3")
		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Cloak)
		}
	}
	return lists
}

func GetEnchatmentLists_Armor_TypeLegend(c *gin.Context) map[string][]string {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}

func GetEnchatmentLists_Armor_ValuesLegend(c *gin.Context) map[string][]float32 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type4")
		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Cloak)
		}
	}
	return lists
}

func GetEnchatmentLists_Armor_TypeUnique(c *gin.Context) map[string][]string {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Armor_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
			c.Query("enchantment_" + slot + "type4"), // Legend
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}

func GetEnchatmentLists_Armor_ValuesUnique(c *gin.Context) map[string][]float32 {
	slots := []string{"helmet", "chest", "gloves", "pants", "boots", "cloak"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type5")
		if slots[i] == "helmet" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Helmet)
		}
		if slots[i] == "chest" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Chest)
		}
		if slots[i] == "gloves" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Gloves)
		}
		if slots[i] == "pants" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Pants)
		}
		if slots[i] == "boots" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Boots)
		}
		if slots[i] == "cloak" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Cloak)
		}
	}
	return lists
}

/////////\\\\\\\\\ --------------> ENCHCANTMENT LISTS ACCESSORY <------------------ //////////\\\\\\\\\

func GetEnchatmentLists_Accessory_ExceptionBase(c *gin.Context) map[string][]string {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		if slots[i] == "necklace" {
			lists[slots[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Necklace), ItemsByNameAccesory(c.Query("item"+slots[i])))
		}
		if slots[i] == "ring" {
			lists[slots[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Ring), ItemsByNameAccesory(c.Query("item"+slots[i])))
		}
		if slots[i] == "ringtwo" {
			lists[slots[i]] = EnchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Ring), ItemsByNameAccesory(c.Query("item"+slots[i])))
		}
	}
	return lists

}

func GetEnchantmentLists_Accessory_ValuesUncommon(c *gin.Context) map[string][]float32 {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := map[string][]float32{}
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type")
		if slots[i] == "necklace" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Necklace)
		}
		if slots[i] == "ring" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
		if slots[i] == "ringtwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
	}
	return lists
}

func GetEnchantmentLists_Accessory_TypeRare(c *gin.Context) map[string][]string {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"), // Uncommon
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}
	return lists
}

func GetEnchantmentLists_Accessory_ValuesRare(c *gin.Context) map[string][]float32 {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type2")
		if slots[i] == "necklace" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Necklace)
		}
		if slots[i] == "ring" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
		if slots[i] == "ringtwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeEpic(c *gin.Context) map[string][]string {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchantmentLists_Accessory_ValuesEpic(c *gin.Context) map[string][]float32 {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type3")
		if slots[i] == "necklace" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Necklace)
		}
		if slots[i] == "ring" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
		if slots[i] == "ringtwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeLegend(c *gin.Context) map[string][]string {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchantmentLists_Accessory_ValuesLegend(c *gin.Context) map[string][]float32 {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type4")
		if slots[i] == "necklace" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Necklace)
		}
		if slots[i] == "ring" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
		if slots[i] == "ringtwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
	}
	return lists
}
func GetEnchantmentLists_Accessory_TypeUnique(c *gin.Context) map[string][]string {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]string)

	baseLists := GetEnchatmentLists_Accessory_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
			c.Query("enchantment_" + slot + "type4"), // Legend
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists
}
func GetEnchantmentLists_Accessory_ValuesUnique(c *gin.Context) map[string][]float32 {
	slots := []string{"necklace", "ring", "ringtwo"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type5")
		if slots[i] == "necklace" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Necklace)
		}
		if slots[i] == "ring" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
		if slots[i] == "ringtwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.Ring)
		}
	}
	return lists
}

/////////\\\\\\\\\ --------------> ENCHCANTMENT LISTS WEAPON <------------------ //////////\\\\\\\\\

func GetEnchantmentLists_Weapon_ExceptionBase(c *gin.Context) map[string][]string {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := map[string][]string{}
	for i := 0; i < len(slots); i++ {
		if slots[i] == "pwo" {
			if ItemsByNameWeapon(c.Query("item"+slots[i])).SlotType == "Main Hand" {
				lists[slots[i]] = EnchamentbySlot(Enchantments.WeaponOneHand)
			}
		}
		/*
			if slots[i] == "pwt" {
				lists[slots[i]] = EnchamentbySlot(Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchamentbySlot(Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchamentbySlot(Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists

}
func GetEnchantmentLists_Weapon_ValuesUncommon(c *gin.Context) map[string][]float32 {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := map[string][]float32{}
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type")
		if slots[i] == "pwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeRare(c *gin.Context) map[string][]string {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"), // Uncommon
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesRare(c *gin.Context) map[string][]float32 {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type2")
		if slots[i] == "pwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeEpic(c *gin.Context) map[string][]string {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesEpic(c *gin.Context) map[string][]float32 {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type3")
		if slots[i] == "pwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeLegend(c *gin.Context) map[string][]string {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesLegend(c *gin.Context) map[string][]float32 {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type4")
		if slots[i] == "pwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
func GetEnchantmentLists_Weapon_TypeUnique(c *gin.Context) map[string][]string {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]string)

	baseLists := GetEnchantmentLists_Weapon_ExceptionBase(c)

	for _, slot := range slots {
		// Pass all previous selections up to Epic
		previousSelections := []string{
			c.Query("enchantment_" + slot + "type"),  // Uncommon
			c.Query("enchantment_" + slot + "type2"), // Rare
			c.Query("enchantment_" + slot + "type3"), // Epic
			c.Query("enchantment_" + slot + "type4"), // Legend
		}
		lists[slot] = EnchantTypeExeption(baseLists[slot], previousSelections)
	}

	return lists

}
func GetEnchantmentLists_Weapon_ValuesUnique(c *gin.Context) map[string][]float32 {
	slots := []string{"pwo", "pwt", "swo", "swt"}
	lists := make(map[string][]float32)
	for i := 0; i < len(slots); i++ {
		Query := c.Query("enchantment_" + slots[i] + "type5")
		if slots[i] == "pwo" {
			lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.WeaponOneHand)
		}
		/*
			if slots[i] == "pwt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.PrimaryWeaponTwoHand)
			}
			if slots[i] == "swo" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponOneHand)
			}
			if slots[i] == "swt" {
				lists[slots[i]] = EnchantValuesCalc(Query, Enchantments.SecondaryWeaponTwoHand)
			}
		*/
	}
	return lists
}
