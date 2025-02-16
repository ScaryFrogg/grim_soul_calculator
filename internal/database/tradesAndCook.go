package database

import (
	"log"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
)

func (s *service) GetTrades() []types.Trade {
	q := `SELECT give_id, give_quantity, give.name, get_id, get_quantity, get.name FROM trade
		INNER JOIN item give ON give_id = give.id
		INNER JOIN item get ON get_id = get.id`

	trades := make([]types.Trade, 0)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting trades data failed: %v", err)
		return trades
	}
	defer rows.Close()
	for rows.Next() {
		var req types.Trade
		if err := rows.Scan(&req.GiveId, &req.GiveQuantity, &req.GiveName, &req.GetId, &req.GetQuantity, &req.GetName); err != nil {
			log.Printf("Failed Parsing trades row: %v", err)
			return trades
		}
		trades = append(trades, req)
	}
	return trades
}

func (s *service) GetTradesForItem(id string) []types.Trade {
	q := `SELECT give_quantity, give.name, get_quantity, get.name FROM trade
		INNER JOIN item give ON give_id = give.id
		INNER JOIN item get ON get_id = get.id
		WHERE give_id = ? OR get_id = ?`

	trades := make([]types.Trade, 0)
	rows, err := s.db.Query(q, id, id)
	if err != nil {
		log.Printf("Getting trades data failed: %v", err)
		return trades
	}
	defer rows.Close()
	for rows.Next() {
		var req types.Trade
		if err := rows.Scan(&req.GiveQuantity, &req.GiveName, &req.GetQuantity, &req.GetName); err != nil {
			log.Printf("Failed Parsing trades row: %v", err)
			return trades
		}
		trades = append(trades, req)
	}
	return trades
}

func (s *service) GetCookData() []types.Recipe {
	q := "SELECT ing1,ing1quantity,ing2,ing2Quantity,result,resultQuantity FROM recipe"
	enemies := make([]types.Recipe, 0, 50)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting enemies data failed: %v", err)
		return enemies
	}
	defer rows.Close()
	for rows.Next() {
		var recipe types.Recipe
		if err := rows.Scan(&recipe.Ing1, &recipe.Ing1Quantity, &recipe.Ing2, &recipe.Ing2Quantity, &recipe.Result, &recipe.ResultQuantity); err != nil {
			log.Printf("Failed Parsing recipes row: %v", err)
			return enemies
		}
		enemies = append(enemies, recipe)
	}
	return enemies
}
