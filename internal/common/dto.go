package dto

type MaterialInfo struct {
	Name    string             `json:"name"`
	Designs []BuildRequirement `json:"designs"`
}

type BuildRequirement struct {
	Design   string `json:"Design"`
	DesignId int    `json:"DesignId"`
	Quantity int    `json:"Quantity"`
}
