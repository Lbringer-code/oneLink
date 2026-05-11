package db

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

const createMigrationsTable = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	filename TEXT PRIMARY KEY ,
	applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)`

func RunMigrations( db *sqlx.DB , migrationsDir string , logger *slog.Logger ) error {
	_ , err := db.Exec(createMigrationsTable)
	if err != nil {
		return fmt.Errorf("create schema_migrations table: %w" , err)
	}

	entries , err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("read migrations dir: %w" , err)
	}

	files := make( []string , 0 , len(entries) )
	for _ , e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name() , ".sql") {
			files = append(files , e.Name())
		}
	}
	sort.Strings(files)

	applied := make(map[string]bool)
	var names []string
	err = db.Select( &names ,  `SELECT filename FROM schema_migrations`)
	if err != nil {
		return fmt.Errorf("query applied migrations: %w" , err)
	}
	for _ , n := range names {
		applied[n] = true
	}

	for _ , file := range files {
		if applied[file] {
			continue
		}

		path := filepath.Join( migrationsDir , file )
		content , err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w" , file , err)
		}

		logger.Info("applying migration" , "file" , file)

		tx , err := db.Beginx()
		if err != nil {
			return fmt.Errorf("begin tx for %s: %w" , file , err)
		}

		 _ , err = tx.Exec(string(content))
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("execute %s: %w" , file , err)
		}

		 _ , err = tx.Exec(`INSERT INTO schema_migrations (filename) VALUES ($1)` , file)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("record %s: %w" , file , err)
		}

		err = tx.Commit()
		if err != nil {
			return fmt.Errorf("commit %s: %w" , file , err)
		}
	}
	return nil
}