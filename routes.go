package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func charBuilderHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/*
func testHandler(c *gin.Context) {
	printClass := InttoClass(c.Param("classSelection"))

	c.HTML(http.StatusOK, "index.html", gin.H{
		"printclass": printClass,
	})
}*/

func updateStatsHandler(c *gin.Context) {

	//class query
	class := c.Param("classSelection")
	classStatSelect := SelectClass(class)
	//printClass := InttoClass(c.Param("classSelection"))

	//rarityList := []string{"1", "2", "3", "4", "5", "6", "7"}

	//itemLists_Armor := GetItemLists_Armor(c, class) // list of armor of class selected
	//itemLists_Accessory := GetItemLists_Accesory(c) // list of accessory all classes

	//ratingLists_Armor := GetRatingLists_Armor(c) // list rating result of armor selected

	itemsSelected_Armor := GetSelectedItems_Armor(c) //Items query selection Armor
	raritySelected_Armor := GetSelectedRarities_Armor(c)
	ratingSelected_Armor := GetSelectedRatings_Armor(c)

	itemsSelected_Accessory := GetSelectedItems_Accessory(c) //Items query selection Accessory
	raritySelected_Accessory := GetSelectedRarities_Accessory(c)

	//Enchatment query selected Armor
	enchantmentSelectedOther_ArmorUncommon := GetSelectedEnchantmentsOther_ArmorUncommon(c)
	enchantmentSelectedOther_ArmorRare := GetSelectedEnchantmentsOther_ArmorRare(c)
	enchantmentSelectedOther_ArmorEpic := GetSelectedEnchantmentsOther_ArmorEpic(c)
	enchantmentSelectedOther_ArmorLegend := GetSelectedEnchantmentsOther_ArmorLegend(c)
	enchantmentSelectedOther_ArmorUnique := GetSelectedEnchantmentsOther_ArmorUnique(c)

	enchantmentSelected_ArmorUncommon := GetSelectedEnchantmentsBase_ArmorUncommon(c)
	enchantmentSelected_ArmorRare := GetSelectedEnchantmentsBase_ArmorRare(c)
	enchantmentSelected_ArmorEpic := GetSelectedEnchantmentsBase_ArmorEpic(c)
	enchantmentSelected_ArmorLegend := GetSelectedEnchantmentsBase_ArmorLegend(c)
	enchantmentSelected_ArmorUnique := GetSelectedEnchantmentsBase_ArmorUnique(c)

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

	//list accessory query by class

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

	//enchamentList_Armor_TypesUncommon := GetEnchatmentLists_Armor_Exception(c)
	//enchamentList_Armor_ValuesUncommon := GetEnchatmentLists_Armor_ValuesUncommon(c)
	//enchantmentList_Armor_TypeRare := GetEnchatmentLists_Armor_TypeRare(c)
	//enchantmentList_Armor_ValueRare := GetEnchatmentLists_Armor_ValuesRare(c)
	//enchantmentList_Armor_TypeEpic := GetEnchatmentLists_Armor_TypeEpic(c)
	//enchantmentList_Armor_ValueEpic := GetEnchatmentLists_Armor_ValuesEpic(c)
	//enchantmentList_Armor_TypeLegend := GetEnchatmentLists_Armor_TypeLegend(c)
	//enchantmentList_Armor_ValueLegend := GetEnchatmentLists_Armor_ValuesLegend(c)
	//enchantmentList_Armor_TypeUnique := GetEnchatmentLists_Armor_TypeUnique(c)
	//enchantmentList_Armor_ValueUnique := GetEnchatmentLists_Armor_ValuesUnique(c)

	//enchantmentList_Accessory_TypeUncommon := GetEnchantmentLists_Accessory_Exception(c)
	//enchantmentList_Accessory_ValuesUncommon := GetEnchantmentLists_Accessory_ValuesUncommon(c)
	//enchantmentList_Accessory_TypeRare := GetEnchantmentLists_Accessory_TypeRare(c)
	//enchantmentList_Accessory_ValuesRare := GetEnchantmentLists_Accessory_ValuesRare(c)
	//enchantmentList_Accessory_TypeEpic := GetEnchantmentLists_Accessory_TypeEpic(c)
	//enchantmentList_Accessory_ValuesEpic := GetEnchantmentLists_Accessory_ValuesEpic(c)
	//enchantmentList_Accessory_TypeLegend := GetEnchantmentLists_Accessory_TypeLegend(c)
	//enchantmentList_Accessory_ValuesLegend := GetEnchantmentLists_Accessory_ValuesLegend(c)
	//enchantmentList_Accessory_TypeUnique := GetEnchantmentLists_Accessory_TypeUnique(c)
	//enchantmentList_Accessory_ValuesUnique := GetEnchantmentLists_Accessory_ValuesUnique(c)

	computedStatsOther_Armor := Computed_Stats{}
	computedStatsOther_Armor = computedStatsOther_Armor.AddEnchant(EnchantComputedOthers(enchantmentSelectedOther_ArmorUncommon), EnchantComputedOthers(enchantmentSelectedOther_ArmorRare), EnchantComputedOthers(enchantmentSelectedOther_ArmorEpic), EnchantComputedOthers(enchantmentSelectedOther_ArmorLegend), EnchantComputedOthers(enchantmentSelectedOther_ArmorUnique))
	computedStatsOther_Accesory := Computed_Stats{}
	computedStatsOther_Accesory = computedStatsOther_Accesory.AddEnchant(EnchantComputedOthers(enchantmentSelectedOther_AccessoryUncommon), EnchantComputedOthers(enchantmentSelectedOther_AccessoryRare), EnchantComputedOthers(enchantmentSelectedOther_AccessoryEpic), EnchantComputedOthers(enchantmentSelectedOther_AccessoryLegend), EnchantComputedOthers(enchantmentSelectedOther_AccessoryUnique))

	totalRating := RatingCalc(ratingSelected_Armor)
	totalSpeed := SpeedCalc(itemsSelected_Armor, raritySelected_Armor)
	//totalRatingWeapon := RatingCalc(ratingSelected_Weapon)

	updatedStatsArmor := SetItemStats(classStatSelect, itemsSelected_Armor, raritySelected_Armor)
	updatedStatsArmor = updatedStatsArmor.AddStats(setEnchantStats(enchantmentSelected_ArmorUncommon), setEnchantStats(enchantmentSelected_ArmorRare), setEnchantStats(enchantmentSelected_ArmorEpic), setEnchantStats(enchantmentSelected_ArmorLegend), setEnchantStats(enchantmentSelected_ArmorUnique))

	updatedStatsAccessory := SetItemStatsAccessory(characterHolder, itemsSelected_Accessory, raritySelected_Accessory)
	updatedStatsAccessory = updatedStatsAccessory.AddStats(setEnchantStats(enchantmentSelected_AccessoryUncommon), setEnchantStats(enchantmentSelected_AccessoryRare), setEnchantStats(enchantmentSelected_AccessoryEpic), setEnchantStats(enchantmentSelected_AccessoryLegend), setEnchantStats(enchantmentSelected_AccessoryUnique))

	updatedStats := Stats{}
	updatedStats = updatedStats.AddStats(updatedStatsArmor, updatedStatsAccessory)

	computedStatsCurve := CalculateComputedValues(updatedStats, totalRating, totalSpeed, computedStatsOther_Armor, computedStatsOther_Accesory) // rating variable armor rating and speed

	computedStats := ComputedTotal(computedStatsCurve, computedStatsOther_Armor, computedStatsOther_Accesory)

	/*
		c.HTML(http.StatusOK, "charbuilder.html", gin.H{

			"classpick":                class,
			"raritylistpick":           rarityList,
			"stats":                    updatedStats,
			"computedstats":            computedStats,
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
		"stats":         updatedStats,
		"computedstats": computedStats,
		//"test":          GetItemLists_Accesory_Json(c)["Necklace"],
		//"test2":         ItemsByNameAccessory("Ring Of Courage"),
	})

}

//////////\\\\\\\\\ ------->  LISTS ITEMS HANDLER <------- //////////\\\\\\\\\

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
		"list": imagePants},
	)
}

func Boots_List_Handler(c *gin.Context) {
	class := c.Param("classSelection")
	bootsList := GetItemLists_Armor_Json(c, class)["Feet"]
	imageBoots := ImageLocation("Feet", bootsList)

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

func RingTwo_List_Handler(c *gin.Context) {
	ringTwoList := GetItemLists_Accesory_Json(c)["RingTwo"]
	imageRingTwo := ImageLocation("RingTwo", ringTwoList)

	c.JSON(http.StatusOK, gin.H{
		"list": imageRingTwo},
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
	/*

		helmetEnchantList_Uncommon := GetEnchatmentLists_Armor_Exception(c)["helmet"]
		helmetEnchantList_ValuesUncommon := GetEnchatmentLists_Armor_ValuesUncommon(c)["helmet"]

		helmetEnchantList_Rare := GetEnchatmentLists_Armor_TypeRare(c)["helmet"]
		helmetEnchantList_ValuesRare := GetEnchatmentLists_Armor_ValuesRare(c)["helmet"]

		helmetEnchantList_Epic := GetEnchatmentLists_Armor_TypeEpic(c)["helmet"]
		helmetEnchantList_ValuesEpic := GetEnchatmentLists_Armor_ValuesEpic(c)["helmet"]

		helmetEnchantList_Legend := GetEnchatmentLists_Armor_TypeLegend(c)["helmet"]
		helmetEnchantList_ValuesLegend := GetEnchatmentLists_Armor_ValuesLegend(c)["helmet"]

		helmetEnchantList_Unique := GetEnchatmentLists_Armor_TypeUnique(c)["helmet"]
		helmetEnchantList_ValuesUnique := GetEnchatmentLists_Armor_ValuesUnique(c)["helmet"]

	*/
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
	r.GET("/ringtwolist/", RingTwo_List_Handler)

	// RATING LISTS ENDPOINTS
	r.GET("/helmetratinglist/", Helmet_RatingList_Handler)
	r.GET("/chestratinglist/", Chest_RatingList_Handler)
	r.GET("/glovesratinglist/", Gloves_RatingList_Handler)
	r.GET("/pantsratinglist/", Pants_RatingList_Handler)
	r.GET("/bootsratinglist/", Boots_RatingList_Handler)
	r.GET("/cloakratinglist/", Cloak_RatingList_Handler)

	//ENCHANTMENT LISTS ENDPOINTS
	r.GET("/enchamentlistarmor/", Helmet_EnchantmentList_Handler)
	r.GET("/enchamentlistnecklace/", Necklace_EnchantmentList_Handler)
	//r.GET("/enchamentlistaccessory/:classSelection", Enchantment_List_Handler_Accessory)

	// THE ONE
	r.GET("/charbuilder/:classSelection", updateStatsHandler)

	//r.GET("/charbuilder/", charBuilderHandler)
	//r.GET("/charbuilder/:classSelection", updateStatsHandler)
}
