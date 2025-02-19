package database

import (
	"log"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
)

func (s *service) GetEnemiesInfos() []types.EnemyListInfo {
	q := "SELECT id,name FROM enemy GROUP BY id"
	enemies := make([]types.EnemyListInfo, 0)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting enemy data failed: %v", err)
		return enemies
	}
	defer rows.Close()
	for rows.Next() {
		var info types.EnemyListInfo
		if err := rows.Scan(&info.Id, &info.Name); err != nil {
			log.Printf("Failed Parsing enemy row: %v", err)
			return enemies
		}
		enemies = append(enemies, info)
	}
	return enemies
}

func (s *service) GetEnemyById(id string) []types.Enemy {
	q := `SELECT  enemy.id, difficulty, enemy.name ,health, armor, attack, elem_att, et.name as elemAttType,  special,
			Soul_Eater, tgk, type, Instakill, fracture_armor, fracture_weapon, Pierce, Cold, Fire,
			Reap, Poison, Frostbite, Bleed, Ignite,   Blind,  IcySlow,  ShriekingSlow,  Fear,
			Exile,  Freeze,  Stun,   soc1,  soc2,   soc3,  Trap
		FROM enemy
		LEFT JOIN elemental_type  et ON et.id = elem_att_type
		WHERE enemy.id = ?`
	enemies := make([]types.Enemy, 0, 5)
	rows, err := s.db.Query(q, id)
	if err != nil {
		log.Printf("Getting enemies data failed: %v", err)
		return enemies
	}
	defer rows.Close()
	for rows.Next() {
		var enemy types.Enemy
		if err := rows.Scan(&enemy.Id, &enemy.Difficulty, &enemy.Name, &enemy.Health, &enemy.Armor, &enemy.Attack, &enemy.ElemAtt,
			&enemy.ElemAttType, &enemy.Special, &enemy.SoulEater, &enemy.Tgk, &enemy.HorM, &enemy.Instakill,
			&enemy.FractureArmor, &enemy.FractureWeapon, &enemy.Pierce, &enemy.Cold, &enemy.Fire,
			&enemy.Reap, &enemy.Poison, &enemy.Frostbite, &enemy.Bleed, &enemy.Ignite, &enemy.Blind,
			&enemy.IcySlow, &enemy.ShriekingSlow, &enemy.Fear, &enemy.Exile, &enemy.Freeze, &enemy.Stun,
			&enemy.Soc1, &enemy.Soc2, &enemy.Soc3, &enemy.Trap); err != nil {
			log.Printf("Failed Parsing enemies row: %v", err)
			return enemies
		}
		enemies = append(enemies, enemy)
	}
	return enemies
}
