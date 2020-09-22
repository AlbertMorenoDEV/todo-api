package query

type Query interface {
	QueryID() string
}

type Response interface {
	ResponseID() string
}

type Handler interface {
	Handle(q Query) (*Response, error)
}
