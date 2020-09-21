package command

type Command interface {
	CommandID() string
}

type Handler interface {
	Handle(Command) error
}
