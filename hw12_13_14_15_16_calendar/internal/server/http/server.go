package internalhttp

import (
	"context"
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

type Application interface { // TODO
}

func NewServer(logger *zap.Logger, app Application, address string) *Server {
	s := &Server{
		address: address,
		logger:  logger,
		app:     app,
	}

	return s
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			"started",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("IP", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Info(
			"completed",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("IP", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
			zap.Time("request_datetime", start),
			zap.Duration("duration", time.Since(start)),
		)
	})
}

func echoHelloWorld(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World"))
}

func (s *Server) Start(ctx context.Context) error {
	// Создаем новый роутер
	r := mux.NewRouter()

	// Регистрируем обработчики
	r.HandleFunc("/", echoHelloWorld).Methods("GET")
	r.HandleFunc("/hello", echoHelloWorld).Methods("GET")

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
