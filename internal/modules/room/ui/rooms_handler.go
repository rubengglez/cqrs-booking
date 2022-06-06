package ui

import (
	"encoding/json"
	"log"
	"net/http"

	application "github.com/rubengglez/cqrs-booking/internal/modules/room/application"
	"github.com/rubengglez/cqrs-booking/internal/modules/room/domain"
)

type handler func(w http.ResponseWriter, r *http.Request)

func NewRoomsHandler(repo domain.RoomRepository) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		roomsResponse, err := application.GetRooms(repo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		log.Println(roomsResponse.Rooms)
		json.NewEncoder(w).Encode(roomsResponse)
	}
}
