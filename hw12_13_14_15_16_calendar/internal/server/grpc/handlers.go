package internalgrpc

import (
	"context"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"go.uber.org/zap"
	"time"
)

func (s *EventsServiceServer) CreateEvent(_ context.Context, req *RequestCreateEvent) (*Event, error) {
	var event model.Event
	event.Title = req.Title
	event.Description = req.Description
	event.UserID = int(req.UserID)

	err := s.app.CreateEvent(&event)
	if err != nil {
		s.logger.Error("Failed to create event", zap.Error(err))
		return nil, err
	}

	pbEvent := Event{
		ID:          int64(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserID:      int64(event.UserID),
	}

	return &pbEvent, nil
}

func (s *EventsServiceServer) GetEvent(_ context.Context, req *RequestGetEvent) (*Event, error) {
	event, err := s.app.GetEvent(int(req.ID))
	if err != nil {
		s.logger.Error("Failed to get event", zap.Error(err))
		return nil, err
	}

	pbEvent := Event{
		ID:          int64(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserID:      int64(event.UserID),
	}

	return &pbEvent, nil
}

func (s *EventsServiceServer) UpdateEvent(_ context.Context, req *RequestUpdateEvent) (*Event, error) {
	event, err := s.app.GetEvent(int(req.ID))
	if err != nil {
		s.logger.Error("Failed to get event", zap.Error(err))
		return nil, err
	}

	var updatedEvent model.Event
	updatedEvent.ID = int(req.ID)
	updatedEvent.Title = req.Title
	updatedEvent.Description = req.Description
	updatedEvent.UserID = int(req.UserID)

	err = s.app.UpdateEvent(event.ID, &updatedEvent)
	if err != nil {
		s.logger.Error("Failed to update event", zap.Error(err))
		return nil, err
	}

	pbEvent := Event{
		ID:          int64(event.ID),
		Title:       event.Title,
		Description: event.Description,
		UserID:      int64(event.UserID),
	}

	return &pbEvent, nil
}

func (s *EventsServiceServer) DeleteEvent(_ context.Context, req *RequestDeleteEvent) (*Event, error) {
	_, err := s.app.GetEvent(int(req.ID))
	if err != nil {
		s.logger.Error("Event not found", zap.Error(err))
		return nil, err
	}

	err = s.app.DeleteEvent(int(req.ID))
	if err != nil {
		s.logger.Error("Failed to delete event", zap.Error(err))
		return nil, err
	}
	pbEvent := Event{}

	return &pbEvent, nil
}

func (s *EventsServiceServer) ListByDay(_ context.Context, req *RequestListByDate) (*ResponseListByDate, error) {
	date, err := time.Parse("2006-01-02", req.DateTme)
	if err != nil {
		s.logger.Error("Failed to parse date", zap.Error(err))
		return nil, err
	}

	events, err := s.app.ListByDay(date)
	if err != nil {
		s.logger.Error("Failed to get events", zap.Error(err))
		return nil, err
	}

	pbEvents := ResponseListByDate{}

	for _, event := range events {
		pbEvents.Events = append(pbEvents.Events, &Event{
			ID:          int64(event.ID),
			Title:       event.Title,
			Description: event.Description,
			UserID:      int64(event.UserID),
		})
	}

	return &pbEvents, nil
}

func (s *EventsServiceServer) ListByWeek(_ context.Context, req *RequestListByDate) (*ResponseListByDate, error) {
	date, err := time.Parse("2006-01-02", req.DateTme)
	if err != nil {
		s.logger.Error("Failed to parse date", zap.Error(err))
		return nil, err
	}

	events, err := s.app.ListByWeek(date)
	if err != nil {
		s.logger.Error("Failed to get events", zap.Error(err))
		return nil, err
	}

	pbEvents := ResponseListByDate{}

	for _, event := range events {
		pbEvents.Events = append(pbEvents.Events, &Event{
			ID:          int64(event.ID),
			Title:       event.Title,
			Description: event.Description,
			UserID:      int64(event.UserID),
		})
	}

	return &pbEvents, nil
}

func (s *EventsServiceServer) ListByMonth(_ context.Context, req *RequestListByDate) (*ResponseListByDate, error) {
	date, err := time.Parse("2006-01-02", req.DateTme)
	if err != nil {
		s.logger.Error("Failed to parse date", zap.Error(err))
		return nil, err
	}

	events, err := s.app.ListByMonth(date)
	if err != nil {
		s.logger.Error("Failed to get events", zap.Error(err))
		return nil, err
	}

	pbEvents := ResponseListByDate{}

	for _, event := range events {
		pbEvents.Events = append(pbEvents.Events, &Event{
			ID:          int64(event.ID),
			Title:       event.Title,
			Description: event.Description,
			UserID:      int64(event.UserID),
		})
	}

	return &pbEvents, nil
}
