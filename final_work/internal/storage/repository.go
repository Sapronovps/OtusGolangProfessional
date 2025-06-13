package storage

import "github.com/Sapronovps/OtusGolangProfessional/final_work/internal/model"

type BannerRepository interface {
	CreateSlot(s *model.Slot) error          // Создание слота
	GetSlot(id int) (*model.Slot, error)     // Получение слота
	CreateBanner(b *model.Banner) error      // Создание баннера
	GetBanner(id int) (*model.Banner, error) // Получение баннера
	UpdateBanner(b *model.Banner) error      // Обновление баннера
	DeleteBanner(id int) error               // Удаление баннера
}
