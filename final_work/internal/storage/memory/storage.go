package memory

import (
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/model"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/storage"
)

type Storage struct {
	bannerRepository storage.BannerRepository
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Banner() storage.BannerRepository {
	if s.bannerRepository != nil {
		return s.bannerRepository
	}
	s.bannerRepository = &BannerRepository{
		Slots:   make(map[int]*model.Slot),
		Banners: make(map[int]*model.Banner),
	}
	return s.bannerRepository
}
