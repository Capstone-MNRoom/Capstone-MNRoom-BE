package roomfacilitys

import (
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
	"context"
	"time"
)

type Core struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      users.Core
	Rooms     rooms.Core
	Facilitys facilitys.Core
}

type Business interface {
	GetData(ctx context.Context, id int) (data []Core, err error)
	GetDataRow(ctx context.Context, id int) (row int, err error)
}

type Data interface {
	GetData(ctx context.Context, id int) (data []Core, err error)
	GetDataRow(ctx context.Context, id int) (row int, err error)
}
