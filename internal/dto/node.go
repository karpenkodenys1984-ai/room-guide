package dto

type Position struct {
	X int
	Y int
}

type RoomType string

const (
	RoomTypeRoom     RoomType = "room"
	RoomTypeEntrance RoomType = "entrance"
	RoomTypeExit     RoomType = "exit"
)

type RoomData struct {
	Label string
	Type  RoomType
}

type Node struct {
	Id          int16
	Type        string
	Position    Position
	Data        RoomData
	Draggable   *bool
	Selectable  *bool
	Connectable *bool
}
