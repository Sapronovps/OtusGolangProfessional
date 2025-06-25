package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/app"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/server/http"
	storage2 "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage/sql"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	_ = config

	logg := logger.New(config.Logger.Level, config.Logger.File)
	defer logg.Sync()

	var storage storage2.Storage
	if config.DB.InMemory {
		storage = memorystorage.New()
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DB.Host, config.DB.Username, config.DB.Password, config.DB.DBName)
		db, err := app.NewDB(dsn)
		if err != nil {
			panic(fmt.Errorf("connet to db: %w", err))
		}
		storage = sqlstorage.New(db)
	}

	calendar := app.New(logg, storage)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	defer cancel()

	if config.Server.IsHTTP {
		serverAddress := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
		server := internalhttp.NewServer(logg, calendar, serverAddress)

		go func() {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err := server.Stop(ctx); err != nil {
				logg.Error("failed to stop http server: " + err.Error())
			}
		}()

		logg.Info("calendar is running...")

		if err := server.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}
	server := internalgrpc.NewEventsServiceServer(logg, calendar, config.Server.AddressGrpc)

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()
	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start grpc server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
