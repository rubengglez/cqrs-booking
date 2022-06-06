package domain

type RoomRepository interface {
	AllAvailable() Rooms
}
