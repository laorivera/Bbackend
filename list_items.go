package main

type Item struct {
	Armet          Item_Armor
	CrusaderHelm   Item_Armor
	Doublet        Item_Armor
	RivetedGloves  Item_Armor
	OccultistPants Item_Armor
	LightfootBoots Item_Armor
	RadiantCloak   Item_Armor
	AdventureCloak Item_Armor
	ArmingSword    Item_Weapon
	Bardiche       Item_Weapon
	FoxPendant     Item_Accessory
	BarbutaHelm    Item_Armor
	ChapelDeFer    Item_Armor
}

var item = Item{

	Armet:          CreateItemArmor("./data/itemArmet2.json"),
	CrusaderHelm:   CreateItemArmor("./data/itemCrusaderHelm.json"),
	Doublet:        CreateItemArmor("./data/itemDoublet.json"),
	RivetedGloves:  CreateItemArmor("./data/itemRivetedGloves.json"),
	OccultistPants: CreateItemArmor("./data/itemOccultistPants.json"),
	LightfootBoots: CreateItemArmor("./data/itemLightfootBoots.json"),
	RadiantCloak:   CreateItemArmor("./data/itemRadiantCloak.json"),
	AdventureCloak: CreateItemArmor("./data/itemAdventureCloak.json"),
	ArmingSword:    CreateItemWeapon("./data/itemArmingSword.json"),
	Bardiche:       CreateItemWeapon("./data/itemBardiche.json"),
	FoxPendant:     CreateItemAccessory("./data/itemFoxPendant.json"),
	BarbutaHelm:    CreateItemArmor("./data/itemBarbutaHelm.json"),
	ChapelDeFer:    CreateItemArmor("./data/itemChapelDeFer.json"),
}

type List_Items struct {
	ItemsArmor     []Item_Armor
	ItemsWeapon    []Item_Weapon
	ItemsAccessory []Item_Accessory
}

var Items = List_Items{
	ItemsArmor: []Item_Armor{
		item.Armet,
		item.CrusaderHelm,
		item.Doublet,
		item.RivetedGloves,
		item.OccultistPants,
		item.LightfootBoots,
		item.RadiantCloak,
		item.AdventureCloak,
		item.BarbutaHelm,
		item.ChapelDeFer,
	},
	ItemsWeapon: []Item_Weapon{
		item.ArmingSword,
		item.Bardiche,
	},
	ItemsAccessory: []Item_Accessory{
		item.FoxPendant,
	},
}
