package memory

import (
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/model"
	"sync"
)

type BannerRepository struct {
	mu      sync.RWMutex
	Slots   map[int]*model.Slot
	Banners map[int]*model.Banner
}

func (b *BannerRepository) CreateSlot(slot *model.Slot) error {
	slot.ID = len(b.Slots) + 1
	b.mu.Lock()
	b.Slots[slot.ID] = slot
	b.mu.Unlock()
	return nil
}

func (b *BannerRepository) GetSlot(id int) (*model.Slot, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	slot, ok := b.Slots[id]
	if !ok {
		return nil, fmt.Errorf("no slot with id %d", id)
	}
	return slot, nil
}

func (b *BannerRepository) CreateBanner(banner *model.Banner) error {
	banner.ID = len(b.Banners) + 1
	b.mu.Lock()
	b.Banners[banner.ID] = banner
	b.mu.Unlock()
	return nil
}

func (b *BannerRepository) GetBanner(id int) (*model.Banner, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	banner, ok := b.Banners[id]
	if !ok {
		return nil, fmt.Errorf("no banner with id %d", id)
	}
	return banner, nil
}

func (b *BannerRepository) UpdateBanner(banner *model.Banner) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Banners[banner.ID] = banner
	return nil
}

func (b *BannerRepository) DeleteBanner(id int) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.Banners, id)
	return nil
}
