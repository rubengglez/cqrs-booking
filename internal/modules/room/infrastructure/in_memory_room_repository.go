package infrastructure

import "github.com/rubengglez/cqrs-booking/internal/modules/room/domain"

type InMemoryRoomRepository struct{}

func (r InMemoryRoomRepository) AllAvailable() domain.Rooms {
	result := make([]domain.Room, 0)
	result = append(result, domain.Room{
		Number: 101,
	})
	result = append(result, domain.Room{
		Number: 102,
	})
	return result
}
