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

func (g AssetConfigGenerator) getCategorySettings(categories []Category) []CategoryConfig {
	var counter int16 = 0
	var categoryConfigs = make([]CategoryConfig, len(categories))
	var totalRarities float32
	for i := 0; i < len(categories); i++ {
		totalRarities += categories[i].Rarity
	}
	fmt.Println(categories)
	for i := 0; i < len(categories); i++ {
		category := categories[i]
		total := int16(math.Ceil(
			float64(category.Rarity) / float64(totalRarities) * float64(g.maxAmount),
		))
		categoryConfigs[i] = CategoryConfig{
			name:     category.Name,
			starting: counter,
			total:    total,
			ending:   counter + total,
			rarity:   category.Rarity,
		}
		counter = counter + total
	}
	return categoryConfigs
}

func (g AssetConfigGenerator) restartCountersIfNeeded() {
	haveAllCountersReachedTheirMaximum := g.counters.arches == g.maxAmount &&
		g.counters.aura == g.maxAmount &&
		g.counters.blips == g.maxAmount &&
		g.counters.gems == g.maxAmount &&
		g.counters.hands == g.maxAmount &&
		g.counters.stairs == g.maxAmount &&
		g.counters.watchers == g.maxAmount

	if haveAllCountersReachedTheirMaximum {
		g.counters.aura = 0
		g.counters.blips = 0
		g.counters.gems = 0
		g.counters.hands = 0
		g.counters.stairs = 0
		g.counters.watchers = 0
	}
}

func (g AssetConfigGenerator) initialiseCountersAndSettingForAttribute(attribute AttributeType) {

	switch attribute {
	case HANDS:
		g.settings.hands.categories = g.getCategorySettings(g.NamedManifest.hands.Categories)
	case AURA:
		g.settings.aura.categories = g.getCategorySettings(g.NamedManifest.aura.Categories)
	case WATCHERS:
		g.settings.watchers.categories = g.getCategorySettings(g.NamedManifest.watchers.Categories)
	case STAIRS:
		g.settings.stairs.categories = g.getCategorySettings(g.NamedManifest.stairs.Categories)
	case ARCHES:
		g.settings.arches.categories = g.getCategorySettings(g.NamedManifest.arches.Categories)
	case GEMS:
		g.settings.gems.categories = g.getCategorySettings(g.NamedManifest.gems.Categories)
	case BLIPS:
		g.settings.blips.categories = g.getCategorySettings(g.NamedManifest.blips.Categories)
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
