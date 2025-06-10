package internalgrpc

import (
	"context"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"net"
	"os"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Application interface {
	CreateEvent(event *model.Event) error
	GetEvent(id int) (model.Event, error)
	UpdateEvent(id int, event *model.Event) error
	DeleteEvent(id int) error
	ListByDay(date time.Time) ([]model.Event, error)
	ListByWeek(date time.Time) ([]model.Event, error)
	ListByMonth(date time.Time) ([]model.Event, error)
}

type EventsServiceServer struct {
	logger  *zap.Logger
	app     Application
	address string
	server  *grpc.Server
	UnimplementedEventServiceServer
}

func NewEventsServiceServer(logger *zap.Logger, app Application, address string) *EventsServiceServer {
	return &EventsServiceServer{
		logger:  logger,
		app:     app,
		address: address,
	}
}

func (s *EventsServiceServer) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		s.logger.Fatal("failed to listen", zap.Error(err))
	}

	s.server = grpc.NewServer(grpc.UnaryInterceptor(s.LoggerInterceptor))
	RegisterEventServiceServer(s.server, s)

	s.logger.Info("gRPC server started", zap.String("address", s.address))
	if err := s.server.Serve(listener); err != nil {
		s.logger.Fatal("failed to serve", zap.Error(err))
	}

	<-ctx.Done()
	return nil
}

func (s *EventsServiceServer) Stop(_ context.Context) error {
	s.server.GracefulStop()

	os.Exit(1)
	return nil
}
