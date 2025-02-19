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
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Designs     []BuildRequirement `json:"designs"`
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

type Recipe struct {
	Ing1           string  `json:"ing1"`
	Ing2           *string `json:"ing2,omitempty"`
	Result         string  `json:"result"`
	Ing1Quantity   int     `json:"ing1Quantity"`
	Ing2Quantity   *int    `json:"ing2Quantity,omitempty"`
	ResultQuantity int     `json:"resultQuantity"`
}

type ArmorData struct {
	Name           string  `json:"name"`
	ProtectionType *string `json:"protectionType,omitempty"`
	Crafting       *string `json:"crafting,omitempty"`
	Effect         *string `json:"effect,omitempty"`
	Id             int     `json:"id"`
	Armor          int     `json:"armor"`
	Protection     *int    `json:"protection,omitempty"`
	Durability     *int    `json:"durability,omitempty"`
	Slot           *int    `json:"slot,omitempty"`
	SetId          *int    `json:"setId,omitempty"`
}

type Shield struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Armor int    `json:"armor"`
}

type EnemyListInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Enemy struct {
	Id             string  `json:"id"`
	Difficulty     string  `json:"difficulty"`
	Name           string  `json:"name"`
	Special        string  `json:"special"`
	ElemAttType    *string `json:"elem_att_type,omitempty"`
	HorM           *string `json:"horM,omitempty"`
	FractureArmor  *bool   `json:"fractureArmor,omitempty"`
	FractureWeapon *bool   `json:"fractureWeapon,omitempty"`
	Pierce         *string `json:"pierce,omitempty"`
	Cold           *string `json:"cold,omitempty"`
	Fire           *string `json:"fire,omitempty"`
	Health         int     `json:"health"`
	Armor          int     `json:"armor"`
	Attack         int     `json:"attack"`
	ElemAtt        *int    `json:"elem_att,omitempty"`
	SoulEater      *int    `json:"soulEater,omitempty"`
	Tgk            *int    `json:"tgk,omitempty"`
	Instakill      *bool   `json:"instakill,omitempty"`
	Reap           *string `json:"reap,omitempty"`
	Poison         *string `json:"poison,omitempty"`
	Frostbite      *string `json:"frostbite,omitempty"`
	Bleed          *string `json:"bleed,omitempty"`
	Ignite         *string `json:"ignite,omitempty"`
	Blind          *bool   `json:"blind,omitempty"`
	IcySlow        *bool   `json:"icySlow,omitempty"`
	ShriekingSlow  *bool   `json:"shriekingSlow,omitempty"`
	Fear           *bool   `json:"fear,omitempty"`
	Exile          *bool   `json:"exile,omitempty"`
	Freeze         *bool   `json:"freeze,omitempty"`
	Stun           *bool   `json:"stun,omitempty"`
	Soc1           *bool   `json:"soc1,omitempty"`
	Soc2           *bool   `json:"soc2,omitempty"`
	Soc3           *bool   `json:"soc3,omitempty"`
	Trap           *bool   `json:"trap,omitempty"`
}
