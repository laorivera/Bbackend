package main

type Item struct {
	Armet          Item_Armor
	CrusaderHelm   Item_Armor
	ChapelDeFer    Item_Armor
	Doublet        Item_Armor
	RivetedGloves  Item_Armor
	OccultistPants Item_Armor
	LightfootBoots Item_Armor
	RadiantCloak   Item_Armor
	AdventureCloak Item_Armor
	BarbutaHelm    Item_Armor
	ArmingSword    Item_Weapon
	Bardiche       Item_Weapon
	LooseTrousers  Item_Armor
	FoxPendant     Item_Accessory
	RingOfCourage  Item_Accessory
}

var item = Item{

	Armet:          CreateItemArmor("./data/Armet.json"),
	BarbutaHelm:    CreateItemArmor("./data/BarbutaHelm.json"),
	ChapelDeFer:    CreateItemArmor("./data/ChapelDeFer.json"),
	CrusaderHelm:   CreateItemArmor("./data/CrusaderHelm.json"),
	Doublet:        CreateItemArmor("./data/Doublet.json"),
	RivetedGloves:  CreateItemArmor("./data/RivetedGloves.json"),
	LooseTrousers:  CreateItemArmor("./data/LooseTrousers.json"),
	OccultistPants: CreateItemArmor("./data/OccultistPants.json"),
	LightfootBoots: CreateItemArmor("./data/LightfootBoots.json"),
	RadiantCloak:   CreateItemArmor("./data/RadiantCloak.json"),
	AdventureCloak: CreateItemArmor("./data/AdventureCloak.json"),
	ArmingSword:    CreateItemWeapon("./data/ArmingSword.json"),
	Bardiche:       CreateItemWeapon("./data/Bardiche.json"),
	FoxPendant:     CreateItemAccessory("./data/FoxPendant.json"),
	RingOfCourage:  CreateItemAccessory("./data/RingOfCourage.json"),
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
		item.LooseTrousers,
	},
	ItemsWeapon: []Item_Weapon{
		item.ArmingSword,
		item.Bardiche,
	},
	ItemsAccessory: []Item_Accessory{
		item.FoxPendant,
		item.RingOfCourage,
	},
}
