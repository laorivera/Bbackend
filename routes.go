package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	printClass := InttoClass(c.Param("classSelection"))

	rarityList := []string{"1", "2", "3", "4", "5", "6", "7"}

	//list equipement query by class
	helmetList := ItemsBySlotType(c.Param("classSelection"), "Head")
	chestList := ItemsBySlotType(c.Param("classSelection"), "Chest")
	glovesList := ItemsBySlotType(c.Param("classSelection"), "Hands")
	pantsList := ItemsBySlotType(c.Param("classSelection"), "Legs")
	bootsList := ItemsBySlotType(c.Param("classSelection"), "Foot")
	cloakList := ItemsBySlotType(c.Param("classSelection"), "Back")

	//list weapon query by class
	primaryWeaponList := WeaponsBySlotType(c.Param("classSelection"), "Main Hand")

	//list accessory query by class
	necklaceList := AccessoryBySlotType(c.Param("classSelection"), "Necklace")

	//List items query selected
	itemsSelected_Armor := []Item_Armor{
		ItemsByNameArmor(c.Query("itemhelmet")),
		ItemsByNameArmor(c.Query("itemchest")),
		ItemsByNameArmor(c.Query("itemgloves")),
		ItemsByNameArmor(c.Query("itempants")),
		ItemsByNameArmor(c.Query("itemboots")),
		ItemsByNameArmor(c.Query("itemcloak")),
	}

	itemsSelected_Accessory := []Item_Accessory{
		ItemsByNameAccessory(c.Query("itemnecklace")),
	}

	raritySelected_Accessory := []int{
		StringtoInt(c.Query("rarityselect_necklace")),
	}

	enchantmentSelectedBase_Accessory := []map[string]int{
		Enchantattrib(c.Query("enchantment_necklacetype"), c.Query("enchantment_necklacevalue")),
	}

	enchantmentSelectedOther_Accessory := []map[string]float64{
		Enchantother(c.Query("enchantment_necklacetype"), c.Query("enchantment_necklacevalue")),
	}
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

	raritySelected_Armor := []int{
		StringtoInt(c.Query("rarityselect_helmet")),
		StringtoInt(c.Query("rarityselect_chest")),
		StringtoInt(c.Query("rarityselect_gloves")),
		StringtoInt(c.Query("rarityselect_pants")),
		StringtoInt(c.Query("rarityselect_boots")),
		StringtoInt(c.Query("rarityselect_cloak")),
	}

	ratingSelected_Armor := []int{
		StringtoInt(c.Query("armorrating_helmet")),
		StringtoInt(c.Query("armorrating_chest")),
		StringtoInt(c.Query("armorrating_gloves")),
		StringtoInt(c.Query("armorrating_pants")),
		StringtoInt(c.Query("armorrating_boots")),
		StringtoInt(c.Query("armorrating_cloak")),
	}

	enchantmentSelectedOther_Armor := []map[string]float64{
		Enchantother(c.Query("enchantment_helmettype"), c.Query("enchantment_helmetvalue")),
		Enchantother(c.Query("enchantment_chesttype"), c.Query("enchantment_chestvalue")),
		Enchantother(c.Query("enchantment_glovestype"), c.Query("enchantment_glovesvalue")),
		Enchantother(c.Query("enchantment_pantstype"), c.Query("enchantment_pantsvalue")),
		Enchantother(c.Query("enchantment_bootstype"), c.Query("enchantment_bootsvalue")),
		Enchantother(c.Query("enchantment_cloaktype"), c.Query("enchantment_cloakvalue")),
	}

	enchantmentSelectedBase_Armor := []map[string]int{

		Enchantattrib(c.Query("enchantment_helmettype"), c.Query("enchantment_helmetvalue")),
		Enchantattrib(c.Query("enchantment_chesttype"), c.Query("enchantment_chestvalue")),
		Enchantattrib(c.Query("enchantment_glovestype"), c.Query("enchantment_glovesvalue")),
		Enchantattrib(c.Query("enchantment_pantstype"), c.Query("enchantment_pantsvalue")),
		Enchantattrib(c.Query("enchantment_bootstype"), c.Query("enchantment_bootsvalue")),
		Enchantattrib(c.Query("enchantment_cloaktype"), c.Query("enchantment_cloakvalue")),
	}

	ratingListPrimaryWeapon := ItemsByNameWeapon(c.Query("itemprimaryweapon")).DamageRatings[StringtoInt(c.Query("rarityselect_weapon"))]

	// Lists rating values
	ratingListHelmet := ItemsByNameArmor(c.Query("itemhelmet")).ArmorRatings[StringtoInt(c.Query("rarityselect_helmet"))]
	ratingListChest := ItemsByNameArmor(c.Query("itemchest")).ArmorRatings[StringtoInt(c.Query("rarityselect_chest"))]
	ratingListGloves := ItemsByNameArmor(c.Query("itemgloves")).ArmorRatings[StringtoInt(c.Query("rarityselect_gloves"))]
	ratingListPants := ItemsByNameArmor(c.Query("itempants")).ArmorRatings[StringtoInt(c.Query("rarityselect_pants"))]
	ratingListBoots := ItemsByNameArmor(c.Query("itemboots")).ArmorRatings[StringtoInt(c.Query("rarityselect_boots"))]
	ratingListCloak := ItemsByNameArmor(c.Query("itemcloak")).ArmorRatings[StringtoInt(c.Query("rarityselect_cloak"))]

	enchamentListHead := EchantBaseAttribExeption(EnchamentbySlot(Enchantments.Helmet), ItemsByNameArmor(c.Query("itemhelmet")))
	enchamentListHeadValue := EnchantValuesCalc(c.Query("enchantment_helmettype"), Enchantments.Helmet)

	enchantmentListNecklace := EchantBaseAttribExeptionAcc(EnchamentbySlot(Enchantments.Necklace), ItemsByNameAccessory(c.Query("itemnecklace")))
	//enchantmentListHeadTwo := EnchantTypeExeption(enchamentListHead, c.Query("enchantment_helmettype"))

	computedStatsOther := EnchantComputedOthers(enchantmentSelectedOther_Armor)
	computedStatsOtherAcc := EnchantComputedOthers(enchantmentSelectedOther_Accessory)

	computedStatsOther = computedStatsOther.AddEnchant(computedStatsOtherAcc)

	totalRating := RatingCalc(ratingSelected_Armor)
	totalSpeed := SpeedCalc(itemsSelected_Armor, raritySelected_Armor)
	//totalRatingWeapon := RatingCalc(ratingSelected_Weapon)

	updatedStats := SetItemStats(classStatSelect, itemsSelected_Armor, raritySelected_Armor, enchantmentSelectedBase_Armor)
	updatedStatsAccessory := SetItemStatsAccessory(characterStatsx, itemsSelected_Accessory, raritySelected_Accessory, enchantmentSelectedBase_Accessory)

	updatedStats = updatedStats.AddStats(updatedStatsAccessory)

	computedStatsCurve := CalculateComputedValues(updatedStats, totalRating, totalSpeed, computedStatsOther) // rating variable armor rating and speed

	computedStats := ComputedTotal(computedStatsCurve, computedStatsOther)

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

}

func Helmet_List_Handler(c *gin.Context) {
	class := "1"
	helmetList := ItemsBySlotType_Json((class), "Head")
	imageHelmet := ImageLocation("Head", helmetList)

	c.JSON(http.StatusOK, gin.H{
		"character": imageHelmet},
	)
}

func setupRoutes(r *gin.Engine) {
	r.GET("/helmetlist/", Helmet_List_Handler)
	r.GET("/charbuilder/", charBuilderHandler)
	r.GET("/charbuilder/:classSelection", updateStatsHandler)
	//r.GET("/charbuilder/:classSelection", updateStatsHandler)
}
