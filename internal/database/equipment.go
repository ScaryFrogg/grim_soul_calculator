package database

import (
	"log"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
)

func (s *service) ShieldsData() []types.Shield {
	q := "SELECT id,name,armor FROM shield"
	shields := make([]types.Shield, 0)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting shield data failed: %v", err)
		return shields
	}
	defer rows.Close()
	for rows.Next() {
		var shield types.Shield
		if err := rows.Scan(&shield.Id, &shield.Name, &shield.Armor); err != nil {
			log.Printf("Failed Parsing shield row: %v", err)
			return shields
		}
		shields = append(shields, shield)
	}
	return shields
}

func (s *service) WeaponData() []types.Weapon {
	q := "SELECT name, damage, attack_speed as attackSpeed, s1,s2,s3,s4,s5 FROM weapon LEFT JOIN sharpen ON weapon.id = sharpen.weapon_id"
	weapons := make([]types.Weapon, 0, 32)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting weapons data failed: %v", err)
		return weapons
	}
	defer rows.Close()
	for rows.Next() {
		var weapon types.Weapon
		if err := rows.Scan(&weapon.Name, &weapon.Damage, &weapon.AttackSpeed, &weapon.S1, &weapon.S2, &weapon.S3, &weapon.S4, &weapon.S5); err != nil {
			log.Printf("Failed Parsing weapons row: %v", err)
			return weapons
		}
		weapons = append(weapons, weapon)
	}
	return weapons
}

func (s *service) ArmorPerSlot() [5][]types.ArmorData {
	q := `SELECT armor.id,armor.name,armor,protection,prot.name AS protectionType,durability,crafting,effect,slot
		FROM armor 
		LEFT JOIN protection_type prot ON armor.protection_type = prot.id
		ORDER BY slot`
	var result [5][]types.ArmorData

	allArmor := make([]types.ArmorData, 0)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting sets data failed: %v", err)
		return result
	}
	defer rows.Close()
	for rows.Next() {
		var req types.ArmorData
		if err := rows.Scan(&req.Id, &req.Name, &req.Armor, &req.Protection, &req.ProtectionType, &req.Durability, &req.Crafting, &req.Effect, &req.Slot); err != nil {
			log.Printf("Failed Parsing sets row: %v", err)
			return result
		}
		allArmor = append(allArmor, req)
	}
	size := len(allArmor) / 5
	for i := 0; i < 5; i++ {
		result[i] = allArmor[i*size : size*(i+1)]
	}
	return result
}

func (s *service) GetSets() []types.ArmorData {
	q := `SELECT sets.id,sets.name,armor,protection,prot.name AS protectionType,durability,crafting,effect FROM sets
		LEFT JOIN protection_type prot ON sets.protection_type = prot.id
	`
	sets := make([]types.ArmorData, 0)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting sets data failed: %v", err)
		return sets
	}
	defer rows.Close()
	for rows.Next() {
		var req types.ArmorData
		if err := rows.Scan(&req.Id, &req.Name, &req.Armor, &req.Protection, &req.ProtectionType, &req.Durability, &req.Crafting, &req.Effect); err != nil {
			log.Printf("Failed Parsing sets row: %v", err)
			return sets
		}
		sets = append(sets, req)
	}
	return sets
}

func (s *service) GetPiecesForSet(id int) []types.ArmorData {
	q := `SELECT armor.id,armor.name,armor,protection,prot.name AS protectionType,durability,crafting,effect FROM armor
		LEFT JOIN protection_type prot ON armor.protection_type = prot.id
		WHERE set_id = ?`
	sets := make([]types.ArmorData, 0)
	rows, err := s.db.Query(q, id)
	if err != nil {
		log.Printf("Getting set pieces data failed: %v", err)
		return sets
	}
	defer rows.Close()
	for rows.Next() {
		var req types.ArmorData
		if err := rows.Scan(&req.Id, &req.Name, &req.Armor, &req.Protection, &req.ProtectionType, &req.Durability, &req.Crafting, &req.Effect); err != nil {
			log.Printf("Failed Parsing set pieces row: %v", err)
			return sets
		}
		sets = append(sets, req)
	}
	return sets
}
