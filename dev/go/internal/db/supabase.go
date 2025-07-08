package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type DBConfig struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewDB(config DBConfig) (*DB, error) {
	// Debug DNS resolution
	log.Printf("🔍 Testing DNS resolution for db.fssbljnxonqzwctasvjk.supabase.co...")

	// Test DNS resolution directly
	ips, err := net.LookupHost("db.fssbljnxonqzwctasvjk.supabase.co")
	if err != nil {
		log.Printf("❌ DNS lookup failed: %v", err)
		log.Printf("🔄 Attempting database connection anyway...")
	} else {
		log.Printf("✅ DNS resolved to: %v", ips)
	}

	// Test TCP connectivity
	log.Printf("🔍 Testing TCP connectivity...")
	conn, err := net.DialTimeout("tcp", "db.fssbljnxonqzwctasvjk.supabase.co:5432", 10*time.Second)
	if err != nil {
		log.Printf("❌ TCP connection failed: %v", err)
	} else {
		log.Printf("✅ TCP connection successful")
		conn.Close()
	}

	log.Printf("🔍 Opening database connection...")
	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	log.Printf("🔍 Testing database ping...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Supabase database connected successfully")
	return &DB{db}, nil
}

func (d *DB) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := d.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

func (d *DB) GetDatabaseStatistics() map[string]interface{} {
	stats := d.Stats()
	return map[string]interface{}{
		"open_connections":    stats.OpenConnections,
		"in_use":              stats.InUse,
		"idle":                stats.Idle,
		"wait_count":          stats.WaitCount,
		"wait_duration":       stats.WaitDuration.String(),
		"max_idle_closed":     stats.MaxIdleClosed,
		"max_lifetime_closed": stats.MaxLifetimeClosed,
	}
}
