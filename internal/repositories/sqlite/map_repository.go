package sqlite

import (
	"database/sql"
	"fmt"
	models "room-guide/internal/models"
	"time"

	_ "modernc.org/sqlite"
)

type MapRepository struct {
	db *sql.DB
}

func NewMapRepository(db *sql.DB) *MapRepository {
	return &MapRepository{db: db}
}

func (r *MapRepository) Save(mapPath string) (*models.MapRecord, error) {
	now := time.Now()
	result, err := r.db.Exec(
		`INSERT INTO maps (map_path, created_timestamp, updated_timestamp) VALUES (?, ?, ?)`,
		mapPath, now, now,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save map: %w", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, fmt.Errorf("failed to get last insert map id: %w", err)
	}

	return &models.MapRecord{
		MapId:            id,
		MapPath:          mapPath,
		CreatedTimestamp: now,
		UpdatedTimestamp: now,
	}, nil
}

func (r *MapRepository) GetLatest() (*models.MapRecord, error) {
	row := r.db.QueryRow(
		`SELECT map_id, map_path, created_timestamp, updated_timestamp 
		 FROM maps 
		 ORDER BY created_timestamp DESC 
		 LIMIT 1`,
	)

	var record models.MapRecord

	err := row.Scan(
		&record.MapId,
		&record.MapPath,
		&record.CreatedTimestamp,
		&record.UpdatedTimestamp,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get latest map: %w", err)
	}

	return &record, nil
}
