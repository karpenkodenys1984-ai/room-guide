package services

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"room-guide/internal/dto"
	"room-guide/internal/models"
	"room-guide/internal/repositories"
	repoSqlite "room-guide/internal/repositories/sqlite"
	"strings"
)

type MapService struct {
	mapsDir         string
	defaultFileName string
	repo            repositories.MapRepository
	nodeRepo        repoSqlite.NodeRepository
}

var allowedMimeTypes = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
	"image/webp": ".webp",
}

func NewMapService(mapsDir string, db *sql.DB) *MapService {
	return &MapService{
		mapsDir:         mapsDir,
		defaultFileName: "map.png",
		repo:            repoSqlite.NewMapRepository(db),
		nodeRepo:        *repoSqlite.NewNodeRepository(db),
	}
}

func (s *MapService) SaveBackground(base64Data string, fileName string) (string, error) {
	if fileName == "" {
		fileName = s.defaultFileName
	}

	mimeType, rawData, err := parseBase64Image(base64Data)
	if err != nil {
		return "", err
	}

	ext, ok := allowedMimeTypes[mimeType]
	if !ok {
		return "", fmt.Errorf("unsupported image type: %s", mimeType)
	}

	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ext

	decoded, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	if err := os.MkdirAll(s.mapsDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create maps dir: %w", err)
	}

	savePath := filepath.Join(s.mapsDir, fileName)
	if err := os.WriteFile(savePath, decoded, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	_, err = s.repo.Save(savePath)

	if err != nil {
		return "", fmt.Errorf("failed to store to db: %w", err)
	}

	return savePath, nil
}

func (s *MapService) GetMapBackround() (string, error) {
	lastRecord, err := s.repo.GetLatest()

	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(lastRecord.MapPath)

	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(lastRecord.MapPath))

	mimeTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".webp": "image/webp",
	}

	mime := mimeTypes[ext]

	return "data:" + mime + ";base64," + base64.StdEncoding.EncodeToString(data), nil

}

func (s *MapService) SaveNode(label string, x float32, y float32) error {
	_, err := s.nodeRepo.Save(label, x, y)

	if err != nil {
		return err
	}

	return nil
}

func (s *MapService) GetAllNodes() ([]dto.Node, error) {
	records, err := s.nodeRepo.FindAll()
	if err != nil {
		return nil, err
	}

	nodes := make([]dto.Node, len(records))
	for i, record := range records {
		nodes[i] = convertNodeRecordToDto(record)
	}

	return nodes, nil
}

func (s *MapService) UpdateNode(nodeId int64, label string, x float32, y float32) error {
	return s.nodeRepo.UpdateNode(nodeId, label, x, y)
}

func parseBase64Image(base64Data string) (mimeType, rawData string, err error) {
	// data:image/png;base64,<data>
	if !strings.HasPrefix(base64Data, "data:") {
		return "", "", fmt.Errorf("invalid base64 format: missing data prefix")
	}

	semicolon := strings.Index(base64Data, ";")
	comma := strings.Index(base64Data, ",")

	if semicolon == -1 || comma == -1 || comma < semicolon {
		return "", "", fmt.Errorf("invalid base64 format: malformed header")
	}

	mimeType = base64Data[5:semicolon]
	rawData = base64Data[comma+1:]

	if rawData == "" {
		return "", "", fmt.Errorf("invalid base64 format: empty data")
	}

	return mimeType, rawData, nil
}

func convertNodeRecordToDto(record *models.NodeRecord) dto.Node {
	return dto.Node{
		Id:   int16(record.NodeId),
		Type: string(dto.RoomTypeRoom),
		Position: dto.Position{
			X: int(record.X),
			Y: int(record.Y),
		},
		Data: dto.RoomData{
			Label: record.Label,
			Type:  dto.RoomTypeRoom,
		},
	}
}
