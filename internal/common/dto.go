package dto

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
