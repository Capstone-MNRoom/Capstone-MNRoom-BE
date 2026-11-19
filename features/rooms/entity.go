package rooms

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/users"
	"context"
	"time"
)

type Core struct {
	ID             int
	ImageRoom      string
	ImagePengelola string
	RoomName       string
	Capacity       int
	HotelName      string
	RentalPrice    int
	Address        string
	City           string
	Facilitys      []int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           users.Core
	Categorys      categorys.Core
}

type CoreRoomFacilitys struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      users.Core
	Rooms     Core
	Facilitys facilitys.Core
}

type Business interface {
	InsertData(ctx context.Context, insert Core) (data Core, err error)
	InsertDataRoomFacilitys(ctx context.Context, insert CoreRoomFacilitys) (row int, err error)
	GetDataAll(ctx context.Context, page int) (data []Core, err error)
	GetData(ctx context.Context, id int) (data Core, err error)
	UpdateData(ctx context.Context, id int, insert Core) (row int, err error)
	DeleteData(ctx context.Context, id int) (row int, err error)
	GetToken(ctx context.Context, id int, idToken int) (data Core, err error)
	GetDataAllUserRoom(ctx context.Context, idToken int) (data []Core, err error)
	GetDataIDRoom(ctx context.Context, id int) (row int, err error)
	GetDataByCategory(ctx context.Context, id int) (data []Core, err error)
}

type Data interface {
	InsertData(ctx context.Context, insert Core) (data Core, err error)
	InsertDataRoomFacilitys(ctx context.Context, insert CoreRoomFacilitys) (row int, err error)
	GetDataAll(ctx context.Context, page int) (data []Core, err error)
	GetData(ctx context.Context, id int) (data Core, err error)
	UpdateData(ctx context.Context, id int, insert Core) (row int, err error)
	DeleteData(ctx context.Context, id int) (row int, err error)
	GetToken(ctx context.Context, id int, idToken int) (data Core, err error)
	GetDataAllUserRoom(ctx context.Context, idToken int) (data []Core, err error)
	GetDataIDRoom(ctx context.Context, id int) (row int, err error)
	GetDataByCategory(ctx context.Context, id int) (data []Core, err error)
}
