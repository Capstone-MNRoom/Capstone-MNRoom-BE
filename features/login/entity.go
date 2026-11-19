package login

import "context"

type Core struct {
	ID       int
	Email    string
	Password string
}

type Business interface {
	Auth(ctx context.Context, data Core) (dataAuth Core, err error)
}

type Data interface {
	Auth(ctx context.Context, data Core) (dataAuth Core, err error)
}
