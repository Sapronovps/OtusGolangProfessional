package internalhttp

import (
	"context"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"io"
	"net/http"
	"os"
)

type Server struct {
	Address string
	Logger
	Application
}

type Logger interface {
	Error(msg string)
	Warning(msg string)
	Info(msg string)
	Debug(msg string)
}

type Application interface {
	AddEvent(ctx context.Context, title string) (*model.Event, error)
}

func NewServer(logger Logger, app Application) *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	s.Logger.Info("starting http server: " + s.Address)

	// Обработчик для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("handle func /")
		res, err := s.AddEvent(ctx, "iiio")
		if err != nil {
			s.Logger.Error("could not add event to /:" + err.Error())
			return
		}

		_, err = io.WriteString(w, fmt.Sprintf("Hi otus, %d:%s", res.ID, res.Title))
		if err != nil {
			s.Logger.Error("could not write to /:" + err.Error())
		}
	})

	err := http.ListenAndServe(s.Address, nil)
	if err != nil {
		return fmt.Errorf("could not start http server: %w", err)
	}

	go func() {
		<-ctx.Done()
		err := s.Stop(ctx)
		if err != nil {
			s.Logger.Error("could not stop http server: " + err.Error())
		}
	}()

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	os.Exit(1)
	return nil
}

// TODO
