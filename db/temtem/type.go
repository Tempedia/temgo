package temtem

import (
	"time"

	"github.com/uptrace/bun"
)

type TemtemType struct {
	bun.BaseModel `bun:"table:temtem_type"`

	Name               string   `bun:"name,notnull,pk" json:"name"`
	Icon               string   `bun:"icon,notnull,nullzero" json:"icon"`
	Comment            string   `bun:"comment,notnull,nullzero" json:"comment"`
	Trivia             []string `bun:"trivia,array,notnull,nullzero" json:"trivia"`
	EeffectiveAgainst  []string `bun:"effective_against,array,notnull,nullzero" json:"effective_against"`
	IneffectiveAgainst []string `bun:"ineffective_against,array,notnull,nullzero" json:"ineffective_against"`
	ResistantTo        []string `bun:"resistant_to,array,notnull,nullzero" json:"resistant_to"`
	WeakTo             []string `bun:"weak_to,array,notnull,nullzero" json:"weak_to"`

	Color string `bun:"color,notnull,nullzero" json:"color"`

	Sort int `bun:"sort,notnull,nullzero" json:"sort"`
}

type TemtemGenderRatio struct {
	Male   int `json:"male" jsonb:"male"`
	Female int `json:"female" jsonb:"female"`
}

type TemtemDescription struct {
	PhysicalAppearance string `json:"Physical Appearance" jsonb:"Physical Appearance"`
	Tempedia           string `json:"Tempedia" jsonb:"Tempedia"`
}

type TemtemTVYield struct {
	HP    int `json:"HP" jsonb:"HP"`
	ATK   int `json:"ATK" jsonb:"ATK"`
	DEF   int `json:"DEF" jsonb:"DEF"`
	SPD   int `json:"SPD" jsonb:"SPD"`
	STA   int `json:"STA" jsonb:"STA"`
	SPATK int `json:"SPATK" jsonb:"SPATK"`
	SPDEF int `json:"SPDEF" jsonb:"SPDEF"`
}

type TemtemEvolvesTo struct {
	To     string `json:"to" jsonb:"to"`
	Method string `json:"method" jsonb:"method"`
	Level  int    `json:"level" jsonb:"level"`
	Gender string `json:"gender" jsonb:"gender"`
	Place  string `json:"place" jsonb:"place"`
	TV     int    `json:"tv" jsonb:"tv"`
}

type IntRange struct {
	From int `json:"from" jsonb:"from"`
	To   int `json:"to" jsonb:"to"`
}

type TemtemStatsItem struct {
	Base int      `json:"base" jsonb:"base"`
	L50  IntRange `json:"50" jsonb:"50"`
	L100 IntRange `json:"100" jsonb:"100"`
}

type TemtemStats struct {
	HP    TemtemStatsItem `json:"HP" jsonb:"HP"`
	ATK   TemtemStatsItem `json:"ATK" jsonb:"ATK"`
	DEF   TemtemStatsItem `json:"DEF" jsonb:"DEF"`
	SPD   TemtemStatsItem `json:"SPD" jsonb:"SPD"`
	STA   TemtemStatsItem `json:"STA" jsonb:"STA"`
	SPATK TemtemStatsItem `json:"SPATK" jsonb:"SPATK"`
	SPDEF TemtemStatsItem `json:"SPDEF" jsonb:"SPDEF"`
}

type TemtemGallery struct {
	Text   string `json:"text" jsonb:"text"`
	FileID string `json:"fileid" jsonb:"fileid"`
	Group  string `json:"group" jsonb:"group"`
}

type TemtemSubspecie struct {
	Type     string `json:"type" jsonb:"type"`
	Icon     string `json:"icon" jsonb:"icon"`
	LumaIcon string `json:"luma_icon" jsonb:"luma_icon"`
}

type Temtem struct {
	bun.BaseModel           `bun:"table:temtem"`
	NO                      int               `bun:"no,pk" json:"no"`
	Name                    string            `bun:"name,notnull" json:"name"`
	Type                    []string          `bun:"type,notnull,nullzero,array" json:"type"`
	CatchRate               float64           `bun:"catch_rate,notnull,nullzero" json:"catch_rate"`
	GenderRatio             TemtemGenderRatio `bun:"gender_ratio,notnull,type:jsonb" json:"gender_ratio"`
	ExperienceYieldModifier float64           `bun:"experience_yield_modifier,notnull,nullzero" json:"experience_yield_modifier"`
	Icon                    string            `bun:"icon,notnull,nullzero" json:"icon"`
	LumaIcon                string            `bun:"luma_icon,notnull,nullzero" json:"luma_icon"`
	Traits                  []string          `bun:"traits,notnull,nullzero,array" json:"traits"`
	Description             TemtemDescription `bun:"description,notnull,type:jsonb" json:"description"`
	Cry                     string            `bun:"cry,notnull,nullzero" json:"cry"`

	Height float64 `bun:"height,notnull,nullzero" json:"height"`
	Weight float64 `bun:"weight,notnull,nullzero" json:"weight"`

	TVYield TemtemTVYield `bun:"tv_yield,notnull,type:jsonb" json:"tv_yield"`

	EvolvesTo   []TemtemEvolvesTo        `bun:"evolves_to,notnull,nullzero,type:jsonb" json:"evolves_to"`
	Stats       TemtemStats              `bun:"stats,notnull,type:jsonb" json:"stats"`
	TypeMatchup []map[string]interface{} `bun:"type_matchup,notnull,nullzero,type:jsonb" json:"type_matchup"`

	Trivia []string `bun:"trivia,notnull,nullzero,array" json:"trivia"`

	Subspecies []TemtemSubspecie `bun:"subspecies,notnull,nullzero,type:jsonb" json:"subspecies"`

	Gallery []TemtemGallery `bun:"gallery,notnull,nullzero,type:jsonb" json:"gallery"`
	Renders []TemtemGallery `bun:"renders,notnull,nullzero,type:jsonb" json:"renders"`

	Subspecie *TemtemSubspecie `bun:"-" json:"subspecie,omitempty"`

	Techniques map[string]interface{} `bun:"-" json:"techniques,omitempty"`
}

type TemtemTrait struct {
	bun.BaseModel `bun:"table:temtem_trait"`
	Name          string `bun:"name,notnull,pk" json:"name"`
	Description   string `bun:"description,notnull,nullzero" json:"description"`
	Impact        string `bun:"impact,notnull,nullzero" json:"impact"`
	Trigger       string `bun:"trigger,notnull,nullzero" json:"trigger"`
	Effect        string `bun:"effect,notnull,nullzero" json:"effect"`
}

type TemtemTechnique struct {
	bun.BaseModel `bun:"table:temtem_technique"`
	Name          string `bun:"name,notnull,pk" json:"name"`
	Type          string `bun:"type,notnull" json:"type"`
	Class         string `bun:"class,notnull" json:"class"`
	Damage        int    `bun:"damage,notnull,nullzero" json:"damage"`
	STACost       int    `bun:"sta_cost,notnull,nullzero" json:"sta_cost"`
	Hold          int    `bun:"hold,notnull,nullzero" json:"hold"`
	Priority      int    `bun:"priority,notnull,nullzero" json:"priority"`
	Targeting     string `bun:"targeting,notnull,nullzero" json:"targeting"`
	Description   string `bun:"description,notnull,nullzero" json:"description"`
	Video         string `bun:"video,notnull,nullzero" json:"video"`

	SynergyType        string `bun:"synergy_type,notnull,nullzero" json:"synergy_type"`
	SynergyDescription string `bun:"synergy_description,notnull,nullzero" json:"synergy_description"`
	SynergyEffects     string `bun:"synergy_effects,notnull,nullzero" json:"synergy_effects"`
	SynergyDamage      int    `bun:"synergy_damage,notnull,nullzero" json:"synergy_damage"`
	SynergySTACost     int    `bun:"synergy_sta_cost,notnull,nullzero" json:"synergy_sta_cost"`
	SynergyPriority    int    `bun:"synergy_priority,notnull,nullzero" json:"synergy_priority"`
	SynergyTargeting   string `bun:"synergy_targeting,notnull,nullzero" json:"synergy_targeting"`
	SynergyVideo       string `bun:"synergy_video,notnull,nullzero" json:"synergy_video"`
}

type TemtemCourseTechnique struct {
	bun.BaseModel `bun:"table:temtem_course_technique"`
	ID            int64  `bun:"id,notnull,pk" json:"-"`
	Temtem        string `bun:"temtem,notnull" json:"-"`
	Stab          bool   `bun:"stab,notnull,nullzero" json:"stab"`
	Course        string `bun:"course,notnull" json:"course"`
	TechniqueName string `bun:"technique_name,notnull" json:"technique_name"`

	Technique    *TemtemTechnique `bun:"rel:belongs-to,join:technique_name=name" json:"technique"`
	TemtemObject *Temtem          `bun:"rel:belongs-to,join:temtem=name" json:"temtem"`
}

type TemtemLevelingUpTechnique struct {
	bun.BaseModel `bun:"table:temtem_leveling_up_technique"`
	ID            int64  `bun:"id,notnull,pk" json:"-"`
	Temtem        string `bun:"temtem,notnull" json:"-"`
	Stab          bool   `bun:"stab,notnull,nullzero" json:"stab"`
	Level         int    `bun:"level,notnull" json:"level"`
	TechniqueName string `bun:"technique_name,notnull" json:"technique_name"`

	Group string `bun:"group,notnull,nullzero" json:"group"`

	Technique    *TemtemTechnique `bun:"rel:belongs-to,join:technique_name=name" json:"technique"`
	TemtemObject *Temtem          `bun:"rel:belongs-to,join:temtem=name" json:"temtem"`
}

type TemtemBreedingTechniqueParent struct {
	Name string `json:"name" jsonb:"name"`
	Hint string `json:"hint" jsonb:"hint"`
}

type TemtemBreedingTechnique struct {
	bun.BaseModel `bun:"table:temtem_breeding_technique"`
	ID            int64                           `bun:"id,notnull,pk" json:"-"`
	Temtem        string                          `bun:"temtem,notnull" json:"-"`
	Stab          bool                            `bun:"stab,notnull,nullzero" json:"stab"`
	Parents       []TemtemBreedingTechniqueParent `bun:"parents,notnull,type:jsonb" json:"parents"`
	TechniqueName string                          `bun:"technique_name,notnull" json:"technique_name"`

	Technique    *TemtemTechnique `bun:"rel:belongs-to,join:technique_name=name" json:"technique"`
	TemtemObject *Temtem          `bun:"rel:belongs-to,join:temtem=name" json:"temtem"`
}

type TemtemLocation struct {
	bun.BaseModel      `bun:"table:temtem_location"`
	Name               string   `bun:"name,notnull,pk" json:"name"`
	Description        string   `bun:"description,notnull,nullzero" json:"description"`
	Island             string   `bun:"island,notnull,nullzero" json:"island"`
	Image              string   `bun:"image,notnull,nullzero" json:"image"`
	Comment            string   `bun:"comment,notnull,nullzero" json:"comment"`
	ConnectedLocations []string `bun:"connected_locations,notnull,nullzero,array" json:"connected_locations"`
}

type TemtemLocationAreaTemtem struct {
	Name string `json:"name" jsonb:"name"`
	Odds []struct {
		Odds string `json:"odds" jsonb:"odds"`
		Desc string `json:"desc" jsonb:"desc"`
	} `json:"odds" jsonb:"odds"`
	Level struct {
		From int  `json:"from" jsonb:"from"`
		To   int  `json:"to" jsonb:"to"`
		Egg  bool `json:"egg" jsonb:"egg"`
	} `json:"level" jsonb:"level"`
}

type TemtemLocationArea struct {
	bun.BaseModel `bun:"table:temtem_location_area"`
	ID            int64                      `bun:"id,notnull,pk" json:"-"`
	Name          string                     `bun:"name,notnull" json:"name"`
	Location      string                     `bun:"location,notnull" json:"location"`
	Image         string                     `bun:"image,notnull,nullzero" json:"image"`
	Temtems       []TemtemLocationAreaTemtem `bun:"temtems,notnull,nullzero,type:jsonb" json:"temtems"`
}

type TemtemStatusCondition struct {
	bun.BaseModel `bun:"table:temtem_status_condition"`
	// ID            int64  `bun:"id,notnull,pk" json:"-"`
	Name        string `bun:"name,notnull,pk" json:"name"`
	Icon        string `bun:"icon,notnull,nullzero" json:"icon"`
	Description string `bun:"description,notnull,nullzero" json:"description"`
	Group       string `bun:"group,notnull,nullzero" json:"group"`

	Techniques []string `bun:"techniques,notnull,nullzero,array" json:"techniques"`
	Traits     []string `bun:"traits,notnull,nullzero,array" json:"traits"`
}

type TemtemCourseItem struct {
	bun.BaseModel `bun:"table:temtem_course_item"`
	NO            string `bun:"no,notnull,pk" json:"no"`
	TechniqueName string `bun:"technique,notnull,nullzero" json:"-"`
	Source        string `bun:"source,notnull,nullzero" json:"source"`

	Technique *TemtemTechnique `bun:"rel:belongs-to,join:technique=name" json:"technique"`
}

type TemtemItemCategory struct {
	bun.BaseModel `bun:"table:temtem_item_category"`
	Name          string `bun:"name,notnull,pk" json:"name"`
	ParentName    string `bun:"parent,notnull,nullzero" json:"-"`
	Sort          int    `bun:"sort,notnull,nullzero" json:"-"`

	Parent *TemtemItemCategory `bun:"rel:belongs-to,join:parent=name" json:"parent,omitempty"`
}

type TemtemItemExtra struct {
	Source       string `json:"Source,omitempty" jsonb:"Source"`
	Location     string `json:"Location,omitempty" jsonb:"Location"`
	CaptureBonus string `json:"Capture Bonus,omitempty" jsonb:"Capture Bonus"`
	Quest        string `json:"Quest,omitempty" jsonb:"Quest"`
}

type TemtemItem struct {
	bun.BaseModel `bun:"table:temtem_item"`
	Name          string          `bun:"name,notnull,pk" json:"name"`
	Icon          string          `bun:"icon,notnull,nullzero" json:"icon"`
	Description   string          `bun:"description,notnull,nullzero" json:"description"`
	Tradable      bool            `bun:"tradable,notnull,nullzero" json:"tradable"`
	BuyPrice      string          `bun:"buy_price,notnull,nullzero" json:"buy_price"`
	SellPrice     string          `bun:"sell_price,notnull,nullzero" json:"sell_price"`
	CategoryName  string          `bun:"category,notnull" json:"-"`
	Extra         TemtemItemExtra `bun:"extra,notnull,nullzero,type:jsonb" json:"extra"`
	Sort          int             `bun:"sort,notnull,nullzero" json:"-"`

	Category *TemtemItemCategory `bun:"rel:belongs-to,join:category=name" json:"category,omitempty"`
}

type TemtemUserTeam struct {
	bun.BaseModel `bun:"table:temtem_user_team"`
	ID            string      `bun:"id,pk,default:gen_random_uuid()" json:"id"`
	Name          string      `bun:"name,notnull" json:"name"`
	Temtems       interface{} `bun:"temtems,notnull,type:jsonb" json:"temtems"`

	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
}
