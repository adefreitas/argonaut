package main

type AttributeType string

const (
	AURAS             AttributeType = "00_Auras"
	WATCHERS                        = "01_Watchers"
	GEMS                            = "02_Gems"
	STAIRS                          = "03_Stairs"
	BLIPS                           = "05_Blips"
	BLIPS_AURA                      = "06_Blip_Aura"
	ARCHES                          = "07_Arches"
	HAND_TOP_LEFT                   = "07_Hand_Top_Left"
	HAND_TOP_RIGHT                  = "08_Hand_Top_Right"
	HAND_BOTTOM_LEFT                = "09_Hand_Bottom_Left"
	HAND_BOTTOM_RIGHT               = "10_Hand_Bottom_Right"
	ELEMENTS                        = "11_Elements"
	MUSIC                           = "12_Music"
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
	auras           AttributeManifest
	watchers        AttributeManifest
	gems            AttributeManifest
	stairs          AttributeManifest
	blips           AttributeManifest
	blipAura        AttributeManifest
	arches          AttributeManifest
	handTopLeft     AttributeManifest
	handTopRight    AttributeManifest
	handBottomLeft  AttributeManifest
	handBottomRight AttributeManifest
	elements        AttributeManifest
	music           AttributeManifest
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

type CategoryAudioDetails struct {
	name           string
	audioInputPath string
	rarity         float32
}

type AssetConfigGeneratorSetting struct {
	categories []CategoryConfig
}

type AssetConfigGeneratorSettings struct {
	auras           AssetConfigGeneratorSetting
	watchers        AssetConfigGeneratorSetting
	gems            AssetConfigGeneratorSetting
	stairs          AssetConfigGeneratorSetting
	blips           AssetConfigGeneratorSetting
	blipAura        AssetConfigGeneratorSetting
	arches          AssetConfigGeneratorSetting
	handTopLeft     AssetConfigGeneratorSetting
	handTopRight    AssetConfigGeneratorSetting
	handBottomLeft  AssetConfigGeneratorSetting
	handBottomRight AssetConfigGeneratorSetting
	elements        AssetConfigGeneratorSetting
	music           AssetConfigGeneratorSetting
}

type AttributeFrames []string

type Frames struct {
	auras           AttributeFrames
	watchers        AttributeFrames
	stairs          AttributeFrames
	blips           AttributeFrames
	blipAura        AttributeFrames
	arches          AttributeFrames
	handTopLeft     AttributeFrames
	handTopRight    AttributeFrames
	handBottomLeft  AttributeFrames
	handBottomRight AttributeFrames
	gems            AttributeFrames
	elements        AttributeFrames
}

type ManifestAttributeData struct {
	Name   string  `json:"name"`
	Rarity float32 `json:"rariry"`
}

type ManifestData struct {
	Auras           ManifestAttributeData
	Watchers        ManifestAttributeData
	Stairs          ManifestAttributeData
	Blips           ManifestAttributeData
	BlipAura        ManifestAttributeData
	Arches          ManifestAttributeData
	HandTopLeft     ManifestAttributeData
	HandTopRight    ManifestAttributeData
	HandBottomLeft  ManifestAttributeData
	HandBottomRight ManifestAttributeData
	Gems            ManifestAttributeData
	Elements        ManifestAttributeData
	Music           ManifestAttributeData
}

type GenerationData struct {
	frames         Frames
	data           ManifestData
	audioInputPath string
}
