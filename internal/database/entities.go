package database

import "database/sql"

type Weapon struct {
	AttackSpeed float32 `json:"attackSpeed"`
	Name        string  `json:"name"`
	Damage      int     `json:"damage"`
	S1          int     `json:"s1"`
	S2          int     `json:"s2"`
	S3          int     `json:"s3"`
	S4          int     `json:"s4"`
	S5          int     `json:"s5"`
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
