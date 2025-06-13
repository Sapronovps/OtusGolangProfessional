package app

import (
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/model"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/storage"
	"go.uber.org/zap"
	"sync"
	"time"
)

type App struct {
	mu      sync.Mutex
	logger  *zap.Logger
	storage storage.Storage
}

func NewApp(logger *zap.Logger, storage storage.Storage) *App {
	return &App{logger: logger, storage: storage}
}

func (a *App) AddSlot(slot *model.Slot) error {
	return a.storage.Banner().CreateSlot(slot)
}

func (a *App) GetSlot(id int) (slot *model.Slot, err error) {
	return a.storage.Banner().GetSlot(id)
}

func (a *App) AddBanner(banner *model.Banner) error {
	return a.storage.Banner().CreateBanner(banner)
}

func (a *App) GetBanner(id int) (banner *model.Banner, err error) {
	return a.storage.Banner().GetBanner(id)
}

func (a *App) RemoveBanner(id int) error {
	return a.storage.Banner().DeleteBanner(id)
}

func (a *App) AttachBannerToSlot(slotID, bannerID int) error {
	slot, err := a.GetSlot(slotID)
	if err != nil {
		return fmt.Errorf("could not find slot with id %d: %w", slotID, err)
	}

	banner, err := a.GetBanner(bannerID)
	if err != nil {
		return fmt.Errorf("could not find banner with id %d: %w", bannerID, err)
	}
	if slot.Banners == nil {
		slot.Banners = make(map[int]*model.Banner)
	}

	slot.Banners[bannerID] = banner
	return nil
}

func (a *App) DetachBannerFromSlot(slotID, bannerID int) error {
	slot, err := a.GetSlot(slotID)
	if err != nil {
		return fmt.Errorf("could not find slot with id %d: %w", slotID, err)
	}
	_, ok := slot.Banners[bannerID]
	if !ok {
		return fmt.Errorf("no banner found with id %d", bannerID)
	}
	delete(slot.Banners, bannerID)
	return nil
}

func (a *App) RegisterClick(bannerID int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	banner, err := a.GetBanner(bannerID)
	if err != nil {
		return fmt.Errorf("could not find banner with id %d: %w", bannerID, err)
	}

	banner.Clicks++
	banner.Weight = float64(banner.Clicks) / float64(banner.Shows) // Обновляем CTR
	return nil
}

func (a *App) calculateFatigue(slotId, bannerId int) (float64, error) {
	slot, err := a.GetSlot(slotId)
	if err != nil {
		return 0, fmt.Errorf("could not find slot with id %d: %w", slotId, err)
	}
	history := slot.ShowHistory[bannerId]
	now := time.Now()
	recentShows := 0

	// Учитываем только показы за последние N часов (или другой интервал)
	for _, t := range history {
		if now.Sub(t) < time.Hour*24 {
			recentShows++
		}
	}

	return float64(recentShows), nil
}
