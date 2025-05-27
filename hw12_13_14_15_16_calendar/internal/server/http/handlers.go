package internalhttp

import (
	"encoding/json"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World"))
}

func createEvent(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.CreateEvent(&event)
	if err != nil {
		http.Error(w, "Error create event "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func getEvent(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	event, err := app.GetEvent(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func updateEvent(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	_, err = app.GetEvent(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	var updatedEvent model.Event
	err = json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	updatedEvent.ID = id
	err = app.UpdateEvent(id, &updatedEvent)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error update event:"+err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEvent)
}

func deleteEvent(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	_, err = app.GetEvent(id)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	err = app.DeleteEvent(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error delete event:"+err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func listByDay(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	dateStr := vars["date"]

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Please use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	events, err := app.ListByDay(date)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error list by day:"+err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func listByWeek(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	dateStr := vars["date"]

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Please use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	events, err := app.ListByWeek(date)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error list by day:"+err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func listByMonth(w http.ResponseWriter, r *http.Request, app Application) {
	lock.Lock()
	defer lock.Unlock()

	vars := mux.Vars(r)
	dateStr := vars["date"]

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Please use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	events, err := app.ListByMonth(date)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error list by day:"+err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
