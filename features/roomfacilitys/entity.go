package roomfacilitys

import (
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
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
	GetData(id int) (data []Core, err error)
}

type Data interface {
	GetData(id int) (data []Core, err error)
}
