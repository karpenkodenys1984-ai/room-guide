package repositories

import (
	"room-guide/internal/models"
)

type NodeRepository interface {
	Save(label string, x float32, y float32) (*models.NodeRecord, error)
	FindAll() ([]*models.NodeRecord, error)
	UpdateNode(nodeId int64, label string, x float32, y float32) error
	DeleteNode(nodeId int64) error
}
