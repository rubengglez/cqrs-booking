package infrastructure

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rubengglez/cqrs-booking/internal/modules/room/infrastructure"
	room_ui "github.com/rubengglez/cqrs-booking/internal/modules/room/ui"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func Run() {
	r := mux.NewRouter()
	roomRepository := infrastructure.InMemoryRoomRepository{}
	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/rooms", room_ui.NewRoomsHandler(roomRepository))

	http.ListenAndServe("localhost:10001", r)
}
