package repositories

import (
	"room-guide/internal/models"
)

type NodeRepository interface {
	Save(label string, x int16, y int16) (*models.NodeRecord, error)
	FindAll() ([]*models.NodeRecord, error)
}
