package identifier

type Identifier struct {
	value string
}

func New(value string) (Identifier, error) {
	var vo Identifier
	vo.value = value
	return vo, nil
}

func (vo Identifier) String() string {
	return vo.value
}

func (vo Identifier) EqualsTo(other Identifier) bool {
	return vo.String() == other.String()
}
