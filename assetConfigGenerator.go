package main

import (
	"fmt"
	"math"
)

type AssetConfigGenerator struct {
	settings      AssetConfigGeneratorSettings
	counters      AssetConfigGeneratorCounters
	NamedManifest NamedManifest
	maxAmount     int16
}

func (g AssetConfigGenerator) getCategorySettings(attributeManifest AttributeManifest, totalRarities float32) []CategoryConfig {
	var counter int16 = 0
	var categoryConfigs = make([]CategoryConfig, len(attributeManifest.Categories))
	fmt.Println(attributeManifest.Categories)
	for i := 0; i < len(attributeManifest.Categories); i++ {
		category := attributeManifest.Categories[i]
		total := int16(math.Ceil(
			float64(category.Rarity) / float64(totalRarities) * float64(g.maxAmount),
		))
		categoryConfigs[i] = CategoryConfig{name: category.Name, starting: counter, total: total, ending: counter + total, rarity: category.Rarity}
		counter = counter + total
	}
	return categoryConfigs
}

func (g AssetConfigGenerator) initialiseCountersAndSettingForAttribute(attribute AttributeType) {
	var categories []Category
	var totalRarities float32
	for i := 0; i < len(categories); i++ {
		totalRarities += categories[i].Rarity
	}

	switch attribute {
	case HANDS:
		categories = g.NamedManifest.hands.Categories
		g.counters.hands = 0
		g.settings.hands.categories = g.getCategorySettings(g.NamedManifest.hands, totalRarities)
	case AURA:
		categories = g.NamedManifest.aura.Categories
		g.counters.aura = 0
		g.settings.aura.categories = g.getCategorySettings(g.NamedManifest.aura, totalRarities)
	case WATCHERS:
		categories = g.NamedManifest.watchers.Categories
		g.counters.watchers = 0
		g.settings.watchers.categories = g.getCategorySettings(g.NamedManifest.watchers, totalRarities)
	case STAIRS:
		categories = g.NamedManifest.stairs.Categories
		g.counters.stairs = 0
		g.settings.stairs.categories = g.getCategorySettings(g.NamedManifest.stairs, totalRarities)
	case ARCHES:
		categories = g.NamedManifest.arches.Categories
		g.counters.arches = 0
		g.settings.arches.categories = g.getCategorySettings(g.NamedManifest.arches, totalRarities)
	case GEMS:
		categories = g.NamedManifest.gems.Categories
		g.counters.gems = 0
		g.settings.gems.categories = g.getCategorySettings(g.NamedManifest.gems, totalRarities)
	case BLIPS:
		categories = g.NamedManifest.blips.Categories
		g.counters.blips = 0
		g.settings.blips.categories = g.getCategorySettings(g.NamedManifest.blips, totalRarities)
	default:
		fmt.Println("Unknown attribute", attribute)
	}

}

func (g AssetConfigGenerator) initialiseCountersAndSettings() {
	g.initialiseCountersAndSettingForAttribute(HANDS)
	g.initialiseCountersAndSettingForAttribute(AURA)
	g.initialiseCountersAndSettingForAttribute(WATCHERS)
	g.initialiseCountersAndSettingForAttribute(STAIRS)
	g.initialiseCountersAndSettingForAttribute(ARCHES)
	g.initialiseCountersAndSettingForAttribute(GEMS)
	g.initialiseCountersAndSettingForAttribute(BLIPS)
}

func (g AssetConfigGenerator) init(maxAmount int16, namedManifest NamedManifest) {
	g.maxAmount = maxAmount
	g.NamedManifest = namedManifest
	g.initialiseCountersAndSettings()
}
