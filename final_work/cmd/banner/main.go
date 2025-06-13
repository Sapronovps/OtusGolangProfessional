package main

import (
	"flag"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/app"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/logger"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/model"
	"github.com/Sapronovps/OtusGolangProfessional/final_work/internal/storage/memory"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)
	logg := logger.New(config.Logger.Level, config.Logger.File)
	storage := memory.New()

	application := app.NewApp(logg, storage)

	newSlot := model.Slot{
		ID:          1,
		Description: "Hello world",
	}

	err := application.AddSlot(&newSlot)

	if err != nil {
		panic("Failed to add new slot")
	}

	slot, err := application.GetSlot(1)
	if err != nil {
		panic("Failed to get slot")
	}
	_ = slot

	newBanner := model.Banner{
		Description: "New Banner",
	}

	err = application.AddBanner(&newBanner)

	if err != nil {
		panic("Failed to add new banner")
	}

	banner, err := application.GetBanner(1)
	if err != nil {
		panic("Failed to get banner")
	}

	err = application.AttachBannerToSlot(slot.ID, banner.ID)
	if err != nil {
		panic("Failed to attach banner")
	}

	err = application.RegisterClick(banner.ID)
	err = application.RegisterClick(banner.ID)
	err = application.RegisterClick(banner.ID)
	if err != nil {
		panic("Failed to register click")
	}

	fmt.Println(banner.Shows)
}
