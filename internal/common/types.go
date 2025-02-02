package types

import "database/sql"

type Weapon struct {
	AttackSpeed float32 `json:"attackSpeed"`
	Name        string  `json:"name"`
	Damage      int     `json:"damage"`
	S1          *int32  `json:"s1,omitempty"`
	S2          *int32  `json:"s2,omitempty"`
	S3          *int32  `json:"s3,omitempty"`
	S4          *int32  `json:"s4,omitempty"`
	S5          *int32  `json:"s5,omitempty"`
}

type Enemy struct {
	Name   string `json:"name"`
	Health int    `json:"health"`
	Armor  int    `json:"armor"`
}

type Blueprint struct {
	Name         string        `json:"name"`
	Requirements []Requirement `json:"requirements"`
	Id           int           `json:"id"`
	Next         sql.NullInt32 `json:"next"`
}

type Requirement struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Quantity int    `json:"quantity"`
}

type Trade struct {
	GiveName     string `json:"giveName"`
	GetName      string `json:"getName"`
	GiveId       int    `json:"giveId"`
	GetId        int    `json:"getId"`
	GiveQuantity int    `json:"giveQuantity"`
	GetQuantity  int    `json:"getQuantity"`
}

type BuildBaseInfo struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type MaterialInfo struct {
	Name    string             `json:"name"`
	Designs []BuildRequirement `json:"designs"`
}

type BuildRequirement struct {
	Design   string `json:"design"`
	DesignId int    `json:"designId"`
	Quantity int    `json:"quantity"`
}

type TradeForItem struct {
	ObtaiedFrom []Trade `json:"obtainedFrom"`
	TradeFor    []Trade `json:"tradeFor"`
}
