package internalhttp

import (
	"context"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	address string
	logger  *zap.Logger
	app     Application
	server  *http.Server
}

type Application interface {
	CreateEvent(event *model.Event) error
	GetEvent(id int) (model.Event, error)
	UpdateEvent(id int, event *model.Event) error
	DeleteEvent(id int) error
	ListByDay(date time.Time) ([]model.Event, error)
	ListByWeek(date time.Time) ([]model.Event, error)
	ListByMonth(date time.Time) ([]model.Event, error)
}

func NewServer(logger *zap.Logger, app Application, address string) *Server {
	s := &Server{
		address: address,
		logger:  logger,
		app:     app,
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	// Создаем новый роутер
	r := mux.NewRouter()

	// Регистрируем обработчики
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/hello", home).Methods("GET")
	r.HandleFunc("/events", func(writer http.ResponseWriter, request *http.Request) {
		createEvent(writer, request, s.app)
	}).Methods("POST")
	r.HandleFunc("/events/{id}", func(writer http.ResponseWriter, request *http.Request) {
		getEvent(writer, request, s.app)
	}).Methods("GET")
	r.HandleFunc("/events/{id}", func(writer http.ResponseWriter, request *http.Request) {
		updateEvent(writer, request, s.app)
	}).Methods("PUT")
	r.HandleFunc("/events/{id}", func(writer http.ResponseWriter, request *http.Request) {
		deleteEvent(writer, request, s.app)
	}).Methods("DELETE")
	r.HandleFunc("/listByDay/{date}", func(writer http.ResponseWriter, request *http.Request) {
		listByDay(writer, request, s.app)
	}).Methods("GET")
	r.HandleFunc("/listByWeek/{date}", func(writer http.ResponseWriter, request *http.Request) {
		listByWeek(writer, request, s.app)
	}).Methods("GET")
	r.HandleFunc("/listByMonth/{date}", func(writer http.ResponseWriter, request *http.Request) {
		listByMonth(writer, request, s.app)
	}).Methods("GET")

	// Добавляем middleware для логирования
	r.Use(s.loggingMiddleware)

	// Настраиваем HTTP сервер
	s.server = &http.Server{
		Addr:         s.address,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Запускаем сервер
	s.logger.Info("Server is running", zap.String("address", s.address))
	err := s.server.ListenAndServe()
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	os.Exit(1)
	return nil
}

// TODO
