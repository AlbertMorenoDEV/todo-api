package identifier

type Identifier string

func New(value string) (Identifier, error) {
	return Identifier(value), nil
}

func (vo Identifier) String() string {
	return string(vo)
}

func (vo Identifier) EqualsTo(other Identifier) bool {
	return vo.String() == other.String()
}
