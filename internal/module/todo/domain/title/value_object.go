package title

type Title struct {
	value string
}

func New(value string) (Title, error) {
	var vo Title
	vo.value = value
	return vo, nil
}

func (vo Title) String() string {
	return vo.value
}

func (vo Title) EqualsTo(other Title) bool {
	return vo.String() == other.String()
}
