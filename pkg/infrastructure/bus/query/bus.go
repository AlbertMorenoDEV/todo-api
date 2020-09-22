package query

import "fmt"

type Bus struct {
	handlersMap map[string]Handler
}

func NewBus() Bus {
	return Bus{handlersMap: make(map[string]Handler)}
}

func (b *Bus) RegisterHandler(q Query, qh Handler) error {
	qrID := q.QueryID()

	_, ok := b.handlersMap[qrID]
	if ok {
		return fmt.Errorf("the Query %s is already register", qrID)
	}
	b.handlersMap[qrID] = qh
	return nil
}

func (b Bus) Ask(q Query) (*Response, error) {
	qrID := q.QueryID()

	qh, ok := b.handlersMap[qrID]
	if !ok {
		return nil, fmt.Errorf("there not any QueryHandler associate to Query %s", qrID)
	}
	return qh.Handle(q)
}
