package routes

import (
	"develop/dev11/task11/calendar"
	"develop/dev11/task11/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	c calendar.Calendar
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)full body"}`))
		return
	}
	fields := make(map[string]string)
	json.Unmarshal(body, &fields)

	var event models.Event
	log.Println(fields["date"])
	event.EventDate, err = time.Parse("2006.01.02", fields["date"])

	if err != nil {
		log.Fatalf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)time"}`))
		return
	}
	event.Description = fields["event"]

	h.c.CreateEvent(event)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result":"Event created"}`))
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}
	fields := make(map[string]string)
	json.Unmarshal(body, &fields)

	var event models.Event
	event.ID, err = strconv.Atoi(fields["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}
	event.EventDate, err = time.Parse("2006.01.02", fields["date"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}
	event.Description = fields["event"]

	h.c.UpdateEvent(event)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result":"Event updated"}`))
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}
	fields := make(map[string]string)
	json.Unmarshal(body, &fields)
	var event models.Event
	event.ID, err = strconv.Atoi(fields["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}

	h.c.DeleteEvent(event)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result":"Event deleted"}`))
}

func (h *Handler) ShowEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}

	date, err := time.Parse("2006.01.02", r.URL.Query().Get("date"))
	if err != nil {
		date = time.Now()
	}

	events := h.c.ShowEvents(date, r.URL.Path[16:])
	result, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"json"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
