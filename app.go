package main

import (
	"context"
	"database/sql"
	"log/slog"
	"room-guide/internal/dto"
	"room-guide/internal/services"
)

type App struct {
	ctx        context.Context
	mapService *services.MapService
	log        *slog.Logger
}

func NewApp(log *slog.Logger) *App {

	db, err := sql.Open("sqlite", "maps.db")

	if err != nil {
		log.Error("failed to open database", slog.String("error", err.Error()))
		panic(err)
	}

	return &App{
		mapService: services.NewMapService("maps", db),
		log:        log,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.log.Info("App started")
}

func (a *App) SaveMapBackground(base64Data string, fileName string) (string, error) {
	savePath, err := a.mapService.SaveBackground(base64Data, fileName)

	if err != nil {
		a.log.Error("failed to save map background",
			slog.String("file", fileName),
			slog.String("error", err.Error()),
		)
		return "", err
	}

	return savePath, nil
}

func (a *App) GetMapBackground() (string, error) {
	return a.mapService.GetMapBackround()
}

func (a *App) SaveNode(label string, x float32, y float32) (int64, error) {
	return a.mapService.SaveNode(label, x, y)
}

func (a *App) GetAllNodes() ([]dto.Node, error) {
	return a.mapService.GetAllNodes()
}

func (a *App) UpdateNode(nodeId int64, label string, x float32, y float32) error {
	return a.mapService.UpdateNode(nodeId, label, x, y)
}

func (a *App) DeleteNode(nodeId int64) error {
	return a.mapService.DeleteNode(nodeId)
}
