package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateStatsHandler(c *gin.Context) {

	//class query
	class := c.Param("classSelection")
	classStatSelect := SelectClass(class)
	//printClass := InttoClass(c.Param("classSelection"))

	itemsSelected_Armor := GetSelectedItems_Armor(c) //Items query selection Armor
	raritySelected_Armor := GetSelectedRarities_Armor(c)
	ratingSelected_Armor := GetSelectedRatings_Armor(c)

	itemsSelected_Accessory := GetSelectedItems_Accessory(c) //Items query selection Accessory
	raritySelected_Accessory := GetSelectedRarities_Accessory(c)

	//Enchatment query selected Armor
	enchantmentSelected_ArmorUncommon := GetSelectedEnchantmentsBase_ArmorUncommon(c)
	enchantmentSelected_ArmorRare := GetSelectedEnchantmentsBase_ArmorRare(c)
	enchantmentSelected_ArmorEpic := GetSelectedEnchantmentsBase_ArmorEpic(c)
	enchantmentSelected_ArmorLegend := GetSelectedEnchantmentsBase_ArmorLegend(c)
	enchantmentSelected_ArmorUnique := GetSelectedEnchantmentsBase_ArmorUnique(c)

	enchantmentSelectedOther_ArmorUncommon := GetSelectedEnchantmentsOther_ArmorUncommon(c)
	enchantmentSelectedOther_ArmorRare := GetSelectedEnchantmentsOther_ArmorRare(c)
	enchantmentSelectedOther_ArmorEpic := GetSelectedEnchantmentsOther_ArmorEpic(c)
	enchantmentSelectedOther_ArmorLegend := GetSelectedEnchantmentsOther_ArmorLegend(c)
	enchantmentSelectedOther_ArmorUnique := GetSelectedEnchantmentsOther_ArmorUnique(c)

	// Enchatment query selected Accesory
	enchantmentSelected_AccessoryUncommon := GetSelectedEnchantmentsBase_AccessoryUncommon(c)
	enchantmentSelected_AccessoryRare := GetSelectedEnchantmentsBase_AccessoryRare(c)
	enchantmentSelected_AccessoryEpic := GetSelectedEnchantmentsBase_AccessoryEpic(c)
	enchantmentSelected_AccessoryLegend := GetSelectedEnchantmentsBase_AccessoryLegend(c)
	enchantmentSelected_AccessoryUnique := GetSelectedEnchantmentsBase_AccessoryUnique(c)

	enchantmentSelectedOther_AccessoryUncommon := GetSelectedEnchantmentsOther_AccessoryUncommon(c)
	enchantmentSelectedOther_AccessoryRare := GetSelectedEnchantmentsOther_AccessoryRare(c)
	enchantmentSelectedOther_AccessoryEpic := GetSelectedEnchantmentsOther_AccessoryEpic(c)
	enchantmentSelectedOther_AccessoryLegend := GetSelectedEnchantmentsOther_AccessoryLegend(c)
	enchantmentSelectedOther_AccessoryUnique := GetSelectedEnchantmentsOther_AccessoryUnique(c)

	//list weapon query by class
	//primaryWeaponList := WeaponsBySlotType(c.Param("classSelection"), "Main Hand")

	/*
		itemsSelected_Weapon := []Item_Weapon{
				ItemsByNameWeapon(c.Query("itemweapon")),
			}

		raritySelected_Weapon := []int{
				StringtoInt(c.Query("rarityselect_weapon")),
			}

		ratingSelected_Weapon := []int{
				StringtoInt(c.Query("weaponrating")),
			}
	*/

	//ratingListPrimaryWeapon := ItemsByNameWeapon(c.Query("itemprimaryweapon")).DamageRatings[StringtoInt(c.Query("rarityselect_weapon"))]

	computedStatsEnchant_Other_Armor := Computed_Stats{}
	computedStatsEnchant_Other_Armor = computedStatsEnchant_Other_Armor.AddEnchant(EnchantComputedOthers(enchantmentSelectedOther_ArmorUncommon),
		EnchantComputedOthers(enchantmentSelectedOther_ArmorRare), EnchantComputedOthers(enchantmentSelectedOther_ArmorEpic),
		EnchantComputedOthers(enchantmentSelectedOther_ArmorLegend), EnchantComputedOthers(enchantmentSelectedOther_ArmorUnique))

	computedStatsEnchant_Other_Accesory := Computed_Stats{}
	computedStatsEnchant_Other_Accesory = computedStatsEnchant_Other_Accesory.AddEnchant(EnchantComputedOthers(enchantmentSelectedOther_AccessoryUncommon),
		EnchantComputedOthers(enchantmentSelectedOther_AccessoryRare), EnchantComputedOthers(enchantmentSelectedOther_AccessoryEpic),
		EnchantComputedOthers(enchantmentSelectedOther_AccessoryLegend), EnchantComputedOthers(enchantmentSelectedOther_AccessoryUnique))

	totalRating := RatingCalc(ratingSelected_Armor)

	totalBaseItem := BaseItemCalc(itemsSelected_Armor, raritySelected_Armor)

	totalSpeed := SpeedCalc(itemsSelected_Armor, raritySelected_Armor)

	updatedBaseEchant_StatsArmor := SetItemStats(classStatSelect, itemsSelected_Armor, raritySelected_Armor)
	updatedBaseEchant_StatsArmor = updatedBaseEchant_StatsArmor.AddStats(setEnchantStats(enchantmentSelected_ArmorUncommon), setEnchantStats(enchantmentSelected_ArmorRare),
		setEnchantStats(enchantmentSelected_ArmorEpic), setEnchantStats(enchantmentSelected_ArmorLegend), setEnchantStats(enchantmentSelected_ArmorUnique))

	updatedBaseEnchant_StatsAccessory := SetItemStatsAccessory(characterHolder, itemsSelected_Accessory, raritySelected_Accessory)
	updatedBaseEnchant_StatsAccessory = updatedBaseEnchant_StatsAccessory.AddStats(setEnchantStats(enchantmentSelected_AccessoryUncommon), setEnchantStats(enchantmentSelected_AccessoryRare),
		setEnchantStats(enchantmentSelected_AccessoryEpic), setEnchantStats(enchantmentSelected_AccessoryLegend), setEnchantStats(enchantmentSelected_AccessoryUnique))

	updatedTotalStats := Stats{}
	updatedTotalStats = updatedTotalStats.AddStats(updatedBaseEchant_StatsArmor, updatedBaseEnchant_StatsAccessory)

	computedStatsCurve := CalculateComputedValues(updatedTotalStats, totalRating, totalSpeed, computedStatsEnchant_Other_Armor, computedStatsEnchant_Other_Accesory, totalBaseItem)

	computedStatsTotal := ComputedTotal(computedStatsCurve, computedStatsEnchant_Other_Armor, computedStatsEnchant_Other_Accesory, totalBaseItem)

	/*
		c.HTML(http.StatusOK, "charbuilder.html", gin.H{

			"classpick":                class,
			"raritylistpick":           rarityList,
			"stats":                    updatedTotalStats,
			"computedstats":            computedStatsTotal,
			"totalrating":              totalRating,
			"helmetlist":               helmetList,
			"chestlist":                chestList,
			"gloveslist":               glovesList,
			"pantslist":                pantsList,
			"bootslist":                bootsList,
			"cloaklist":                cloakList,
			"necklacelist":             necklaceList,
			"primaryweaponlist":        primaryWeaponList,
			"printclass":               printClass,
			"helmet_ratinglist":        ratingListHelmet,
			"chest_ratinglist":         ratingListChest,
			"gloves_ratinglist":        ratingListGloves,
			"pants_ratinglist":         ratingListPants,
			"boots_ratinglist":         ratingListBoots,
			"cloak_ratinglist":         ratingListCloak,
			"primaryweapon_ratinglist": ratingListPrimaryWeapon,
			"helmet_enchantmenttype":   enchamentListHead,
			"helmet_enchantmentvalue":  enchamentListHeadValue,
			"necklace_enchantmenttype": enchantmentListNecklace,
			//"helmet_enchantmenttypetwo": enchantmentListHeadTwo,
			"test": helmetList,
		})
	*/

	c.JSON(http.StatusOK, gin.H{
		"stats":         updatedTotalStats,
		"computedstats": computedStatsTotal,
	})

}

//////////\\\\\\\\\ ------->  LISTS ITEMS HANDLER <------- //////////\\\\\\\\\

func itemDisplayHandler(c *gin.Context) {
	selecteditem := c.Param("item")

	item := ItemsByNameArmor(selecteditem)
	itemAccesory := ItemsByNameAccesory(selecteditem)

	c.JSON(http.StatusOK, gin.H{
		"itemdata":    item,
		"itemdataacc": itemAccesory,
	})
}

func Helmet_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	helmetList := GetItemLists_Armor_Json(c, class)["Head"]
	imageHelmet := ImageLocation("Head", helmetList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageHelmet},
	)
}

func Chest_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	chestList := GetItemLists_Armor_Json(c, class)["Chest"]

	imageChest := ImageLocation("Chest", chestList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageChest},
	)
}

func Gloves_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	glovesList := GetItemLists_Armor_Json(c, class)["Hands"]
	imageGloves := ImageLocation("Hands", glovesList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageGloves},
	)
}

func Pants_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	pantsList := GetItemLists_Armor_Json(c, class)["Legs"]
	imagePants := ImageLocation("Legs", pantsList)

	c.JSON(http.StatusOK, gin.H{
		"list": imagePants,
	})

}

func Boots_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	bootsList := GetItemLists_Armor_Json(c, class)["Foot"]
	imageBoots := ImageLocation("Foot", bootsList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageBoots},
	)
}

func Cloak_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	cloakList := GetItemLists_Armor_Json(c, class)["Back"]
	imageCloak := ImageLocation("Back", cloakList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageCloak},
	)
}

func Necklace_List_Handler(c *gin.Context) {
	necklaceList := GetItemLists_Accesory_Json(c)["Necklace"]
	imageNecklace := ImageLocation("Necklace", necklaceList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageNecklace},
	)
}

func Ring_List_Handler(c *gin.Context) {
	ringList := GetItemLists_Accesory_Json(c)["Ring"]
	imageRing := ImageLocation("Ring", ringList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageRing},
	)
}

//////////\\\\\\\\\ ------->   RATING LISTS HANDLER <------- //////////\\\\\\\\\

func Helmet_RatingList_Handler(c *gin.Context) {

	helmetRatingList := GetRatingLists_Armor(c)["helmet"]

	c.JSON(http.StatusOK, gin.H{
		"list": helmetRatingList},
	)
}

func Chest_RatingList_Handler(c *gin.Context) {
	chestRatingList := GetRatingLists_Armor(c)["chest"]
	c.JSON(http.StatusOK, gin.H{
		"list": chestRatingList},
	)
}

func Gloves_RatingList_Handler(c *gin.Context) {
	glovesRatingList := GetRatingLists_Armor(c)["gloves"]
	c.JSON(http.StatusOK, gin.H{
		"list": glovesRatingList},
	)
}

func Pants_RatingList_Handler(c *gin.Context) {
	pantsRatingList := GetRatingLists_Armor(c)["pants"]
	c.JSON(http.StatusOK, gin.H{
		"list": pantsRatingList},
	)
}

func Boots_RatingList_Handler(c *gin.Context) {
	bootsRatingList := GetRatingLists_Armor(c)["boots"]
	c.JSON(http.StatusOK, gin.H{
		"list": bootsRatingList},
	)
}

func Cloak_RatingList_Handler(c *gin.Context) {
	cloakRatingList := GetRatingLists_Armor(c)["cloak"]
	c.JSON(http.StatusOK, gin.H{
		"list": cloakRatingList},
	)
}

//////////\\\\\\\\\ ------->   ENCHANTMENT LISTS HANDLER   <------- //////////\\\\\\\\\

func Helmet_EnchantmentList_Handler(c *gin.Context) {
	helmet_EnchantmentName := map[string][]string{}
	helmet_EnchantmentValues := map[string][]float32{}

	helmet_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_ExceptionBase(c)["helmet"]
	helmet_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["helmet"]
	helmet_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["helmet"]
	helmet_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["helmet"]
	helmet_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["helmet"]

	helmet_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["helmet"]
	helmet_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["helmet"]
	helmet_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["helmet"]
	helmet_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["helmet"]
	helmet_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["helmet"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  helmet_EnchantmentName["Uncommon"],
		"listvalue_uncommon": helmet_EnchantmentValues["Uncommon"],
		"listname_rare":      helmet_EnchantmentName["Rare"],
		"listvalue_rare":     helmet_EnchantmentValues["Rare"],
		"listname_epic":      helmet_EnchantmentName["Epic"],
		"listvalue_epic":     helmet_EnchantmentValues["Epic"],
		"listname_legend":    helmet_EnchantmentName["Legend"],
		"listvalue_legend":   helmet_EnchantmentValues["Legend"],
		"listname_unique":    helmet_EnchantmentName["Unique"],
		"listvalue_unique":   helmet_EnchantmentValues["Unique"],
	})

}

func Chest_EnchantmentList_Handler(c *gin.Context) {
	chest_EnchantmentName := map[string][]string{}
	chest_EnchantmentValues := map[string][]float32{}

	chest_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_ExceptionBase(c)["chest"]
	chest_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["chest"]
	chest_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["chest"]
	chest_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["chest"]
	chest_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["chest"]

	chest_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["chest"]
	chest_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["chest"]
	chest_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["chest"]
	chest_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["chest"]
	chest_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["chest"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  chest_EnchantmentName["Uncommon"],
		"listvalue_uncommon": chest_EnchantmentValues["Uncommon"],
		"listname_rare":      chest_EnchantmentName["Rare"],
		"listvalue_rare":     chest_EnchantmentValues["Rare"],
		"listname_epic":      chest_EnchantmentName["Epic"],
		"listvalue_epic":     chest_EnchantmentValues["Epic"],
		"listname_legend":    chest_EnchantmentName["Legend"],
		"listvalue_legend":   chest_EnchantmentValues["Legend"],
		"listname_unique":    chest_EnchantmentName["Unique"],
		"listvalue_unique":   chest_EnchantmentValues["Unique"],
	})

}

func Gloves_EnchantmentList_Handler(c *gin.Context) {
	gloves_EnchantmentName := map[string][]string{}
	gloves_EnchantmentValues := map[string][]float32{}

	gloves_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_ExceptionBase(c)["gloves"]
	gloves_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["gloves"]
	gloves_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["gloves"]
	gloves_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["gloves"]
	gloves_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["gloves"]

	gloves_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["gloves"]
	gloves_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["gloves"]
	gloves_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["gloves"]
	gloves_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["gloves"]
	gloves_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["gloves"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  gloves_EnchantmentName["Uncommon"],
		"listvalue_uncommon": gloves_EnchantmentValues["Uncommon"],
		"listname_rare":      gloves_EnchantmentName["Rare"],
		"listvalue_rare":     gloves_EnchantmentValues["Rare"],
		"listname_epic":      gloves_EnchantmentName["Epic"],
		"listvalue_epic":     gloves_EnchantmentValues["Epic"],
		"listname_legend":    gloves_EnchantmentName["Legend"],
		"listvalue_legend":   gloves_EnchantmentValues["Legend"],
		"listname_unique":    gloves_EnchantmentName["Unique"],
		"listvalue_unique":   gloves_EnchantmentValues["Unique"],
	})

}

func Pants_EnchantmentList_Handler(c *gin.Context) {
	pants_EnchantmentName := map[string][]string{}
	pants_EnchantmentValues := map[string][]float32{}

	pants_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_ExceptionBase(c)["pants"]
	pants_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["pants"]
	pants_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["pants"]
	pants_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["pants"]
	pants_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["pants"]

	pants_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["pants"]
	pants_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["pants"]
	pants_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["pants"]
	pants_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["pants"]
	pants_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["pants"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  pants_EnchantmentName["Uncommon"],
		"listvalue_uncommon": pants_EnchantmentValues["Uncommon"],
		"listname_rare":      pants_EnchantmentName["Rare"],
		"listvalue_rare":     pants_EnchantmentValues["Rare"],
		"listname_epic":      pants_EnchantmentName["Epic"],
		"listvalue_epic":     pants_EnchantmentValues["Epic"],
		"listname_legend":    pants_EnchantmentName["Legend"],
		"listvalue_legend":   pants_EnchantmentValues["Legend"],
		"listname_unique":    pants_EnchantmentName["Unique"],
		"listvalue_unique":   pants_EnchantmentValues["Unique"],
	})

}

func Boots_EnchantmentList_Handler(c *gin.Context) {
	boots_EnchantmentName := map[string][]string{}
	boots_EnchantmentValues := map[string][]float32{}

	boots_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_ExceptionBase(c)["boots"]
	boots_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["boots"]
	boots_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["boots"]
	boots_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["boots"]
	boots_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["boots"]

	boots_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["boots"]
	boots_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["boots"]
	boots_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["boots"]
	boots_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["boots"]
	boots_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["boots"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  boots_EnchantmentName["Uncommon"],
		"listvalue_uncommon": boots_EnchantmentValues["Uncommon"],
		"listname_rare":      boots_EnchantmentName["Rare"],
		"listvalue_rare":     boots_EnchantmentValues["Rare"],
		"listname_epic":      boots_EnchantmentName["Epic"],
		"listvalue_epic":     boots_EnchantmentValues["Epic"],
		"listname_legend":    boots_EnchantmentName["Legend"],
		"listvalue_legend":   boots_EnchantmentValues["Legend"],
		"listname_unique":    boots_EnchantmentName["Unique"],
		"listvalue_unique":   boots_EnchantmentValues["Unique"],
	})
}

func Cloak_EnchantmentList_Handler(c *gin.Context) {
	cloak_EnchantmentName := map[string][]string{}
	cloak_EnchantmentValues := map[string][]float32{}

	cloak_EnchantmentName["Uncommon"] = GetEnchatmentLists_Armor_ExceptionBase(c)["cloak"]
	cloak_EnchantmentName["Rare"] = GetEnchatmentLists_Armor_TypeRare(c)["cloak"]
	cloak_EnchantmentName["Epic"] = GetEnchatmentLists_Armor_TypeEpic(c)["cloak"]
	cloak_EnchantmentName["Legend"] = GetEnchatmentLists_Armor_TypeLegend(c)["cloak"]
	cloak_EnchantmentName["Unique"] = GetEnchatmentLists_Armor_TypeUnique(c)["cloak"]

	cloak_EnchantmentValues["Uncommon"] = GetEnchatmentLists_Armor_ValuesUncommon(c)["cloak"]
	cloak_EnchantmentValues["Rare"] = GetEnchatmentLists_Armor_ValuesRare(c)["cloak"]
	cloak_EnchantmentValues["Epic"] = GetEnchatmentLists_Armor_ValuesEpic(c)["cloak"]
	cloak_EnchantmentValues["Legend"] = GetEnchatmentLists_Armor_ValuesLegend(c)["cloak"]
	cloak_EnchantmentValues["Unique"] = GetEnchatmentLists_Armor_ValuesUnique(c)["cloak"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  cloak_EnchantmentName["Uncommon"],
		"listvalue_uncommon": cloak_EnchantmentValues["Uncommon"],
		"listname_rare":      cloak_EnchantmentName["Rare"],
		"listvalue_rare":     cloak_EnchantmentValues["Rare"],
		"listname_epic":      cloak_EnchantmentName["Epic"],
		"listvalue_epic":     cloak_EnchantmentValues["Epic"],
		"listname_legend":    cloak_EnchantmentName["Legend"],
		"listvalue_legend":   cloak_EnchantmentValues["Legend"],
		"listname_unique":    cloak_EnchantmentName["Unique"],
		"listvalue_unique":   cloak_EnchantmentValues["Unique"],
	})

}

func Necklace_EnchantmentList_Handler(c *gin.Context) {
	necklace_EnchantmentName := map[string][]string{}
	necklace_EnchantmentValues := map[string][]float32{}

	necklace_EnchantmentName["Uncommon"] = GetEnchatmentLists_Accessory_ExceptionBase(c)["necklace"]

	necklace_EnchantmentName["Rare"] = GetEnchantmentLists_Accessory_TypeRare(c)["necklace"]
	necklace_EnchantmentName["Epic"] = GetEnchantmentLists_Accessory_TypeEpic(c)["necklace"]
	necklace_EnchantmentName["Legend"] = GetEnchantmentLists_Accessory_TypeLegend(c)["necklace"]
	necklace_EnchantmentName["Unique"] = GetEnchantmentLists_Accessory_TypeUnique(c)["necklace"]

	necklace_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Accessory_ValuesUncommon(c)["necklace"]
	necklace_EnchantmentValues["Rare"] = GetEnchantmentLists_Accessory_ValuesRare(c)["necklace"]
	necklace_EnchantmentValues["Epic"] = GetEnchantmentLists_Accessory_ValuesEpic(c)["necklace"]
	necklace_EnchantmentValues["Legend"] = GetEnchantmentLists_Accessory_ValuesLegend(c)["necklace"]
	necklace_EnchantmentValues["Unique"] = GetEnchantmentLists_Accessory_ValuesUnique(c)["necklace"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  necklace_EnchantmentName["Uncommon"],
		"listvalue_uncommon": necklace_EnchantmentValues["Uncommon"],
		"listname_rare":      necklace_EnchantmentName["Rare"],
		"listvalue_rare":     necklace_EnchantmentValues["Rare"],
		"listname_epic":      necklace_EnchantmentName["Epic"],
		"listvalue_epic":     necklace_EnchantmentValues["Epic"],
		"listname_legend":    necklace_EnchantmentName["Legend"],
		"listvalue_legend":   necklace_EnchantmentValues["Legend"],
		"listname_unique":    necklace_EnchantmentName["Unique"],
		"listvalue_unique":   necklace_EnchantmentValues["Unique"],
	})

}

func Ring_EnchantmentList_Handler(c *gin.Context) {
	ring_EnchantmentName := map[string][]string{}
	ring_EnchantmentValues := map[string][]float32{}

	ring_EnchantmentName["Uncommon"] = GetEnchatmentLists_Accessory_ExceptionBase(c)["ring"]
	ring_EnchantmentName["Rare"] = GetEnchantmentLists_Accessory_TypeRare(c)["ring"]
	ring_EnchantmentName["Epic"] = GetEnchantmentLists_Accessory_TypeEpic(c)["ring"]
	ring_EnchantmentName["Legend"] = GetEnchantmentLists_Accessory_TypeLegend(c)["ring"]
	ring_EnchantmentName["Unique"] = GetEnchantmentLists_Accessory_TypeUnique(c)["ring"]

	ring_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Accessory_ValuesUncommon(c)["ring"]
	ring_EnchantmentValues["Rare"] = GetEnchantmentLists_Accessory_ValuesRare(c)["ring"]
	ring_EnchantmentValues["Epic"] = GetEnchantmentLists_Accessory_ValuesEpic(c)["ring"]
	ring_EnchantmentValues["Legend"] = GetEnchantmentLists_Accessory_ValuesLegend(c)["ring"]
	ring_EnchantmentValues["Unique"] = GetEnchantmentLists_Accessory_ValuesUnique(c)["ring"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  ring_EnchantmentName["Uncommon"],
		"listvalue_uncommon": ring_EnchantmentValues["Uncommon"],
		"listname_rare":      ring_EnchantmentName["Rare"],
		"listvalue_rare":     ring_EnchantmentValues["Rare"],
		"listname_epic":      ring_EnchantmentName["Epic"],
		"listvalue_epic":     ring_EnchantmentValues["Epic"],
		"listname_legend":    ring_EnchantmentName["Legend"],
		"listvalue_legend":   ring_EnchantmentValues["Legend"],
		"listname_unique":    ring_EnchantmentName["Unique"],
		"listvalue_unique":   ring_EnchantmentValues["Unique"],
	})

}

func RingTwo_EnchantmentList_Handler(c *gin.Context) {
	ringTwo_EnchantmentName := map[string][]string{}
	ringTwo_EnchantmentValues := map[string][]float32{}

	ringTwo_EnchantmentName["Uncommon"] = GetEnchatmentLists_Accessory_ExceptionBase(c)["ringtwo"]
	ringTwo_EnchantmentName["Rare"] = GetEnchantmentLists_Accessory_TypeRare(c)["ringtwo"]
	ringTwo_EnchantmentName["Epic"] = GetEnchantmentLists_Accessory_TypeEpic(c)["ringtwo"]
	ringTwo_EnchantmentName["Legend"] = GetEnchantmentLists_Accessory_TypeLegend(c)["ringtwo"]
	ringTwo_EnchantmentName["Unique"] = GetEnchantmentLists_Accessory_TypeUnique(c)["ringtwo"]

	ringTwo_EnchantmentValues["Uncommon"] = GetEnchantmentLists_Accessory_ValuesUncommon(c)["ringtwo"]
	ringTwo_EnchantmentValues["Rare"] = GetEnchantmentLists_Accessory_ValuesRare(c)["ringtwo"]
	ringTwo_EnchantmentValues["Epic"] = GetEnchantmentLists_Accessory_ValuesEpic(c)["ringtwo"]
	ringTwo_EnchantmentValues["Legend"] = GetEnchantmentLists_Accessory_ValuesLegend(c)["ringtwo"]
	ringTwo_EnchantmentValues["Unique"] = GetEnchantmentLists_Accessory_ValuesUnique(c)["ringtwo"]

	c.JSON(http.StatusOK, gin.H{
		"listname_uncommon":  ringTwo_EnchantmentName["Uncommon"],
		"listvalue_uncommon": ringTwo_EnchantmentValues["Uncommon"],
		"listname_rare":      ringTwo_EnchantmentName["Rare"],
		"listvalue_rare":     ringTwo_EnchantmentValues["Rare"],
		"listname_epic":      ringTwo_EnchantmentName["Epic"],
		"listvalue_epic":     ringTwo_EnchantmentValues["Epic"],
		"listname_legend":    ringTwo_EnchantmentName["Legend"],
		"listvalue_legend":   ringTwo_EnchantmentValues["Legend"],
		"listname_unique":    ringTwo_EnchantmentName["Unique"],
		"listvalue_unique":   ringTwo_EnchantmentValues["Unique"],
	})
}

//////////\\\\\\\\\ ------->   ROUTES  <------- //////////\\\\\\\\\

func setupRoutes(r *gin.Engine) {
	// LISTS ITEMS ENDPOINTS
	r.GET("/helmetlist/:classSelection", Helmet_List_Handler)
	r.GET("/chestlist/:classSelection", Chest_List_Handler)
	r.GET("/gloveslist/:classSelection", Gloves_List_Handler)
	r.GET("/pantslist/:classSelection", Pants_List_Handler)
	r.GET("/bootslist/:classSelection", Boots_List_Handler)
	r.GET("/cloaklist/:classSelection", Cloak_List_Handler)
	r.GET("/necklacelist/", Necklace_List_Handler)
	r.GET("/ringlist/", Ring_List_Handler)
	//r.GET("/ringtwolist/", RingTwo_List_Handler)

	// RATING LISTS ENDPOINTS
	r.GET("/helmetratinglist/", Helmet_RatingList_Handler)
	r.GET("/chestratinglist/", Chest_RatingList_Handler)
	r.GET("/glovesratinglist/", Gloves_RatingList_Handler)
	r.GET("/pantsratinglist/", Pants_RatingList_Handler)
	r.GET("/bootsratinglist/", Boots_RatingList_Handler)
	r.GET("/cloakratinglist/", Cloak_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.GET("/enchantmentlisthelmet/", Helmet_EnchantmentList_Handler)
	r.GET("/enchantmentlistchest/", Chest_EnchantmentList_Handler)
	r.GET("/enchantmentlistgloves/", Gloves_EnchantmentList_Handler)
	r.GET("/enchantmentlistpants/", Pants_EnchantmentList_Handler)
	r.GET("/enchantmentlistboots/", Boots_EnchantmentList_Handler)
	r.GET("/enchantmentlistcloak/", Cloak_EnchantmentList_Handler)
	r.GET("/enchantmentlistnecklace/", Necklace_EnchantmentList_Handler)
	r.GET("/enchantmentlistring/", Ring_EnchantmentList_Handler)
	r.GET("/enchantmentlistringtwo/", RingTwo_EnchantmentList_Handler)

	// ITEM FIXED ATTRIBUTES ENDPOINT
	r.GET("/itemdisplay/:item", itemDisplayHandler)

	// Calculate stats endpoint
	r.GET("/charbuilder/:classSelection", updateStatsHandler)

}
