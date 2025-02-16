package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
)

func (s *service) GetDesigns() []types.BuildBaseInfo {
	q := "select name, id from design"
	requiremets := make([]types.BuildBaseInfo, 0)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting designs data failed: %v", err)
		return requiremets
	}
	defer rows.Close()
	for rows.Next() {
		var info types.BuildBaseInfo
		if err := rows.Scan(&info.Name, &info.Id); err != nil {
			log.Printf("Failed Parsing designgs row: %v", err)
			return requiremets
		}
		requiremets = append(requiremets, info)
	}
	return requiremets
}

func (s *service) GetDesignsForItem(materialId string) []types.BuildRequirement {
	q := `	SELECT design.name, design_id, quantity
		FROM craft_requirement
		INNER JOIN design ON design.id = design_id
		WHERE material_id = ?`

	requiremets := make([]types.BuildRequirement, 0)
	rows, err := s.db.Query(q, materialId)
	if err != nil {
		log.Printf("Getting  dati failed: %v", err)
		return requiremets
	}
	defer rows.Close()
	for rows.Next() {
		var req types.BuildRequirement
		if err := rows.Scan(&req.Design, &req.DesignId, &req.Quantity); err != nil {
			log.Printf("Failed Parsing requiremets row: %v", err)
			return requiremets
		}
		requiremets = append(requiremets, req)
	}
	return requiremets
}

func (s *service) GetItem(materialId string) (*types.MaterialInfo, error) {
	q := `	SELECT name, description 
		FROM item
		WHERE id = ?`

	row := s.db.QueryRow(q, materialId)
	var material types.MaterialInfo
	if err := row.Scan(&material.Name, &material.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Not Found Material with id %v", materialId)
		}
		log.Printf("design ById %v: %v", materialId, err)
		return nil, fmt.Errorf("Failed QUERY")
	}
	return &material, nil
}

func (s *service) GetRequirements(design string) []types.Requirement {
	q := `SELECT item.id, item.name, quantity
		FROM craft_requirement
		INNER JOIN design on design.id = design_id
		INNER JOIN item on item.id = material_id
		where design.id = ?`
	requiremets := make([]types.Requirement, 0, 10)
	rows, err := s.db.Query(q, design)
	if err != nil {
		log.Printf("Getting requirement data failed: %v", err)
		return requiremets
	}
	defer rows.Close()
	for rows.Next() {
		var req types.Requirement
		if err := rows.Scan(&req.Id, &req.Name, &req.Quantity); err != nil {
			log.Printf("Failed Parsing requiremets row: %v", err)
			return requiremets
		}
		requiremets = append(requiremets, req)
	}
	return requiremets
}

func (s *service) GetDesign(id string) types.Blueprint {
	q := "select id, name, next from design where id = ?"
	var design types.Blueprint
	row := s.db.QueryRow(q, id)
	if err := row.Scan(&design.Id, &design.Name, &design.Next); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No Design found with id [%v]", id)
			return design
		}
		log.Printf("design ById %v: %v", id, err)
		return design
	}
	return design
}
