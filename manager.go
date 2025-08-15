package main

type ItemManager interface {
	CreateItemArmor(item string) Item_Armor
	CreateItemAccessory(item string) Item_Accessory
	CreateItemWeapon(item string) Item_Weapon
}

func (m Item_Armor) CreateItemArmor(item string) Item_Armor {
	return Item_Armor{}
}

func (m Item_Accessory) CreateItemAccessory(item string) Item_Accessory {
	return Item_Accessory{}
}
func (m Item_Weapon) CreateItemWeapon(item string) Item_Weapon {
	return Item_Weapon{}
}
