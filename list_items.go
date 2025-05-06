package main

type Item struct {
	Armet                Item_Armor
	CrusaderHelm         Item_Armor
	ChapelDeFer          Item_Armor
	Doublet              Item_Armor
	RivetedGloves        Item_Armor
	OccultistPants       Item_Armor
	LightfootBoots       Item_Armor
	RadiantCloak         Item_Armor
	AdventurerCloak      Item_Armor
	BarbutaHelm          Item_Armor
	ArmingSword          Item_Weapon
	Bardiche             Item_Weapon
	LooseTrousers        Item_Armor
	FoxPendant           Item_Accessory
	RingOfCourage        Item_Accessory
	AdventurerBoots      Item_Armor
	ArcaneHood           Item_Armor
	CeremonialHeaddress  Item_Armor
	Chaperon             Item_Armor
	Coif                 Item_Armor
	DarkgroveHood        Item_Armor
	ElkwoodCrown         Item_Armor
	FeatheredHat         Item_Armor
	ForestHood           Item_Armor
	Gjermundbu           Item_Armor
	GreatHelm            Item_Armor
	Hounskull            Item_Armor
	KettleHat            Item_Armor
	LeatherBonnet        Item_Armor
	LeatherCap           Item_Armor
	NormanNasalHelm      Item_Armor
	OpenSallet           Item_Armor
	RangerHood           Item_Armor
	RogueCowl            Item_Armor
	Sallet               Item_Armor
	ShadowHood           Item_Armor
	ShadowMask           Item_Armor
	Spangenhelm          Item_Armor
	StrawHat             Item_Armor
	Topfhelm             Item_Armor
	VikingHelm           Item_Armor
	OccultistHood        Item_Armor
	AdventurerTunic      Item_Armor
	ArcaneGrab           Item_Armor
	ArcaneGloves         Item_Armor
	BardicPants          Item_Armor
	BloodwovenGloves     Item_Armor
	ChampionArmor        Item_Armor
	CrusaderArmor        Item_Armor
	BloodwovenRobe       Item_Armor
	BuckledBoots         Item_Armor
	ButtonedBoots        Item_Armor
	ButtonedLeggings     Item_Armor
	ClothPants           Item_Armor
	CuffedBoots          Item_Armor
	DarkCuirass          Item_Armor
	DarkgroveRobe        Item_Armor
	DarkleafBoots        Item_Armor
	DarkLeatherLeggings  Item_Armor
	DarkPlateArmor       Item_Armor
	DashingBoots         Item_Armor
	DreadHood            Item_Armor
	ElkwoodGloves        Item_Armor
	FineCuirass          Item_Armor
	ForestBoots          Item_Armor
	Frock                Item_Armor
	GlovesOfUtility      Item_Armor
	GrandBrigandine      Item_Armor
	WandererAttire       Item_Armor
	WatchmanCloak        Item_Armor
	WizardHat            Item_Armor
	WizardShoes          Item_Armor
	HeavyBoots           Item_Armor
	HeavyGambeson        Item_Armor
	HeavyGauntlets       Item_Armor
	HeavyLeggings        Item_Armor
	HeavyLeatherLeggings Item_Armor
	LacedTurnshoe        Item_Armor
	LeatherGloves        Item_Armor
	LeatherLeggings      Item_Armor
	LightAketon          Item_Armor
	LightGauntlets       Item_Armor
	LowBoots             Item_Armor
	MarauderOutfit       Item_Armor
	MercurialCloak       Item_Armor
	MysticGloves         Item_Armor
	MysticVestments      Item_Armor
	NorthernFullTunic    Item_Armor
	OccultistRobe        Item_Armor
	OccultistTunic       Item_Armor
	OldShoes             Item_Armor
	Pourpoint            Item_Armor
	TemplarArmor         Item_Armor
	BadgerPendant        Item_Accessory
	BearPendant          Item_Accessory
	FrostAmulet          Item_Accessory
	MonkeyPendant        Item_Accessory
	NecklaceOfPeace      Item_Accessory
	OwlPendant           Item_Accessory
	OxPendant            Item_Accessory
	PhoenixChoker        Item_Accessory
	RatPendant           Item_Accessory
	TorqOfSoul           Item_Accessory
	WindLocket           Item_Accessory
	FangsOfDeathNecklace Item_Accessory
	RingOfFinesse        Item_Accessory
	RingOfQuickness      Item_Accessory
	RingOfSurvival       Item_Accessory
	RingOfVitality       Item_Accessory
	RingOfWisdom         Item_Accessory
	RingOfResolve        Item_Accessory
	TatteredCloak        Item_Armor
	SplendidCloak        Item_Armor
}

var item = Item{
	Armet:                CreateItemArmor("./data/Armet.json"),
	BarbutaHelm:          CreateItemArmor("./data/BarbutaHelm.json"),
	ChapelDeFer:          CreateItemArmor("./data/ChapelDeFer.json"),
	CrusaderHelm:         CreateItemArmor("./data/CrusaderHelm.json"),
	Doublet:              CreateItemArmor("./data/Doublet.json"),
	RivetedGloves:        CreateItemArmor("./data/RivetedGloves.json"),
	OccultistPants:       CreateItemArmor("./data/OccultistPants.json"),
	LooseTrousers:        CreateItemArmor("./data/LooseTrousers.json"),
	LightfootBoots:       CreateItemArmor("./data/LightfootBoots.json"),
	RadiantCloak:         CreateItemArmor("./data/RadiantCloak.json"),
	AdventurerCloak:      CreateItemArmor("./data/AdventurerCloak.json"),
	ArmingSword:          CreateItemWeapon("./data/ArmingSword.json"),
	Bardiche:             CreateItemWeapon("./data/Bardiche.json"),
	FoxPendant:           CreateItemAccessory("./data/FoxPendant.json"),
	RingOfCourage:        CreateItemAccessory("./data/RingOfCourage.json"),
	AdventurerBoots:      CreateItemArmor("./data/AdventurerBoots.json"),
	ArcaneHood:           CreateItemArmor("./data/ArcaneHood.json"),
	CeremonialHeaddress:  CreateItemArmor("./data/CeremonialHeaddress.json"),
	Chaperon:             CreateItemArmor("./data/Chaperon.json"),
	Coif:                 CreateItemArmor("./data/Coif.json"),
	DarkgroveHood:        CreateItemArmor("./data/DarkgroveHood.json"),
	ElkwoodCrown:         CreateItemArmor("./data/ElkwoodCrown.json"),
	FeatheredHat:         CreateItemArmor("./data/FeatheredHat.json"),
	ForestHood:           CreateItemArmor("./data/ForestHood.json"),
	Gjermundbu:           CreateItemArmor("./data/Gjermundbu.json"),
	GreatHelm:            CreateItemArmor("./data/GreatHelm.json"),
	Hounskull:            CreateItemArmor("./data/Hounskull.json"),
	KettleHat:            CreateItemArmor("./data/KettleHat.json"),
	LeatherBonnet:        CreateItemArmor("./data/LeatherBonnet.json"),
	LeatherCap:           CreateItemArmor("./data/LeatherCap.json"),
	NormanNasalHelm:      CreateItemArmor("./data/NormanNasalHelm.json"),
	OpenSallet:           CreateItemArmor("./data/OpenSallet.json"),
	RangerHood:           CreateItemArmor("./data/RangerHood.json"),
	RogueCowl:            CreateItemArmor("./data/RogueCowl.json"),
	Sallet:               CreateItemArmor("./data/Sallet.json"),
	ShadowHood:           CreateItemArmor("./data/ShadowHood.json"),
	ShadowMask:           CreateItemArmor("./data/ShadowMask.json"),
	Spangenhelm:          CreateItemArmor("./data/Spangenhelm.json"),
	StrawHat:             CreateItemArmor("./data/StrawHat.json"),
	Topfhelm:             CreateItemArmor("./data/Topfhelm.json"),
	VikingHelm:           CreateItemArmor("./data/VikingHelm.json"),
	OccultistHood:        CreateItemArmor("./data/OccultistHood.json"),
	AdventurerTunic:      CreateItemArmor("./data/AdventurerTunic.json"),
	ArcaneGrab:           CreateItemArmor("./data/ArcaneGrab.json"),
	ArcaneGloves:         CreateItemArmor("./data/ArcaneGloves.json"),
	BardicPants:          CreateItemArmor("./data/BardicPants.json"),
	BloodwovenGloves:     CreateItemArmor("./data/BloodwovenGloves.json"),
	ChampionArmor:        CreateItemArmor("./data/ChampionArmor.json"),
	CrusaderArmor:        CreateItemArmor("./data/CrusaderArmor.json"),
	BloodwovenRobe:       CreateItemArmor("./data/BloodwovenRobe.json"),
	BuckledBoots:         CreateItemArmor("./data/BuckledBoots.json"),
	ButtonedBoots:        CreateItemArmor("./data/ButtonedBoots.json"),
	ButtonedLeggings:     CreateItemArmor("./data/ButtonedLeggings.json"),
	ClothPants:           CreateItemArmor("./data/ClothPants.json"),
	CuffedBoots:          CreateItemArmor("./data/CuffedBoots.json"),
	DarkCuirass:          CreateItemArmor("./data/DarkCuirass.json"),
	DarkgroveRobe:        CreateItemArmor("./data/DarkgroveRobe.json"),
	DarkleafBoots:        CreateItemArmor("./data/DarkleafBoots.json"),
	DarkLeatherLeggings:  CreateItemArmor("./data/DarkLeatherLeggings.json"),
	DarkPlateArmor:       CreateItemArmor("./data/DarkPlateArmor.json"),
	DashingBoots:         CreateItemArmor("./data/DashingBoots.json"),
	DreadHood:            CreateItemArmor("./data/DreadHood.json"),
	ElkwoodGloves:        CreateItemArmor("./data/ElkwoodGloves.json"),
	FineCuirass:          CreateItemArmor("./data/FineCuirass.json"),
	ForestBoots:          CreateItemArmor("./data/ForestBoots.json"),
	Frock:                CreateItemArmor("./data/Frock.json"),
	GlovesOfUtility:      CreateItemArmor("./data/GlovesOfUtility.json"),
	GrandBrigandine:      CreateItemArmor("./data/GrandBrigandine.json"),
	WandererAttire:       CreateItemArmor("./data/WandererAttire.json"),
	WatchmanCloak:        CreateItemArmor("./data/WatchmanCloak.json"),
	WizardHat:            CreateItemArmor("./data/WizardHat.json"),
	WizardShoes:          CreateItemArmor("./data/WizardShoes.json"),
	HeavyBoots:           CreateItemArmor("./data/HeavyBoots.json"),
	HeavyGambeson:        CreateItemArmor("./data/HeavyGambeson.json"),
	HeavyGauntlets:       CreateItemArmor("./data/HeavyGauntlets.json"),
	HeavyLeggings:        CreateItemArmor("./data/HeavyLeggings.json"),
	HeavyLeatherLeggings: CreateItemArmor("./data/HeavyLeatherLeggings.json"),
	LacedTurnshoe:        CreateItemArmor("./data/LacedTurnshoe.json"),
	LeatherGloves:        CreateItemArmor("./data/LeatherGloves.json"),
	LeatherLeggings:      CreateItemArmor("./data/LeatherLeggings.json"),
	LightAketon:          CreateItemArmor("./data/LightAketon.json"),
	LightGauntlets:       CreateItemArmor("./data/LightGauntlets.json"),
	LowBoots:             CreateItemArmor("./data/LowBoots.json"),
	MarauderOutfit:       CreateItemArmor("./data/MarauderOutfit.json"),
	MercurialCloak:       CreateItemArmor("./data/MercurialCloak.json"),
	MysticGloves:         CreateItemArmor("./data/MysticGloves.json"),
	MysticVestments:      CreateItemArmor("./data/MysticVestments.json"),
	NorthernFullTunic:    CreateItemArmor("./data/NorthernFullTunic.json"),
	OccultistRobe:        CreateItemArmor("./data/OccultistRobe.json"),
	OccultistTunic:       CreateItemArmor("./data/OccultistTunic.json"),
	OldShoes:             CreateItemArmor("./data/OldShoes.json"),
	Pourpoint:            CreateItemArmor("./data/Pourpoint.json"),
	TemplarArmor:         CreateItemArmor("./data/TemplarArmor.json"),
	BadgerPendant:        CreateItemAccessory("./data/BadgerPendant.json"),
	BearPendant:          CreateItemAccessory("./data/BearPendant.json"),
	FrostAmulet:          CreateItemAccessory("./data/FrostAmulet.json"),
	MonkeyPendant:        CreateItemAccessory("./data/MonkeyPendant.json"),
	NecklaceOfPeace:      CreateItemAccessory("./data/NecklaceOfPeace.json"),
	OwlPendant:           CreateItemAccessory("./data/OwlPendant.json"),
	OxPendant:            CreateItemAccessory("./data/OxPendant.json"),
	PhoenixChoker:        CreateItemAccessory("./data/PhoenixChoker.json"),
	RatPendant:           CreateItemAccessory("./data/RatPendant.json"),
	TorqOfSoul:           CreateItemAccessory("./data/TorqOfSoul.json"),
	WindLocket:           CreateItemAccessory("./data/WindLocket.json"),
	FangsOfDeathNecklace: CreateItemAccessory("./data/FangsOfDeathNecklace.json"),
	RingOfFinesse:        CreateItemAccessory("./data/RingOfFinesse.json"),
	RingOfQuickness:      CreateItemAccessory("./data/RingOfQuickness.json"),
	RingOfSurvival:       CreateItemAccessory("./data/RingOfSurvival.json"),
	RingOfVitality:       CreateItemAccessory("./data/RingOfVitality.json"),
	RingOfWisdom:         CreateItemAccessory("./data/RingOfWisdom.json"),
	RingOfResolve:        CreateItemAccessory("./data/RingOfResolve.json"),
	TatteredCloak:        CreateItemArmor("./data/TatteredCloak.json"),
	SplendidCloak:        CreateItemArmor("./data/SplendidCloak.json"),
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
		item.LooseTrousers,
		item.LightfootBoots,
		item.RadiantCloak,
		item.AdventurerCloak,
		item.BarbutaHelm,
		item.ChapelDeFer,
		item.OccultistPants,
		item.AdventurerBoots,
		item.ArcaneHood,
		item.CeremonialHeaddress,
		item.Chaperon,
		item.Coif,
		item.DarkgroveHood,
		item.ElkwoodCrown,
		item.FeatheredHat,
		item.ForestHood,
		item.Gjermundbu,
		item.GreatHelm,
		item.Hounskull,
		item.KettleHat,
		item.LeatherBonnet,
		item.LeatherCap,
		item.NormanNasalHelm,
		item.OpenSallet,
		item.RangerHood,
		item.RogueCowl,
		item.Sallet,
		item.ShadowHood,
		item.ShadowMask,
		item.Spangenhelm,
		item.StrawHat,
		item.Topfhelm,
		item.VikingHelm,
		item.OccultistHood,
		item.AdventurerTunic,
		item.ArcaneGrab,
		item.ArcaneGloves,
		item.BardicPants,
		item.BloodwovenGloves,
		item.ChampionArmor,
		item.CrusaderArmor,
		item.BloodwovenRobe,
		item.BuckledBoots,
		item.ButtonedBoots,
		item.ButtonedLeggings,
		item.ClothPants,
		item.CuffedBoots,
		item.DarkCuirass,
		item.DarkgroveRobe,
		item.DarkleafBoots,
		item.DarkLeatherLeggings,
		item.DarkPlateArmor,
		item.DashingBoots,
		item.DreadHood,
		item.ElkwoodGloves,
		item.FineCuirass,
		item.ForestBoots,
		item.Frock,
		item.GlovesOfUtility,
		item.GrandBrigandine,
		item.WandererAttire,
		item.WatchmanCloak,
		item.WizardHat,
		item.WizardShoes,
		item.HeavyBoots,
		item.HeavyGambeson,
		item.HeavyGauntlets,
		item.HeavyLeggings,
		item.HeavyLeatherLeggings,
		item.LacedTurnshoe,
		item.LeatherGloves,
		item.LeatherLeggings,
		item.LightAketon,
		item.LightGauntlets,
		item.LowBoots,
		item.MarauderOutfit,
		item.MercurialCloak,
		item.MysticGloves,
		item.MysticVestments,
		item.NorthernFullTunic,
		item.OccultistRobe,
		item.OccultistTunic,
		item.OldShoes,
		item.Pourpoint,
		item.TemplarArmor,
		item.TatteredCloak,
		item.SplendidCloak,
	},
	ItemsWeapon: []Item_Weapon{
		item.ArmingSword,
		item.Bardiche,
	},
	ItemsAccessory: []Item_Accessory{
		item.FoxPendant,
		item.RingOfCourage,
		item.BadgerPendant,
		item.BearPendant,
		item.FrostAmulet,
		item.MonkeyPendant,
		item.NecklaceOfPeace,
		item.OwlPendant,
		item.OxPendant,
		item.PhoenixChoker,
		item.RatPendant,
		item.TorqOfSoul,
		item.WindLocket,
		item.FangsOfDeathNecklace,
		item.RingOfFinesse,
		item.RingOfQuickness,
		item.RingOfSurvival,
		item.RingOfVitality,
		item.RingOfWisdom,
		item.RingOfResolve,
	},
}
