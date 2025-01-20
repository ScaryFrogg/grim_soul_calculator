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
	WeaponData() []Weapon
	GetEnemies() []Enemy
	GetDesigns() []string
	GetDesign(id string) Blueprint
	GetMaterial(materialId string) (name string, err error)
	GetDesignsForMaterial(materialId string) []dto.BuildRequirement
	GetRequirements(design string) []Requirement
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
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func (s *service) GetMaterial(materialId string) (name string, err error) {
	q := `	SELECT name 
		FROM material
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

func (s *service) GetDesignsForMaterial(materialId string) []dto.BuildRequirement {
	q := `	SELECT design.name, quantity, design_id
		FROM craft_requirement
		INNER JOIN design ON design.id = design_id
		WHERE material_id = ?`

	requiremets := make([]dto.BuildRequirement, 0)
	rows, err := s.db.Query(q, materialId)
	if err != nil {
		log.Printf("Getting  dati failed: %v", err)
		return requiremets
	}
	defer rows.Close()
	for rows.Next() {
		var req dto.BuildRequirement
		if err := rows.Scan(&req.Design, &req.DesignId, &req.Quantity); err != nil {
			log.Printf("Failed Parsing requiremets row: %v", err) // Log the error and terminate the program
			return requiremets
		}
		requiremets = append(requiremets, req)
	}
	return requiremets
}
func (s *service) GetDesigns() []string {
	q := "select name from design"
	requiremets := make([]string, 0, 10)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting designs data failed: %v", err) // Log the error and terminate the program
		return requiremets
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Printf("Failed Parsing designgs row: %v", err) // Log the error and terminate the program
			return requiremets
		}
		requiremets = append(requiremets, name)
	}
	return requiremets
}

func (s *service) GetRequirements(design string) []Requirement {
	q := `SELECT material.name, quantity
		FROM craft_requirement
		INNER JOIN design on design.id = design_id
		INNER JOIN material on material.id = material_id
		where design.id = ?`
	requiremets := make([]Requirement, 0, 10)
	rows, err := s.db.Query(q, design)
	if err != nil {
		log.Printf("Getting requirement data failed: %v", err) // Log the error and terminate the program
		return requiremets
	}
	defer rows.Close()
	for rows.Next() {
		var req Requirement
		if err := rows.Scan(&req.Name, &req.Quantity); err != nil {
			log.Printf("Failed Parsing requiremets row: %v", err) // Log the error and terminate the program
			return requiremets
		}
		requiremets = append(requiremets, req)
	}
	return requiremets
}

func (s *service) WeaponData() []Weapon {
	q := "SELECT name, damage, attack_speed as attackSpeed, s1,s2,s3,s4,s5 FROM weapon INNER JOIN sharpen ON weapon.id = sharpen.weapon_id"
	weapons := make([]Weapon, 0, 32)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting weapons data failed: %v", err) // Log the error and terminate the program
		return weapons
	}
	defer rows.Close()
	for rows.Next() {
		var weapon Weapon
		if err := rows.Scan(&weapon.Name, &weapon.Damage, &weapon.AttackSpeed, &weapon.S1, &weapon.S2, &weapon.S3, &weapon.S4, &weapon.S5); err != nil {
			log.Printf("Failed Parsing weapons row: %v", err) // Log the error and terminate the program
			return weapons
		}
		weapons = append(weapons, weapon)
	}
	return weapons
}

func (s *service) GetEnemies() []Enemy {
	q := "select name, health, armor from enemy"
	enemies := make([]Enemy, 0, 200)
	rows, err := s.db.Query(q)
	if err != nil {
		log.Printf("Getting enemies data failed: %v", err) // Log the error and terminate the program
		return enemies
	}
	defer rows.Close()
	for rows.Next() {
		var enemy Enemy
		if err := rows.Scan(&enemy.Name, &enemy.Health, &enemy.Armor); err != nil {
			log.Printf("Failed Parsing enemies row: %v", err) // Log the error and terminate the program
			return enemies
		}
		enemies = append(enemies, enemy)
	}
	return enemies
}

func (s *service) GetDesign(id string) Blueprint {
	q := "select id, name, next from design where id = ?"
	var design Blueprint
	row := s.db.QueryRow(q, id)
	if err := row.Scan(&design.Id, &design.Name, &design.Next); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No Design found with id [%v]", id) // Log the error and terminate the program
			return design
		}
		log.Printf("design ById %v: %v", id, err)
		return design
	}
	return design
}

func New() Service {
	// Reuse Connection
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
		log.Printf("db down: %v", err) // Log the error and terminate the program
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
