package main

type AttributeType string

const (
	HANDS    AttributeType = "Hands"
	AURA                   = "Aura"
	WATCHERS               = "Watchers"
	STAIRS                 = "Stairs"
	ARCHES                 = "Arches"
	GEMS                   = "Gems"
	BLIPS                  = "Blips"
)

type Category struct {
	Name   string  `json:"name"`
	Rarity float32 `json:"rarity"`
}

type AttributeManifest struct {
	Attribute  AttributeType `json:"attribute"`
	Categories []Category    `json:"categories"`
}

type Manifest []AttributeManifest

type NamedManifest struct {
	hands    AttributeManifest
	aura     AttributeManifest
	watchers AttributeManifest
	stairs   AttributeManifest
	arches   AttributeManifest
	gems     AttributeManifest
	blips    AttributeManifest
}

type CategoryConfig struct {
	name     string
	rarity   float32
	starting int16
	ending   int16
	total    int16
}

type CategoryRenderingDetails struct {
	name   string
	files  []string
	rarity float32
}

type AssetConfigGeneratorSetting struct {
	categories []CategoryConfig
}

type AssetConfigGeneratorSettings struct {
	hands    AssetConfigGeneratorSetting
	aura     AssetConfigGeneratorSetting
	watchers AssetConfigGeneratorSetting
	stairs   AssetConfigGeneratorSetting
	arches   AssetConfigGeneratorSetting
	gems     AssetConfigGeneratorSetting
	blips    AssetConfigGeneratorSetting
}

type AssetConfigGeneratorCounters struct {
	hands    int16
	aura     int16
	watchers int16
	stairs   int16
	arches   int16
	gems     int16
	blips    int16
}

type AttributeFrames []string

type Frames struct {
	hands    AttributeFrames
	aura     AttributeFrames
	watchers AttributeFrames
	stairs   AttributeFrames
	arches   AttributeFrames
	gems     AttributeFrames
	blips    AttributeFrames
}

type ManifestAttributeData struct {
	name   string
	rarity float32
}

type ManifestData struct {
	hands    ManifestAttributeData
	aura     ManifestAttributeData
	watchers ManifestAttributeData
	stairs   ManifestAttributeData
	arches   ManifestAttributeData
	gems     ManifestAttributeData
	blips    ManifestAttributeData
}

type GenerationData struct {
	frames Frames
	data   ManifestData
}
