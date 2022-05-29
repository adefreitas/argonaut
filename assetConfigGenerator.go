package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type AssetConfigGenerator struct {
	settings      AssetConfigGeneratorSettings
	NamedManifest NamedManifest
	maxAmount     int16
}

func (g *AssetConfigGenerator) getCategorySettings(categories []Category, attribute AttributeType) []CategoryConfig {
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
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(categoryConfigs), func(i, j int) {
		categoryConfigs[i], categoryConfigs[j] = categoryConfigs[j], categoryConfigs[i]
	})
	return categoryConfigs
}

func (g *AssetConfigGenerator) initialiseSettingForAttribute(attribute AttributeType) {
	switch attribute {
	case HAND_TOP_LEFT:
		g.settings.handTopLeft.categories = g.getCategorySettings(g.NamedManifest.handTopLeft.Categories, HAND_TOP_LEFT)
	case HAND_TOP_RIGHT:
		g.settings.handTopRight.categories = g.getCategorySettings(g.NamedManifest.handTopRight.Categories, HAND_TOP_RIGHT)
	case HAND_BOTTOM_LEFT:
		g.settings.handBottomLeft.categories = g.getCategorySettings(g.NamedManifest.handBottomLeft.Categories, HAND_BOTTOM_LEFT)
	case HAND_BOTTOM_RIGHT:
		g.settings.handBottomRight.categories = g.getCategorySettings(g.NamedManifest.handBottomRight.Categories, HAND_BOTTOM_RIGHT)
	case BLIPS_AURA:
		g.settings.blipAura.categories = g.getCategorySettings(g.NamedManifest.blipAura.Categories, BLIPS_AURA)
	case ELEMENTS:
		g.settings.elements.categories = g.getCategorySettings(g.NamedManifest.elements.Categories, ELEMENTS)
	case AURAS:
		g.settings.auras.categories = g.getCategorySettings(g.NamedManifest.auras.Categories, AURAS)
	case WATCHERS:
		g.settings.watchers.categories = g.getCategorySettings(g.NamedManifest.watchers.Categories, WATCHERS)
	case STAIRS:
		g.settings.stairs.categories = g.getCategorySettings(g.NamedManifest.stairs.Categories, STAIRS)
	case ARCHES:
		g.settings.arches.categories = g.getCategorySettings(g.NamedManifest.arches.Categories, ARCHES)
	case GEMS:
		g.settings.gems.categories = g.getCategorySettings(g.NamedManifest.gems.Categories, GEMS)
	case BLIPS:
		g.settings.blips.categories = g.getCategorySettings(g.NamedManifest.blips.Categories, BLIPS)
	case MUSIC:
		g.settings.music.categories = g.getCategorySettings(g.NamedManifest.music.Categories, MUSIC)
	default:
		fmt.Println("Unknown attribute", attribute)
	}

}

func (g *AssetConfigGenerator) initialiseSettings() {
	g.initialiseSettingForAttribute(HAND_BOTTOM_LEFT)
	g.initialiseSettingForAttribute(HAND_BOTTOM_RIGHT)
	g.initialiseSettingForAttribute(HAND_TOP_LEFT)
	g.initialiseSettingForAttribute(HAND_TOP_RIGHT)
	g.initialiseSettingForAttribute(ELEMENTS)
	g.initialiseSettingForAttribute(BLIPS_AURA)
	g.initialiseSettingForAttribute(AURAS)
	g.initialiseSettingForAttribute(WATCHERS)
	g.initialiseSettingForAttribute(STAIRS)
	g.initialiseSettingForAttribute(ARCHES)
	g.initialiseSettingForAttribute(GEMS)
	g.initialiseSettingForAttribute(BLIPS)
	g.initialiseSettingForAttribute(MUSIC)
}

func (g *AssetConfigGenerator) getCategoriesForAttribute(attribute AttributeType) []CategoryConfig {
	switch attribute {
	case BLIPS_AURA:
		return g.settings.blipAura.categories
	case ELEMENTS:
		return g.settings.elements.categories
	case HAND_TOP_LEFT:
		return g.settings.handTopLeft.categories
	case HAND_TOP_RIGHT:
		return g.settings.handTopRight.categories
	case HAND_BOTTOM_LEFT:
		return g.settings.handBottomLeft.categories
	case HAND_BOTTOM_RIGHT:
		return g.settings.handBottomRight.categories
	case AURAS:
		return g.settings.auras.categories
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
	case MUSIC:
		return g.settings.music.categories
	default:
		fmt.Println("Attribute not found when getting counter")
		return make([]CategoryConfig, 0)
	}
}

func getFiles(attribute string, category string) []string {
	images := make([]string, 200)
	capitalisedAttributeName := strings.Title(attribute)
	// fmt.Println("Getting files for", attribute, category)
	for i := 0; i < 200; i++ {
		frameNumber := fmt.Sprintf("%05d", i)
		images[i] = INPUT_ATTRIBUTES_DIR + "/" + capitalisedAttributeName + "/" + category + "/" + category + "_" + frameNumber + ".png"
	}
	return images
}

func (g *AssetConfigGenerator) findAttributeCategoryByCounter(attribute AttributeType, counter int16) CategoryRenderingDetails {
	var category CategoryConfig
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

func (g *AssetConfigGenerator) findAudioAttributeCategoryByCounter(attribute AttributeType, counter int16) CategoryAudioDetails {
	var category CategoryConfig
	categories := g.getCategoriesForAttribute(attribute)
	for i := 0; i < len(categories); i++ {
		if counter >= categories[i].starting && counter <= categories[i].ending {
			category = categories[i]
		}
	}
	audioInputPath := fmt.Sprintf("%s/%s/%s/%s.wav", INPUT_ATTRIBUTES_DIR, attribute, category.name, category.name)
	return CategoryAudioDetails{
		name:           category.name,
		audioInputPath: audioInputPath,
		rarity:         category.rarity,
	}
}

func (g *AssetConfigGenerator) init(maxAmount int16, namedManifest NamedManifest) {
	g.maxAmount = maxAmount
	g.NamedManifest = namedManifest
	g.initialiseSettings()
}

func (g *AssetConfigGenerator) generate(counter int16) GenerationData {
	var data ManifestData
	var frames Frames
	// TODO: Look into using https://pkg.go.dev/github.com/fatih/structs@v1.1.0?utm_source=gopls
	handTopLeft := g.findAttributeCategoryByCounter(HAND_TOP_LEFT, counter)
	handTopRight := g.findAttributeCategoryByCounter(HAND_TOP_RIGHT, counter)
	handBottomLeft := g.findAttributeCategoryByCounter(HAND_BOTTOM_LEFT, counter)
	handBottomRight := g.findAttributeCategoryByCounter(HAND_BOTTOM_RIGHT, counter)
	elements := g.findAttributeCategoryByCounter(ELEMENTS, counter)
	blipAura := g.findAttributeCategoryByCounter(BLIPS_AURA, counter)
	auras := g.findAttributeCategoryByCounter(AURAS, counter)
	watchers := g.findAttributeCategoryByCounter(WATCHERS, counter)
	stairs := g.findAttributeCategoryByCounter(STAIRS, counter)
	arches := g.findAttributeCategoryByCounter(ARCHES, counter)
	gems := g.findAttributeCategoryByCounter(GEMS, counter)
	blips := g.findAttributeCategoryByCounter(BLIPS, counter)
	music := g.findAudioAttributeCategoryByCounter(MUSIC, counter)
	data.HandTopLeft = ManifestAttributeData{Name: handTopLeft.name, Rarity: handTopLeft.rarity}
	data.HandTopRight = ManifestAttributeData{Name: handTopRight.name, Rarity: handTopRight.rarity}
	data.HandBottomLeft = ManifestAttributeData{Name: handBottomLeft.name, Rarity: handBottomLeft.rarity}
	data.HandBottomRight = ManifestAttributeData{Name: handBottomRight.name, Rarity: handBottomRight.rarity}
	data.Elements = ManifestAttributeData{Name: elements.name, Rarity: elements.rarity}
	data.BlipAura = ManifestAttributeData{Name: blipAura.name, Rarity: blipAura.rarity}
	data.Auras = ManifestAttributeData{Name: auras.name, Rarity: auras.rarity}
	data.Watchers = ManifestAttributeData{Name: watchers.name, Rarity: watchers.rarity}
	data.Stairs = ManifestAttributeData{Name: stairs.name, Rarity: stairs.rarity}
	data.Arches = ManifestAttributeData{Name: arches.name, Rarity: arches.rarity}
	data.Gems = ManifestAttributeData{Name: gems.name, Rarity: gems.rarity}
	data.Blips = ManifestAttributeData{Name: blips.name, Rarity: blips.rarity}
	data.Music = ManifestAttributeData{Name: music.name, Rarity: music.rarity}

	frames.handTopLeft = handTopLeft.files
	frames.handTopRight = handTopRight.files
	frames.handBottomLeft = handBottomLeft.files
	frames.handBottomRight = handBottomRight.files
	frames.elements = elements.files
	frames.blipAura = blipAura.files
	frames.auras = auras.files
	frames.watchers = watchers.files
	frames.stairs = stairs.files
	frames.arches = arches.files
	frames.gems = gems.files
	frames.blips = blips.files

	return GenerationData{frames: frames, data: data, audioInputPath: music.audioInputPath}
}
