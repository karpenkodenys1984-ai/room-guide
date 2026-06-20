package sqlite

import (
	"database/sql"
	"fmt"
	models "room-guide/internal/models"
	"time"

	_ "modernc.org/sqlite"
)

type NodeRepository struct {
	db *sql.DB
}

func NewNodeRepository(db *sql.DB) *NodeRepository {
	return &NodeRepository{db: db}
}

func (r *NodeRepository) Save(label string, x int16, y int16) (*models.NodeRecord, error) {
	now := time.Now()

	result, err := r.db.Exec(
		`INSERT INTO nodes(label, x, y, created_timestamp, updated_timestamp) VALUES(?,?,?,?,?)`,
		label,
		x,
		y,
		now,
		now,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to save node: %w", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, fmt.Errorf("failed to get last insert node id: %w", err)
	}

	return &models.NodeRecord{
		NodeId:           id,
		Label:            label,
		X:                x,
		Y:                y,
		CreatedTimestamp: now,
		UpdatedTimestamp: now,
	}, nil
}

func (r *NodeRepository) FindAll() ([]*models.NodeRecord, error) {
	rows, err := r.db.Query(
		`SELECT node_id, label, x, y, created_timestamp, updated_timestamp FROM nodes`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query nodes: %w", err)
	}
	defer rows.Close()

	var nodes []*models.NodeRecord

	for rows.Next() {
		node := &models.NodeRecord{}
		err := rows.Scan(
			&node.NodeId,
			&node.Label,
			&node.X,
			&node.Y,
			&node.CreatedTimestamp,
			&node.UpdatedTimestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan node row: %w", err)
		}
		nodes = append(nodes, node)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return nodes, nil
}
