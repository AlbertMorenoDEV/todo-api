package jsonapi

type Error struct {
	ID     string `json:"id"`
	Status int64  `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type Errors []Error
