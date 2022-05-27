package main

import (
	"fmt"
	"math"
	"strings"
)

type AssetConfigGenerator struct {
	settings      AssetConfigGeneratorSettings
	counters      AssetConfigGeneratorCounters
	NamedManifest NamedManifest
	maxAmount     int16
}

func (g *AssetConfigGenerator) getCategorySettings(categories []Category) []CategoryConfig {
	var counter int16 = 0
	var categoryConfigs = make([]CategoryConfig, len(categories))
	var totalRarities float32
	for i := 0; i < len(categories); i++ {
		totalRarities += categories[i].Rarity
	}
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

func (g *AssetConfigGenerator) restartCountersIfNeeded() {
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

func (g *AssetConfigGenerator) initialiseCountersAndSettingForAttribute(attribute AttributeType) {
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

func (g *AssetConfigGenerator) initialiseCountersAndSettings() {
	g.initialiseCountersAndSettingForAttribute(HANDS)
	g.initialiseCountersAndSettingForAttribute(AURA)
	g.initialiseCountersAndSettingForAttribute(WATCHERS)
	g.initialiseCountersAndSettingForAttribute(STAIRS)
	g.initialiseCountersAndSettingForAttribute(ARCHES)
	g.initialiseCountersAndSettingForAttribute(GEMS)
	g.initialiseCountersAndSettingForAttribute(BLIPS)
}

func getNextAttribute(attribute AttributeType) AttributeType {
	switch attribute {
	case HANDS:
		return AURA
	case AURA:
		return WATCHERS
	case WATCHERS:
		return STAIRS
	case STAIRS:
		return ARCHES
	case ARCHES:
		return GEMS
	case GEMS:
		return BLIPS
	case BLIPS:
		return HANDS
	default:
		fmt.Println("Attribute not found when getting next attribute")
		return HANDS
	}
}

func (g *AssetConfigGenerator) getCounterForAttribute(attribute AttributeType) int16 {
	switch attribute {
	case HANDS:
		return g.counters.hands
	case AURA:
		return g.counters.aura
	case WATCHERS:
		return g.counters.watchers
	case STAIRS:
		return g.counters.stairs
	case ARCHES:
		return g.counters.arches
	case GEMS:
		return g.counters.gems
	case BLIPS:
		return g.counters.blips
	default:
		fmt.Println("Attribute not found when getting counter")
		return 0
	}
}

func (g *AssetConfigGenerator) getCategoriesForAttribute(attribute AttributeType) []CategoryConfig {
	switch attribute {
	case HANDS:
		return g.settings.hands.categories
	case AURA:
		return g.settings.aura.categories
	case WATCHERS:
		return g.settings.watchers.categories
	case STAIRS:
		return g.settings.stairs.categories
	case ARCHES:
		return g.settings.arches.categories
	case GEMS:
		return g.settings.gems.categories
	case BLIPS:
		return g.settings.blips.categories
	default:
		fmt.Println("Attribute not found when getting counter")
		return make([]CategoryConfig, 0)
	}
}

func (g *AssetConfigGenerator) incrementCounterForAttribute(attribute AttributeType) {
	switch attribute {
	case HANDS:
		g.counters.hands++
	case AURA:
		g.counters.aura++
	case WATCHERS:
		g.counters.watchers++
	case STAIRS:
		g.counters.stairs++
	case ARCHES:
		g.counters.arches++
	case GEMS:
		g.counters.gems++
	case BLIPS:
		g.counters.blips++
	default:
		fmt.Println("Attribute not found when increasing counter")
	}
}

func (g *AssetConfigGenerator) setCounterForAttribute(attribute AttributeType, value int16) {
	switch attribute {
	case HANDS:
		g.counters.hands = value
	case AURA:
		g.counters.aura = value
	case WATCHERS:
		g.counters.watchers = value
	case STAIRS:
		g.counters.stairs = value
	case ARCHES:
		g.counters.arches = value
	case GEMS:
		g.counters.gems = value
	case BLIPS:
		g.counters.blips = value
	default:
		fmt.Println("Attribute not found when increasing counter")
	}
}

func (g *AssetConfigGenerator) updateCounters() {
	g.counters.hands++
	g.counters.aura++
	g.counters.watchers++
	g.counters.stairs++
	g.counters.arches++
	g.counters.gems++
	g.counters.blips++
	g.counters.arches++
	g.restartCountersIfNeeded()
}

func getFiles(attribute string, category string) []string {
	images := make([]string, 200)
	capitalisedAttributeName := strings.Title(attribute)
	for i := 0; i < 200; i++ {
		frameNumber := fmt.Sprintf("%05d", i)
		images[i] = INPUT_ATTRIBUTES_DIR + "/" + capitalisedAttributeName + "/" + category + "/" + category + "_" + frameNumber + ".png"
	}
	return images
}

func (g *AssetConfigGenerator) findAttributeCategoryByCounter(attribute AttributeType) CategoryRenderingDetails {
	var category CategoryConfig
	counter := g.getCounterForAttribute(attribute)
	categories := g.getCategoriesForAttribute(attribute)
	for i := 0; i < len(categories); i++ {
		if counter >= categories[i].starting && counter <= categories[i].ending {
			category = categories[i]
		}
	}
	return CategoryRenderingDetails{
		name:   category.name,
		files:  getFiles(string(attribute), category.name),
		rarity: category.rarity,
	}
}

func (g *AssetConfigGenerator) init(maxAmount int16, namedManifest NamedManifest) {
	g.maxAmount = maxAmount
	g.NamedManifest = namedManifest
	g.initialiseCountersAndSettings()
}

func (g *AssetConfigGenerator) generate() GenerationData {
	var data ManifestData
	var frames Frames
	// TODO: Look into using https://pkg.go.dev/github.com/fatih/structs@v1.1.0?utm_source=gopls
	hands := g.findAttributeCategoryByCounter(HANDS)
	aura := g.findAttributeCategoryByCounter(AURA)
	watchers := g.findAttributeCategoryByCounter(WATCHERS)
	stairs := g.findAttributeCategoryByCounter(STAIRS)
	arches := g.findAttributeCategoryByCounter(ARCHES)
	gems := g.findAttributeCategoryByCounter(GEMS)
	blips := g.findAttributeCategoryByCounter(BLIPS)
	data.Hands = ManifestAttributeData{Name: hands.name, Rarity: hands.rarity}
	data.Aura = ManifestAttributeData{Name: aura.name, Rarity: aura.rarity}
	data.Watchers = ManifestAttributeData{Name: watchers.name, Rarity: watchers.rarity}
	data.Stairs = ManifestAttributeData{Name: stairs.name, Rarity: stairs.rarity}
	data.Arches = ManifestAttributeData{Name: arches.name, Rarity: arches.rarity}
	data.Gems = ManifestAttributeData{Name: gems.name, Rarity: gems.rarity}
	data.Blips = ManifestAttributeData{Name: blips.name, Rarity: blips.rarity}

	frames.hands = hands.files
	frames.aura = aura.files
	frames.watchers = watchers.files
	frames.stairs = stairs.files
	frames.arches = arches.files
	frames.gems = gems.files
	frames.blips = blips.files

	g.updateCounters()
	return GenerationData{frames: frames, data: data}
}
