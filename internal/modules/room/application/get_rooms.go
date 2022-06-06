package application

import domain "github.com/rubengglez/cqrs-booking/internal/modules/room/domain"

type RoomResponse struct {
	Number int
}

type RoomsResponse struct {
	Rooms []RoomResponse
}

func GetRooms(r domain.RoomRepository) (*RoomsResponse, error) {
	rooms := r.AllAvailable()

	return toRoomsResponse(rooms), nil
}

func toRoomsResponse(rooms domain.Rooms) *RoomsResponse {
	roomsResponse := RoomsResponse{}
	roomsResponse.Rooms = make([]RoomResponse, 0)
	for _, room := range rooms {
		response := toRoomResponse(&room)
		roomsResponse.Rooms = append(roomsResponse.Rooms, *response)
	}
	return &roomsResponse
}

func toRoomResponse(r *domain.Room) *RoomResponse {
	return &RoomResponse{
		Number: r.Number,
	}
}
