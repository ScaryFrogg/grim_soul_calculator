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

type Service interface {
	WeaponData() []types.Weapon
	GetEnemies() []types.Enemy
	GetDesigns() []types.BuildBaseInfo
	GetDesign(id string) types.Blueprint
	GetItem(materialId string) (*types.MaterialInfo, error)
	GetDesignsForItem(materialId string) []types.BuildRequirement
	GetRequirements(design string) []types.Requirement
	GetSets() []types.ArmorData
	GetPiecesForSet(id int) []types.ArmorData
	GetTrades() []types.Trade
	GetTradesForItem(id string) []types.Trade
	GetCookData() []types.Recipe
	Health() map[string]string
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dburl      string
	dbInstance *service
)

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
