package title

type Title string

func New(value string) (Title, error) {
	return Title(value), nil
}

func (vo Title) String() string {
	return string(vo)
}

func (vo Title) EqualsTo(other Title) bool {
	return vo.String() == other.String()
}
