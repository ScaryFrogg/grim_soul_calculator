package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// Service represents a service that interacts with a database.
type Service interface {
	WeaponData() []types.Weapon
	GetEnemies() []types.Enemy
	GetDesigns() []types.BuildBaseInfo
	GetTrades() []types.Trade
	GetTradesForItem(id string) []types.Trade
	GetDesign(id string) types.Blueprint
	GetItem(materialId string) (name string, err error)
	GetDesignsForItem(materialId string) []types.BuildRequirement
	GetRequirements(design string) []types.Requirement
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dburl      string
	dbInstance *service
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

func (s *service) GetItem(materialId string) (name string, err error) {
	q := `	SELECT name 
		FROM item
		WHERE id = ?`

	row := s.db.QueryRow(q, materialId)
	if err := row.Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Not Found Material with id %v", materialId)
		}
		log.Printf("design ById %v: %v", materialId, err)
		return "", fmt.Errorf("Failed QUERY")
	}
	return name, nil
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

func (s *service) GetEnemies() []types.Enemy {
	q := "select name, health, armor from enemy"
	enemies := make([]types.Enemy, 0, 200)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting enemies data failed: %v", err)
		return enemies
	}
	defer rows.Close()
	for rows.Next() {
		var enemy types.Enemy
		if err := rows.Scan(&enemy.Name, &enemy.Health, &enemy.Armor); err != nil {
			log.Printf("Failed Parsing enemies row: %v", err)
			return enemies
		}
		enemies = append(enemies, enemy)
	}
	return enemies
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

func New() Service {
	dburl = os.Getenv("BLUEPRINT_DB_URL") // Reuse Connection
	if dburl == "" {
		log.Fatal("missing BLUEPRINT_DB_URL environment variable")
	}
	if dbInstance != nil {
		return dbInstance
	}
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("db down: %v", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}
