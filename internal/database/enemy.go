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
	q := `SELECT id, difficulty, name ,health, armor, attack, elem_att, elem_att_type,  special,
	Soul_Eater, tgk, 'h/m', Instakill, fracture_armor, fracture_weapon, Pierce, Cold, Fire
	FROM enemy
	WHERE id = ?
	`
	// Reap, Poison, Frostbite, Bleed, Ignite,   Blind,  'Icy Slow',  'Shrieking Slow',  Fear,
	// Exile,  Freeze,  Stun,   soc1,  soc2,   soc3,  Trap
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
			&enemy.FractureArmor, &enemy.FractureWeapon, &enemy.Pierce, &enemy.Cold, &enemy.Fire); err != nil {
			log.Printf("Failed Parsing enemies row: %v", err)
			return enemies
		}
		enemies = append(enemies, enemy)
	}
	return enemies
}
