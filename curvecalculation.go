package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func CalculateComputedValues(stats_compute Stats, rating_armor int, rating_speed int, enchant Computed_Stats, enchantacc Computed_Stats, baseitem Computed_Stats) Computed_Stats {

	var result = Computed_Stats{

		Health:                  runcurvehybrid(stats_compute.Strength, stats_compute.Vigor, "./calc/curvemaxhealth.json"),
		MemoryCapacity:          runcurve(stats_compute.Knowledge, "./calc/curvememorycapacity.json"),
		HealthRecovery:          runcurve(stats_compute.Vigor, "./calc/curvehealthrecovery.json") * 100,                                                  //calc display adjusment
		ActionSpeed:             runcurvehybrid(stats_compute.Agility, stats_compute.Dexterity, "./calc/curvesactionspeed.json") * 100,                   //calc display adjusment
		RegularInteractionSpeed: runcurvehybridx(stats_compute.Agility, stats_compute.Resourcefulness, "./calc/curveregularinteractionspeed.json") * 100, //calc display adjusment
		MoveSpeed:               runcurve(stats_compute.Agility, "./calc/curvemovespeed.json") + 300 + float64(rating_speed),                             //calc display adjusment
		PhysicalPower:           runcurve(stats_compute.Strength+int(enchant.PhysicalPower+enchantacc.PhysicalPower), "./calc/curvephysicalpower.json"),
		MagicalPower:            runcurve(stats_compute.Will+int(enchant.MagicalPower+enchantacc.MagicalPower), "./calc/curvemagicalpower.json"),                          //calc display adjusment
		ManualDexterity:         runcurve(stats_compute.Dexterity, "./calc/curvemanualdexterity.json") * 100,                                                              //calc display adjusment
		EquipSpeed:              runcurve(stats_compute.Dexterity, "./calc/curveequipspeed.json") * 100,                                                                   //calc display adjusment
		BuffDuration:            runcurve(stats_compute.Will, "./calc/curvebuffduration.json") * 100,                                                                      //calc display adjusment
		DebuffDuration:          runcurve(stats_compute.Will, "./calc/curvedebuffduration.json") * 100,                                                                    //calc display adjusment
		MagicRating:             runcurve(stats_compute.Will, "./calc/curvemagicresistance.json") + (enchant.MagicRating + enchantacc.MagicRating + baseitem.MagicRating), //calc display adjusment
		SpellRecovery:           runcurve(stats_compute.Knowledge, "./calc/curvespellrecovery.json") * 100,
		SpellCastingSpeed:       runcurve(stats_compute.Knowledge, "./calc/curvespellcastingspeed.json") * 100,
		MagicalInteractionSpeed: runcurve(stats_compute.Will, "./calc/curvemagicalinteractionspeed.json") * 100,
		Persuasiveness:          runcurve(stats_compute.Resourcefulness, "./calc/curvepersuasiveness.json"),
		CooldownReduction:       runcurve(stats_compute.Resourcefulness, "./calc/curvecooldownreduction.json") * 100,
		PhysicalDamageReduction: runcurve(rating_armor+(enchant.FromArmorRating+enchantacc.FromArmorRating), "./calc/curvephysicaldamagereduction.json") * 100, // 0 defense value without equipement
	}

	result.PhysicalPowerBonus = runcurve(int(result.PhysicalPower), "./calc/curvephysicaldamage.json")*100 + enchant.PhysicalPowerBonus
	result.MagicalPowerBonus = runcurve(int(result.MagicalPower), "./calc/curvemagicaldamage.json")*100 + enchant.MagicalPowerBonus
	result.Health = result.Health + (result.Health * (enchant.MaxHealthBonus + enchantacc.MaxHealthBonus) / 100)
	result.MagicalDamageReduction = runcurve(int(result.MagicRating), "./calc/curvemagicaldamagereduction2.json") * 100
	result.MoveSpeedBonus = result.MoveSpeed / 3
	result.FromArmorRating = rating_armor
	result.MemoryCapacity = result.MemoryCapacity + (result.MemoryCapacity * (enchant.MemoryCapacityBonus + enchantacc.MemoryCapacityBonus) / 100)

	return result

}

// struct for curve values
type Key struct {
	Time  float64 `json:"Time"`
	Value float64 `json:"Value"`
}

// struct for curve values
type Curve struct {
	Name string `json:"Name"`
	Keys []Key  `json:"Keys"`
}

// CARGA LA CURVA JSON
func LoadCurveData(filename string) (Curve, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Curve{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Curve{}, fmt.Errorf("failed to read file: %w", err)
	}

	var curve Curve
	err = json.Unmarshal(data, &curve)
	if err != nil {
		return Curve{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return curve, nil
}

// EXTRAE VALOR DE CURVA
func GetCurveValue(curve Curve, rating float64) float64 {
	if len(curve.Keys) == 0 {
		return 0 // Return 0 if there's no data
	}

	// Handle edge cases where rating is out of bounds
	if rating <= curve.Keys[0].Time {
		return curve.Keys[0].Value
	}
	if rating >= curve.Keys[len(curve.Keys)-1].Time {
		return curve.Keys[len(curve.Keys)-1].Value
	}

	// Find the correct interval in the curve keys
	for i := 0; i < len(curve.Keys)-1; i++ {
		if rating >= curve.Keys[i].Time && rating <= curve.Keys[i+1].Time {
			// Linear interpolation between Keys[i] and Keys[i+1]
			t1, v1 := curve.Keys[i].Time, curve.Keys[i].Value
			t2, v2 := curve.Keys[i+1].Time, curve.Keys[i+1].Value
			// Interpolate
			return v1 + (rating-t1)*(v2-v1)/(t2-t1)
		}
	}

	return 0 // Default case, should not reach here
}

// EJECUTA CALCULO CURVA HYBRIDA 2
func runcurvehybridx(varstatone int, varstattwo int, jsonfile string) float64 {
	curve, err := LoadCurveData(jsonfile)
	if err != nil {
		fmt.Println("Error loading curve data:", err)
		return 0
	}
	rating := float64(varstatone)*0.4 + float64(varstattwo)*0.6
	resultcurve := GetCurveValue(curve, rating)
	return resultcurve
}

// EJECUTA CALCULO CURVA HYBRIDA
func runcurvehybrid(varstatone int, varstattwo int, jsonfile string) float64 {
	curve, err := LoadCurveData(jsonfile)
	if err != nil {
		fmt.Println("Error loading curve data:", err)
		return 0
	}
	rating := float64(varstatone)*0.25 + float64(varstattwo)*0.75
	resultcurve := GetCurveValue(curve, rating)
	return resultcurve
}

// EJECUTA CALCULO CURVA
func runcurve(varstatone int, jsonfile string) float64 {
	curve, err := LoadCurveData(jsonfile)
	if err != nil {
		fmt.Println("Error loading curve data:", err)
		return 0
	}
	resultcurve := GetCurveValue(curve, float64(varstatone))
	return resultcurve
}
