package identifier

type Identifier string

func New(value string) (Identifier, error) {
	var vo Identifier
	vo = Identifier(value)
	return vo, nil
}

func (vo Identifier) String() string {
	return string(vo)
}

func (vo Identifier) EqualsTo(other Identifier) bool {
	return vo.String() == other.String()
}
