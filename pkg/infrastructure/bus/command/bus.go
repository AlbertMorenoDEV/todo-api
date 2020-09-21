package command

import "fmt"

type Bus struct {
	handlersMap map[string]Handler
}

func NewBus() Bus {
	return Bus{handlersMap: make(map[string]Handler)}
}

func (b *Bus) RegisterHandler(c Command, ch Handler) error {
	cmdID := c.CommandID()

	_, ok := b.handlersMap[cmdID]
	if ok {
		return fmt.Errorf("the Command %s is already register", cmdID)
	}
	b.handlersMap[cmdID] = ch
	return nil
}

func (b Bus) Dispatch(c Command) error {
	cmdID := c.CommandID()

	ch, ok := b.handlersMap[cmdID]
	if !ok {
		return fmt.Errorf("there not any CommandHandler associate to Command %s", cmdID)
	}
	return ch.Handle(c)
}
