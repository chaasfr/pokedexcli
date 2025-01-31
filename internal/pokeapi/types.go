package pokeapi

const LocationAPIBaseUrl = "https://pokeapi.co/api/v2/location-area/"
const PokemonAPIBaseUrl = "https://pokeapi.co/api/v2/pokemon/"

type LocationBulkResult struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResult struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int64                 `json:"game_index"`
	ID                   int64                 `json:"id"`
	Location             Location              `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod Location                           `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64    `json:"rate"`
	Version Location `json:"version"`
}

type Name struct {
	Language Location `json:"language"`
	Name     string   `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        Location                        `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int64             `json:"max_chance"`
	Version          Location          `json:"version"`
}

type EncounterDetail struct {
	Chance          int64         `json:"chance"`
	ConditionValues []interface{} `json:"condition_values"`
	MaxLevel        int64         `json:"max_level"`
	Method          Location      `json:"method"`
	MinLevel        int64         `json:"min_level"`
}

type Pokemon struct {
	ID                     int64       `json:"id"`
	Name                   string      `json:"name"`
	BaseExperience         int64       `json:"base_experience"`
	Height                 int64       `json:"height"`
	IsDefault              bool        `json:"is_default"`
	Order                  int64       `json:"order"`
	Weight                 int64       `json:"weight"`
	Abilities              []Ability   `json:"abilities"`
	Forms                  []Species   `json:"forms"`
	GameIndices            []GameIndex `json:"game_indices"`
	HeldItems              []HeldItem  `json:"held_items"`
	LocationAreaEncounters string      `json:"location_area_encounters"`
	Moves                  []Move      `json:"moves"`
	Species                Species     `json:"species"`
	Sprites                Sprites     `json:"sprites"`
	Cries                  Cries       `json:"cries"`
	Stats                  []Stat      `json:"stats"`
	Types                  []Type      `json:"types"`
	PastTypes              []PastType  `json:"past_types"`
}

type Ability struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int64   `json:"slot"`
	Ability  Species `json:"ability"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type GameIndex struct {
	GameIndex int64   `json:"game_index"`
	Version   Species `json:"version"`
}

type HeldItem struct {
	Item           Species         `json:"item"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type VersionDetail struct {
	Rarity  int64   `json:"rarity"`
	Version Species `json:"version"`
}

type Move struct {
	Move                Species              `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int64   `json:"level_learned_at"`
	VersionGroup    Species `json:"version_group"`
	MoveLearnMethod Species `json:"move_learn_method"`
}

type PastType struct {
	Generation Species `json:"generation"`
	Types      []Type  `json:"types"`
}

type Type struct {
	Slot int64   `json:"slot"`
	Type Species `json:"type"`
}

type GenerationV struct {
	BlackWhite Sprites `json:"black-white"`
}

type GenerationIv struct {
	DiamondPearl        Sprites `json:"diamond-pearl"`
	HeartgoldSoulsilver Sprites `json:"heartgold-soulsilver"`
	Platinum            Sprites `json:"platinum"`
}

type Versions struct {
	GenerationI    GenerationI     `json:"generation-i"`
	GenerationIi   GenerationIi    `json:"generation-ii"`
	GenerationIii  GenerationIii   `json:"generation-iii"`
	GenerationIv   GenerationIv    `json:"generation-iv"`
	GenerationV    GenerationV     `json:"generation-v"`
	GenerationVi   map[string]Home `json:"generation-vi"`
	GenerationVii  GenerationVii   `json:"generation-vii"`
	GenerationViii GenerationViii  `json:"generation-viii"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Sprites         `json:"showdown"`
}

type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
	Other            *Other      `json:"other,omitempty"`
	Versions         *Versions   `json:"versions,omitempty"`
	Animated         *Sprites    `json:"animated,omitempty"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  RedBlue `json:"yellow"`
}

type RedBlue struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Crystal `json:"gold"`
	Silver  Crystal `json:"silver"`
}

type Crystal struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIii struct {
	Emerald          OfficialArtwork `json:"emerald"`
	FireredLeafgreen Crystal         `json:"firered-leafgreen"`
	RubySapphire     Crystal         `json:"ruby-sapphire"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Home struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type GenerationVii struct {
	Icons             DreamWorld `json:"icons"`
	UltraSunUltraMoon Home       `json:"ultra-sun-ultra-moon"`
}

type DreamWorld struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}

type GenerationViii struct {
	Icons DreamWorld `json:"icons"`
}

type Stat struct {
	BaseStat int64   `json:"base_stat"`
	Effort   int64   `json:"effort"`
	Stat     Species `json:"stat"`
}
