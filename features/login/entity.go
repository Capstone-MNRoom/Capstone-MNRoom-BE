package login

type Core struct {
	ID       int
	Email    string
	Password string
}

type Business interface {
	Auth(data Core) (dataAuth Core, err error)
}

type Data interface {
	Auth(data Core) (dataAuth Core, err error)
}
