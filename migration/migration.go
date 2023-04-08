package migration

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func normalizeFileName(s string) string {
	return strings.NewReplacer(
		" ", "-",
		".", "-").
		Replace(s)
}

type Migration struct {
	tx *sql.Tx
}

func NewMigration(tx *sql.Tx) *Migration {
	return &Migration{tx: tx}
}

func (m *Migration) Migrate(migrationType string, limit int) error {
	_, err := m.tx.Exec(
		`CREATE TABLE IF NOT EXISTS migrations(
			migration varchar
			, migration_type varchar(4)
		)`)
	if err != nil {
		return err
	}
	var recentMigration, recentMigrationType string
	err = m.tx.QueryRow(
		`SELECT
			COALESCE(migration,'0')
			, COALESCE(migration_type,'')
			FROM migrations`,
	).Scan(&recentMigration,
		&recentMigrationType,
	)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return err
	}

	var (
		queries,
		migrationPaths []string
		lastMigration string
	)
	err = filepath.Walk("./migration", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		// fmt.Println(fileName, "isSQL", filepath.Ext(fileName) == ".sql")
		if filepath.Ext(fileName) == ".sql" {
			filePaths := strings.Split(fileName, ".")
			// fmt.Println("isHigherFromRecent && same type", filePaths[0] > recentMigration && filePaths[2] == migrationType)
			switch migrationType {
			case "up":
				if filePaths[0] >= recentMigration && filePaths[2] == migrationType {
					migrationPaths = append(migrationPaths, path)
				}
			case "down":
				if filePaths[0] <= recentMigration && filePaths[2] == migrationType {
					migrationPaths = append(migrationPaths, path)
				}
			}
			if filePaths[0] == recentMigration && filePaths[2] == migrationType && migrationType == recentMigrationType {
				migrationPaths = migrationPaths[:len(migrationPaths)-1]
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// down migration start from latest
	if migrationType == "down" {
		reverseArray(migrationPaths)
	}
	migrationCount := len(migrationPaths)

	// if limit not set, migrate all migration file
	if limit > 0 && migrationCount > 0 {
		// handling slice out of bounds
		if limit > migrationCount {
			limit = migrationCount
		}
		migrationPaths = migrationPaths[:limit]
	}

	for _, path := range migrationPaths {
		queryFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer queryFile.Close()
		scanner := bufio.NewScanner(queryFile)
		var query string
		for scanner.Scan() {
			queryRow := scanner.Text()
			query = strings.TrimSpace(fmt.Sprintf(
				"%s\n%s",
				query,
				queryRow))
			if strings.TrimSpace(queryRow) == "" || strings.HasSuffix(queryRow, ";") {
				// add query
				queries = append(queries, query)
				query = ""
			}
		}
		// check if at the end of file still have query
		if strings.TrimSpace(query) != "" {
			queries = append(queries, query)
		}

		//get last migration date
		fileStat, err := queryFile.Stat()
		if err != nil {
			return err
		}
		fileName := fileStat.Name()
		filePaths := strings.Split(fileName, ".")
		lastMigration = filePaths[0]
	}

	// execute query
	for _, query := range queries {
		// in case theres empty string query
		if strings.TrimSpace(query) == "" {
			continue
		}

		_, err = m.tx.Exec(query)
		if err != nil {
			log.Println(query, "\n[FAIL EXECUTED]", err)
			return err
		}
		log.Println(query, "\n[SUCCESSFULLY EXECUTED]")
	}

	// update migrations table
	if len(queries) > 0 && strings.TrimSpace(queries[0]) != "" {
		_, err = m.tx.Exec(
			`DELETE FROM migrations`,
		)
		if err != nil {
			return err
		}
		_, err = m.tx.Exec(
			`INSERT INTO migrations(migration, migration_type) VALUES ($1, $2)`,
			lastMigration,
			migrationType,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Migration) New(migrationName string) (err error) {
	now := time.Now().UTC().Format(`20060102150405`)
	migrationName = normalizeFileName(migrationName)
	fileNameFormat := fmt.Sprintf("./migration/%s.%s.%%s.sql", now, migrationName)
	migrationTypes := []string{"up", "down"}
	for _, migrationType := range migrationTypes {
		fileName := fmt.Sprintf(fileNameFormat, migrationType)
		err = os.WriteFile(fileName, nil, 0600)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return
}

func reverseArray[T any](arr []T) {
	length := len(arr)
	mid := length / 2
	for i := 0; i < mid; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
}
