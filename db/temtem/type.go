package temtem

import "github.com/uptrace/bun"

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

	Sort int `bun:"sort,notnull,nullzero" json:"-"`
}
