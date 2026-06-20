package repositories

import (
	"room-guide/internal/models"
)

type MapRepository interface {
	Save(mapPath string) (*models.MapRecord, error)
	GetLatest() (*models.MapRecord, error)
}
